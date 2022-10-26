package main

import (
	"fmt"
	"gin/cmd/heshui/job"
	"gin/common/env"
	"gin/common/function"
	"github.com/reugn/go-quartz/quartz"
)

func main() {
	env.Active().Value()
	sampleJobs2()
}

func sampleJobs2() {
	sched := quartz.NewStdScheduler()
	sched.Start()
	cronTrigger, err := quartz.NewCronTrigger("0 0/15 * * * *")
	//cronTrigger := quartz.NewRunOnceTrigger(time.Second * 1)

	//cronTrigger, err := quartz.NewCronTrigger("0 0/1 * * * *")
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
