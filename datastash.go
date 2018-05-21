package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	port       = 9999
	appName    = "datastash"
	hostName   = "localhost"
	ipAddr     = "localhost"
	eurekaHost = "http://localhost:8761/eureka"
)

var (
	eureka  fargo.EurekaConnection
	baseUrl = "http://" + ipAddr + ":" + strconv.Itoa(port)
	// 当前微服务实例描述
	instance = &fargo.Instance{
		HostName:         hostName,
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

func main() {
	go enableEurekaClient()
	go listenDestroy()

	r := gin.Default()
	r.GET("/health", health)
	r.GET("/info", info)
	r.Run(":" + strconv.Itoa(port)) // listen and serve on 0.0.0.0:9999
}

// Eureka 健康检查 API Handler
func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}

// Eureka 状态页 API Handler
func info(c *gin.Context) {
	c.JSON(200, gin.H{
		"project": gin.H{
			"version":     "0.0.1-SNAPSHOT",
			"title":       appName,
			"description": "Golang micro service",
		},
	})
}

// 启用 Eureka Client，注册到 Spring Eureka 注册中心
func enableEurekaClient() {
	eureka = fargo.NewConn(eurekaHost)
	eureka.RegisterInstance(instance)
}

// 监听程序结束信号，执行 destroy 方法
func listenDestroy() {
	c := make(chan os.Signal)
	// 监听指定信号 ctrl+c kill
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	// 阻塞直至有信号传入
	s := <-c
	destroy(s)
}

// 反注册 Eureka Client
func destroy(_ os.Signal) {
	eureka.DeregisterInstance(instance)
}
