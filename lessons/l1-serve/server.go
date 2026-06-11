package l1serve

import (
	"context"
	"net"
)

func Serve(ctx context.Context, l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		_ = conn
	}
}
