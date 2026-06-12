package mocks

import "net"

//go:generate mockery --name=Conn --structname=MockConn --testonly --filename=conn.go --outpkg=mocks --output=.
type Conn interface {
	net.Conn
}

//go:generate mockery --name=Listener --structname=MockListener --filename=listener.go --testonly --outpkg=mocks --output=.
type Listener interface {
	net.Listener
}
