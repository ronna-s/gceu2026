package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func Serve(l net.Listener, handle func(net.Conn) error) error {
	var (
		wg   sync.WaitGroup
		err  error
		conn net.Conn
	)

	for {
		conn, err = l.Accept()
		if err != nil {
			break
		}
		acceptedConn := conn
		wg.Go(func() {
			if err := handle(acceptedConn); err != nil {
				log.Println(err)
			}
		})
	}
	wg.Wait()
	return err
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
