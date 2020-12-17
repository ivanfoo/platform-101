FROM golang:1.15.6 AS build

ADD speaker-api /speaker-api
RUN cd /speaker-api && CGO_ENABLED=0 go build -o speaker-api

FROM alpine:3.12
COPY --from=build /speaker-api/speaker-api /speaker-api

EXPOSE 8080
CMD ["/speaker-api"]
