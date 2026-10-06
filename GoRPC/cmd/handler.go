package main

import (
	"context"
	"fmt"
	"os"

	"github.com/IhitplayzYT/GoRPC/gen"
	"github.com/IhitplayzYT/GoRPC/internal/client"
)

func (c *CLI) handlePodCommand(ctx context.Context, client *client.ClusterClient) {
	switch c.action {
	case "create":
		if c.pod_name == "" || c.img == "" {
			fmt.Println("Pod name and image are required for create command")
			usagePod()
			os.Exit(1)
		}
		container := &gen.ContainerSpec{Name:c.pod_name + "-container",Image:c.img,Command:c.command,Args:c.args,CpuT:c.cpu_t,Mem:c.mem}

		pod_id, err := client.CreatePod(ctx, c.pod_name, c.ns, []*gen.ContainerSpec{container})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create pod: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Pod created with id: %s\n", pod_id.PodId)

	case "get":
		if c.pod_id == "" {
			fmt.Println("Pod id is required for get command")
			usagePod()
			os.Exit(1)
		}
		pod, err := client.GetPod(ctx, c.pod_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get pod: %v\n", err)
			os.Exit(1)
		}
		printPod(pod)

	case "delete":
		if c.pod_id == "" {
			fmt.Println("Pod id is required for delete command")
			usagePod()
			os.Exit(1)
		}
		resp, err := client.DeletePod(ctx, c.pod_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete pod: %v\n", err)
			os.Exit(1)
		}
		if resp.Success {
 			fmt.Printf("Pod:%v deleted successfully\n",c.pod_id)
		} else {
			fmt.Printf("Failed to delete pod:%v\n",c.pod_id)
		}

	case "list":
		resp, err := client.ListPods(ctx, c.ns)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list pods: %v\n", err)
			os.Exit(1)
		}
		if len(resp.Pods) == 0 {
			fmt.Println("No pods found")
			return
		}
		fmt.Println("Pods: ",len(resp.Pods))
		for _, pod := range resp.Pods {
			printPod(pod)
		}
		fmt.Println()

	default:
		usagePod()
		os.Exit(1)
	}
}

func (c *CLI) handleNodeCommand(ctx context.Context, client *client.ClusterClient) {
	switch c.action {
	case "register":
		if c.node_addr == "" {
			fmt.Println("Node addr is required for register command")
			usageNode()
			os.Exit(1)
		}
		node_id, err := client.RegisterNode(ctx, c.node_addr, c.cpu_t, c.mem, c.labels)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register node: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Node registered with id: %s\n", node_id.NodeId)

	case "get":
		if c.node_id == "" {
			fmt.Println("Node id is required for get command")
			usageNode()
			os.Exit(1)
		}
		node, err := client.GetNode(ctx, c.node_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get node: %v\n", err)
			os.Exit(1)
		}
		printNode(node)

	case "list":
		resp, err := client.ListNodes(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list nodes: %v\n", err)
			os.Exit(1)
		}
		if len(resp.Nodes) == 0 {
			fmt.Println("No nodes found")
			return
		}
		fmt.Println("Nodes: ",len(resp.Nodes))
		for _, node := range resp.Nodes {
			printNode(node)
		}
		fmt.Println()

	default:
		usageNode()
		os.Exit(1)
	}
}
