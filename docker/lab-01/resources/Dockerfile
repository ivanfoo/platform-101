FROM golang:1.15.6 AS build

ADD chef /chef
RUN cd /chef && CGO_ENABLED=0 go build -o chef

FROM alpine:3.12
COPY --from=build /chef/chef /chef

ENTRYPOINT ["/chef"]
