package contention

import (
	"testing"
)

func BenchmarkMixed(b *testing.B) {
	type funcs struct {
		read  func() int64
		write func() int64
	}
	b.Run("mutex", func(b *testing.B) {
		var i Incr
		b.RunParallel(func(pb *testing.PB) {
			op := 0
			for pb.Next() {
				if op%10 == 0 {
					i.IncrMutex()
				} else {
					_ = i.ReadMutex()
				}
				op++
			}
		})
	})

	b.Run("rwmutex", func(b *testing.B) {
		var i Incr

		b.RunParallel(func(pb *testing.PB) {
			op := 0
			for pb.Next() {
				if op%10 == 0 {
					i.IncrRWMutex()
				} else {
					_ = i.ReadRWMutex()
				}
				op++
			}
		})
	})
	b.Run("atomic", func(b *testing.B) {
		var i Incr

		b.RunParallel(func(pb *testing.PB) {
			op := 0
			for pb.Next() {
				if op%10 == 0 {
					i.IncrAtomic()
				} else {
					_ = i.ReadAtomic()
				}
				op++
			}
		})
	})
}
