package pushDingding

import (
	"gin/common/alarm"
	"gin/common/function"
	"gin/config"
	"github.com/tidwall/gjson"
	"time"
)

type PushDingding struct {
}

func (pushDingding PushDingding) Run(topic string, pushTime int64, data string) {
	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			alarm.DingDing(topic + "执行异常")
		}
	}()

	content := gjson.Get(data, "content")
	title := gjson.Get(data, "title")

	if !content.Exists() {
		return
	}

	contentString := content.String()
	titleString := title.String()

	time.Sleep(time.Duration(3) * time.Second)

	function.SendToDingDing(config.Config.DINGding.Xueqiucommentpush, titleString, contentString)
	function.SendToWeixin(config.Config.Weixin.Xueqiucommentpush, contentString)

}
