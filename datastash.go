package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"
	"strconv"
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
	go listenDestroy()
	connectMongo(mongoUrl)
}

func main() {
	r := gin.Default()
	r.GET("/health", health)
	r.GET("/info", info)
	r.POST("/rpc/stash", stash)
	r.Run(":" + strconv.Itoa(port)) // listen and serve on 0.0.0.0:9999
}
