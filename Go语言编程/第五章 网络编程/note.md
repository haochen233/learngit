## 5.2 HTTP 编程
### 5.2.1 HTTP客户端
&emsp;&emsp;Go内置的net/http包提供了最简洁的HTTP客户端实现，无需借助第三方的网络通信库，就可以直接使用HTTP中用的最多的GET和POST方式请求数据。
1. 基本方法  
   net/http包的Client类型提供了如下几个方法，让我们可以用最简洁的方式实现HTTP请求：  
   ```go
   func (c *Client)Get(url string)(r *Response, err error)
   func (c *Client)Post(url string, bodyType string, io.Reader)(r *Response, err error)
   func (c *Client)PostFrom(url string, data url.Values)(r *Response, err error)
   func (c *Client)Head(url string)(r *Response, err error)
   func (c *Client)Do(req *Request)(resp *Response, err error)
    ```
下面概要介绍这几个方法。
- http.Get()  
要请求一个资源，只需调用Get方法。
- http.Post()
要以POST的方式发送数据,只需调用Post方法并依次传递下面三个参数即可：
  - 请求的目标URL
  - 将要POST数据的资源类型
  - 数据的比特流
- http.PostFrom()
该方法实现了标准编码格式为application/x-www-from-urlencoded的表单提交。
- http.Head()
HTTP中的Head请求方式表明只请求目标URL的头部信息，即HTTP Header而不返回HTTP Body。Go内置的net/http包同样也提供了http.Head()函数。
- (*http.Client).Do()
在多数情况下，http.Get()和http.PostFrom()就可以满足需求。但是如果我们发起的HTTP请求需要更多的定制信息，我们希望设定一些自定义的Http Header字段，比如：
  - 设定自定义的"User-Agent"，而不是默认的 "GO http package"
  - 传递Cookie
&emsp;&emsp;此时可以使用该方法来实现。

2. 高级封装
&emsp;&emsp;除了基本的HTTP操作，Go语言标准库也暴露了比较底层的HTTP相关库，让开发者可以基于这些灵活定制HTTP服务器和使用HTTP服务。
+ 自定义http.Client
Go语言提供的HTTp Client被设计成上下两层结构，一层是上述提到的http.Client结构及其封装的基础方法，不妨将其称为“业务层”。之所以称为业务层，是因为调用放通常只需要关心请求的业务逻辑本身，而无需关心非业务相关的技术细节，这些细节包括：  
  - HTTP底层传输细节
  - HTTP代理
  - gzip压缩
  - 连接池及其管理
  - 认证（SSL或其他认证方式）
  之所以HTTTP Client可以做到这么好的封装性，是因为HTTP Client在底层抽象了http.RoundTripper接口，而http.Transport实现了该接口，从而能够处理更多的细节，我们不妨将其称为“传输层”

### 5.2.2 HTTP服务器端
1. 处理HTTP请求  
   使用net/http包提供的http.ListenAndServer()函数，可以在指定的地址进行
   监听，开启一个HTTP，服务端该方法的原型如下：  
   `func ListenAndServer(addr string, handler Handler) error`
该方法用于在指定的TCP网络地址 **addr** 进行监听，然后调用服务端处理程序来处理传入的连接请求。该函数有两个参数：第一个参数addr即监听地址；第二个参数表示服务端处理程序，通常为空，这意味着服务端调用`http.DefaultServeMux` 进行处理，而服务端编写的业务逻辑处理程序`http.Handle()`或`http.HandleFunc`默认注入`http.DefaultServeMux`中，具体代码如下：  
详见Go标准库`server.go`。