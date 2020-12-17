FROM golang:1.15.6 AS build

ADD speaker /speaker
RUN cd /speaker && CGO_ENABLED=0 go build -o speaker

FROM alpine:3.12
COPY --from=build /speaker/speaker /speaker

CMD ["/speaker"]
