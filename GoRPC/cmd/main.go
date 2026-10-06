package main

import (
	"fmt"
	"os"
	"time"

	"github.com/IhitplayzYT/GoRPC/internal/client"
)


func main() {
	cli := &CLI{}
	cli.ParseArgs()

	conn, err := client.New(cli.srvr_addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to server@%v: %v\n",cli.srvr_addr, err)
		os.Exit(1)
	}

	defer conn.Close()
	ctx, ctx_cancel := conn.WithTimeout(10 * time.Second) // Add a Timeout of 10 seconds on each call
	defer ctx_cancel()

	switch cli.cmd {
	case Pod:
		cli.handlePodCommand(ctx, conn)
	case Node:
		cli.handleNodeCommand(ctx, conn)
	default:
		usage()
		os.Exit(1)
	}
}
