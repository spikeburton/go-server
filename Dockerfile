ARG BASE=1.17.5

FROM golang:${BASE} as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -ldflags="-w -s"

FROM scratch

COPY --from=builder /build/go-server /

EXPOSE 8080
ENTRYPOINT [ "/go-server" ]
