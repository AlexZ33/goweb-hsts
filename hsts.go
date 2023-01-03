//  Package goweb_hsts provides HTTP Strict Transport Security.
// HTTP Strict Transport Security 通知浏览器一个网站
// 只能使用 HTTPS 而不是 HTTP 访问。该插件通过以下方式强制执行 HSTS
// 将所有 HTTP 流量重定向到 HTTPS 并通过设置
// 所有 HTTPS 响应上的 Strict-Transport-Security header。请注意，这
// 仅适用于框架未在开发模式下运行的情况。
//
// More info:
//   - MDN: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
//   - Wikipedia: https://en.wikipedia.org/wiki/HTTP_Strict_Transport_Security
//   - RFC 6797: https://tools.ietf.org/html/rfc6797
//
// # Usage
// 要使用安全的默认设置构建插件，请使用 Default。否则，自己创建拦截器

package goweb_hsts

import "time"

// 拦截器实现自动 HSTS 功能
// 有关详细信息，请参阅 https://tools.ietf.org/html/rfc6797

type Interceptor struct {
	// MaxAge 是浏览器应该记住的持续时间
	//站点只能使用 HTTPS 访问。 最大age必须是正数。使用前将四舍五入到秒。
	MaxAge time.Duration

	// DisableIncludeSubDomains 禁用 includeSubDomains 指令。
	// 当 DisableIncludeSubDomains 为 false 时，所有子域也将添加托管此服务的域到浏览器的 HSTS 列表。
	DisableIncludeSubDomains bool

	// Preload 启用 preload 指令。
	// 仅当此站点应启用时才应启用添加到浏览器 HSTS 预加载列表，支持 // 所有主流浏览器。请参阅 https://hstspreload.org/ 了解
	Preload bool

	// 控制插件在 HTTPS 方面的行为方式。 如果此服务器位于终止 HTTPS 流量的代理之后，则应启用此功能。
	//如果启用此功能，则插件将始终发送 Strict-Transport-Security 标头，并且不会将 HTTP 流量重定向到 HTTPS 流量。
	BehindProxy bool
}

// 默认创建一个具有安全默认值的新 HSTS 拦截器。
//  安全默认值是
//- max-age 设置为 2 年,
//- includeSubDomains  已启用,
//- preload 被禁用.
func Default() Interceptor {
	return Interceptor{
		MaxAge: 63072000 * time.Second,
	}
}

//Before 应该在请求被发送到处理程序之前执行。
func (it Interceptor) Before(rw ResponseWriter , ir * )
