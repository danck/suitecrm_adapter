package hawaicrm

import (
	"os/exec"
	"testing"
)

func TestRunHelp(t *testing.T) {
	cmd := exec.Command("../hawai-crm", "-help")
	err := cmd.Run()
	if err.Error() != "exit status 2" { // naughty and not portable...
		t.Errorf("Expected \"exit status 2\" but got \"%s\"", err.Error())
	}
}
