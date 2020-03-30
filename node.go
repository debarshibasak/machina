package machina

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"errors"

	osType "github.com/debarshibasak/machina/ostype"

	"github.com/debarshibasak/machina/sshclient"
)

type Node struct {
	username           string
	ipOrHost           string
	osType             string
	privateKeyLocation string
	verboseMode        bool
	clientID           string
}

func NewNode(username string, ip string, privateKeyLocation string) *Node {
	return &Node{
		username:           username,
		ipOrHost:           ip,
		privateKeyLocation: privateKeyLocation,
		clientID:           uuid.New().String(),
	}
}

func (n *Node) SetVerboseMode(mode bool) *Node {
	n.verboseMode = mode
	return n
}

func (n *Node) String() string {
	return fmt.Sprintf("ip=%v username=%v key=%v", n.ipOrHost, n.username, n.privateKeyLocation)
}

func (n *Node) determineOS() (osType.OsType, error) {

	client := n.sshClient()
	out, err := client.Collect("uname -a")
	if err != nil {
		return nil, err
	}

	if strings.Contains(out, "Ubuntu") {
		return &osType.Ubuntu{}, err
	}

	if err := client.Run([]string{"ls /etc/centos-release"}); err == nil {
		return &osType.Centos{}, err
	}

	if err := client.Run([]string{"ls /etc/redhat-release"}); err == nil {
		return &osType.Centos{}, err
	}

	return &osType.Unknown{}, errors.New("unknown os type")
}

func (n *Node) sshClient() *sshclient.SSHConnection {
	return &sshclient.SSHConnection{
		Username:    n.username,
		IP:          n.ipOrHost,
		KeyLocation: n.privateKeyLocation,
		VerboseMode: n.verboseMode,
		ClientID:    n.clientID,
	}
}

func (n *Node) sshClientWithTimeout(duration time.Duration) *sshclient.SSHConnection {
	return &sshclient.SSHConnection{
		Username:    n.username,
		IP:          n.ipOrHost,
		KeyLocation: n.privateKeyLocation,
		VerboseMode: n.verboseMode,
		Timeout:     duration,
		ClientID:    n.clientID,
	}
}
