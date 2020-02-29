package main

import (
	"context"
	"flag"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/bilibili/kratos/pkg/naming"
	"github.com/bilibili/kratos/pkg/naming/discovery"
	"kratos-server/internal/di"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func registerWarden() context.CancelFunc {

	ip := "127.0.0.1"   // NOTE: 必须拿到您实例节点的真实IP，
	port := "9000" // NOTE: 必须拿到您实例grpc监听的真实端口，warden默认监听9000
	hn, _ := os.Hostname()
	dis := discovery.New(nil)
	ins := &naming.Instance{
		Zone:     env.Zone,
		Env:      env.DeployEnv,
		AppID:    "server.service",
		Hostname: hn,
		Addrs: []string{
			"grpc://" + ip + ":" + port,
		},
	}
	cancel, err := dis.Register(context.Background(), ins)
	if err != nil {
		panic(err)
	}

	// 省略...

	// 特别注意！！！
	// cancel必须在进程退出时执行！！！
	return cancel
}

func main() {
	flag.Parse()
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("kratos-server start")
	paladin.Init()
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	cancel := registerWarden()

	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("kratos-server exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
