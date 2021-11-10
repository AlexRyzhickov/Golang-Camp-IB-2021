FROM golang:1.17

WORKDIR /workdir
ADD . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app cmd/server/main.go

EXPOSE 8080

ENTRYPOINT ["/workdir/bin/app"]

#FROM golang
#
#WORKDIR /src
#ADD . .
#RUN CGO_ENABLED=0 go build -o /bin/app cmd/main.go
#
#FROM alpine
#
#WORKDIR /bin
#COPY --from=0 /bin/app .
#
#ENTRYPOINT ["app"]