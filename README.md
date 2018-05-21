# datastash

Go 语言编写的一个微服务，用于接收业务数据日志，便于以后数据分析

### Technology

- [glide](https://github.com/Masterminds/glide) Package Management for Golang
- [gin](https://gin-gonic.github.io/gin/) Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.
- [fargo](https://github.com/hudl/fargo) Golang client for Netflix Eureka
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) The Go driver for MongoDB

### Glide Mirrors

```yml
repos:
- original: https://golang.org/x/net/ipv4
  repo: https://github.com/golang/net.git
- original: https://golang.org/x/net/ipv6
  repo: https://github.com/golang/net.git
- original: https://golang.org/x/crypto/ed25519
  repo: https://github.com/golang/crypto.git
- original: https://golang.org/x/sys/unix
  repo: https://github.com/golang/sys.git
```