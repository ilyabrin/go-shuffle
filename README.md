# go-shuffle

Shuffle text just for fun

[![CodeQL](https://github.com/ilyabrin/go-shuffle/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/ilyabrin/go-shuffle/actions/workflows/codeql-analysis.yml)

### Install

```sh
>$ go get -u "github.com/ilyabrin/go-shuffle"
```

### Use
```go
package main

import (
    goshuffle "github.com/ilyabrin/go-shuffle"
)

func main() {
    println(goshuffle.New("Random string anyway"))
}
```
