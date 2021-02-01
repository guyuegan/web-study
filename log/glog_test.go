package log

import (
	"flag"
	"github.com/golang/glog"
	"testing"
)

const AppName = "web-study"

//program arg: -v=4 -alsologtostderr -log_backtrace_at=glog_test.go:42
// -v=4: 日志级别小于或等于 4 的日志将被打印出来

func TestGlog(t *testing.T) {
	_ = flag.Set("log_dir", "/data/logs/"+AppName)
	flag.Parse()
	if flag.NFlag() < 1 { //NFlag返回解析时进行了设置的flag的数量。
		flag.Usage() //当参数小于1个打印flag的默认提示
		return
	}
	//程序退出时确保写入文件中
	defer glog.Flush()

	glog.Info("this is info msg")
	glog.Infof("this is info msg: %v", 12345)
	glog.InfoDepth(1, "this is info msg", 12345)

	glog.Warning("this is warning msg")
	glog.Warningf("this is warning msg: %v", 12345)
	glog.WarningDepth(1, "this is warning msg", 12345)

	glog.Error("this is error msg")
	glog.Errorf("this is error msg: %v", 12345)
	glog.ErrorDepth(1, "this is error msg", 12345)

	// 致命错误日志，打印完日志后程序将会推出（os.Exit()）
	//glog.Fatal("this is fatal msg")
	//glog.Fatalf("this is fatal msg: %v", 12345)
	//glog.FatalDepth(1, "this is fatal msg", 12345)

	//	v level
	glog.V(3).Info("level 3 msg")
	glog.V(4).Info("level 4 msg")
	glog.V(5).Info("level 5 msg")
	glog.V(6).Info("level 6 msg")
}
