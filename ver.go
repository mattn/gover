package gover

import (
	"os/exec"
	"strings"
	"time"
)

var vers = [][2]string{
	{`2012-03-28 23:41:59 +1100`, `go1`},
	{`2012-04-26 12:50:14 -0700`, `go1.0.1`},
	{`2012-06-14 12:50:42 +1000`, `go1.0.2`},
	{`2012-09-21 17:10:44 -0500`, `go1.0.3`},
	{`2013-05-13 13:03:09 -0700`, `go1.1`},
	{`2013-06-13 12:49:43 +1000`, `go1.1.1`},
	{`2013-08-13 15:33:06 +1000`, `go1.1.2`},
	{`2013-11-29 08:32:31 +1100`, `go1.2`},
	{`2014-03-03 11:53:08 +1100`, `go1.2.1`},
	{`2014-05-05 10:43:37 -0700`, `go1.2.2`},
	{`2014-06-19 10:20:01 +1000`, `go1.3`},
	{`2014-08-13 13:18:02 +1000`, `go1.3.1`},
	{`2014-09-25 22:32:20 +0000`, `go1.3.2`},
	{`2014-10-01 11:20:15 +1000`, `go1.3.3`},
	{`2014-12-11 11:27:56 +1100`, `go1.4`},
	{`2015-01-15 21:04:23 +0000`, `go1.4.1`},
	{`2015-02-18 04:24:51 +0000`, `go1.4.2`},
	{`2015-09-23 04:20:05 +0000`, `go1.4.3`},
	{`2015-08-19 05:04:37 +0000`, `go1.5`},
	{`2015-09-09 00:52:53 +0000`, `go1.5.1`},
	{`2016-02-17 19:53:47 +0000`, `go1.6`},
}

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
