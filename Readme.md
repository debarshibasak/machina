# Machina

Machina is an SDK for running remote commands in VM. 
This could be used for orchestrating installations, remote code execution over ssh etc.
Consider it like ansible but with SDK and golang.


##### SSH into a machines
```

node := machina.NewNode(username, ip, privateKeyLocation)

err := node.SSHClient().Run(instructions)
if err != nil { 
  log.Fatal(err)
}

```

#####  SSH with timeout

```

node := machina.NewNode(username, ip, privateKeyLocation)

err := node.SSHClientWithTimeout(2*time.Minutes).Run(instructions)
if err != nil { 
  log.Fatal(err)
}

```

##### Collect the output of instruction

```

node := machina.NewNode(username, ip, privateKeyLocation)

err := node.SSHClientWithTimeout(2*time.Minutes).Collect("cat /tmp/file")
if err != nil { 
  log.Fatal(err)
}

```


##### Determine OS of the machines
```

node := machina.NewNode(username, ip, privateKeyLocation)

osType, err := node.DetermineOS()
if err != nil { 
  log.Fatal(err)
}

log.Println(osType)

```

##### Running commands in parallel

```

node := machina.NewNode(username, ip, privateKeyLocation)

err := node.SSHClient().RunParallel([]string{"rm /tmp/file", "touch /tmp/example"})
if err != nil { 
  log.Fatal(err)
}

```

##### Creating node groups and running tasks in parallel across the nodes

```
	node1 := machina.NewNode("username", "ip", "key")
	node2 := machina.NewNode("username", "ip", "key")
	node3 := machina.NewNode("username", "ip", "key")

	nodeGroupManager := nodegroup.NewNodeGroup(node1, node2, node3)

	err := nodeGroupManager.Run("sudo apt-get install postgres")
	if err != nil {
		log.Fatal(err)
	}

```


#### Roadmap

- Creating orchectration for installing popular software that involves complex orchestrations
- Possible examples for docker, mariadb installation etc.