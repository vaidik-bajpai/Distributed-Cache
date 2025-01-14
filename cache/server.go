package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	cache "github.com/vaidik-bajpai/d-cache/cache/lru_cache"
	pb "github.com/vaidik-bajpai/d-cache/common/api"
	"google.golang.org/grpc"
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

func (srv *Server) awaitShutdown() error {
	quitch := make(chan os.Signal, 1)
	signal.Notify(quitch, syscall.SIGINT, syscall.SIGTERM)
	<-quitch
	return srv.Stop()
}

func (srv *Server) Stop() error {
	srv.cancel()

	time.Sleep(5 * time.Second)
	log.Printf("the cache server has stopped")
	return nil
}

func (srv *Server) Put(ctx context.Context, req *pb.PutReq) (*pb.PutRes, error) {
	srv.cache.Set(req.GetKey(), req.GetValue())
	return &pb.PutRes{}, nil
}

func (srv *Server) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	val := srv.cache.Get(req.GetKey())
	return &pb.GetRes{Value: val}, nil
}
