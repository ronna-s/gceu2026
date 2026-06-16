package contention

import (
	"testing"
)

func BenchmarkMixed(b *testing.B) {
	b.Run("mutex", func(b *testing.B) {
		var i Incr

		b.RunParallel(func(pb *testing.PB) {
			op := 0
			_ = op // remove me...
			_ = i  // remove me...
		})
	})

	b.Run("rwmutex", func(b *testing.B) {
		var i Incr

		b.RunParallel(func(pb *testing.PB) {
			op := 0
			_ = op // remove me...
			_ = i  // remove me...
		})
	})
	b.Run("atomic", func(b *testing.B) {
		var i Incr

		b.RunParallel(func(pb *testing.PB) {
			op := 0
			_ = op // remove me...
			_ = i  // remove me...
		})
	})
}
