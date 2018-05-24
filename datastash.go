package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"
	"strconv"
	"os"
	"os/signal"
	"syscall"
)

const (
	port       = 9999
	appName    = "datastash"
	eurekaHost = "http://localhost:8761/eureka"
	mongoUrl   = "mongodb://localhost:27017"
)

var (
	ipAddr  = getLocalIP()
	eureka  fargo.EurekaConnection
	baseUrl = "http://" + ipAddr + ":" + strconv.Itoa(port)
	// 当前微服务实例描述
	instance = &fargo.Instance{
		HostName:         ipAddr,
		Port:             port,
		App:              appName,
		IPAddr:           ipAddr,
		VipAddress:       ipAddr,
		SecureVipAddress: ipAddr,
		HealthCheckUrl:   baseUrl + "/health",
		StatusPageUrl:    baseUrl + "/info",
		HomePageUrl:      baseUrl,
		Status:           fargo.UP,
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		LeaseInfo:        fargo.LeaseInfo{RenewalIntervalInSecs: 1},
	}
)

func init() {
	go enableEurekaClient()
	go listenNotify()
	connectMongo(mongoUrl)
}

func main() {
	r := gin.Default()
	r.GET("/health", health)
	r.GET("/info", info)
	r.POST("/rpc/stash", stash)
	r.Run(":" + strconv.Itoa(port)) // listen and serve on 0.0.0.0:9999
}

// 监听程序结束信号，执行 destroy 方法
func listenNotify() {
	c := make(chan os.Signal)
	// 监听指定信号 ctrl+c kill
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	// 阻塞直至有信号传入
	s := <-c
	onDestroy(s)
}

// 反注册 Eureka Client
func onDestroy(_ os.Signal) {
	eureka.DeregisterInstance(instance)
}