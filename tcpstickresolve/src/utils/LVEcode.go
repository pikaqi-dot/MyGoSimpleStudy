package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	//读取消息的长度，转化成int32类型（占4个字节）
	var length = int32(len(message))
	//开启一个缓冲区
	var pkg = new(bytes.Buffer)
	//写入消息头(将length写入pkg)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	//读取消息的长度
	//读取前4个字节的数据（int32占4个字节）
	lengthByte, _ := reader.Peek(4)
	//读取lengthByte并返回句柄
	lengthBuff := bytes.NewReader(lengthByte)
	var length int32
	//将lengthBuff中的内容写入length
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	//Buffered返回缓冲中现有的可读取的字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	//读取真正的消息数据
	pack := make([]byte, int(4+length))
	//将reader中的数据读到pack切片中
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	//返回有效数据的切片
	return string(pack[4:]), nil
}
