# lab-01: building images

## Context

The engineering team is developing a service that makes food suggestions based on a complex algorithm.
We want to use containers for both local development and production deployment.

### Build the production image using existing Dockerfile

First, inspect the Dockerfile

```
FROM golang:1.15.6 AS build

# Build the binary
ADD chef /chef
RUN cd /chef && CGO_ENABLED=0 go build -o chef

# Copy the compiled binary that was built in the previous stage
FROM alpine:3.12
COPY --from=build /chef/chef /chef

ENTRYPOINT ["/chef"]
```

After that, just build the image

```
docker build -t platform101/chef
```

### Customize the Dockerfile for local development usage

This should be the very first step, but let's keep it easier for learning purposes.
