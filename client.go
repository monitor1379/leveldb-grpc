/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-15 10:41:58
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-15 11:07:10
 */

package leveldbgrpc

import (
	"context"
	"errors"

	"github.com/monitor1379/leveldb-grpc/proto"
	"google.golang.org/grpc"
)

type Client struct {
	databaseClient proto.DatabaseClient
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

func Dial(address string) (*Client, error) {
	grpcClientConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	databaseClient := proto.NewDatabaseClient(grpcClientConn)

	client := &Client{
		databaseClient: databaseClient,
	}
	return client, nil
}

func (c *Client) Set(key, value []byte) error {
	req := proto.OperationSetRequest{Key: key, Value: value}
	resp, err := c.databaseClient.Set(context.Background(), &req)
	if err != nil {
		return err
	}

	if !resp.GetOk() {
		return errors.New(resp.GetErrorMessage())
	}
	return nil
}

func (c *Client) Get(key []byte) ([]byte, error) {
	req := proto.OperationGetRequest{Key: key}
	resp, err := c.databaseClient.Get(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	if !resp.GetOk() {
		if resp.GetErrorMessage() == "leveldb: not found" {
			return nil, ErrRecordNotFound
		}
		return nil, errors.New(resp.GetErrorMessage())
	}

	value := resp.GetValue()
	return value, nil
}
