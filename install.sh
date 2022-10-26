xiangmu='gin-hong-api'
xiangmupath='/golang/go/src/'
xiangmubinpath='/golang/go/bin'
cd /data/git/$xiangmu/
cd $xiangmupath$xiangmu/
go install .
cd $xiangmupath$xiangmu/cmd/xueqiu
go install .
cd $xiangmupath$xiangmu/cmd/xueqiuv1.1
go install .
cd $xiangmupath$xiangmu/cmd/xueqiuCombinationUpdate
go install .
cd $xiangmupath$xiangmu/cmd/xueqiuCombination
go install .
cd $xiangmupath$xiangmu/cmd/xueqiuUgc
go install .
cd $xiangmupath$xiangmu/cmd/queue
go install .
cd $xiangmupath$xiangmu/cmd/pushRank
go install .
cd $xiangmupath$xiangmu/cmd/heshui
go install .
cd $xiangmupath$xiangmu/cmd/xueqiuUgcV2
go install .
#运行任务
#nohup $xiangmubinpath/gin -env pro > /tmp/gin.log &
#nohup $xiangmubinpath/xueqiu -env pro > /tmp/xueqiu.log &
#nohup $xiangmubinpath/xueqiuv1.1 -env pro > /tmp/xueqiuv1.1.log &

#重启所有命令
supervisorctl reload
