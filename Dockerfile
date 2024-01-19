# Build
FROM golang:1.21.5-alpine as builder

WORKDIR /app

COPY go.mod                                 ./go.mod
COPY go.sum                                 ./go.sum
COPY main.go                                ./main.go
COPY pkg                                    ./pkg
COPY schema.prisma                          ./schema.prisma
COPY ./client/packages/client/dist          ./dist

RUN go mod download
RUN go run github.com/steebchen/prisma-client-go generate
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/main /app/main.go

# Production image
FROM gcr.io/distroless/static
# ARG COMMIT_SHA
# LABEL GIT_SHA=$COMMIT_SHA
 
COPY --from=builder /app/build/main     /main
COPY --from=builder /app/db             /db
COPY --from=builder /app/dist           /dist

CMD ["/main"]

