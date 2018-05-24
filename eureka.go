package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"
	"net"
	"time"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
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
	if err := eureka.RegisterInstance(instance); err != nil {
		log.Panic(err.Error())
	}
	go startHeartBeat()
}

func startHeartBeat() {
	for {
		if err := eureka.HeartBeatInstance(instance); err != nil {
			eureka.ReregisterInstance(instance)
		}
		time.Sleep(10 * time.Second)
	}
}
