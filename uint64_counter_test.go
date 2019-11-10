package counters

import "testing"

func TestUint64Counter_Inc(t *testing.T) {
	c := Uint64Counter(42)
	c.Inc()
	if c != 43 {
		t.Fail()
	}
}
func TestUint64Counter_Add(t *testing.T) {
	c := Uint64Counter(42)
	c.Add(13)
	if c != 55 {
		t.Fail()
	}
}
func TestUint64Counter_Get(t *testing.T) {
	c := Uint64Counter(42)
	if c.Get() != 42 {
		t.Fail()
	}
}

func TestUint64Counter_Set(t *testing.T) {
	c := Uint64Counter(11)
	c.Set(42)
	if c.Get() != 42 {
		t.Fail()
	}
}

func BenchmarkUint64Counter_Get(b *testing.B) {
	c := Uint64Counter(64)
	for i := 0; i < b.N; i++ {
		if c.Get() != 64 {
			b.Fail()
		}
	}
}

func BenchmarkUint64Counter_Set(b *testing.B) {
	c := Uint64Counter(64)
	for i := 0; i < b.N; i++ {
		c.Set(1024)
	}
}

func BenchmarkUint64Counter_Inc(b *testing.B) {
	c := Uint64Counter(64)
	for i := 0; i < b.N; i++ {
		c.Inc()
	}
}

func BenchmarkUint64Counter_Add(b *testing.B) {
	c := Uint64Counter(64)
	for i := 0; i < b.N; i++ {
		c.Add(16)
	}
}
