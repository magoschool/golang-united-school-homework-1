package solution

import (
	"testing"

	"github.com/kyokomi/emoji"
)

func TestPrintMain(t *testing.T) {
	testValue := GetMessage()
	expectedValue := emoji.Sprint("Hello :world_map:!")

	if testValue != expectedValue {
		t.Errorf("PrintMain failed: '%v' expected but '%v' found", expectedValue, testValue)
	}
}
