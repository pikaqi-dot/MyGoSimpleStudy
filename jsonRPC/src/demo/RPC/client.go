package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"pikaqi_dot/jsonRPC/src/demo/RPC/utils"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8080", 30*time.Second) // 30秒超时时间
	if err != nil {
		log.Fatalf("客户端发起连接失败：%v", err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var item utils.Item
	client.Call("ServiceHandler.GetName", 1, &item)
	log.Printf("ServiceHandler.GetName 返回结果：%v\n", item)

	var resp utils.Response
	item = utils.Item{2, "学院君2"}
	client.Call("ServiceHandler.SaveName", item, &resp)
	log.Printf("ServiceHandler.SaveName 返回结果：%v\n", resp)
}
