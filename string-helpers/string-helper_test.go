package stringHelper

import (
	"testing"
)

func TestParseStringToFlagValid(t *testing.T) {
	flagStr := "-src=/home/user/images/test.jpg"
	res, _ := ParseStrToFlag(flagStr)
	actualFlag := res.Name
	actualVal := res.Value
	expectedFlag := "src"
	expectedVal := "/home/user/images/test.jpg"
	if actualFlag != expectedFlag {
		t.Errorf("ParseStringToFlag failed \r\n expected: %v \r\n got: %v", expectedFlag, actualFlag)
	} else {
		t.Logf("success")
	}

	if actualVal != expectedVal {
		t.Errorf("ParseStringToFlag failed \r\n expected: %v \r\n got: %v", expectedVal, actualVal)
	} else {
		t.Logf("success")
	}
}

func TestParseStringToFlagInvalid(t *testing.T) {
	flagStr := "-src"

	if _, err := ParseStrToFlag(flagStr); err != nil {
		t.Logf("success")
	} else {
		t.Errorf("ParseStringToFlag failed, epxected error")
	}
}

func TestParseStringToFlagInvalidInput(t *testing.T) {
	flagStr := "=ew23"
	if _, err := ParseStrToFlag(flagStr); err != nil {
		t.Logf("success")
	} else {
		t.Errorf("ParseStringToFlag failed, epxected error")
	}
}
