package nodegroup_test

import (
	"testing"

	"github.com/debarshibasak/machina"
	"github.com/debarshibasak/machina/nodegroup"
)

func TestNodeGroup(t *testing.T) {

	t.SkipNow()

	node1 := machina.NewNode("username", "ip", "key")
	node2 := machina.NewNode("username", "ip", "key")
	node3 := machina.NewNode("username", "ip", "key")

	nodeGroupManager := nodegroup.NewNodeGroup(node1, node2, node3)

	err := nodeGroupManager.Run("sudo apt-get install postgres")
	if err != nil {
		t.Fatal(err)
	}

}
