### 疑惑
当我们希望对http Client或者Server做一些更加定制化的配置时，就会覆盖掉http库的默认行为，从而导致无法启用HTTP/2协议。


## server
#### 制作证书

首先服务器需要一套公私钥以及证书，这里自签名即可。
生成私钥
`openssl genrsa -out default.key 2048`
生成cert
`openssl req -new -x509 -key default.key -out default.pem -days 3650`

#### 启动服务
除了使用ServeTLS来启动支持HTTPS/2特性的服务端之外，还可以通过http2.ConfigureServer来为http.Server启动HTTPS/2特性并直接使用Serve来启动服务。


### 客户端
因为demo的服务端证书是自签名的，所以需要关闭服务端的证书有效性校验

客户端最重要的是配置Transport，所谓Transport就是底层的连接管理器，包括了协议的处理能力。
因为我们有很多定制化Client配置的需求，所以我们自己生成了一个Transport而不是内置的


总体需要注意一一点的是：
当我们没有使用默认的http配置时，我们需要通过http2.ConfigureXXX重新配置启用HTTP2/S特性。