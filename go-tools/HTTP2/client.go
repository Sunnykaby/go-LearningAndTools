package HTTP2

import (
	"crypto/tls"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	//创建Transport，负责底层的连接与协议管理
	transport := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // 不校验服务端证书
		MaxIdleConns:        G_config.GatewayMaxConnection,
		MaxIdleConnsPerHost: G_config.GatewayMaxConnection,
		IdleConnTimeout:     time.Duration(G_config.GatewayIdleTimeout) * time.Second, // 连接空闲超时
	}

	// 启动HTTP/2协议
	// 这里说明一下：
	// 类似于服务端，因为我们没有使用内置的Client Transport，所以我们需要使用http2.ConfigureTransport来启动HTTPS/2特性
	http2.ConfigureTransport(transport)
	// HTTP/2 客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(G_config.GatewayTimeout) * time.Millisecond, // 请求超时
	}
}
