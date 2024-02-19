package main

import (
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	mapv1 "github.com/sunsunskibiz/protobuf/gen/map/v1"
	simplev1 "github.com/sunsunskibiz/protobuf/gen/simple/v1"
	validatesimplev1 "github.com/sunsunskibiz/protobuf/gen/validatesimple/v1"
)

func doSimple() *simplev1.Simple {
	return &simplev1.Simple{
		Id:         1,
		IsSimple:   true,
		Name:       "Sun",
		SimpleList: []int32{1, 2, 3, 4},
	}
}

func doMap() *mapv1.MapExample {
	return &mapv1.MapExample{
		Ids: map[string]*mapv1.IdWrapper{
			"1": {Id: 1},
			"2": {Id: 2},
			"3": {Id: 3},
		},
	}
}

func doValidate() {
	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	msg := validatesimplev1.Simple{
		Id: 123, // Modify here to get validate failed
	}

	if err = v.Validate(&msg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doMap())
	
	doValidate()
}
