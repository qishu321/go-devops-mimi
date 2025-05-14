package cronjob

import (
	"go-devops-mimi/server/logic/exec_logic"
	"go-devops-mimi/server/model/exec"
	"go-devops-mimi/server/public/common"
	"log"
)

func InitCronJobs() {
	registerCronJob()
}

// 初始化定时任务
func registerCronJob() {
	var crons []exec.Cron

	// 只加载“启用”状态的任务
	if err := common.DB.Where("enable = ?", 1).Find(&crons).Error; err != nil {
		log.Fatalf("初始化定时任务失败：无法从数据库读取 cron 记录: %v", err)
	}

	log.Printf("初始化定时任务：共加载 %d 条记录\n", len(crons))

	// 正确遍历取指针：按索引拿地址，避免 &c 指向同一个循环变量
	for i := range crons {
		c := &crons[i]
		log.Printf("正在注册 cron ID=%d, 类型=%s\n", c.ID, c.CronType)
		exec_logic.SendCron(c)
	}

	log.Println("定时任务初始化完成")

}
