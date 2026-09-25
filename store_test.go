package main

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
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
