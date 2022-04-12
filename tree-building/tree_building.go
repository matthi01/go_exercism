package tree

import "fmt"

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	for _, record := range records {
		fmt.Println(record)
	}
	return &Node{}, nil
}
