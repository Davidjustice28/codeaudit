package tests

import (
	"codeAudit/utils"
	"testing"
)

func TestIsConfigOptionHappyPath(t *testing.T) {
	mockTerminalArg := "r=true"
	valueIsConfig := utils.IsConfigOption(mockTerminalArg)
	if !valueIsConfig {
		t.Fatalf("The argument used wasn't a valid config options argument. Found %s", mockTerminalArg)
	}
}

func TestIsConfigOptionFailure(t *testing.T) {
	mockTerminalArg := "-f"
	valueIsConfig := utils.IsConfigOption(mockTerminalArg)
	if valueIsConfig {
		t.Fatalf("The argument used was unexpectedly a valid config options argument. Found %s", mockTerminalArg)
	}

}
