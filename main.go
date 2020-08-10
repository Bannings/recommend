package main

import (
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"os"
	"recommend/golbal"
	"recommend/log"
	"recommend/router"
	"recommend/router/middleware"
	"recommend/util/db"
	"recommend/util/timed_task"
)

var (
	logLevel = flag.String("L", "info", "log level: info, debug, warn, error, fatal")
	logFile  = flag.String("logfile", "", "log file path")
)

func init() {
	lvl, err := log.ParseLevel(*logLevel)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	if len(*logFile) > 0 {
		log.Std = log.NewRotate(*logFile, "", log.Ldefault, lvl)
	} else {
		log.Std = log.New(os.Stderr, "", log.Ldefault, lvl)
	}
	if _, err := golbal.LoadConfig("prod.json"); err != nil {
		log.Fatal(err)
	}
	if log.DataLogger, err = seelog.LoggerFromConfigAsFile("seelog_homepage_request.xml"); err != nil {
		log.Fatal(err)
	}

	db.InitRedis()
	timed_task.ReloadBannerData()
	timed_task.ReloadTimedTaskData()
	timed_task.LoadNewDescription()

}

func main() {
	go timed_task.ExcuteTimedTask()
	go middleware.HandleError()
	g := router.InitRouter()
	conf := golbal.GetConfig()
	err := g.Run(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		fmt.Printf("start sever err: %v", err)
	}
}
