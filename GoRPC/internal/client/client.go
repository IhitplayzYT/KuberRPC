package client

import (
	"context"
	"fmt"
	"time"

	"github.com/IhitplayzYT/GoRPC/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClusterClient struct {
	conn   *grpc.ClientConn
	client gen.ClusterApiClient
}

func New(addr string) (*ClusterClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // No TLS
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to server: %w", err)
	}
	return &ClusterClient{conn:conn,client:gen.NewClusterApiClient(conn)}, nil
}

func (c *ClusterClient) Close() error {
	return c.conn.Close()
}

func (c *ClusterClient) CreatePod(ctx context.Context, name, namespace string, containers []*gen.ContainerSpec) (*gen.PodId, error) {
	return c.client.CreatePod(ctx,&gen.CreatePodRequest{Name:name,Namespace:namespace,Containers:containers})
}

func (c *ClusterClient) GetPod(ctx context.Context, podID string) (*gen.Pod, error) {
	return c.client.GetPod(ctx,&gen.PodId{PodId:podID})
}

func (c *ClusterClient) DeletePod(ctx context.Context, podID string) (*gen.DeletePodResponse, error) {
	return c.client.DeletePod(ctx,&gen.PodId{PodId:podID})
}

func (c *ClusterClient) ListPods(ctx context.Context, namespace string) (*gen.ListPodsResponse, error) {
	return c.client.ListPods(ctx,&gen.ListPodsRequest{Namespace:namespace})
}

func (c *ClusterClient) RegisterNode(ctx context.Context, addr string, cpu, mem uint64, labels map[string]string) (*gen.NodeId, error) {
	return c.client.RegisterNode(ctx,&gen.RegisterNodeRequest{Addr:addr,CpuT:cpu,Mem:mem,Labels:labels})
}

func (c *ClusterClient) GetNode(ctx context.Context, nodeID string) (*gen.Node, error) {
	return c.client.GetNode(ctx,&gen.NodeId{NodeId:nodeID})
}

func (c *ClusterClient) ListNodes(ctx context.Context) (*gen.ListNodesResponse, error) {
	return c.client.ListNodes(ctx,&gen.ListNodesRequest{})
}

func (c *ClusterClient) ReportPodStatus(ctx context.Context, status *gen.PodStatus) (*gen.Ack, error) {
	return c.client.ReportPodStatus(ctx, status)
}

func (c *ClusterClient) WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
