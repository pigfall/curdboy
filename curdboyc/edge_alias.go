package curdboyc

import (
	ent "github.com/pigfall/ent_utils"
)

type EdgesOfNode struct {
	From    *ent.Type
	Aliases []*EdgeAlias
}

func NewEdgesOfNode(from *ent.Type) *EdgesOfNode {
	return &EdgesOfNode{
		From:    from,
		Aliases: []*EdgeAlias{},
	}
}

func (this *EdgesOfNode) AddEdgeAlias(alias string, to *ent.Edge) {
	// check repeat
	for _, e := range this.Aliases {
		if e.Alias == alias {
			panic("TODO")
		}
	}

	this.Aliases = append(this.Aliases, &EdgeAlias{Alias: alias, To: to})
}

type EdgeAlias struct {
	Alias string
	To    *ent.Edge
}
