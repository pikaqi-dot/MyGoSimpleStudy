package utils

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Ok      bool   `json:"ok"`
	Id      int    `json:"id"`
	Message string `json:"msg"`
}

//这里我们对参数字段进行额外的描述，这样，jsonrpc 包会在序列化 JSON 时，将该聚合字段命名为指定的字符串。
