package leveldbgrpc

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-05-15 10:27:39
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-05-15 11:21:11
 */

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/monitor1379/leveldb-grpc/proto"
	"github.com/syndtr/goleveldb/leveldb"
)

type DatabaseServer struct {
	proto.UnimplementedDatabaseServer
	db *leveldb.DB
}

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(db *leveldb.DB) *Server {
	grpcServer := grpc.NewServer()
	proto.RegisterDatabaseServer(grpcServer, &DatabaseServer{db: db})

	server := &Server{
		grpcServer: grpcServer,
	}

	return server
}

func (s *Server) Serve(listener net.Listener) error {
	return s.grpcServer.Serve(listener)
}

func (s *Server) GracefulStop() {
	s.grpcServer.GracefulStop()
}

func (s *DatabaseServer) Set(ctx context.Context, req *proto.OperationSetRequest) (*proto.OperationSetResponse, error) {
	resp := &proto.OperationSetResponse{Ok: true, ErrorMessage: ""}

	err := s.db.Put(req.GetKey(), req.GetValue(), nil)
	if err != nil {
		resp.Ok = false
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	return resp, nil
}

func (s *DatabaseServer) Get(ctx context.Context, req *proto.OperationGetRequest) (*proto.OperationGetResponse, error) {
	resp := &proto.OperationGetResponse{Ok: true, ErrorMessage: "", Value: nil}

	value, err := s.db.Get(req.Key, nil)
	if err != nil {
		resp.Ok = false
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	resp.Value = value
	return resp, nil
}
