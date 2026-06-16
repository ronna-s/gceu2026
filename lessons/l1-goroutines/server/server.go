package main

import (
	"fmt"
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
	fmt.Println(string(b))
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		conn, err := net.Dial("tcp", l.Addr().String())
		if err != nil {
			log.Printf("unexpected error establishing a connection: %s", err.Error())
		}
		conn.Write([]byte(output))
		conn.Close()

		time.Sleep(time.Second * 5)
		l.Close()
	}()
	log.Fatal(Serve(l, HandleConnection))
}

const output = "        ,_---~~~~~----._\n  _,,_,*^____      _____``*g*\\\"*,\n / __/ /'     ^.  /      \\ ^@q   f\n[  @f | @))    |  | @))   l  0 _/\n \\`/   \\~____ / __ \\_____/    \\\n  |           _l__l_           I\n  }          [______]           I\n  ]            | | |            |\n  ]             ~ ~             |\n  |                            |\n   |                           |"
