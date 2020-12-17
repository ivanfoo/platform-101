# Kubernetes 101

## Labs

- [lab-01](./lab-01)
- [lab-02](./lab-02)

## Requirements

- Basic linux commands
- Docker installed
- Kind (see [instructions](#kind installation))
- Basic Docker knowledge is recommended

### Kind installation

Kind is a kubernetes implementation in docker to run a local cluster. It will be our learning play field.

Download the compiled binary:
```
curl -Lo ./kind "https://kind.sigs.k8s.io/dl/v0.9.0/kind-$(uname)-amd64"
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

Create a local Kubernetes cluster!
```
kind create cluster
```
