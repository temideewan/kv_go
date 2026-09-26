package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore()
	store.Set("Temi", "dayo")
	store.Set("Jonas", "samson")
	store.Set("Clarion", "dangana")

	got := store.Keys()
	want := []string{"Clarion", "Jonas", "Temi"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("the keys are different. Keys: %v, got: %v", got, want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()
	got := store.Keys()

	if len(got) != 0 {
		t.Errorf("Keys() on empty store = %v, we want an empty slice", got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore()
	store.Set("Hello", "World")
	want := "World"

	if got, ok := store.Get("Hello"); ok != nil || got != want {
		t.Errorf("Get() failed")
	}
	if got, ok := store.Get("something"); ok == nil || got != "" {
		t.Errorf("Get() failed")
	}
}

// Table driven test
func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{
		{"deletes existing key", map[string]string{"a": "1", "b": "2"}, "a", 1},
		{"no-op missing key", map[string]string{"a": "1"}, "x", 1},
		{"no-op on empty store", map[string]string{}, "x", 0},
		{"deletes the last remaining key", map[string]string{"only": "value"}, "only", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore()
			for k, v := range tc.setup {
				store.Set(k, v)
			}
			store.Delete(tc.deleteKey)

			if got := len(store.Keys()); got != tc.wantLen {
				t.Errorf("after Delete(%q): Len() = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}
			if _, ok := store.Get(tc.deleteKey); ok == nil {
				t.Errorf("after Delete(%q), key still present", tc.deleteKey)
			}
		})
	}
}
