package main

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/gin-gonic/gin"
)

type TransferConfig struct {
	Port    int `json:"port"`
	Timeout int `json:"timeout"`
}

const (
	timeout     = 60 * time.Second
	defaultPort = 8080
)

func main() {
	args := os.Args[1:]
	port := defaultPort
	if len(args) == 1 {
		port, _ = strconv.Atoi(args[0])
		if port <= 0 {
			port = defaultPort
		}
	}
	app := gin.Default()
	app.GET("/*request", transfer)
	app.Run(":" + strconv.Itoa(port))
}

func transfer(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	requestURL := strings.TrimLeft(ctx.Param("request"), " /")
	requestURL = strings.Replace(requestURL, ":/", "://", 1)
	req := httplib.Get(requestURL).SetTimeout(timeout, timeout*120)
	if strings.HasPrefix(requestURL, "https://") {
		req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	for k, v := range ctx.Request.Header {
		req.Header(k, v[0])
	}
	if u, err := url.Parse(requestURL); err == nil {
		req.Header("host", u.Host)
	}
	resp, err := req.Bytes()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error(), "request": requestURL})
		return
	}
	ctx.Writer.Write(resp)
}
