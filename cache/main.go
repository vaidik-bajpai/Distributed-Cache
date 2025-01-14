package main

import (
	"flag"

	cache "github.com/vaidik-bajpai/d-cache/cache/lru_cache"
)

type config struct {
	serviceName string
	serverAddr  string
	consulAddr  string
}

var (
	cacheCap = flag.Int64("cache-cap", 1000, "item no. limit for the cache")
)

func main() {
	var cfg config
	flag.StringVar(&cfg.consulAddr, "consul-addr", "localhost:8500", "consul address")
	flag.StringVar(&cfg.serviceName, "s-name", "cache", "name of the service")
	flag.StringVar(&cfg.serverAddr, "s-addr", "localhost:8080", "server address for the cache server")
	flag.Parse()

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

	srv.Start()
}
