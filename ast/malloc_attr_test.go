package ast

import (
	"testing"
)

func TestMallocAttr(t *testing.T) {
	nodes := map[string]Node{
		`0x7fc0a69091d1 <line:11:7, line:18:7>`: &MallocAttr{
			Address:  "0x7fc0a69091d1",
			Position: "line:11:7, line:18:7",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
