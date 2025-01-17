package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	cache "github.com/vaidik-bajpai/d-cache/cache/lru_cache"
	pb "github.com/vaidik-bajpai/d-cache/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type Server struct {
	pb.UnimplementedCacheServer
	config     config
	lis        net.Listener
	grpcServer *grpc.Server
	cache      *cache.LRUCache
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewServer(cfg config, cache *cache.LRUCache) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		config: cfg,
		cache:  cache,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (srv *Server) Start() error {
	l, err := net.Listen("tcp", srv.config.serverAddr)
	if err != nil {
		log.Fatalf("could not setup the listener: %v", err)
	}
	srv.lis = l

	srv.grpcServer = grpc.NewServer()
	pb.RegisterCacheServer(srv.grpcServer, srv)

	log.Printf("Starting server at address %s\n", srv.config.serverAddr)
	go func() {
		if err := srv.grpcServer.Serve(srv.lis); err != nil {
			log.Fatalf("could not start the grpc server: %v", err)
		}
	}()

	return srv.awaitShutdown()
}

func (srv *Server) TestStart() error {
	srv.lis = bufconn.Listen(1024 * 1024)

	srv.grpcServer = grpc.NewServer()

	pb.RegisterCacheServer(srv.grpcServer, srv)

	go func() {
		log.Println("Starting server")
		if err := srv.grpcServer.Serve(srv.lis); err != nil {
			log.Fatal("could not start the server")
		}
	}()

	return srv.awaitShutdown()
}

func (srv *Server) awaitShutdown() error {
	quitch := make(chan os.Signal, 1)
	signal.Notify(quitch, syscall.SIGINT, syscall.SIGTERM)
	<-quitch
	return srv.Stop()
}

func (srv *Server) Stop() error {
	srv.cancel()
	log.Printf("the cache server has stopped")
	return nil
}

func (srv *Server) Put(ctx context.Context, req *pb.PutReq) (*pb.PutRes, error) {
	log.Println("put called")
	err := srv.cache.Set(req.GetKey(), req.GetValue())
	if err != nil {
		return nil, fmt.Errorf("could not perform put operation on the store: %s", err.Error())
	}
	return &pb.PutRes{}, nil
}

func (srv *Server) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	log.Printf("key %s inside server get\n", string(req.Key))
	val, err := srv.cache.Get(req.Key)
	if err != nil {
		return nil, fmt.Errorf("error while performing get on the store: %s", err)
	}
	return &pb.GetRes{Value: val}, nil
}

func (srv *Server) Join(ctx context.Context, req *pb.JoinReq) (*pb.JoinRes, error) {
	log.Println("join called")
	err := srv.cache.Join(req.GetNodeID(), req.GetAddr())
	if err != nil {
		return &pb.JoinRes{Success: false}, err
	}

	return &pb.JoinRes{Success: true}, nil
}
