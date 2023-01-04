# goweb-hsts
 enforcing HSTS support

Go (Golang) middleware which redirects users from HTTP to HTTPS and adds the HSTS header.

what is [HSTS](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security)

> HTTP Strict-Transport-Security 响应标头（通常缩写为 HSTS）通知浏览器该站点只能使用 HTTPS 访问，并且将来使用 HTTP 访问它的任何尝试都应自动转换为 HTTPS。
> 这比在服务器上简单地配置 HTTP 到 HTTPS (301) 重定向更安全，因为在服务器上初始 HTTP 连接仍然容易受到中间人攻击。

在 Golang 项目中，HTTP Strict Transport Security 是一种在访问 Web 应用程序时将流量导向 HTTPS URL 选项的功能。

# [Using the Forwarded header](https://www.nginx.com/resources/wiki/start/topics/examples/forwarded/)
传统上，HTTP 反向代理使用非标准标头来通知上游服务器用户的 IP 地址和其他请求属性：

```yaml
X-Forwarded-For: 12.34.56.78, 23.45.67.89
X-Real-IP: 12.34.56.78
X-Forwarded-Host: example.com
X-Forwarded-Proto: https
```
NGINX 甚至提供了一个 $proxy_add_x_forwarded_for 变量来自动将 $remote_addr 附加到任何传入的 X-Forwarded-For 标头。


[RFC 7239](https://www.rfc-editor.org/rfc/rfc7239) 标准化了一个新的 Forwarded 标头，以更有条理的方式携带此信息：

```yaml
Forwarded: for=12.34.56.78;host=example.com;proto=https, for=23.45.67.89
```
Forwarded 的主要好处是可扩展性。
例如，对于 X-Forwarded-For，如果没有硬编码规则（例如“取倒数第二个 IP 地址，但前提是请求来自 10.0.0.0/8”），您不知道要信任哪个 IP 地址。

而对于 Forwarded，您信任的前端代理可以包含一个秘密令牌来标识自己：

```yaml
Forwarded: for=12.34.56.78, for=23.45.67.89;secret=egah2CGj55fSJFs, for=10.1.2.3
```

# X-Forwarded-Proto
- [X-Forwarded-Proto](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/X-Forwarded-Proto) 
  - X-Forwarded-Proto (XFP) 是一个事实上的标准首部，用来确定客户端与代理服务器或者负载均衡服务器之间的连接所采用的传输协议（HTTP 或 HTTPS）。在服务器的访问日志中记录的是负载均衡服务器与服务器之间的连接所使用的传输协议，而非客户端与负载均衡服务器之间所使用的协议。为了确定客户端与负载均衡服务器之间所使用的协议，X-Forwarded-Proto 就派上了用场。

## 问题描述
网站使用了nginx做反向代理，设置了所有http请求全部跳转为https。浏览器和nginx之间走的是https，nginx到tomcat走的是http。网站服务端使用的是springboot、springsecurity，认证授权使用的springsecurity。网站首页地址是/，登录页地址是/login。在首次打开网站地址https://aaa.com/时，框架拦截自动跳转至http://aaa.com/login（不是https//aaa.com/login），而nginx又配置了http自动跳转至https，这样导致的情况是，我本来访问的是https://aaa.com/，结果发起了3次请求：https://aaa.com/ —> http://aaa.com/login —> https://aaa.com/login。

## 解决办法
在nginx服务器中，加入
```nginx configuration
proxy_set_header	X-Forwarded-Proto $scheme;

```
在项目yml文件中，加入

```yaml
server:
  tomcat:
    remoteip:
      protocol-header: X-Forwarded-Proto

```

### 原因分析

X-Forwarded-Proto（XFP）报头是用于识别协议（HTTP 或
HTTPS），其中使用的客户端连接到代理或负载平衡器一个事实上的标准报头。要确定客户端和负载均衡器之间使用的协议，X-Forwarded-Proto可以使用请求标头


# Reference 

- https://github.com/a-h/hsts/blob/master/handler.go
