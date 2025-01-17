package main

import (
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerConnectionHandling(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", ":3000")
			assert.NoError(t, err)
			t.Cleanup(func() {
				conn.Close()
			})
			log.Println("connection established")
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
}

func TestServerReadHandling(t *testing.T) {
	conn, err := net.Dial("tcp", ":3000")
	assert.NotNil(t, conn)
	assert.NoError(t, err)

	t.Cleanup(func() {
		conn.Close()
	})
	log.Println("sending request to the client")
	conn.Write([]byte("GET key_1"))
	time.Sleep(5 * time.Second)
}

func TestPutHandling(t *testing.T) {
	conn, err := net.Dial("tcp", ":3000")
	assert.NotNil(t, conn)
	assert.NoError(t, err)

	t.Cleanup(func() {
		conn.Close()
	})

	log.Println("connection established")
	conn.Write([]byte("PUT key_1 val_1"))
	time.Sleep(5 * time.Second)
}
