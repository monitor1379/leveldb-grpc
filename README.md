# leveldb-grpc

LevelDB gRPC Server


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


## Benchmark


```bash
go test -bench=. -run=none client_test.go
```

```
goos: linux
goarch: amd64
BenchmarkServerSet-8               10000            173876 ns/op
BenchmarkServerGet-8               20000             92282 ns/op
BenchmarkServerSetParallel-8       50000             78767 ns/op
BenchmarkServerGetParallel-8       50000             31610 ns/op
```

Set QPS: about 12000
Get QPS: about 31645