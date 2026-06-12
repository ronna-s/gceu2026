package main

// type MockConn struct {
// 	outgoing  bytes.Buffer
// 	incoming  bytes.Buffer
// 	rdeadline time.Time
// 	wdeadline time.Time
// 	closed    bool
// }

// func NewFakeConnection() *MockConn {
// 	return &MockConn{
// 		outgoing: *bytes.NewBufferString("hello\n world"),
// 	}
// }

// // Close implements [net.Conn].
// func (conn *MockConn) Close() error {
// 	conn.closed = true
// 	return nil
// }

// // LocalAddr implements [net.Conn].
// func (conn *MockConn) LocalAddr() net.Addr {
// 	return &net.TCPAddr{
// 		IP:   net.ParseIP("127.0.0.1"),
// 		Port: 8080,
// 	}
// }

// // Read implements [net.Conn].
// func (conn *MockConn) Read(b []byte) (n int, err error) {
// 	if conn.closed {
// 		return 0, net.ErrClosed
// 	}
// 	if !conn.rdeadline.IsZero() && conn.rdeadline.Before(time.Now()) {
// 		return 0, os.ErrDeadlineExceeded
// 	}
// 	return conn.outgoing.Read(b)
// }

// // RemoteAddr implements [net.Conn].
// func (conn *MockConn) RemoteAddr() net.Addr {
// 	return &net.TCPAddr{
// 		IP:   net.ParseIP("127.0.0.1"),
// 		Port: 8081,
// 	}
// }

// // SetDeadline implements [net.Conn].
// func (conn *MockConn) SetDeadline(t time.Time) error {
// 	conn.rdeadline = t
// 	conn.wdeadline = t
// 	return nil
// }

// // SetReadDeadline implements [net.Conn].
// func (conn *MockConn) SetReadDeadline(t time.Time) error {
// 	conn.rdeadline = t
// 	return nil
// }

// // SetWriteDeadline implements [net.Conn].
// func (conn *MockConn) SetWriteDeadline(t time.Time) error {
// 	conn.wdeadline = t
// 	return nil
// }

// // Write implements [net.Conn].
// func (conn *MockConn) Write(b []byte) (n int, err error) {
// 	if conn.closed {
// 		return 0, net.ErrClosed
// 	}
// 	if !conn.wdeadline.IsZero() && conn.wdeadline.Before(time.Now()) {
// 		return 0, os.ErrDeadlineExceeded
// 	}
// 	return conn.incoming.Write(b)
// }

// var _ net.Conn = &MockConn{}

// type MockListener struct {
// 	closed bool
// }

// // Accept implements [net.Listener].
// func (m *MockListener) Accept() (net.Conn, error) {
// 	if m.closed {
// 		return nil, net.ErrClosed
// 	}
// 	return NewFakeConnection(), nil
// }

// // Addr implements [net.Listener].
// func (m *MockListener) Addr() net.Addr {
// 	return &net.TCPAddr{
// 		IP:   net.ParseIP("127.0.0.1"),
// 		Port: 8080,
// 	}
// }

// // Close implements [net.Listener].
// func (m *MockListener) Close() error {
// 	m.closed = true
// 	return nil
// }

// var _ net.Listener = &MockListener{}
