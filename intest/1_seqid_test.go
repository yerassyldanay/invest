package intest

import (
	"invest/model"
	"testing"
)

func TestSequenceId(t *testing.T) {
	err := model.PrepareSequenceId()
	if err != nil {
		t.Error("expected no err, but got ", err)
	}
}
