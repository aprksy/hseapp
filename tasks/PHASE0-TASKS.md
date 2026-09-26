# HSE App - Task Breakdown (Phase 0: Foundation)

## Overview

This document contains detailed, actionable tasks for **Phase 0: Foundation & Setup**. Each task is designed to be completed independently with clear acceptance criteria.

**Phase 0 Duration**: 1-2 weeks  
**Goal**: Establish development environment, tooling, and project scaffolding

---

## Task 0.1: Repository Structure Setup

**ID**: TASK-001  
**Priority**: CRITICAL  
**Estimated Time**: 2 hours  
**Dependencies**: None  

### Description
Create the foundational directory structure for the monorepo-style repository containing backend, web, mobile, infrastructure, and documentation.

### Actions
1. Create directory structure as defined in plan
2. Add `.gitkeep` files to empty directories
3. Create root `.gitignore` file
4. Add initial directory README files

### Acceptance Criteria
- [ ] All directories from plan created
- [ ] `.gitignore` configured for Go, SvelteKit, Flutter
- [ ] Directory structure committed to git
- [ ] Team can navigate structure clearly

### Files to Create
```
/workspace/
├── .gitignore
├── docs/.gitkeep
├── backend/.gitkeep
├── web/.gitkeep
├── mobile/.gitkeep
├── infra/.gitkeep
├── scripts/.gitkeep
└── tasks/.gitkeep
```

---

## Task 0.2: Root Makefile Creation

**ID**: TASK-002  
**Priority**: CRITICAL  
**Estimated Time**: 3 hours  
**Dependencies**: TASK-001  

### Description
Create the root Makefile that orchestrates development operations across all projects (backend, web, mobile).

### Actions
1. Define common targets: `help`, `setup`, `dev`, `build`, `test`, `clean`
2. Add recursive make calls for subprojects
3. Include Docker commands for local services
4. Add linting and formatting targets
5. Document all targets in help

### Acceptance Criteria
- [ ] `make help` shows all available commands
- [ ] `make setup` initializes all subprojects
- [ ] `make dev` starts all development servers
- [ ] `make test` runs tests across all projects
- [ ] Makefile follows best practices (phony targets, etc.)

### Example Targets
```makefile
.PHONY: help setup dev build test clean docker-up docker-down

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Initialize all projects
	@echo "Setting up backend..."
	@$(MAKE) -C backend setup
	@echo "Setting up web frontend..."
	@$(MAKE) -C web setup
	@echo "Setting up mobile app..."
	@$(MAKE) -C mobile setup

dev: ## Start all development servers
	@echo "Starting backend..."
	@$(MAKE) -C backend dev &
	@echo "Starting web frontend..."
	@$(MAKE) -C web dev &
	@wait

build: ## Build all projects
	@$(MAKE) -C backend build
	@$(MAKE) -C web build
	@$(MAKE) -C mobile build

test: ## Run all tests
	@$(MAKE) -C backend test
	@$(MAKE) -C web test
	@$(MAKE) -C mobile test

clean: ## Clean all build artifacts
	@$(MAKE) -C backend clean
	@$(MAKE) -C web clean
	@$(MAKE) -C mobile clean

docker-up: ## Start Docker services (Supabase local, etc.)
	docker-compose -f infra/docker/docker-compose.yml up -d

docker-down: ## Stop Docker services
	docker-compose -f infra/docker/docker-compose.yml down
```

---

## Task 0.3: Backend Project Initialization (Go)

**ID**: TASK-003  
**Priority**: CRITICAL  
**Estimated Time**: 4 hours  
**Dependencies**: TASK-001, TASK-002  

### Description
Initialize the Go backend project with proper structure, dependencies, and basic configuration.

### Actions
1. Initialize Go module (`go mod init`)
2. Create standard Go project structure
3. Add main.go with health check endpoint
4. Configure logging (zap)
5. Set up environment variable management (viper)
6. Create backend Makefile
7. Add initial unit test

