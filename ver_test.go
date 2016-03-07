package gover

import (
	"regexp"
	"testing"
)

func TestSimple(t *testing.T) {
	paniced := false
	version := ""
	func() {
		defer func() {
			if recover() != nil {
				paniced = true
			}
		}()
		version = Version()
	}()
	if paniced {
		t.Fatal("maybe bug")
	}
	if !regexp.MustCompile(`^go\d+.\d+(.\d+)?$`).MatchString(version) {
		t.Fatal("invalid version:", version)
	}
}
