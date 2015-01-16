package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Prefix struct {
	Network string
	Mask    int
}

func main() {

	valueStruct()
	time.Sleep(time.Second * 1)
	literalStruct()
	time.Sleep(time.Second * 1)
	pointerStruct()
}

func valueStruct() {
	// struct as a value
	var nw Prefix
	nw.Network = "10.1.1.0"
	nw.Mask = 24
	fmt.Println("### struct as a pointer ###")
	PrettyPrint(&nw)
}

func literalStruct() {
	// literal structs are the shortest LOC
	nw2 := &Prefix{"10.1.2.0", 30}
	fmt.Println("### struct as a literal ###")
	PrettyPrint(nw2)
}

func pointerStruct() {
	// struct as a pointer
	nw3 := new(Prefix)
	// very similar to setters/getters in OOP
	nw3.Network = "10.1.1.0"
	// or even like so
	(*nw3).Mask = 28
	fmt.Println("### struct as a pointer ###")
	PrettyPrint(nw3)
}

// print the contents of the network obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
