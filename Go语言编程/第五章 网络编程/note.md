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
  要以POST的方式发送数据
   

   