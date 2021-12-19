package main

import (
	"encoding/binary"
	"fmt"
)

const pkgHeadLen = 16

func main() {
	var sendData = "实现一个从 socket connection 中解码出 goim 协议的解码器。"
	data := encoder(sendData)
	decoder(data)
}

func decoder(data []byte) {
	if len(data) <= pkgHeadLen {
		fmt.Printf("data len < %d.", pkgHeadLen)
		return
	}

	packetLen := binary.BigEndian.Uint32(data[:4])
	fmt.Printf("packetLen:%v\n", packetLen)

	headerLen := binary.BigEndian.Uint16(data[4:6])
	fmt.Printf("headerLen:%v\n", headerLen)

	version := binary.BigEndian.Uint16(data[6:8])
	fmt.Printf("version:%v\n", version)

	operation := binary.BigEndian.Uint32(data[8:12])
	fmt.Printf("operation:%v\n", operation)

	sequence := binary.BigEndian.Uint32(data[12:16])
	fmt.Printf("sequence:%v\n", sequence)

	body := string(data[16:])
	fmt.Printf("decoder body：%v\n", body)
}

func encoder(body string) []byte {
	pkgLen := len(body) + pkgHeadLen
	ret := make([]byte, pkgLen)

	binary.BigEndian.PutUint32(ret[:4], uint32(pkgLen))
	binary.BigEndian.PutUint16(ret[4:6], uint16(pkgHeadLen))
	binary.BigEndian.PutUint16(ret[6:8], uint16(5))
	binary.BigEndian.PutUint32(ret[8:12], uint32(6))
	binary.BigEndian.PutUint32(ret[12:16], uint32(7))

	copy(ret[16:], []byte(body))
	return ret
}