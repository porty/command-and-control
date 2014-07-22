package multiconnector

import (
	"testing"
)

func TestGetListenString(t *testing.T) {
	if getListenString() != ":5647" {
		t.Error("Bad listen string")
	}
}
