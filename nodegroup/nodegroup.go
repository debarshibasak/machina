package nodegroup

import (
	"github.com/debarshibasak/machina"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

type NodeGroup struct {
	nodes []*machina.Node
}

func NewNodeGroup(nodes ...*machina.Node) *NodeGroup {
	return &NodeGroup{nodes: nodes}
}

func (t *NodeGroup) Run(commands ...string) error {

	g, _ := errgroup.WithContext(context.Background())

	for _, node := range t.nodes {
		g.Go(func() error {
			return node.SSHClient().Run(commands...)
		})
	}

	return g.Wait()
}

func (t *NodeGroup) RunParallel(commands []string) error {

	g, _ := errgroup.WithContext(context.Background())

	for _, node := range t.nodes {
		g.Go(func() error {
			return node.SSHClient().RunParallel(commands)
		})
	}

	return g.Wait()
}
