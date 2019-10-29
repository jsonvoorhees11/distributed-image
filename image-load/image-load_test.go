package main

import "testing"

func TestGetSplittedDirectory(t *testing.T) {
	imagePath := "/home/gocode/src/github.com/jsonvoorhees11/distributed-image/images/test.jpg"
	expected := "/home/gocode/src/github.com/jsonvoorhees11/distributed-image/splitted/"
	actual := getSplittedDirectory(imagePath)

	if actual != expected {
		t.Errorf("getSplittedDirectory failed \r\n expected: %v \r\n got: %v", expected, actual)
	} else {
		t.Logf("success")
	}

}
