package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPromises(t *testing.T) {

	t.Log("initially empty")
	{
		a := NewPromises()
		assert.Equal(t, a.Size(), 0)
	}

	t.Log("does not panic on nil")
	{
		var s *Promises
		s.Add("X")
		s.Remove("X")
	}
}

func TestPromises_Add(t *testing.T) {
	s := NewPromises()
	s.Add("A", "B", "C", "D", "X", "Y", "E", "F")
	s.Add("G")
	actualOutput := s.String()
	expectedOutput := "[A,B,C,D,X,Y,E,F,G]"
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPromises_Remove(t *testing.T) {
	s := NewPromises()
	s.Add("A", "B", "C", "D", "X", "Y", "E", "F")

	s.Remove("X", "Y")

	actualOutput := s.String()
	expectedOutput := "[A,B,C,D,E,F]"
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPromises_Contains(t *testing.T) {
	s := NewPromises()
	s.Add("A", "B", "C", "D", "X", "Y", "E", "F")
	s.Add("G")

	table := []struct {
		input          []string
		expectedOutput bool
	}{
		{[]string{"X", "Y"}, true},
		{[]string{"H"}, false},
	}

	for _, test := range table {
		actualOutput := s.Contains(test.input...)
		assert.Equal(t, test.expectedOutput, actualOutput)
	}
}

func TestPromises_Values(t *testing.T) {
	s := NewPromises()
	s.Add("A", "B", "C", "D", "X", "Y", "E", "F")
	s.Add("G")

	actualOutput := s.String()
	expectedOutput := "[A,B,C,D,X,Y,E,F,G]"

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPromises_Size(t *testing.T) {
	s := NewPromises()
	require.Equal(t, s.Size(), 0)

	s.Add("A", "B", "C", "D", "X", "Y", "E", "F")
	assert.Equal(t, s.Size(), 8)

	s.Add("A", "B", "C", "D", "X", "Y", "E", "F", "G")
	assert.Equal(t, s.Size(), 9)

	s.Remove("A", "B", "C", "D", "X", "Y", "E", "F", "G")
	assert.Equal(t, s.Size(), 0)
}

func TestPromises_String(t *testing.T) {
	t.Log("empty")
	{
		s := NewPromises()
		require.Equal(t, s.String(), "[]")
	}

	t.Log("non empty")
	{
		s := NewPromises()
		s.Add("A", "B")
		require.Equal(t, s.String(), "[A,B]")
	}
}

func BenchmarkPromises_Add(b *testing.B) {
	s := NewPromises()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(fmt.Sprintf("%d", i))
	}
}

func BenchmarkPromises_Remove(b *testing.B) {
	s := NewPromises()
	for i := 0; i < 1000; i++ {
		s.Add(fmt.Sprintf("%d", i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(fmt.Sprintf("%d", i))
	}
}

func BenchmarkPromises_Contains(b *testing.B) {
	s := NewPromises()
	for i := 0; i < 1000; i++ {
		s.Add(fmt.Sprintf("%d", i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Contains(fmt.Sprintf("%d", i))
	}
}

func BenchmarkPromises_Size(b *testing.B) {
	s := NewPromises()
	for i := 0; i < 1000; i++ {
		s.Add(fmt.Sprintf("%d", i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Size()
	}
}

func BenchmarkPromises_String(b *testing.B) {
	s := NewPromises()
	for i := 0; i < 1000; i++ {
		s.Add(fmt.Sprintf("%d", i))
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.String()
	}
}
