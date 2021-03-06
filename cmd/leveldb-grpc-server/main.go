package main

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-15 10:37:17
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-15 12:33:45
 */

import (
	"fmt"
	"log"
	"net"

	"github.com/monitor1379/leveldb-grpc"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	host = "0.0.0.0"
	port = 1379
	path = "./db"
)

func main() {
	// open local leveldb
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	// net.Listen
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		return
	}

	// listen and serve
	fmt.Printf("Listening: %s\n", address)
	server := leveldbgrpc.NewServer(db)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
