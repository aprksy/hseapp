# Scripts

Utility scripts for development, testing, and operations.

## Available Scripts

### Development Scripts

- `dev-setup.sh` - One-time development environment setup
- `db-seed.sh` - Seed database with sample data
- `gen-protos.sh` - Generate protobuf files for all services
- `gen-clients.sh` - Generate API clients (OpenAPI, gRPC)

### Testing Scripts

- `test-unit.sh` - Run unit tests across all services
- `test-integration.sh` - Run integration tests
- `test-e2e.sh` - Run end-to-end tests
- `test-coverage.sh` - Generate test coverage reports

### Operations Scripts

- `deploy.sh` - Deploy to specified environment
- `rollback.sh` - Rollback to previous version
- `backup-db.sh` - Backup database
- `restore-db.sh` - Restore database from backup
- `migrate.sh` - Run database migrations
- `health-check.sh` - Check service health

### Utility Scripts

- `clean.sh` - Clean build artifacts
- `lint.sh` - Run linters on all code
- `format.sh` - Format code across all services
- `env-check.sh` - Validate environment configuration

## Usage

Most scripts can be run from the root directory:

```bash
./scripts/dev-setup.sh
./scripts/test-unit.sh
./scripts/deploy.sh dev
```

Or via Makefile:

```bash
make setup
make test
make deploy ENV=dev
```

## Script Conventions

All scripts should:
- Be written in bash (with shebang `#!/usr/bin/env bash`)
- Use `set -euo pipefail` for safety
- Accept environment name as first argument (default: `dev`)
- Exit with non-zero status on failure
- Log actions to stdout/stderr
- Support `--help` flag for usage information

## Adding New Scripts

1. Create script in `scripts/` directory
2. Make it executable: `chmod +x scripts/your-script.sh`
3. Add corresponding Makefile target if frequently used
4. Document in this README

## Related Documentation

- [Makefile](../Makefile)
- [CI/CD Configuration](../.github/workflows/)
- [Architecture](../docs/02-ARCHITECTURE.md)
