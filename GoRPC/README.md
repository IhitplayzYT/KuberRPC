# GoRPC - KubeR CLI Tool

GoRPC is a command-line interface for interacting with KubeR, a tiny Kubernetes implementation. It uses gRPC to communicate with the KubeR control plane.

## Building

```bash
go build -o kuberctl cmd/*
```

## Usage

```bash
./kuberctl [global flags] <resource> <action> [flags]
```

### Global Flags

- `-server string`: gRPC server address (default "[::1]:50051")

### Resources

#### Pod Commands

**Create a pod:**
```bash
./kuberctl pod create <name> --image <image> [flags]
```

Flags:
- `-namespace string`: Namespace for the pod (default "default")
- `-image string`: Container image (required)
- `-cpu uint`: CPU units in millicores (default 1000)
- `-memory uint`: Memory in MB (default 512)

Example:
```bash
./kuberctl pod create my-pod --image nginx --cpu 2000 --memory 1024
```

**List pods:**
```bash
./kuberctl pod list [--namespace <namespace>]
```

**Get pod details:**
```bash
./kuberctl pod get <pod-id>
```

**Delete a pod:**
```bash
./kuberctl pod delete <pod-id>
```

#### Node Commands

**Register a node:**
```bash
./kuberctl node register <addr> [flags]
```

Flags:
- `-cpu uint`: CPU units in millicores (default 1000)
- `-memory uint`: Memory in MB (default 512)

Example:
```bash
./kuberctl node register 192.168.1.100:50052 --cpu 4000 --memory 8192
```

**List nodes:**
```bash
./kuberctl node list
```

**Get node details:**
```bash
./kuberctl node get <node-id>
```

## Protocol Buffer Compilation

To regenerate the protobuf code:

```bash
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/cluster.proto
```

Or use the provided compile command:
```bash
./compile_command
```

## Requirements
- Go 1.26.5 or later
- gRPC server (KubeR) running on the specified address

## Connection

The CLI connects to the KubeR gRPC server using insecure credentials by default. Ensure the KubeR server is running before using the CLI.

Default server address: `[::1]:50051`

You can specify a different server using the `-server` flag:

```bash
./kuberctl -server 192.168.1.100:50051 pod list
```
