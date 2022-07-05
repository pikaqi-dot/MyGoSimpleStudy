//自定义编解码接口实现原理
//上篇教程我们介绍了 Go 语言内置的数据序列化工具 —— Gob，但是 Gob 只能在 Go 语言内部使用，不支持跨语言 RPC 调用，如果要实现这
//一功能，就需要对 RPC 接口的编解码实现进行自定义。
//
//Go 的 net/rpc 实现很灵活，它在数据传输前后实现了编码解码器的接口定义，这意味着，开发者可以自定义数据的传输方式以及 RPC 服务端和
//客户端之间的交互行为。PRC 客户端和服务端提供的编码解码器接口如下：
//type ClientCodec interface {
//    WriteRequest(*Request, interface{}) error
//    ReadResponseHeader(*Response) error
//    ReadResponseBody(interface{}) error
//    Close() error
//}
//
//type ServerCodec interface {
//    ReadRequestHeader(*Request) error
//    ReadRequestBody(interface{}) error
//    WriteResponse(*Response, interface{}) error
//    Close() error
//}

//接口 ClientCodec 定义了 RPC 客户端如何在一个 RPC 会话中发送请求和读取响应。客户端程序通过 WriteRequest() 方法将一个请求写
//入到 RPC 连接中，并通过 ReadResponseHeader() 和 ReadResponseBody() 读取服务端的响应信息。当整个过程执行完毕后，再通过
//Close() 方法来关闭该连接。
//
//接口 ServerCodec 定义了 RPC 服务端如何在一个 RPC 会话中接收请求并发送响应。服务端程序通过 ReadRequestHeader() 和
//ReadRequestBody() 方法从一个 RPC 连接中读取请求信息，然后再通过 WriteResponse() 方法向该连接中的 RPC 客户端发送响应。
//当完成该过程后，通过 Close() 方法来关闭连接。
//
//通过实现上述接口，我们可以自定义数据传输前后的编码解码方式，而不仅仅局限于 Gob。实际上，Go 标准库提供的 net/rpc/jsonrpc 包，
//就是一套实现了 rpc.ClientCodec 和 rpc.ServerCodec 接口的 JSON-RPC 模块。
//
//基于 jsonrpc 包对传输数据进行编解码
//接下来，我们就来演示如何基于内置 jsonrpc 包通过 JSON 对 RPC 传输数据进行编解码。

