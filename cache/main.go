package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	cache "github.com/vaidik-bajpai/d-cache/cache/lru_cache"
	pb "github.com/vaidik-bajpai/d-cache/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DefaultRaftAddress = "localhost:11000"
)

type raftConfig struct {
	isInmem  bool
	addr     string
	nodeID   string
	joinAddr string
}

type config struct {
	serviceName string
	serverAddr  string
	consulAddr  string
	raft        raftConfig
}

func NewDefaultConfig() config {
	return config{
		serviceName: "cache",
		serverAddr:  ":8080",
		raft: raftConfig{
			isInmem:  false,
			addr:     DefaultRaftAddress,
			nodeID:   "node0",
			joinAddr: "",
		},
	}
}

var (
	cacheCap = flag.Int64("cache-cap", 1000, "item no. limit for the cache")
)

func main() {
	var cfg config
	flag.StringVar(&cfg.consulAddr, "consul-addr", "localhost:8500", "consul address")
	flag.StringVar(&cfg.serviceName, "s-name", "cache", "name of the service")
	flag.StringVar(&cfg.serverAddr, "s-addr", "localhost:8080", "server address for the cache server")
	flag.StringVar(&cfg.raft.addr, "raft-addr", DefaultRaftAddress, "raft bind address")
	flag.StringVar(&cfg.raft.nodeID, "id", "", "raft node id, if not set same as raft bind address")
	flag.StringVar(&cfg.raft.joinAddr, "j-addr", "", "set join address, if any")
	flag.BoolVar(&cfg.raft.isInmem, "inmem", false, "use in-memory storage for raft")
	flag.Parse()

	fmt.Println(cfg.raft.joinAddr)

	/* registry, err := consul.NewRegistry(cfg.consulAddr)
	if err != nil {
		log.Fatalf("could not instantiate the registry: %v", err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(cfg.serviceName)
	err = registry.Register(ctx, instanceID, cfg.serviceName, cfg.serverAddr)
	if err != nil {
		log.Fatal("could not register the service with the registry")
	}
	defer registry.Deregister(ctx, instanceID, cfg.serviceName)

	go func() {
		for {
			err := registry.HealthCheck(instanceID, cfg.serviceName)
			if err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}() */

	cache := cache.NewLRUCache(uint64(*cacheCap))
	srv := NewServer(cfg, cache)

	srv.cache.RaftBind = cfg.raft.addr

	err := srv.cache.Open(cfg.raft.joinAddr == "", cfg.raft.nodeID)
	if err != nil {
		panic(err)
	}

	if cfg.raft.joinAddr != "" {
		go func() {
			time.Sleep(10 * time.Second)
			log.Println("start join")

			err := join(cfg.raft.joinAddr, cfg.raft.addr, cfg.raft.nodeID)
			if err != nil {
				log.Printf("join err: %v", err)
			}
		}()
	}

	srv.Start()
}

func join(joinAddr, raftAddr, nodeID string) error {
	conn, err := grpc.NewClient(joinAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := pb.NewCacheClient(conn)
	_, err = client.Join(context.Background(), &pb.JoinReq{
		Addr:   raftAddr,
		NodeID: nodeID,
	})
	if err != nil {
		log.Printf("join failed: %s", err)
		return err
	}

	return nil
}
