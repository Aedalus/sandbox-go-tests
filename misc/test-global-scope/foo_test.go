package test_global_scope

import "testing"

var (
	counter = 0
)

func TestGlobalVars(t *testing.T) {

	t.Run("test a", func(t *testing.T) {
		assertCounter(t, 0)
		counter++
		assertCounter(t, 1)
	})

	t.Run("test b", func(t *testing.T) {
		assertCounter(t, 0)
		counter++
		assertCounter(t, 1)
	})
}

func TestGlobalVars2(t *testing.T) {
	t.Run("test a", func(t *testing.T) {
		assertCounter(t, 0)
		counter++
		assertCounter(t, 1)
	})

	t.Run("test b", func(t *testing.T) {
		assertCounter(t, 0)
		counter++
		assertCounter(t, 1)
	})
}

func assertCounter(t *testing.T, want int) {
	t.Helper()
	if counter != want {
		t.Errorf("got %d, want %d", counter, want)
	}
}
