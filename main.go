package main

import (
	"fmt"

	simplePB "github.com/sunsunskibiz/protobuf/gen/simple/v1"
	mapPB "github.com/sunsunskibiz/protobuf/gen/map/v1"
)

func doSimple() *simplePB.Simple {
	return &simplePB.Simple{
		Id:         1,
		IsSimple:   true,
		Name:       "Sun",
		SimpleList: []int32{1, 2, 3, 4},
	}
}

func doMap() *mapPB.MapExample {
	return &mapPB.MapExample{
		Ids: map[string]*mapPB.IdWrapper{
			"1": {Id: 1},
			"2": {Id: 2},
			"3": {Id: 3},
		},
	}
}

func main() {
	fmt.Println(doSimple())
	fmt.Println(doMap())
}
