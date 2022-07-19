package unit1a

import (
	"bytes"
	"encoding/binary"
	"log"
)

func InttoHex(target int64) []byte {

	//创建buff缓冲区
	buff := new(bytes.Buffer)

	//将target以大端的方式写入 buff
	err := binary.Write(buff, binary.BigEndian, target) //这种写入流，target的位数必须确定(如果你传入int或者string，就是参数报错)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
