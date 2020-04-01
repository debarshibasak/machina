package machina

import (
	"fmt"
	"time"

	"errors"

	osType "github.com/debarshibasak/machina/ostype"

	"github.com/debarshibasak/machina/sshclient"
)

type Node struct {
	username           string
	ip                 string
	osType             string
	privateKeyLocation string
	verboseMode        bool
}

func NewNode(username string, ip string, privateKeyLocation string) *Node {
	return &Node{
		username:           username,
		ip:                 ip,
		privateKeyLocation: privateKeyLocation,
	}
}

func (n *Node) GetUsername() string {
	return n.username
}

func (n *Node) GetIP() string {
	return n.ip
}

func (n *Node) GetPrivateKey() string {
	return n.privateKeyLocation
}

func (n *Node) SetVerboseMode(mode bool) *Node {
	n.verboseMode = mode
	return n
}

func (n *Node) String() string {
	return fmt.Sprintf("ip=%v username=%v key=%v", n.ip, n.username, n.privateKeyLocation)
}

func (n *Node) DetermineOS() (osType.OsType, error) {

	client := n.SSHClient()

	if err := client.Run("ls /etc/lsb-release"); err == nil {
		return &osType.Ubuntu{}, err
	}

	if err := client.Run("ls /etc/centos-release"); err == nil {
		return &osType.Centos{}, err
	}

	if err := client.Run("ls /etc/redhat-release"); err == nil {
		return &osType.Centos{}, err
	}

	return &osType.Unknown{}, errors.New("unknown os type")
}

func (n *Node) SSHClient() *sshclient.SSHConnection {
	return &sshclient.SSHConnection{
		Username:    n.username,
		IP:          n.ip,
		KeyLocation: n.privateKeyLocation,
		VerboseMode: n.verboseMode,
	}
}

func (n *Node) SSHClientWithTimeout(duration time.Duration) *sshclient.SSHConnection {
	return &sshclient.SSHConnection{
		Username:    n.username,
		IP:          n.ip,
		KeyLocation: n.privateKeyLocation,
		VerboseMode: n.verboseMode,
		Timeout:     duration,
	}
}
