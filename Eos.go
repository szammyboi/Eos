package main

import (
	"fmt"

	"github.com/szammyboi/BitExport"
	"github.com/szammyboi/LifxBridge"
)

func main() {
	fmt.Println("Eos starting...")
	LifxBridge.Discovery()

	header := LifxBridge.Header{}
	payload := LifxBridge.LightColorState{}
	bytes := BitExport.MultipleToBytes(header, payload)
	fmt.Println(bytes)
}
