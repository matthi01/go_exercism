package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

var (
	ErrNoRootNode         = errors.New("no root node")
	ErrDuplicateRootNode  = errors.New("duplciate root")
	ErrIdOutOfBounds      = errors.New("provided id is out of bounds")
	ErrRootHasParent      = errors.New("root has parent")
	ErrDuplicateNode      = errors.New("duplicate node")
	ErrHigherIdThanParent = errors.New("higher id parent of lower id")
)

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	var idTopLimit = len(records) - 1
	var foundRoot bool
	var existingIds []int
	for _, r := range records {
		if r.ID == 0 {
			if foundRoot {
				return nil, ErrDuplicateRootNode
			}
			if r.Parent != 0 {
				return nil, ErrRootHasParent
			}
			foundRoot = true
		}
		if r.ID > idTopLimit || r.Parent > idTopLimit {
			return nil, ErrIdOutOfBounds
		}
		if r.ID <= r.Parent && r.ID != 0 {
			return nil, ErrHigherIdThanParent
		}
	}
	if !foundRoot {
		return nil, ErrNoRootNode
	}

	// sort by parent id / subsort by id
	sort.Slice(records, func(i, j int) bool {
		if records[i].Parent != records[j].Parent {
			return records[i].Parent < records[j].Parent
		}
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))
	for i, r := range records {
		if contains(existingIds, r.ID) {
			return nil, ErrDuplicateNode
		}
		nodes[i] = &Node{ID: r.ID}
		existingIds = append(existingIds, r.ID)
		if i != 0 {
			parent := nodes[r.Parent]
			parent.Children = append(parent.Children, nodes[i])
		}
	}
	return nodes[0], nil
}

func contains(list []int, item int) bool {
	for _, n := range list {
		if n == item {
			return true
		}
	}
	return false
}
