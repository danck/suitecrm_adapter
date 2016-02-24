package connector

import (
	"testing"
)

func setEntryBeforeConnected(t *testing.T) {
	_, err := SetEntry("valid", nil)
	if err == nil {
		t.Errorf("Expected error but got %s", err.Error())
	}
}
