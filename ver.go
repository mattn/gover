package gover

import (
	"os/exec"
	"strings"
	"time"
)

func Version() string {
	b, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		panic(err.Error())
	}
	token := strings.Split(string(b), " ")
	if strings.HasPrefix(token[2], "go") {
		return token[2]
	}
	if !strings.HasPrefix(token[2], "devel") {
		panic("Can not detect go version")
	}

	d := strings.Join(token[4:len(token)-1], " ")
	t0, err := time.Parse(`Mon Jan _2 15:04:05 2006 -0700`, d)
	if err != nil {
		panic("Can not detect go version")
	}
	for i, v := range vers {
		t1, err := time.Parse(`2006-01-2 15:04:05 -0700`, v[0])
		if err != nil {
			continue
		}
		if t1.After(t0) {
			return vers[i-1][1]
		}
	}
	return vers[len(vers)-1][1]
}
