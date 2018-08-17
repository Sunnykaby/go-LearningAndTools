package HTTP2

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	serverPem := "./creds/server.pem"
	serverKey := "./creds/server.key"
	serverReadTimeout := 300000
	serverWriteTimeout := 300000

	// TLS证书解析验证
	if _, err := tls.LoadX509KeyPair(serverPem, serverKey); err != nil {
		fmt.Println("Wrong server credentials")
		os.Exit(1)
	}

	// HTTP/2 TLS服务
	//
	server := &http.Server{
		ReadTimeout:  time.Duration(serverReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(serverWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}

	// 监听端口
	if listener, err := net.Listen("tcp", ":"+strconv.Itoa(G_config.ServicePort)); err != nil {
		return
	}

	// 拉起服务
	go server.ServeTLS(listener, G_config.ServerPem, G_config.ServerKey)
}
