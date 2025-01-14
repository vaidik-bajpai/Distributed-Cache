package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type config struct {
	serverPort string
	consulAddr string
}

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
	time.Sleep(10 * time.Second)
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

		go s.handleReads(conn)
	}
}

func (s *Server) handleReads(conn net.Conn) {
	buf := make([]byte, 1024*1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("error could not decode your request: %v", err)
			continue
		}

		if n == 0 {
			continue
		}

		s.processRequest(conn, buf[:n])
	}
}

func (s *Server) processRequest(conn net.Conn, req []byte) {
	reqStr := string(req)

	reqParts := strings.Split(reqStr, " ")

	switch reqParts[0] {
	case "GET":
		if len(reqParts) != 2 {
			conn.Write([]byte("invalid request"))
			return
		}
		conn.Write([]byte("a get request was made"))

	case "PUT":
		if len(reqParts) != 3 {
			conn.Write([]byte("invalid request"))
			return
		}
		conn.Write([]byte("a put req was made"))

	default:
		conn.Write([]byte("invalid command"))
	}
}

func main() {
	var cfg config

	flag.StringVar(&cfg.serverPort, "server-addr", ":3000", "tcp server address")
	flag.Parse()

	srv := NewServer(cfg)
	log.Printf("starting the tcp server on port %s", srv.cfg.serverPort)

	// test
	/* go func() {
		time.Sleep(2 * time.Second)

		log.Println("dialing to the tcp server running on port :3000")
		conn, err := net.Dial("tcp", srv.cfg.serverPort)
		if err != nil {
			log.Println("could not connect to the tcp server on port :3000")
			return
		}
		log.Println("connection established")

		go func() {
			buf := make([]byte, 1024)

			for {
				n, err := conn.Read(buf)
				if err != nil {
					log.Println("error while reading message from the client")
				}
				fmt.Println(string(buf[:n]))
			}
		}()
		time.Sleep(2 * time.Second)
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			conn.Write([]byte(fmt.Sprintf("%s key_%s val_%s", "PUT", string(i), string(i))))
		}
	}() */

	srv.Start()

}
