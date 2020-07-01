package main

import "fmt"

type user struct {
	name map[uint32]uint32
}

type teacher struct {
	hh map[uint32]uint32
}

func main()  {
	a := &user{
		name: map[uint32]uint32{},
	}
	b:= teacher{}
	b.hh = a.name
	b.hh[1] = 1
	fmt.Println(b)
}