### Acceptance Criteria
- [ ] `go run cmd/server/main.go` starts server
- [ ] Health check endpoint responds at `/health`
- [ ] Logging configured and working
- [ ] Environment variables loaded from `.env`
- [ ] `make dev` starts server with hot reload
- [ ] Unit test passes

### Files to Create
```
/backend/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   └── health.go
│   └── middleware/
│       └── logger.go
├── pkg/
│   └── utils/
│       └── utils.go
├── go.mod
├── go.sum
├── .env.example
├── Makefile
└── README.md
```

### Key Code Snippets

**main.go**:
```go
package main

import (
    "log"
    "net/http"
    "github.com/hseapp/backend/internal/config"
    "github.com/hseapp/backend/internal/handlers"
)

func main() {
    cfg := config.Load()
    
    http.HandleFunc("/health", handlers.HealthCheck)
    
    log.Printf("Server starting on port %s", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
```

**Makefile**:
```makefile
.PHONY: setup dev build test clean

setup:
	go mod download
	go install github.com/cosmtrek/air@latest

dev:
	air -c .air.toml

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/
	go clean
```

---

## Task 0.4: Web Frontend Initialization (SvelteKit)

**ID**: TASK-004  
**Priority**: CRITICAL  
**Estimated Time**: 6 hours  
**Dependencies**: TASK-001, TASK-002  

### Description
Initialize SvelteKit frontend with TailwindCSS, Shadcn-Svelte, i18n, and theme system.

### Actions
1. Create SvelteKit project with TypeScript
2. Install and configure TailwindCSS
3. Install Shadcn-Svelte CLI and components
4. Set up svelte-i18n for internationalization
5. Configure theme system with CSS variables
6. Create base layout with sidebar and topbar
7. Set up Supabase client
8. Create web Makefile
9. Add Vitest configuration

### Acceptance Criteria
- [ ] `npm run dev` starts dev server
- [ ] TailwindCSS classes working
- [ ] Shadcn components importable
- [ ] Language switcher changes UI text
- [ ] Theme toggle switches light/dark
- [ ] Base layout renders correctly
- [ ] Supabase client configured
- [ ] Tests run with Vitest

### Files to Create
```
/web/
├── src/
│   ├── lib/
│   │   ├── components/
│   │   │   ├── ui/          # Shadcn components
│   │   │   ├── layout/
│   │   │   │   ├── Sidebar.svelte
│   │   │   │   ├── TopBar.svelte
│   │   │   │   └── Footer.svelte
│   │   │   └── shared/
│   │   ├── stores/
│   │   │   ├── theme.store.ts
│   │   │   └── locale.store.ts
│   │   ├── i18n/
│   │   │   ├── en.json
│   │   │   ├── id.json
│   │   │   └── index.ts
│   │   ├── supabase/
│   │   │   └── client.ts
│   │   └── utils/
│   │       └── cn.ts
│   ├── routes/
│   │   ├── +layout.svelte
│   │   └── +page.svelte
│   ├── app.html
│   ├── app.css
│   └── app.d.ts
├── static/
├── tailwind.config.js
├── postcss.config.js
├── vite.config.ts
├── svelte.config.js
├── tsconfig.json
├── package.json
├── .env.example
├── Makefile
└── README.md
```

### Key Configuration

**tailwind.config.js**:
```javascript
/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        border: 'hsl(var(--border))',
        input: 'hsl(var(--input))',
        ring: 'hsl(var(--ring))',
        background: 'hsl(var(--background))',
        foreground: 'hsl(var(--foreground))',
        // ... Shadcn color scheme
      },
    },
  },
  plugins: [require('tailwindcss-animate')],
}
```

**i18n/index.ts**:
```typescript
import { init, register, getLocaleFromNavigator } from 'svelte-i18n';

register('en', () => import('./en.json'));
register('id', () => import('./id.json'));

init({
  fallbackLocale: 'en',
  initialLocale: getLocaleFromNavigator(),
});
```

---

## Task 0.5: Mobile App Initialization (Flutter)

**ID**: TASK-005  
**Priority**: HIGH  
**Estimated Time**: 6 hours  
**Dependencies**: TASK-001, TASK-002  

