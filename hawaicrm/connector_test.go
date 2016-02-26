package hawaicrm

import (
	"testing"
)

func TestSetEntryBeforeConnected(t *testing.T) {
	_, err := SetEntry("valid", nil)
	if err == nil {
		t.Errorf("Expected error but got %s", err.Error())
	}
}
