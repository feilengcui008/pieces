package fuzz

import (
	"fmt"
	"testing"
)

type InnerStruct struct {
	InnerIntValue int
	Arr           []string `fuzz:"100"`
	Float32       float32
	Float64       float64
	ccccc         int
	arrarr        []float64
}

type TestStruct struct {
	ExportedInt            int
	ExportedIntPtr         *int
	unexportedString       string
	unexportedStringPtr    *string
	unexportedFloat64Slice []float64
	SliceWithTag           []int `fuzz:"10"`
	unexportedMapPtr       *map[int8]string
	InnerStruct            InnerStruct
}

func (ts *TestStruct) Print() {
	fmt.Println("***** ExportedInt ", ts.ExportedInt)
	fmt.Println("***** ExportedIntPtr ", *ts.ExportedIntPtr)
	fmt.Println("***** unexportedString", ts.unexportedString)
	fmt.Println("***** unexportedStringPtr", *ts.unexportedStringPtr)
	fmt.Println("***** unexportedFloat64Slice", ts.unexportedFloat64Slice)
	fmt.Println("***** SliceWithTag", ts.SliceWithTag)
	fmt.Println("***** unexportedMapPtr", *ts.unexportedMapPtr)
	fmt.Printf("***** InnerStruct %#v\n", ts.InnerStruct)
}

func TestFuzz(t *testing.T) {
	// fuzz struct
	ts := &TestStruct{}
	Fuzz(ts)
	t.Logf("%#v\n", ts)
}