### Description
Initialize Flutter mobile app with clean architecture, state management, and Firebase integration.

### Actions
1. Create Flutter project
2. Set up directory structure (clean architecture)
3. Configure state management (Riverpod)
4. Set up local database (Hive)
5. Configure HTTP client (Dio)
6. Integrate Firebase Core
7. Configure FCM
8. Create mobile Makefile
9. Set up environment flavors

### Acceptance Criteria
- [ ] `flutter run` launches app on emulator
- [ ] Clean architecture folders created
- [ ] Riverpod providers working
- [ ] Hive database initialized
- [ ] Dio HTTP client configured
- [ ] Firebase initialized
- [ ] Flavors configured (dev, prod)
- [ ] Tests run successfully

### Files to Create
```
/mobile/
├── lib/
│   ├── core/
│   │   ├── constants/
│   │   ├── errors/
│   │   ├── network/
│   │   │   └── dio_client.dart
│   │   ├── theme/
│   │   └── utils/
│   ├── data/
│   │   ├── datasources/
│   │   ├── models/
│   │   └── repositories/
│   ├── domain/
│   │   ├── entities/
│   │   ├── repositories/
│   │   └── usecases/
│   ├── presentation/
│   │   ├── providers/
│   │   ├── screens/
│   │   └── widgets/
│   ├── injection_container.dart
│   └── main.dart
├── android/
├── ios/
├── test/
├── pubspec.yaml
├── .env.example
├── Makefile
└── README.md
```

**pubspec.yaml** (key dependencies):
```yaml
dependencies:
  flutter:
    sdk: flutter
  flutter_riverpod: ^2.4.0
  dio: ^5.3.0
  hive: ^2.2.3
  hive_flutter: ^1.1.0
  firebase_core: ^2.24.0
  firebase_messaging: ^14.7.0
  intl: ^0.18.0
  
dev_dependencies:
  flutter_test:
    sdk: flutter
  hive_generator: ^2.0.1
  build_runner: ^2.4.6
```

---

## Task 0.6: Infrastructure Setup (Docker & Services)

**ID**: TASK-006  
**Priority**: HIGH  
**Estimated Time**: 4 hours  
**Dependencies**: TASK-001  

### Description
Set up local development infrastructure using Docker Compose, including Supabase local instance and other services.

### Actions
1. Create Docker Compose configuration
2. Configure Supabase local (PostgreSQL + Auth + Storage)
3. Set up Redis for caching/sessions (optional)
4. Configure mailcatcher for email testing
5. Create network and volume configurations
6. Add health checks
7. Document setup process

### Acceptance Criteria
- [ ] `docker-compose up` starts all services
- [ ] PostgreSQL accessible from backend
- [ ] Supabase Auth working locally
- [ ] Mailcatcher captures test emails
- [ ] Health checks passing
- [ ] Services persist data in volumes

### Files to Create

**infra/docker/docker-compose.yml**:
```yaml
version: '3.8'

services:
  # PostgreSQL (Supabase compatible)
  postgres:
    image: supabase/postgres:15.1.0.117
    ports:
      - "54322:5432"
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: hse_app
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5

  # Supabase Auth (Kong gateway)
  kong:
    image: kong:2.8.1
    ports:
      - "8000:8000"
    environment:
      KONG_DATABASE: "off"
      KONG_DECLARATIVE_CONFIG: /kong/declarative/kong.yml
    volumes:
      - ./kong:/kong/declarative

  # Redis (optional)
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]

  # Mailcatcher for email testing
  mailcatcher:
    image: schickling/mailcatcher
    ports:
      - "1080:1080"
      - "1025:1025"

volumes:
  postgres_data:
  redis_data:
```

---

## Task 0.7: CI/CD Pipeline Setup

**ID**: TASK-007  
**Priority**: HIGH  
**Estimated Time**: 4 hours  
**Dependencies**: TASK-003, TASK-004, TASK-005  

### Description
Create GitHub Actions workflows for automated testing, linting, and building.

