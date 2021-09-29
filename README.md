# Capitalize
## Get pkg
```zsh
go get github.com/lmllr/capitalize
```

## Examples
```go
package main

import (
        "fmt"

        "github.com/lmllr/capitalize"
)

func main() {
        fmt.Println(capitalize.Str("foobar"))
        fmt.Println(capitalize.Str("fooBAR"))
        fmt.Println(capitalize.Str("FOOBAR"))
}
```
