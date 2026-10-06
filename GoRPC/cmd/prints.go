package main

import (
	"fmt"

	"github.com/IhitplayzYT/GoRPC/gen"
)


func printPod(pod *gen.Pod) {
	fmt.Printf("ID: %s\n", pod.Id)
	fmt.Printf("Name: %s\n", pod.Name)
	fmt.Printf("Namespace: %s\n", pod.Namespace)
	fmt.Printf("Phase: %s\n", pod.Phase)
	fmt.Printf("NodeID: %s\n", pod.NodeId)
	if pod.Status != nil {
		fmt.Printf("Status Phase: %s\n", pod.Status.Phase)
		fmt.Printf("Status Message: %s\n", pod.Status.Mssg)
		ecodes,ready := make([]string,0),make([]string,0)
		for _,cs := range pod.Status.Statuses{
			ecodes = append(ecodes, If(cs.Ecode == 0,"o","x"))
			ready = append(ready, If(cs.Ready,"*","."))	
		}
		fmt.Println("Container Status for the Pod:")
		fmt.Printf("Containers ecode: %v  (o=success,x=fail)\n",ecodes)
		fmt.Printf("Containers ready: %v  (*=ready,.=working)\n",ecodes)
	}
	fmt.Println("---")
}


func printNode(node *gen.Node) {
	fmt.Printf("ID: %s\n", node.Id)
	fmt.Printf("Addr: %s\n", node.Addr)
	fmt.Printf("Status: %s\n", node.Status)
	fmt.Printf("Last Ping: %d\n", node.LastPing)
	if node.Cap != nil {
		fmt.Printf("Capacity CPU: %d ms\n", node.Cap.CpuT)
		fmt.Printf("Capacity Mem: %d MB\n", node.Cap.Mem)
	}
	if node.Availib != nil {
		fmt.Printf("Availib CPU: %d ms\n", node.Availib.CpuT)
		fmt.Printf("Availib Mem: %d MB\n", node.Availib.Mem)
	}
	fmt.Printf("Labels: %v",node.Labels)
	fmt.Println("---")
}



// The usage functions formats are built by AI

func usage() {
	fmt.Println("GoRPC - CLI for KubeR")
	fmt.Println("\nUsage:")
	fmt.Println("  gorpc [global flags] <resource> <action> [flags]")
	fmt.Println("\nResources:")
	fmt.Println("  pod, pods    - Manage pods")
	fmt.Println("  node, nodes  - Manage nodes")
	fmt.Println("\nGlobal Flags:")
	fmt.Println("  -server string")
	fmt.Println("        gRPC server address (default \"[::1]:50051\")")
	fmt.Println("\nExamples:")
	fmt.Println("  gorpc pod create my-pod --image nginx")
	fmt.Println("  gorpc pod list")
	fmt.Println("  gorpc pod get <pod-id>")
	fmt.Println("  gorpc pod delete <pod-id>")
	fmt.Println("  gorpc node list")
	fmt.Println("  gorpc node register <addr>")
}

func usagePod() {
	fmt.Println("Pod Commands:")
	fmt.Println("  create <name>    Create a new pod")
	fmt.Println("  get <id>         Get pod details")
	fmt.Println("  delete <id>      Delete a pod")
	fmt.Println("  list             List all pods")
	fmt.Println("\nPod Flags:")
	fmt.Println("  -namespace string")
	fmt.Println("        Namespace for the pod (default \"default\")")
	fmt.Println("  -image string")
	fmt.Println("        Container image (required for create)")
	fmt.Println("  -cpu uint")
	fmt.Println("        CPU units in millisecs (default 1000)")
	fmt.Println("  -memory uint")
	fmt.Println("        Memory in MB (default 512)")
}

func usageNode() {
	fmt.Println("Node Commands:")
	fmt.Println("  register <addr>  Register a new node")
	fmt.Println("  get <id>         Get node details")
	fmt.Println("  list             List all nodes")
	fmt.Println("\nNode Flags:")
	fmt.Println("  -cpu uint")
	fmt.Println("        CPU units in millisecs (default 1000)")
	fmt.Println("  -memory uint")
	fmt.Println("        Memory in MB (default 512)")
}