### Actions
1. Create workflow directory structure
2. Set up backend CI (Go)
3. Set up web CI (SvelteKit)
4. Set up mobile CI (Flutter)
5. Create combined workflow
6. Add caching for dependencies
7. Configure artifact storage

### Acceptance Criteria
- [ ] Push triggers CI pipeline
- [ ] Backend tests run automatically
- [ ] Web tests run automatically
- [ ] Mobile builds successfully
- [ ] Linting checks pass
- [ ] Caching reduces build time
- [ ] Artifacts stored for downloads

### Files to Create

**.github/workflows/ci.yml**:
```yaml
name: CI

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main, develop]

jobs:
  backend-test:
    runs-on: ubuntu-latest
    defaults:
      run:
        working-directory: ./backend
    steps:
      - uses: actions/checkout@v3
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Cache Go modules
        uses: actions/cache@v3
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
      - name: Install dependencies
        run: go mod download
      - name: Run tests
        run: go test -v ./...
      - name: Build
        run: go build -v ./...

  web-test:
    runs-on: ubuntu-latest
    defaults:
      run:
        working-directory: ./web
    steps:
      - uses: actions/checkout@v3
      - name: Set up Node.js
        uses: actions/setup-node@v3
        with:
          node-version: '18'
      - name: Cache npm dependencies
        uses: actions/cache@v3
        with:
          path: ~/.npm
          key: ${{ runner.os }}-node-${{ hashFiles('**/package-lock.json') }}
      - name: Install dependencies
        run: npm ci
      - name: Run tests
        run: npm test
      - name: Build
        run: npm run build

  mobile-build:
    runs-on: ubuntu-latest
    defaults:
      run:
        working-directory: ./mobile
    steps:
      - uses: actions/checkout@v3
      - name: Set up Flutter
        uses: subosito/flutter-action@v2
        with:
          flutter-version: '3.13.0'
      - name: Cache Pub dependencies
        uses: actions/cache@v3
        with:
          path: ~/.pub-cache
          key: ${{ runner.os }}-pub-${{ hashFiles('**/pubspec.lock') }}
      - name: Install dependencies
        run: flutter pub get
      - name: Analyze
        run: flutter analyze
      - name: Run tests
        run: flutter test
      - name: Build APK
        run: flutter build apk --debug
```

---

## Task 0.8: Documentation Setup

**ID**: TASK-008  
**Priority**: MEDIUM  
**Estimated Time**: 3 hours  
**Dependencies**: TASK-001  

### Description
Create comprehensive documentation including README, development guide, and API standards.

### Actions
1. Update root README.md with project overview
2. Create DEVELOPMENT.md with setup instructions
3. Document API standards and conventions
4. Create CONTRIBUTING.md guide
5. Add CODE_OF_CONDUCT.md
6. Set up API documentation (OpenAPI/Swagger)

### Acceptance Criteria
- [ ] README provides clear project overview
- [ ] DEVELOPMENT.md enables new developer onboarding
- [ ] API standards documented
- [ ] Contribution guidelines clear
- [ ] Code of conduct established
- [ ] API docs accessible

### Files to Create

**README.md** (root):
```markdown
# HSE App

Health, Safety, and Environment management platform for Indonesian regulations with international standards support.

## Quick Links

- [Constitution](docs/01-CONSTITUTION.md)
- [Specifications](specs/00-SPECIFICATIONS.md)
- [Implementation Plan](plans/00-IMPLEMENTATION_PLAN.md)
- [Development Guide](docs/DEVELOPMENT.md)

## Tech Stack

- **Backend**: Go, Supabase (PostgreSQL), Firebase
- **Web**: SvelteKit, Shadcn-Svelte, TailwindCSS
- **Mobile**: Flutter, Hive, Firebase Cloud Messaging
- **Infrastructure**: Docker, GitHub Actions

## Getting Started

```bash
# Clone repository
git clone https://github.com/your-org/hse-app.git
cd hse-app

# Setup all projects
make setup

# Start development servers
make dev

