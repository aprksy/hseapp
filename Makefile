# HSE App Makefile - Development Operations

.PHONY: help dev-backend dev-web dev-mobile test-backend test-web test-mobile build-backend build-web build-mobile clean lint-backend lint-web lint-mobile proto-generate docs-serve setup

# Default target
help:
	@echo "HSE App Development Commands"
	@echo "=============================="
	@echo ""
	@echo "Development:"
	@echo "  dev-backend     - Start backend server (Go)"
	@echo "  dev-web         - Start web dev server (SvelteKit)"
	@echo "  dev-mobile      - Start mobile dev (Flutter)"
	@echo "  dev-all         - Start all services"
	@echo ""
	@echo "Testing:"
	@echo "  test-backend    - Run backend tests"
	@echo "  test-web        - Run web tests"
	@echo "  test-mobile     - Run mobile tests"
	@echo "  test-all        - Run all tests"
	@echo ""
	@echo "Building:"
	@echo "  build-backend   - Build backend binary"
	@echo "  build-web       - Build web production bundle"
	@echo "  build-mobile    - Build mobile APK/IPA"
	@echo "  build-all       - Build all projects"
	@echo ""
	@echo "Code Quality:"
	@echo "  lint-backend    - Lint backend code"
	@echo "  lint-web        - Lint web code"
	@echo "  lint-mobile     - Lint mobile code"
	@echo "  lint-all        - Lint all projects"
	@echo ""
	@echo "Utilities:"
	@echo "  proto-generate  - Generate protobuf files"
	@echo "  docs-serve      - Serve documentation"
	@echo "  clean           - Clean build artifacts"
	@echo "  setup           - Initial project setup"

# Development
dev-backend:
	@echo "Starting backend server..."
	cd backend && go run cmd/main.go

dev-web:
	@echo "Starting web dev server..."
	cd web && npm install && npm run dev -- --host 0.0.0.0

dev-mobile:
	@echo "Starting mobile development..."
	cd mobile && flutter pub get && flutter run

dev-all:
	@echo "Starting all services..."
	@make dev-backend &
	@make dev-web &
	@echo "All services started. Press Ctrl+C to stop."

# Testing
test-backend:
	@echo "Running backend tests..."
	cd backend && go test ./... -v

test-web:
	@echo "Running web tests..."
	cd web && npm run test

test-mobile:
	@echo "Running mobile tests..."
	cd mobile && flutter test

test-all: test-backend test-web test-mobile

# Building
build-backend:
	@echo "Building backend..."
	cd backend && go build -o bin/server cmd/main.go

build-web:
	@echo "Building web..."
	cd web && npm run build

build-mobile:
	@echo "Building mobile..."
	cd mobile && flutter build apk --release

build-all: build-backend build-web build-mobile

# Code Quality
lint-backend:
	@echo "Linting backend..."
	cd backend && gofmt -s -w . && go vet ./...

lint-web:
	@echo "Linting web..."
	cd web && npm run lint

lint-mobile:
	@echo "Linting mobile..."
	cd mobile && flutter analyze

lint-all: lint-backend lint-web lint-mobile

# Utilities
proto-generate:
	@echo "Generating protobuf files..."
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       proto/*.proto

docs-serve:
	@echo "Serving documentation..."
	@python3 -m http.server 8000 --directory docs

clean:
	@echo "Cleaning build artifacts..."
	rm -rf backend/bin
	rm -rf web/build web/.svelte-kit
	rm -rf mobile/build
	find . -name "node_modules" -type d -prune -exec rm -rf '{}' +
	find . -name ".dart_tool" -type d -prune -exec rm -rf '{}' +

setup:
	@echo "Setting up project..."
	@echo "Installing backend dependencies..."
	cd backend && go mod tidy
	@echo "Installing web dependencies..."
	cd web && npm install
	@echo "Installing mobile dependencies..."
	cd mobile && flutter pub get
	@echo "Setup complete!"

# Environment variables
export PORT ?= 8080
export DATABASE_URL ?= postgresql://localhost:5432/hseapp
export JWT_SECRET ?= your-secret-key-change-in-production
export SUPABASE_URL ?= https://your-project.supabase.co
export SUPABASE_KEY ?= your-anon-key
export FIREBASE_PROJECT_ID ?= your-firebase-project
