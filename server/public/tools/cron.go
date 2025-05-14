// tools/cron.go
package tools

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// JobFunc 定义任务函数签名，返回 error
type JobFunc func() error

// Cron 定时任务调度器结构体
type Cron struct {
	// 存储业务函数列表
	absTimeJobs  map[int64][]JobFunc  // 一次性绝对时间任务集合
	intervalJobs map[int][]JobFunc    // 间隔任务集合
	cronExprJobs map[string][]JobFunc // Cron 表达式任务集合

	// 存储 gocron.Job 引用，用于后续移除
	absTimeJobMap  map[int64]*gocron.Job
	intervalJobMap map[int]*gocron.Job
	cronExprJobMap map[string]*gocron.Job

	Scheduler *gocron.Scheduler // gocron 调度器实例
	NotifyCh  chan []JobFunc    // 任务通知通道
}

// NewCron 创建并返回一个新的 Cron 实例
func NewCron() *Cron {
	sch := gocron.NewScheduler(time.Local)
	return &Cron{
		absTimeJobs:    make(map[int64][]JobFunc),
		intervalJobs:   make(map[int][]JobFunc),
		cronExprJobs:   make(map[string][]JobFunc),
		absTimeJobMap:  make(map[int64]*gocron.Job),
		intervalJobMap: make(map[int]*gocron.Job),
		cronExprJobMap: make(map[string]*gocron.Job),
		Scheduler:      sch,
		NotifyCh:       make(chan []JobFunc, 100),
	}
}

// AddAbsTimeJob 添加一次性绝对时间任务
func (cj *Cron) AddAbsTimeJob(startTime time.Time, jFunc JobFunc) {
	ts := startTime.Unix()
	// 如果时间已经过去，直接跳过
	if startTime.Before(time.Now()) {
		return
	}

	if _, exists := cj.absTimeJobs[ts]; !exists {
		job, _ := cj.Scheduler.Every(1).
			LimitRunsTo(1).
			StartAt(time.Unix(ts, 0)).
			Do(func() {
				cj.NotifyCh <- cj.absTimeJobs[ts]
			})
		cj.absTimeJobMap[ts] = job
	}
	cj.absTimeJobs[ts] = append(cj.absTimeJobs[ts], jFunc)
}

// AddIntervalJob 添加间隔任务
func (cj *Cron) AddIntervalJob(seconds int, jFunc JobFunc) {
	if _, exists := cj.intervalJobs[seconds]; !exists {
		// 首次 StartAt 设为 now + interval
		firstRun := time.Now().Add(time.Duration(seconds) * time.Second)
		job, _ := cj.Scheduler.Every(seconds).Seconds().StartAt(firstRun).
			Do(func() {
				cj.NotifyCh <- cj.intervalJobs[seconds]
			})
		cj.intervalJobMap[seconds] = job
	}
	cj.intervalJobs[seconds] = append(cj.intervalJobs[seconds], jFunc)
}

// AddCronExprJob 添加 Cron 表达式任务
func (cj *Cron) AddCronExprJob(expr string, jFunc JobFunc) error {
	if _, exists := cj.cronExprJobs[expr]; !exists {
		job, err := cj.Scheduler.Cron(expr).
			Do(func() {
				cj.NotifyCh <- cj.cronExprJobs[expr]
			})
		if err != nil {
			return fmt.Errorf("无效的 Cron 表达式 %s: %w", expr, err)
		}
		cj.cronExprJobMap[expr] = job
	}
	cj.cronExprJobs[expr] = append(cj.cronExprJobs[expr], jFunc)
	return nil
}

// RemoveAbsTimeJob 移除一次性任务
func (cj *Cron) RemoveAbsTimeJob(ts int64) {
	if job, ok := cj.absTimeJobMap[ts]; ok {
		cj.Scheduler.RemoveByReference(job)
		delete(cj.absTimeJobMap, ts)
		delete(cj.absTimeJobs, ts)
	}
}

// RemoveIntervalJob 移除间隔任务
func (cj *Cron) RemoveIntervalJob(interval int) {
	if job, ok := cj.intervalJobMap[interval]; ok {
		cj.Scheduler.RemoveByReference(job)
		delete(cj.intervalJobMap, interval)
		delete(cj.intervalJobs, interval)
	}
}

// RemoveCronExprJob 移除 Cron 表达式任务
func (cj *Cron) RemoveCronExprJob(expr string) {
	if job, ok := cj.cronExprJobMap[expr]; ok {
		cj.Scheduler.RemoveByReference(job)
		delete(cj.cronExprJobMap, expr)
		delete(cj.cronExprJobs, expr)
	}
}

var CronJob *Cron

func init() {
	CronJob = NewCron()
	CronJob.Scheduler.StartAsync()
	// 后台消费通知并执行
	go func() {
		for funcs := range CronJob.NotifyCh {
			for _, fn := range funcs {
				if err := fn(); err != nil {
					fmt.Println("任务执行出错:", err)
				}
			}
		}
	}()
}
