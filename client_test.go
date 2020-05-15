/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-15 11:16:10
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-15 12:17:07
 */

package leveldbgrpc_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/monitor1379/leveldb-grpc"
	"github.com/syndtr/goleveldb/leveldb"
)

func newServerAndStart(ctx context.Context, address, dbPath string) {

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		return
	}

	server := leveldbgrpc.NewServer(db)

	go func() {
		// fmt.Printf("Listening: %s\n", address)
		err = server.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			// fmt.Println("GracefulStop")
			server.GracefulStop()
			return
		}
	}

}

func TestServerSet(t *testing.T) {
	address := "localhost:1379"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go newServerAndStart(ctx, address, "/tmp/leveldb/db1")

	// dial server
	client, err := leveldbgrpc.Dial(address)
	if err != nil {
		t.Error("dial faied:", err)
		return
	}

	groundTruthValue := time.Now().String()
	err = client.Set([]byte("k1"), []byte(groundTruthValue))
	if err != nil {
		t.Error("set failed:", err)
		return
	}

	value, err := client.Get([]byte("k1"))
	if err != nil {
		t.Error("get failed:", err)
		return
	}

	if string(value) != groundTruthValue {
		t.Errorf("expected %s but got %s", groundTruthValue, string(value))
		return
	}

}

func BenchmarkServerSet(b *testing.B) {
	time.Sleep(2 * time.Second)
	address := "localhost:1379"
	valueLen := 1024

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go newServerAndStart(ctx, address, "./db2")

	// dial server
	client, err := leveldbgrpc.Dial(address)
	_ = client
	if err != nil {
		b.Error("dial faied:", err)
		return
	}

	keys := [][]byte{}
	for i := 0; i < b.N; i++ {
		keys = append(keys, []byte(fmt.Sprintf("%d", i)))
	}

	value := []byte{}
	for i := 0; i < valueLen; i++ {
		value = append(value, byte('1'))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := client.Set(keys[i], value)
		if err != nil {
			b.Error("set failed:", err)
		}
	}

}

func BenchmarkServerGet(b *testing.B) {
	time.Sleep(2 * time.Second)
	address := "localhost:1379"
	valueLen := 1024

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go newServerAndStart(ctx, address, "./db3")

	// dial server
	client, err := leveldbgrpc.Dial(address)
	_ = client
	if err != nil {
		b.Error("dial faied:", err)
		return
	}

	keys := [][]byte{}
	for i := 0; i < b.N; i++ {
		keys = append(keys, []byte(fmt.Sprintf("%d", i)))
	}

	value := []byte{}
	for i := 0; i < valueLen; i++ {
		value = append(value, byte('1'))
	}

	for i := 0; i < b.N; i++ {
		err := client.Set(keys[i], value)
		if err != nil {
			b.Error("set failed:", err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Get(keys[i])
		if err != nil {
			b.Error("get failed:", err)
		}
	}
}

func BenchmarkServerSetParallel(b *testing.B) {
	time.Sleep(2 * time.Second)
	address := "localhost:1379"
	valueLen := 1024

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go newServerAndStart(ctx, address, "./db4")

	// dial server
	client, err := leveldbgrpc.Dial(address)
	_ = client
	if err != nil {
		b.Error("dial faied:", err)
		return
	}

	keys := [][]byte{}
	for i := 0; i < b.N; i++ {
		keys = append(keys, []byte(fmt.Sprintf("%d", i)))
	}

	value := []byte{}
	for i := 0; i < valueLen; i++ {
		value = append(value, byte('1'))
	}

	var waitGroup sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		waitGroup.Add(1)
		go func(i int) {
			err := client.Set(keys[i], value)
			if err != nil {
				b.Error("set failed:", err)
			}
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()

}

func BenchmarkServerGetParallel(b *testing.B) {
	time.Sleep(2 * time.Second)
	address := "localhost:1379"
	valueLen := 1024

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go newServerAndStart(ctx, address, "./db5")

	// dial server
	client, err := leveldbgrpc.Dial(address)
	_ = client
	if err != nil {
		b.Error("dial faied:", err)
		return
	}

	keys := [][]byte{}
	for i := 0; i < b.N; i++ {
		keys = append(keys, []byte(fmt.Sprintf("%d", i)))
	}

	value := []byte{}
	for i := 0; i < valueLen; i++ {
		value = append(value, byte('1'))
	}

	for i := 0; i < b.N; i++ {
		err := client.Set(keys[i], value)
		if err != nil {
			b.Error("set failed:", err)
		}
	}

	var waitGroup sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		waitGroup.Add(1)
		go func(i int) {
			_, err := client.Get(keys[i])
			if err != nil {
				b.Error("get failed:", err)
			}
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}
