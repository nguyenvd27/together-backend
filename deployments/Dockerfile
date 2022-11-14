FROM golang:1.17-alpine3.15 as builder

WORKDIR /together_backend

# download dependent packages
COPY go.mod go.sum /together_backend/
RUN GO111MODULE=on go mod download

COPY . /together_backend
RUN go build -o /together_backend/cmd cmd/main.go

FROM alpine:3.15.0
WORKDIR /
COPY --from=builder /together_backend/cmd/main .
CMD ["/main"]
