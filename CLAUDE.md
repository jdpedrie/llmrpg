# LLMRPG - Build and Style Guide

## Build Commands
- `go build ./...` - Build all packages
- `go build ./cmd/main.go` - Build main binary
- `go run ./cmd/main.go` - Run the application
- `go run ./cmd/main.go serve` - Run the gRPC server

## Test Commands  
- `go test ./...` - Run all tests
- `go test ./game` - Test specific package
- `go test -v ./... -run TestFunctionName` - Run specific test with verbose output

## Code Style
- Use standard Go formatting with `gofmt`
- Follow [Go standard project layout](https://github.com/golang-standards/project-layout)
- Group imports: standard library first, then third-party, then local packages
- Use descriptive variable names in camelCase
- For DB entities, use `DBType()` method to specify the EdgeDB type name
- Error handling: return errors with context, use package `fmt.Errorf("context: %w", err)`
- Model conversions: use FromProto/ToProto methods for proto/model conversion
- Type conversions: use helper functions like setUUID, setOptionalString for geltypes
- Prefer composition over inheritance with small, focused types