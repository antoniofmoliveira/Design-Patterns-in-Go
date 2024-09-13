package main

import "fmt"

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

func main() {
	root := Tree{
		LeafValue: 0,
		Right: &Tree{
			LeafValue: 5,
			Right:     &Tree{6, nil, nil},
			Left:      nil,
		},
		Left: &Tree{4, nil, nil},
	}

	fmt.Println(root.LeafValue)
	fmt.Println(root.Left.LeafValue)
	fmt.Println(root.Right.Right.LeafValue)
}
