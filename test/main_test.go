package test

import (
	"testing"
)

func TestRun(t *testing.T) {
	err := main.run()
	if err != nil {
		t.Error("Failed run(), Error Msg:", err)
	}
}