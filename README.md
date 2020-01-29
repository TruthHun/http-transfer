# HTTP-Transfer


`A`可以访问`B`，`B`可以访问`C`，但是`A`不能访问`C`，那么这时候，如果`A`想要访问`C`，那么只能通过`B`来进行请求。

`HTTP-Transfer`不是一个翻墙程序，只是一个简陋的`HTTP`请求转发服务。

目前只支持`Get`请求，为了[`BookStack`](https://www.bookstack.cn)的采集功能而开发

## 使用

启动服务，监听 8080 端口
```
./http-transfer 8080
```

请求百度的logo

```
http://localhost:8080/https://www.baidu.com/img/superlogo_c4d7df0a003d3db9b65e9ef0fe6da1ec.png?where=super
```