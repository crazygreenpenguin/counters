package counters

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestNewCounter(t *testing.T) {
	cnt := NewCounter()
	if cnt.counter != 0 {
		t.Fail()
	}
}

func TestCounter_Set(t *testing.T) {
	cnt := NewCounter()

	cnt.Set(5)
	if cnt.counter != 5 {
		t.Fail()
	}
}

func BenchmarkCounter_Set(b *testing.B) {
	b.StopTimer()
	cnt := NewCounter()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cnt.Set(5)
		if cnt.counter != 5 {
			b.Fail()
		}
	}
}

func TestCounter_Inc(t *testing.T) {
	cnt := NewCounter()

	cnt.Inc()
	if cnt.counter != 1 {
		t.Fail()
	}
	cnt.Inc()
	if cnt.counter != 2 {
		t.Fail()
	}
}

func BenchmarkCounter_Inc(b *testing.B) {
	b.StopTimer()
	cnt := NewCounter()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cnt.Inc()
		if cnt.counter != uint64(i+1) {
			b.Fail()
		}
	}
}

func TestCounter_Get(t *testing.T) {
	cnt := NewCounter()
	cnt.Set(12)
	if cnt.Get() != 12 {
		t.Fail()
	}
}

func TestCounter_GetWithReset(t *testing.T) {
	cnt := NewCounter()
	cnt.Set(12)
	if cnt.GetWithReset() != 12 {
		t.Fail()
	}
	if cnt.Get() != 0 {
		t.Fail()
	}
}

func TestCounter_GetRPS(t *testing.T) {
	cnt := NewCounter()

	for i := 1; i < 10001; i = i * 10 {
		_ = cnt.GetRPS()
		for j := 0; j < i; j++ {
			cnt.Inc()
		}
		time.Sleep(time.Second)
		rps := cnt.GetRPS()
		dev := math.Abs(float64(i)-rps) * 100 / float64(i)
		fmt.Printf("%9.2f %6d %.2f%%\n", rps, i, dev)
		if dev > 1 {
			t.Error("RPS deviation ", i, " should be <1%, but it is ", dev)
		}
	}
}

func BenchmarkCounter_GetRPS(b *testing.B) {
	b.StopTimer()
	cnt := NewCounter()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if cnt.GetRPS() != 0 {
			b.Fail()
		}
	}
}

func BenchmarkCounter_Get(b *testing.B) {
	b.StopTimer()
	cnt := NewCounter()
	cnt.Set(12)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if cnt.Get() != 12 {
			b.Fail()
		}
	}
}

func BenchmarkCounter_GetWithReset(b *testing.B) {
	b.StopTimer()
	cnt := NewCounter()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = cnt.GetWithReset()
	}
}
