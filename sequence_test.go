package counters

import "testing"

func TestSequence_Set(t *testing.T) {
	s := Sequence(0)
	s.Set(5)
	if s != 5 {
		t.Fail()
	}
}

func TestSequence_Get(t *testing.T) {
	s := Sequence(0)
	s.Set(5)
	if s.Get() != 5 {
		t.Fail()
	}

}

func TestSequence_Next(t *testing.T) {
	s := Sequence(5)
	if s.Next() != 6 {
		t.Fail()
	}
	if s.Next() != 7 {
		t.Fail()
	}
	if s.Next() != 8 {
		t.Fail()
	}
	if s.Next() != 9 {
		t.Fail()
	}
	if s.Next() != 10 {
		t.Fail()
	}
	if s.Next() == 10 {
		t.Fail()
	}
}
func BenchmarkSequence_Set(b *testing.B) {
	s := Sequence(0)
	for i := 0; i < b.N; i++ {
		s.Set(42)
	}
}

func BenchmarkSequence_Get(b *testing.B) {
	s := Sequence(42)
	for i := 0; i < b.N; i++ {
		if s.Get() != 42 {
			b.Fail()
		}
	}
}

func BenchmarkSequence_Next(b *testing.B) {
	s := Sequence(0)
	for i := 0; i < b.N; i++ {
		_ = s.Next()
	}
}
