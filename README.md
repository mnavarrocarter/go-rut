Go Rut
======

A library to work with Chilean RUTs in Go

## Install

```shell
go get github.com/mnavarrocarter/go-rut
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/mnavarrocarter/go-rut"
	"log"
)

func main() {
	r, err := rut.Parse("24. 736  732 2") // Parse a rut, no matter how badly formatted
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println(r.String()) //  Prints: 24.736.732-2
}
```