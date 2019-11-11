package counters

import (
	"testing"
)

func TestMaxTime_Get(t *testing.T) {
	mt := MaxTime(15)
	if mt.Get() != 15 {
		t.Fail()
	}
}

func TestMaxTime_Reset(t *testing.T) {
	mt := MaxTime(15)
	if mt != 15 {
		t.Fail()
	}
	mt.Reset()
	if mt != 0 {
		t.Fail()
	}
}

func TestMaxTime_Store(t *testing.T) {
	mt := MaxTime(15)
	mt.Store(25)
	if mt.Get() != 25 {
		t.Fail()
	}
	mt.Store(17)
	if mt.Get() != 25 {
		t.Fail()
	}
}

func BenchmarkMaxTime_Get(b *testing.B) {
	mt := MaxTime(15)
	for i := 0; i < b.N; i++ {
		if mt.Get() != 15 {
			b.Fail()
		}
	}
}

func BenchmarkMaxTime_Reset(b *testing.B) {
	mt := MaxTime(15)
	for i := 0; i < b.N; i++ {
		mt.Reset()
	}
}

func BenchmarkMaxTime_Store(b *testing.B) {
	mt := MaxTime(0)
	for i := 0; i < b.N; i++ {
		mt.Store(int64(i))
	}
}
