package sshclient

import "testing"

func TestSSHConnection_RunParallel(t *testing.T) {

	var ssh SSHConnection

	ssh.IP = "192.168.64.9"
	ssh.Username = "ubuntu"
	ssh.KeyLocation = "/Users/debarshibasak/.ssh/id_rsa"

	err := ssh.RunParallel([]string{
		"curl -o test.iso http://mirror.ufs.ac.za/linuxmint/stable/14/linuxmint-14-kde-dvd-64bit.iso",
		"ls $pwd",
		"curl -o -vL https://github.com/kubernetes/kubernetes/releases/download/v1.18.0/kubernetes.tar.gz",
	})

	if err != nil {
		t.Fatal(err)
	}

}
