# SecNex Golang Logger

Helper package for logging in Go.

## Installation

```bash
go get github.com/secnex/gologger
```

**go.mod**

```go
require github.com/secnex/gologger v0.0.4
```

## Usage

```go
package main

import (
	"github.com/secnex/logger"
)

func main() {
	logger := logger.NewLogger()
	logger.INFO("Hello, World!")
	logger.ERROR("Hello, World!")
	logger.WARN("Hello, World!")
	logger.DEBUG("Hello, World!")
}
```
