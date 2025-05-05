# go-build-with-flags

```go
go build -ldflags="-s -w -X main.version=0.1 -X main.commit=$(git rev-parse HEAD) -X main.date=$(date +%Y-%m-%dT%H:%M:%S%z)" -o hello.exe .
```