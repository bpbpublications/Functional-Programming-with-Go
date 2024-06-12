package main

import (
 "fmt"
)

type ImmutableArray struct {
 data []int
}

func NewImmutableArray(data []int) *ImmutableArray {
 newArray := make([]int, len(data))
 copy(newArray, data)
 return &ImmutableArray{data: newArray}
}

func (a *ImmutableArray) Set(index int, value int) *ImmutableArray {
 newArray := make([]int, len(a.data))
 copy(newArray, a.data)
 newArray[index] = value
 return &ImmutableArray{data: newArray}
}

func main() {
 arr := NewImmutableArray([]int{1, 2, 3})
 newArr := arr.Set(1, 5)
 fmt.Println(arr.data)    // Output: [1 2 3]
 fmt.Println(newArr.data) // Output: [1 5 3]
}
