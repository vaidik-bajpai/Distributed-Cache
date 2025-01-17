package main

import (
	"flag"
	"log"
)

type config struct {
	serverPort string
	consulAddr string
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
				log.Println("read loop")
				n, err := conn.Read(buf)
				if err != nil {
					log.Println("error while reading message from the client")
				}
				fmt.Println(string(buf[:n]))
			}
		}()

		time.Sleep(2 * time.Second)
		conn.Write([]byte(fmt.Sprintf("%s key_%d val_%d", "PUT", 0, 0))) // Fix here

		time.Sleep(5 * time.Second)
		conn.Write([]byte(fmt.Sprintf("GET key_%d", 0)))

	}() */

	srv.Start()
}
