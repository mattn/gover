# gover

Version detection of go runtime.

## Why?

When using devel version, `go version` doesn't return formatted string. It return devel strings like below.

```
go version devel +71d13a8 Thu Mar 3 05:34:48 2016 +0000 windows/amd64
```

Some Continuous Integration Services (including travis.ci) provide `tip` to use devel version, So we can't use `runtime.Version()` to detect version of go runtime.

## Usage

```go
package main

import (
	"fmt"
	"github.com/mattn/gover"
)

func main() {
	fmt.Println(gover.Version())
}

```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
