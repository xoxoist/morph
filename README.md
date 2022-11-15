# Morph

<img align="right" width="159px" src="https://raw.githubusercontent.com/coffeehaze/asset/main/morph.png">

[![codecov](https://codecov.io/gh/coffeehaze/morph/branch/master/graph/badge.svg)](https://codecov.io/gh/coffeehaze/morph)
[![Go Report Card](https://goreportcard.com/badge/github.com/coffeehaze/morph)](https://goreportcard.com/report/github.com/coffeehaze/morph)
[![GoDoc](https://pkg.go.dev/badge/github.com/coffeehaze/morph?status.svg)](https://pkg.go.dev/github.com/coffeehaze/morph?tab=doc)
[![Join the chat at https://gitter.im/coffeehaze/morph](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/coffeehaze/morph?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Sourcegraph](https://sourcegraph.com/github.com/coffeehaze/morph/-/badge.svg)](https://sourcegraph.com/github.com/coffeehaze/morph?badge)
[![Release](https://img.shields.io/github/release/coffeehaze/morph.svg?style=flat-square)](https://github.com/coffeehaze/morph/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/coffeehaze/morph)](https://www.tickgit.com/browse?repo=github.com/coffeehaze/morph)

Morph is simple tools that helps you work with protoc stub and struct, where you can convert protoc stub to struct, or otherwise,
save your time by copying all attribute data, except (Objects, Slices, Array) to target struct or protoc.

## Contents

- [Morph](#morph)
    - [Contents](#contents)
    - [Installation](#installation)
    - [Quick Start](#quick-start)
    - [API Examples](#api-examples)
        - [Conversion](#conversion)

### Installation

1. Required go installed on your machine

```sh
go version
```

2. Get morph

```sh
go get -u github.com/coffeehaze/morph
```

3. Import morph

```go
import "github.com/coffeehaze/morph"
```

### Quick Start

```go
package main

import (
	"fmt"
	"github.com/coffeehaze/morph"
	"github.com/coffeehaze/morph/example/model"
	pb "github.com/coffeehaze/morph/example/protobuf"
)

func sampleStructToProtoc() *pb.Todo {
	var todoProtocBlank = &pb.Todo{}
	todoStruct := model.Todo{
		ID:            "5b9e1416-1f06-4a61-a30a-0dcff164639b",
		Name:          "Lloyd",
		Completed:     true,
		NumberCode:    129520,
		NumberProduct: 25983578228,
		Codes:         []model.Code{{1}, {4}, {9}},
	}


	for _, c := range todoStruct.Codes {
		todoCodeProtocBlank := &pb.Code{}
		morph.Struct(&c).Protoc(todoCodeProtocBlank)
		todoProtocBlank.Codes = append(todoProtocBlank.Codes, todoCodeProtocBlank)
	}

	return todoProtocBlank
}

func sampleProtocToStruct(todoProtoc *pb.Todo) model.Todo {
	var todo model.Todo
	morph.Protoc(todoProtoc).Struct(&todo)

	for _, i := range todoProtoc.Codes {
		var code model.Code
		morph.Protoc(i).Struct(&code)
		todo.Codes = append(todo.Codes, code)
	}

	return todo
}

func main() {
	todoProtoc := sampleStructToProtoc()
	fmt.Println(fmt.Sprintf("protoc : %+v", todoProtoc))
	todoStruct := sampleProtocToStruct(todoProtoc)
	fmt.Println(fmt.Sprintf("struct : %+v", todoStruct))
}
```

### API Examples

### Conversion

- `morph.Struct(v interface{}) morph.ProtocTransformed`
    ```go
    var todoProtocBlank = &pb.Todo{}
    todoStruct := model.Todo{
        ID:            "5b9e1416-1f06-4a61-a30a-0dcff164639b",
        Name:          "Lloyd",
        Completed:     true,
        NumberCode:    129520,
        NumberProduct: 25983578228,
    }
    // binds all struct attributes to protoc attributes
    morph.Struct(&todoStruct).Protoc(todoProtocBlank)
    ```
  
- `morph.Protoc(v interface{}) morph.StructTransformed`
    ```go
    var todo model.Todo
    todoProtoc := &pb.Todo{
        Id:            "5b9e1416-1f06-4a61-a30a-0dcff164639b",
        Name:          "Lloyd",
        Completed:     true,
        NumberCode:    129520,
        NumberProduct: 25983578228,
    }
    // binds all protoc attributes to struct attributes
    morph.Protoc(todoProtoc).Struct(&todo)
    ```

- `end`