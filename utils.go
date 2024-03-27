package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func Deserialize[T any](data []byte) T {
	var result T

	decoder := gob.NewDecoder(bytes.NewBuffer(data))
	err := decoder.Decode(&result)
	if err != nil {
		log.Panic(err)
	}

	return result
}
