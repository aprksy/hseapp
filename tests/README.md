# Tests

Test suites for HSE App following the testing pyramid.

## Directory Structure

```
tests/
├── unit/               # Unit tests (fast, isolated)
│   ├── backend/        # Go unit tests
│   ├── web/            # SvelteKit component tests
│   └── mobile/         # Flutter widget tests
├── integration/        # Integration tests (service interactions)
│   ├── api/            # API integration tests
│   ├── database/       # Database integration tests
│   └── external/       # External service integration tests
└── e2e/                # End-to-end tests (full user flows)
    ├── web/            # Web E2E tests (Playwright/Cypress)
    └── mobile/         # Mobile E2E tests (Flutter integration_test)
```

## Testing Strategy

### Unit Tests
- Test individual functions, methods, and components
- Mock all external dependencies
- Fast execution (< 1 second per test)
- High coverage target (> 80%)

**Run:**
```bash
make test-unit
# or
./scripts/test-unit.sh
```

### Integration Tests
- Test interactions between components
- Use test database with seeded data
- Test API endpoints, database queries, external services
- Moderate execution time (< 10 seconds per test)

**Run:**
```bash
make test-integration
# or
./scripts/test-integration.sh
```

### E2E Tests
- Test complete user workflows
- Run against staging environment
- Include authentication, navigation, data persistence
- Slower execution (< 60 seconds per test)

**Run:**
```bash
make test-e2e
# or
./scripts/test-e2e.sh
```

## Test Data

Test data is managed through:
- **Fixtures**: Static test data in `tests/fixtures/`
- **Factories**: Dynamic test data generators
- **Seeds**: Database seeding scripts

## Coverage Reports

Generate coverage reports:

```bash
make test-coverage
```

Outputs:
- `coverage/unit.html` - Unit test coverage
- `coverage/integration.html` - Integration test coverage
- `coverage/total.html` - Combined coverage

## CI/CD Integration

Tests run automatically on:
- Pull requests (unit + integration)
- Merges to main (all tests)
- Scheduled nightly runs (E2E tests)

## Writing Tests

### Backend (Go)

```go
// backend/internal/repository/user_test.go
package repository

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
    // Arrange
    repo := NewTestUserRepository()
    user := &domain.User{Email: "test@example.com"}
    
    // Act
    err := repo.Create(user)
    
    // Assert
    assert.NoError(t, err)
    assert.NotEmpty(t, user.ID)
}
```

### Web (SvelteKit + Vitest)

```typescript
// web/src/lib/components/Button.test.ts
import { render, screen } from '@testing-library/svelte'
import Button from './Button.svelte'

describe('Button', () => {
  it('renders with label', () => {
    render(Button, { props: { label: 'Click me' } })
    expect(screen.getByText('Click me')).toBeInTheDocument()
  })
})
```

### Mobile (Flutter)

```dart
// mobile/test/widgets/login_form_test.dart
import 'package:flutter_test/flutter_test.dart'

void main() {
  testWidgets('Login form validates email', (tester) async {
    await tester.pumpWidget(LoginForm())
    
    await tester.enterText(find.byType(TextFormField), 'invalid')
    await tester.tap(find.byType(ElevatedButton))
    await tester.pump()
    
    expect(find.text('Invalid email'), findsOneWidget)
  })
}
```

## Related Documentation

- [Architecture](../docs/02-ARCHITECTURE.md)
- [Development Setup](../docs/03-DEVELOPMENT.md)
- [Makefile](../Makefile)
