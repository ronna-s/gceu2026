package main

import (
	"errors"
	"net"
	"os"
	"sync/atomic"
	"testing"
	"testing/synctest"

	"github.com/ronna-s/gceu2026/mocks"
	"github.com/stretchr/testify/assert"
)

func TestServe(t *testing.T) {
	t.Run("listener returns an error", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			exampleErr := errors.New("error")
			listener := mocks.NewMockListener(t)
			listener.On("Accept").Return((net.Conn)(nil), exampleErr).Once()
			assert.ErrorIs(t, exampleErr, Serve(listener, nil))
		})
	})
	t.Run("handle succeeds", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			listener := mocks.NewMockListener(t)
			conn := mocks.NewMockConn(t)
			listener.On("Accept").Return(conn, nil).Once()
			listener.On("Accept").Return((net.Conn)(nil), os.ErrClosed).Once()
			called := false
			f := func(c net.Conn) error {
				assert.Equal(t, conn, c)
				called = true
				return nil
			}
			assert.ErrorIs(t, os.ErrClosed, Serve(listener, f))
			assert.True(t, called)
		})
	})

	t.Run("handle fails", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			listener := mocks.NewMockListener(t)
			conn := mocks.NewMockConn(t)
			listener.On("Accept").Return(conn, nil).Once()
			listener.On("Accept").Return((net.Conn)(nil), os.ErrClosed).Once()
			called := false
			assert.ErrorIs(t, os.ErrClosed, Serve(listener, func(c net.Conn) error {
				if c != conn {
					t.Error("handle expected the original connection")
				}
				called = true
				return errors.New("some error")
			}))
			synctest.Wait()
			assert.True(t, called)
			listener.AssertExpectations(t)
		})
	})

	t.Run("handle fails then succeeds", func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			listener := mocks.NewMockListener(t)
			conn1 := mocks.NewMockConn(t)
			conn2 := mocks.NewMockConn(t)
			listener.On("Accept").Return(conn1, nil).Once()
			listener.On("Accept").Return(conn2, nil).Once()
			listener.On("Accept").Return((net.Conn)(nil), os.ErrClosed).Once()
			var handleCalled int32
			assert.ErrorIs(t, os.ErrClosed, Serve(listener, func(c net.Conn) error {
				atomic.AddInt32(&handleCalled, 1)
				if c == conn1 {
					return errors.New("some error")
				}
				if c == conn2 {
					return nil
				}
				t.Error("unexpected connection received")
				return nil
			}))
			assert.Equal(t, int32(2), handleCalled, "handle was expected to be called twice")
			listener.AssertExpectations(t)
		})
	})

}
