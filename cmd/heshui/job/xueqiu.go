package job

import (
	"gin/common/function"
	"gin/config"
	"github.com/reugn/go-quartz/quartz"
	"github.com/syyongx/php2go"
	"time"
)

// PrintJob implements the quartz.Job interface.
type Xueqiu struct {
	Desc   string
	Signal chan bool
}

// Description returns the description of the PrintJob.
func (pj *Xueqiu) Description() string {
	return pj.Desc
}

// Key returns the unique PrintJob key.
func (pj *Xueqiu) Key() int {
	return quartz.HashCode(pj.Description())
}

// Execute is called by a Scheduler when the Trigger associated with this job fires.
func (pj *Xueqiu) Execute() {
	currentHour := time.Now().Hour()

	allow := []int{
		8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22,
	}

	title := "该喝水啦~~~"
	if php2go.InArray(currentHour, allow) == true {
		function.SendToDingDingAt(config.Config.DINGding.Heshuipush, title, title, "18163912092")
	}
}
