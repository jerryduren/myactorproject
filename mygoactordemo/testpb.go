package main

import (
	"log"
	"github.com/golang/protobuf/proto"
	"github.com/jerryduren/myactorproject/myinclude"
	"fmt"
)

func main() {
	test := &myinclude.Test {
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &myinclude.Test_OptionalGroup {
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &myinclude.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	// etc.
	
	fmt.Println("Label =",test.GetLabel(),"Type =",test.GetType(),"Reps =",test.GetReps())
}
