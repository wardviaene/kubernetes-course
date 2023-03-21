FROM golang as build-stage
WORKDIR /go/
COPY main.go /go/src
RUN cd /go/src && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kvstore

FROM alpine
COPY --from=build-stage /go/src/kvstore /
EXPOSE 8080
CMD ["/kvstore"]