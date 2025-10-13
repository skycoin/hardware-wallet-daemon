// +build testrunmain

package main

import (
	"os"
	"testing"
)

func TestRunMain(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd"}
	main()
}
