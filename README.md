# leveldb-grpc

LevelDB gRPC Server

## Benchmark

Platform:
- CPU: 8vCPU, Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
- Mem: 8G RAM
- Disk: HDD, ST1000LM035-1RK172


Versions:
- golang: go1.14.3 linux/amd64
- protoc-gen-go: v1.22.0-devel
- protoc: v3.7.1
- google.golang.org/grpc: 1.30.0-dev


```bash
go test -bench=. -run=none client_test.go
```

```
goos: linux
goarch: amd64
BenchmarkServerSet-8               13704            106428 ns/op
BenchmarkServerGet-8               15654             84919 ns/op
BenchmarkServerSetParallel-8       32892             67833 ns/op
BenchmarkServerGetParallel-8       51129             29454 ns/op
```

- Set QPS: about 12000+
- Get QPS: about 30000+

## Installation

```
go get -u -v github.com/monitor1379/leveldb-grpc
```

## Examples

Server:
```go
package main

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

```


Client:
```go
package main

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

```

