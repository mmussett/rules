package rete

import (
	"container/list"
	"context"
	"fmt"
	"strconv"

	"github.com/project-flogo/rules/common/model"
)

//node a building block of the rete network
type node interface {
	abstractNode
	getIdentifiers() []model.TupleType
	getID() int
	addNodeLink(nodeLink)
	assertObjects(ctx context.Context, handles []reteHandle, isRight bool)
}

type nodeImpl struct {
	identifiers []model.TupleType
	nodeLinkVar nodeLink
	id          int
	rule        model.Rule
}

//NewNode ... returns a new node
func newNode(nw Network, rule model.Rule, identifiers []model.TupleType) node {
	n := nodeImpl{}
	n.initNodeImpl(nw, rule, identifiers)
	return &n
}

func (n *nodeImpl) initNodeImpl(nw Network, rule model.Rule, identifiers []model.TupleType) {

	n.id = nw.incrementAndGetId()

	n.identifiers = identifiers
	n.rule = rule
}

func (n *nodeImpl) getIdentifiers() []model.TupleType {
	return n.identifiers
}

func (n *nodeImpl) getID() int {
	return n.id
}

func (n *nodeImpl) addNodeLink(nl nodeLink) {
	n.nodeLinkVar = nl
}

func (n *nodeImpl) String() string {
	str := "id:" + strconv.Itoa(n.id) + ", idrs:"
	for _, nodeIdentifier := range n.identifiers {
		str += string(nodeIdentifier) + ","
	}
	return str
}

//FindSimilarNodes find similar nodes
func findSimilarNodes(nodeSet *list.List) []node {
	if nodeSet.Len() < 2 {
		//TODO: Handle error
		return nil
	}
	maxCommon := 0
	similarNodes := make([]node, 2)
	for e := nodeSet.Front(); e != nil; e = e.Next() {
		node1 := e.Value.(node)
		for j := e.Next(); j != nil; j = j.Next() {
			node2 := j.Value.(node)
			common := len(IntersectionIdentifiers(node1.getIdentifiers(), node2.getIdentifiers()))
			if common > maxCommon {
				maxCommon = common
				similarNodes[0] = node1
				similarNodes[1] = node2
			}
		}
	}
	if maxCommon == 0 {
		similarNodes[0] = nodeSet.Front().Value.(node)
		similarNodes[1] = nodeSet.Front().Next().Value.(node)
	}
	return similarNodes
}

func (n *nodeImpl) assertObjects(ctx context.Context, handles []reteHandle, isRight bool) {
	fmt.Println("Abstract method here.., see filterNodeImpl and joinNodeImpl")
}
