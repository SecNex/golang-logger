# SecNex Golang Logger

## Installation

```bash
go get github.com/secnex/logger
```

**go.mod**

```go
require github.com/secnex/logger v0.0.1
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
