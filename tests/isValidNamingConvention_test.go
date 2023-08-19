package tests

import (
	"codeAudit/utils"
	"testing"
)

func TestWhenPassedAInvalidNamingConvention(t *testing.T) {
	mockNamingConvention := "mock_invalid_convention"
	isValidConvention := utils.IsValidNamingConvention(mockNamingConvention)
	if isValidConvention {
		t.Errorf("Expected to return false but returned true indicating %s was an unexpectedly valid convention", mockNamingConvention)
	}
}

func TestWhenPassedValidNamingConvention(t *testing.T) {
	mockNamingConvention := "camel"
	isValidConvention := utils.IsValidNamingConvention(mockNamingConvention)
	if !isValidConvention {
		t.Errorf("Expected to return true but returned false indicating %s wasn't valid as expected", mockNamingConvention)
	}
}
