package tools

import (
	"errors"
	"sync"
	"time"
)

var ErrExecTimeout = errors.New("pool execution timed out")

type Pool struct {
	taskQueue         chan func()
	minWorkers        int           // 最小启动的 worker 数量
	maxWorkers        int           // 最大 worker 数量
	workerIdleTimeout time.Duration // worker 空闲超时时间

	mu      sync.Mutex // 保护 currentWorkers
	current int        // 当前正在运行的 worker 数量

	done chan struct{} // 用于关闭池的信号
	wg   sync.WaitGroup
}

// NewPool 创建一个协程池，
// minWorkers: 最小长期运行的 worker 数量 (>=0)
// maxWorkers: 最大 worker 数量 (>= minWorkers)
// // queueSize: 任务队列缓冲大小（队列满后才会扩容协程）
// workerIdleTimeout: worker 空闲多久后退出（如果当前 worker 数量超过 minWorkers）
func NewPool(minWorkers, maxWorkers, queueSize int, workerIdleTimeout time.Duration) *Pool {
	if minWorkers < 0 || maxWorkers < minWorkers || queueSize < 0 {
		panic("invalid pool configuration")
	}

	p := &Pool{
		taskQueue:         make(chan func(), queueSize),
		minWorkers:        minWorkers,
		maxWorkers:        maxWorkers,
		workerIdleTimeout: workerIdleTimeout,
		current:           0,
		done:              make(chan struct{}),
	}

	// 预启动 minWorkers 个 worker
	for i := 0; i < p.minWorkers; i++ {
		p.startWorker()
	}

	return p
}

// startWorker 启动一个 worker goroutine
func (p *Pool) startWorker() {
	p.mu.Lock()
	if p.current >= p.maxWorkers {
		p.mu.Unlock()
		return
	}
	p.current++
	p.mu.Unlock()

	p.wg.Add(1)
	go p.worker()
}

// worker 是 worker goroutine 的主体
func (p *Pool) worker() {
	defer p.wg.Done()
	// 定义一个空闲定时器，用于在空闲超时后退出 worker
	idleTimer := time.NewTimer(p.workerIdleTimeout)
	defer idleTimer.Stop()

	for {
		select {
		// 优先从任务队列中获取任务
		case task, ok := <-p.taskQueue:
			if !ok {
				// 任务队列关闭，退出 worker
				return
			}
			// 执行任务
			task()
			// 每次执行完任务后重置定时器
			if !idleTimer.Stop() {
				<-idleTimer.C
			}
			idleTimer.Reset(p.workerIdleTimeout)
		// 如果长时间没有收到任务，则判断是否可以退出（当前 worker 数量大于最小值时退出）
		case <-idleTimer.C:
			p.mu.Lock()
			if p.current > p.minWorkers {
				p.current--
				p.mu.Unlock()
				return
			}
			p.mu.Unlock()
			// 如果达到最小 worker 数量，则重置定时器继续等待任务
			idleTimer.Reset(p.workerIdleTimeout)
		// 收到关闭池的信号，退出 worker
		case <-p.done:
			return
		}
	}
}

// Exec 提交任务到池中（无超时控制）
func (p *Pool) Exec(task func()) error {
	select {
	case p.taskQueue <- task:
		// 任务入队成功
		return nil
	default:
		// 任务队列已满，尝试启动新 worker 以处理任务
		p.mu.Lock()
		if p.current < p.maxWorkers {
			// 启动新的 worker
			p.startWorker()
			p.mu.Unlock()
			// 再次提交任务，这里可以保证任务会被及时处理
			p.taskQueue <- task
			return nil
		}
		p.mu.Unlock()
		// 如果当前 worker 数量已达到最大且队列满，则返回超时错误
		return ErrExecTimeout
	}
}

// ExecTimeout 提交任务，并等待指定超时（注意：这里仅仅等待任务能否入队）
func (p *Pool) ExecTimeout(task func(), timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case p.taskQueue <- task:
		return nil
	case <-timer.C:
		return ErrExecTimeout
	}
}

// Shutdown 优雅关闭池，等待所有任务完成
func (p *Pool) Shutdown() {
	// 关闭 done 信道，用于通知 worker 退出
	close(p.done)
	// 关闭任务队列，通知所有等待任务的 worker
	close(p.taskQueue)
	// 等待所有 worker 退出
	p.wg.Wait()
}
