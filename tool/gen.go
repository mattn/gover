// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	p, err := exec.LookPath("go")
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(filepath.Dir(p))
	b, err := exec.Command("git", "tag").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("table.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, strings.Join([]string{
		`//go:generate go run tool/gen.go`,
		``,
		`package gover`,
		``,
		`var vers = [][2]string{`,
		``}, "\n"))
	for _, s := range strings.Split(string(b), "\n") {
		if !strings.HasPrefix(s, "go") {
			continue
		}
		if strings.Contains(s, "rc") || strings.Contains(s, "beta") {
			continue
		}
		b, err = exec.Command("git", "log", "--format=%ci", fmt.Sprintf("%s^..%s", s, s)).CombinedOutput()
		if err != nil {
			continue
		}
		d := strings.TrimSpace(string(b))
		fmt.Fprintf(f, "	{`%s`, `%s`},\n", d, s)
	}
	fmt.Fprintln(f, "}")
}
