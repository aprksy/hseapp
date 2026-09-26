# Backend Makefile

Makefile for backend development operations.

## Prerequisites

- Go 1.21+
- Docker (for database and integration tests)
- Make 4.0+

## Common Commands

### Development

```bash
make dev          # Start backend in development mode with hot reload
make run          # Run backend once (no hot reload)
make build        # Build production binary
make clean        # Clean build artifacts
```

### Testing

```bash
make test         # Run all tests
make test-unit    # Run unit tests only
make test-integration  # Run integration tests (requires Docker)
make test-coverage     # Run tests with coverage report
make test-race    # Run tests with race detector
```

### Code Quality

```bash
make lint         # Run linters (golangci-lint)
make fmt          # Format code (goimports)
make vet          # Run go vet
make check        # Run all quality checks
```

### Database

```bash
make db-migrate   # Run database migrations
make db-seed      # Seed database with sample data
make db-reset     # Reset database (WARNING: deletes all data)
make db-backup    # Backup database
```

### API

```bash
make api-gen      # Generate OpenAPI specs from code
make api-docs     # Serve API documentation
make proto-gen    # Generate protobuf files
```

## Variables

Override these variables as needed:

```bash
make dev GO_ENV=development
make test GO_TEST_FLAGS="-v -race"
make build LD_FLAGS="-s -w"
```

Common variables:
- `GO_ENV`: Environment (development, staging, production)
- `GO_OS`: Target OS (linux, darwin, windows)
- `GO_ARCH`: Target architecture (amd64, arm64)
- `GO_TEST_FLAGS`: Additional test flags
- `LD_FLAGS`: Linker flags

## Examples

### Run with custom environment
```bash
make dev GO_ENV=staging
```

### Build for Linux ARM64
```bash
make build GO_OS=linux GO_ARCH=arm64
```

### Run specific test package
```bash
make test-unit PKG=./internal/handler
```

### View coverage in browser
```bash
make test-coverage
open coverage.out.html
```

## Internal Targets

These targets are used by other targets and typically not called directly:

- `check-go`: Verify Go is installed
- `check-deps`: Verify dependencies are installed
- `install-tools`: Install development tools
- `deps`: Download Go modules

## CI/CD Integration

Used in GitHub Actions workflows:

```yaml
- name: Test
  run: make test-unit
  
- name: Lint
  run: make lint
  
- name: Build
  run: make build
```

## Related Documentation

- [Root Makefile](../../Makefile)
- [Backend Architecture](../../docs/02-ARCHITECTURE.md)
- [Development Setup](../../docs/03-DEVELOPMENT.md)