# Run tests
make test
```

## Project Structure

- `/backend` - Go backend API
- `/web` - SvelteKit web application
- `/mobile` - Flutter mobile app
- `/docs` - Documentation
- `/specs` - Product specifications
- `/plans` - Implementation plans
- `/tasks` - Task breakdowns
- `/infra` - Infrastructure as code

## License

[Your License Here]
```

**docs/DEVELOPMENT.md**:
```markdown
# Development Guide

## Prerequisites

- Go 1.21+
- Node.js 18+
- Flutter 3.13+
- Docker & Docker Compose
- Make

## Local Setup

### 1. Clone and Initialize

```bash
git clone https://github.com/your-org/hse-app.git
cd hse-app
make setup
```

### 2. Environment Configuration

Copy example env files and fill in values:

```bash
cp backend/.env.example backend/.env
cp web/.env.example web/.env
cp mobile/.env.example mobile/.env
```

### 3. Start Services

```bash
# Start Docker services (PostgreSQL, Redis, etc.)
make docker-up

# Start development servers
make dev
```

### 4. Access Services

- Backend API: http://localhost:8080
- Web App: http://localhost:5173
- Mailcatcher: http://localhost:1080
- PostgreSQL: localhost:54322

## Common Commands

```bash
make help        # Show all commands
make dev         # Start all dev servers
make test        # Run all tests
make build       # Build all projects
make clean       # Clean build artifacts
make docker-up   # Start Docker services
make docker-down # Stop Docker services
```

## Backend Development

```bash
cd backend
make dev    # Start with hot reload
make test   # Run tests
make build  # Build binary
```

## Web Development

```bash
cd web
npm run dev    # Start dev server
npm test       # Run tests
npm run build  # Build for production
```

## Mobile Development

```bash
cd mobile
flutter run           # Run on connected device/emulator
flutter test          # Run tests
flutter build apk     # Build APK
flutter build ios     # Build iOS app
```

## Database

Access PostgreSQL:

```bash
psql -h localhost -p 54322 -U postgres -d hse_app
```

## Troubleshooting

### Common Issues

1. **Port already in use**
   ```bash
   lsof -ti:8080 | xargs kill
   ```

2. **Docker services not starting**
   ```bash
   docker-compose down -v
   docker-compose up
   ```

3. **Node modules issues**
   ```bash
   rm -rf node_modules package-lock.json
   npm install
   ```

## Testing

Run all tests:

```bash
make test
```

Run specific project tests:

```bash
cd backend && make test
cd web && npm test
cd mobile && flutter test
```

## Code Style

- Backend: `gofmt`, `golint`
- Web: ESLint, Prettier
- Mobile: `dart format`, `flutter analyze`

## Deployment

(To be added in Phase 3)
```

---

## Task Summary

| ID | Task | Priority | Est. Time | Status |
|----|------|----------|-----------|--------|
| TASK-001 | Repository Structure | CRITICAL | 2h | ⏳ Pending |
| TASK-002 | Root Makefile | CRITICAL | 3h | ⏳ Pending |
| TASK-003 | Backend Initialization | CRITICAL | 4h | ⏳ Pending |
| TASK-004 | Web Initialization | CRITICAL | 6h | ⏳ Pending |
| TASK-005 | Mobile Initialization | HIGH | 6h | ⏳ Pending |
| TASK-006 | Infrastructure Setup | HIGH | 4h | ⏳ Pending |
| TASK-007 | CI/CD Pipeline | HIGH | 4h | ⏳ Pending |
| TASK-008 | Documentation | MEDIUM | 3h | ⏳ Pending |

**Total Estimated Time**: 32 hours (≈ 4 working days)

---

## Approval Checklist

Before proceeding to Phase 1, verify:

- [ ] All Phase 0 tasks completed
- [ ] Team members can run `make setup` successfully
- [ ] All development servers start with `make dev`
- [ ] CI/CD pipeline green on main branch
- [ ] Documentation complete and accessible
- [ ] No critical bugs or blockers

---

**Document Status**: Draft v1.0  
**Created**: Phase 0 Tasks  
**Next Step**: Begin TASK-001 after approval  
**Approval Required**: Yes - Before starting implementation
