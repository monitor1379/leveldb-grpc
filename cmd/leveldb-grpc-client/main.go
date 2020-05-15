package main

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-15 10:41:51
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-15 11:14:19
 */

import (
	"fmt"
	"log"

	"github.com/monitor1379/leveldb-grpc"
)

var (
	host = "localhost"
	port = 1379
)

func main() {
	address := fmt.Sprintf("%s:%d", host, port)

	// dial server
	client, err := leveldbgrpc.Dial(address)
	if err != nil {
		log.Fatal(err)
		return
	}

	// set k1 v1
	err = client.Set([]byte("k1"), []byte("v1"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// get k1
	value, err := client.Get([]byte("k1"))
	if err != nil {
		log.Fatal(err)
		return
	}
	// print: "v1"
	fmt.Println(string(value))

	// get k2
	value, err = client.Get([]byte("k2"))
	// print: true, true
	fmt.Println(value == nil, err == leveldbgrpc.ErrRecordNotFound)
}
