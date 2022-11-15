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

	morph.Struct(&todoStruct).Protoc(todoProtocBlank)

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
