# kgateway

A fork of [kgateway-dev/kgateway](https://github.com/kgateway-dev/kgateway) — a Kubernetes-native API Gateway built on Envoy Proxy.

> **Personal fork** — I'm using this to learn Kubernetes Gateway API and experiment with Envoy-based routing configs.

## Overview

kgateway is a feature-rich, Kubernetes-native API Gateway that leverages [Envoy Proxy](https://www.envoyproxy.io/) as the data plane. It provides advanced traffic management, security, and observability capabilities for microservices running on Kubernetes.

## Features

- **Kubernetes-native**: Built with Kubernetes Gateway API support
- **Envoy-powered**: Uses Envoy Proxy for high-performance traffic management
- **Advanced routing**: Header-based routing, traffic splitting, and more
- **Security**: JWT authentication, rate limiting, and WAF support
- **Observability**: Metrics, tracing, and logging integrations
- **Extensibility**: Plugin system for custom extensions

## Prerequisites

- Go 1.22+
- Kubernetes 1.28+
- `kubectl` configured to point at your cluster
- `helm` 3.x (for Helm-based installation)

## Getting Started

### Installation

```bash
# Install via Helm
helm repo add kgateway https://kgateway.dev/charts
helm install kgateway kgateway/kgateway \
  --namespace kgateway-system \
  --create-namespace
```

### Quick Start

1. **Deploy a sample application:**

```bash
kubectl apply -f examples/httpbin.yaml
```

2. **Create a Gateway:**

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: my-gateway
  namespace: default
spec:
  gatewayClassName: kgateway
  listeners:
  - name: http
    port: 80
    protocol: HTTP
```

3. **Create an HTTPRoute:**

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: my-route
  namespace: default
spec:
  parentRefs:
  - name: my-gateway
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /
    backendRefs:
    - name: my-service
      port: 8080
```

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/your-org/kgateway.git
cd kgateway

# Install dependencies
go mod download

# Build
make build

# Run tests
make test
```

### Using Dev Container

This project includes a [Dev Container](.devcontainer/devcontainer.json) configuration for a consistent development environment. Open the project in VS Code and select **Reopen in Container**.

### Running Locally

```bash
# Run the controller locally against a cluster
make run
```

## Contributing

We welcome contributions! Please see our [contribution guidelines](CONTRIBUTING.md) and open an issue or pull request.

Before submitting a PR, please ensure:
- All tests pass (`make test`)
- Code is linted (`make lint`)
- Relevant documentation is updated

## License

Apache License 2.0 — see [LICENSE](LICENSE) for details.
