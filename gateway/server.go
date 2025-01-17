package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	pb "github.com/vaidik-bajpai/d-cache/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	cfg      config
	listener net.Listener
}

func NewServer(cfg config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.cfg.serverPort)
	if err != nil {
		log.Fatalf("error creating a listener: %v", err)

	}
	defer listener.Close()

	s.listener = listener

	go s.handleConnections()

	return s.awaitShutdown()
}

func (s *Server) awaitShutdown() error {
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh
	return s.stop()
}

func (s *Server) stop() error {
	time.Sleep(2 * time.Second)
	log.Println("shutting down the server")
	s.listener.Close()
	return nil
}

func (s *Server) handleConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("the connection could not be formed: %v", err)
			continue
		}
		log.Println("conn was made successfully")

		go s.handleReads(conn)
	}
}

func (s *Server) handleReads(conn net.Conn) {
	buf := make([]byte, 1024*1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			switch err {
			case io.EOF:
				log.Println("client disconnected, closing connection.")
				return
			default:
				log.Printf("Error reading from connection: %v", err)
			}
			continue
		}

		log.Printf("Reading from client %s", string(buf[:n]))

		err = s.processRequest(conn, buf[:n])
		if err != nil {
			log.Printf("error occured while processing the request: %v\n", err)
			continue
		}
	}
}

func (s *Server) processRequest(conn net.Conn, req []byte) error {
	reqStr := string(req)

	reqParts := strings.Split(reqStr, " ")

	switch reqParts[0] {
	case "GET":
		log.Printf("%q %q\n", reqParts[0], reqParts[1])
		if len(reqParts) != 2 {
			return errors.New("invalid no of args")
		}

		key := strings.TrimSpace(reqParts[1])
		val, err := s.processGetRequest(key)
		if err != nil {
			return err
		}
		conn.Write([]byte(val + "\n"))

	case "PUT":
		if len(reqParts) != 3 {
			log.Println(200)
			return errors.New("invalid no of args")
		}
		key := strings.TrimSpace(reqParts[1])
		val := strings.TrimSpace(reqParts[2])
		err := s.processPutRequest(key, val)
		if err != nil {
			return err
		}

		conn.Write([]byte(reqParts[1] + "--->" + reqParts[2] + "\n"))
	default:
		return errors.New("invalid command")
	}

	return nil
}

func (s *Server) processGetRequest(key string) (string, error) {
	conn, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewCacheClient(conn)

	log.Printf("key %s inside process get\n", key)
	res, err := client.Get(context.Background(), &pb.GetReq{
		Key: []byte(key),
	})
	if err != nil {
		return "", err
	}

	return string(res.GetValue()), nil
}

func (s *Server) processPutRequest(key, value string) error {
	log.Println("1")
	conn, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	log.Println("2")
	client := pb.NewCacheClient(conn)

	_, err = client.Put(context.Background(), &pb.PutReq{
		Key:   []byte(key),
		Value: []byte(value),
	})

	log.Println("3")

	return err
}
