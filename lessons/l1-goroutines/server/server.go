package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func Serve(l net.Listener, handle func(net.Conn) error) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		wg.Go(func() {
			if err := handle(conn); err != nil {
				log.Println(err)
			}
		})
	}
}

func HandleConnection(conn net.Conn) error {
	conn.SetDeadline(time.Now().Add(time.Second))
	b, err := io.ReadAll(conn)
	if err != nil {
		return err
	}
	log.Println(b)
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(Serve(l, HandleConnection))
}
