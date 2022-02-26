package store_test

import (
	"testing"

	store "go-api/store"
)

var db = store.New()

func TestGet(t *testing.T) {
	key := "1"
	val := db.Get(key)
	if val != "ozge" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", val, "ozge")
	}
}
