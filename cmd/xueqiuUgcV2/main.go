package main

import (
	"fmt"
	"gin/cmd/xueqiuUgcV2/job"
	"gin/common/env"
	"gin/common/function"
	"gin/core"
	"gin/core/redis"
	"github.com/reugn/go-quartz/quartz"
)

func main() {
	core.InitDb()
	redis.InitRedis()
	env.Active().Value()
	sampleJobs2()
}

func sampleJobs2() {
	sched := quartz.NewStdScheduler()
	sched.Start()

	cronTrigger, err := quartz.NewCronTrigger("0 0/10 * * * *")
	//cronTrigger := quartz.NewRunOnceTrigger(time.Second * 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	exitChan := make(chan bool, 1)

	cronJob := job.Xueqiu{"Cron job", exitChan}

	//for {

	start_time := function.ExeStart()
	sched.ScheduleJob(&cronJob, cronTrigger)

	for i := 0; i < 1; i++ {
		<-exitChan
	}

	function.PrintUseTime(start_time)

	sched.DeleteJob(cronJob.Key())

	//}

	sched.Stop()
}
