package tree

import "fmt"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	fmt.Println(records)
	for _, record := range records {
		newNode := Node{ID: record.ID, Children: []*Node{}}
	}

	return nil, nil
}
