# Machina

Machina is an SDK for running remote commands in VM. 
This could be used for orchestrating installations, remote code execution over ssh etc.
Consider it like ansible but with SDK.


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

#### Roadmap

- Creating sdk for installing popular software that involves complex orchestrations
- Possible examples for docker, mariadb installation etc.
- Parallel Command Runners