package main

import (
	"fmt"

	pb "github.com/dshechn/protocol-buffers-demo/pb/demo_type"
	"github.com/golang/protobuf/proto"
)

func main() {
	aa := &pb.Person{
		ValueDouble:   0,
		ValueFloat:    0,
		ValueInt32:    0,
		ValueInt64:    0,
		ValueUint32:   0,
		ValueUint64:   0,
		ValueSint32:   0,
		ValueSint64:   0,
		ValueFixed32:  0,
		ValueFixed64:  0,
		ValueSfixed32: 0,
		ValueSfixed64: 0,
		ValueBool:     false,
		ValueString:   "abcdef",
		ValueBytes:    []byte("abcdef"),
	}
	bb, err := proto.Marshal(aa)
	fmt.Println(err)
	fmt.Println(bb)
}
