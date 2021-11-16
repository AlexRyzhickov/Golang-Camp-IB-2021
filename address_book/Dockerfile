FROM golang

WORKDIR /src
ADD . .
RUN CGO_ENABLED=0 go build -o /bin/app cmd/server/main.go

FROM alpine

WORKDIR /bin
COPY --from=0 /bin/app .

ENTRYPOINT ["app"]