package main

import (
	"io"
	"log"
	"net"
	"time"
)

func Serve(l net.Listener, handle func(net.Conn) error) error {
	// your code goes here
	return nil
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
	go func() {
		time.Sleep(time.Second * 5)
		l.Close()
	}()
	log.Fatal(Serve(l, HandleConnection))
}
