# go-logoutput

A writer that change output file every day

## Exemple

Create a directory and the file `2006-01-02.log` into.

```go
package main

import (
	"github.com/HuguesGuilleus/go-logoutput"
	"log"
)

func main() {
	logoutput.SetLog("log")
	log.Println("Hello World")
}
```
