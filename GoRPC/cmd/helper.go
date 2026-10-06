package main

import (
	"flag"
	"os"
)

type Command int

const (
	Pod Command = iota
	Node
)

type CLI struct {
	cmd Command
	action string
	srvr_addr string
	ns string
	pod_name string
	pod_id string
	node_id string
	node_addr string
	cpu_t uint64
	mem uint64
	img string
	command []string
	args []string
	labels map[string]string
}


func (c *CLI) ParseArgs() {
	flags := flag.NewFlagSet("gorpc", flag.ExitOnError)	
	flags.StringVar(&c.srvr_addr, "server", "[::1]:50051", "gRPC server addr")
	flags.StringVar(&c.ns, "namespace", "default", "Namespace for the pod")
	flags.Uint64Var(&c.cpu_t, "cpu", 2*60*1000, "CPU time in ms")
	flags.Uint64Var(&c.mem, "memory", 512, "Memory in Mb")
	flags.StringVar(&c.img, "image", "", "Container image")
	
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	flags.Parse(os.Args[2:])

	switch cmd {
	case "pod", "pods":
		c.cmd = Pod
		if len(flags.Args()) < 1 {
			usagePod()
			os.Exit(1)
		}
		c.action = flags.Args()[0]
		
		switch c.action {
		case "create":
			if len(flags.Args()) >= 2 {
				c.action = flags.Args()[1]
			}
		case "get", "delete":
			if len(flags.Args()) >= 2 {
				c.pod_id = flags.Args()[1]
			}
		}
		
	case "node", "nodes":
		c.cmd = Node
		if len(flags.Args()) < 1 {
			usageNode()
			os.Exit(1)
		}
		c.action = flags.Args()[0]
		
		switch c.action {
		case "register":
			if len(flags.Args()) >= 2 {
				c.node_addr = flags.Args()[1]
			}
		case "get":
			if len(flags.Args()) >= 2 {
				c.node_id = flags.Args()[1]
			}
		}
	default:
		usage()
		os.Exit(1)
	}
}

func If[T any](cond bool,if_true T,is_false T) T{
	if cond{
		return if_true
	}else{
		return is_false
	}
}