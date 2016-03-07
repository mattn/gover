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
			recover()
			paniced = true
		}()
		version = Version()
	}()
	if !regexp.MustCompile(`^go\d+.\d+(.\d+)?$`).MatchString(version) {
		t.Fatal("invalid version:", version)
	}
}
