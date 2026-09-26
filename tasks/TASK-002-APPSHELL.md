# TASK-002: App Shell Implementation (SPEC-001)

## Overview
Implement the foundational App Shell for Web and Mobile clients, establishing the visual framework, navigation structure, and core layout components that will house all future features.

## Objectives
- Create responsive, multi-tenant aware app shell
- Implement navigation system (sidebar, topbar, mobile menu)
- Set up theme system (light/dark/system)
- Establish i18n foundation (Indonesian + English)
- Create placeholder areas for future features
- Ensure accessibility (WCAG 2.1 AA)

## Scope

### Web Client (SvelteKit + Shadcn-Svelte)
- [ ] AppShell component with slots for content
- [ ] Sidebar navigation with collapsible sections
- [ ] TopBar with user menu, notifications, language/theme toggles
- [ ] Mobile-responsive hamburger menu
- [ ] Breadcrumb navigation
- [ ] Footer with compliance info
- [ ] Theme provider (light/dark/system)
- [ ] i18n setup with id_ID and en_ID locales
- [ ] Multi-tenant branding (logo, colors)
- [ ] Loading skeletons
- [ ] Error boundaries

### Mobile Client (Flutter)
- [ ] AppShell scaffold with BottomNavigationBar
- [ ] Drawer for extended navigation
- [ ] AppBar with user info and notifications
- [ ] Responsive layouts (phone/tablet)
- [ ] Theme system matching web
- [ ] i18n setup matching web
- [ ] Multi-tenant branding
- [ ] Loading indicators
- [ ] Error handling screens

### Shared/Backend
- [ ] User profile API endpoint
- [ ] Tenant configuration API
- [ ] Preference storage (theme, language)
- [ ] Notification count API

## Acceptance Criteria

### Functional
- ✅ User can switch between light/dark/system themes
- ✅ User can switch between Indonesian and English
- ✅ Navigation responds to user roles (role-based menu items)
- ✅ Multi-tenant branding loads correctly
- ✅ Mobile menu works smoothly on devices < 768px
- ✅ All interactive elements are keyboard accessible
- ✅ Loading states show during data fetch
- ✅ Error states display gracefully

### Technical
- ✅ Components follow DDD separation (presentation layer only)
- ✅ No business logic in UI components
- ✅ All text uses i18n keys (no hardcoded strings)
- ✅ Theme values from CSS variables / Flutter ThemeData
- ✅ Navigation config driven by JSON/YAML (not hardcoded)
- ✅ Unit tests for utility functions
- ✅ E2E tests for critical user flows
- ✅ Lighthouse score > 90 for performance/accessibility

## Implementation Phases

### Phase 1: Foundation (2 days)
**Web:**
- Setup i18n with svelte-i18n
- Configure theme provider
- Create base layout components

**Mobile:**
- Setup intl package
- Configure ThemeData for light/dark
- Create base scaffold

**Backend:**
- User profile endpoint
- Tenant config endpoint
- Preferences CRUD

### Phase 2: Navigation (2 days)
**Web:**
- Sidebar component with Shadcn Collapsible
- TopBar with user menu
- Mobile drawer
- Breadcrumb component

**Mobile:**
- BottomNavigationBar
- Drawer with DataTable
- AppBar actions

### Phase 3: Polish & Testing (1 day)
- Loading skeletons
- Error boundaries
- Accessibility audit
- Cross-browser testing
- Device testing (mobile)
- Performance optimization

## File Structure

```
web/src/lib/components/
├── app-shell/
│   ├── AppShell.svelte
│   ├── Sidebar.svelte
│   ├── TopBar.svelte
│   ├── MobileNav.svelte
│   ├── Breadcrumb.svelte
│   └── Footer.svelte
├── theme/
│   ├── ThemeProvider.svelte
│   └── ThemeToggle.svelte
└── i18n/
    ├── index.ts
    ├── locales/id_ID.json
    └── locales/en_ID.json

mobile/lib/
├── core/
│   ├── shell/
│   │   ├── app_shell.dart
│   │   ├── app_drawer.dart
│   │   └── app_bar.dart
│   └── theme/
│       ├── app_theme.dart
│       └── theme_provider.dart
└── l10n/
    ├── app_id.arb
    └── app_en.arb

backend/internal/
├── handlers/
│   └── profile.go
├── services/
│   └── profile_service.go
└── repository/
    └── profile_repo.go
```

## Dependencies

### Web
- svelte-i18n
- @supabase/supabase-js
- shadcn-svelte components
- lucide-svelte (icons)

### Mobile
- intl
- flutter_localizations
- get (state management)
- supabase_flutter

### Backend
- Existing Go modules
- JWT for auth

## Testing Strategy

### Unit Tests
- Theme toggle logic
- Language switching
- Navigation config parsing
- Role-based menu filtering

### Integration Tests
- User preference persistence
- Multi-tenant branding load
- API integration for profile

### E2E Tests
- Login → App Shell load
- Theme switch persists
- Language switch persists
- Navigation to all sections
- Mobile responsive behavior

## Definition of Done
- [ ] All components implemented
- [ ] All acceptance criteria met
- [ ] Unit tests passing (>80% coverage)
- [ ] E2E tests passing
- [ ] Accessibility audit passed
- [ ] Performance budget met (< 100KB initial load)
- [ ] Code reviewed
- [ ] Documentation updated
- [ ] Deployed to staging environment

## Risks & Mitigation

| Risk | Impact | Mitigation |
|------|--------|------------|
| Theme system conflicts | High | Use CSS variables, test early |
| i18n missing keys | Medium | Linting rules, type-safe keys |
| Mobile performance | Medium | Lazy load components, code split |
| Multi-tenant branding complexity | Low | Config-driven approach |

## Estimated Effort
- **Total**: 5 days (40 hours)
- **Web**: 18 hours
- **Mobile**: 14 hours
- **Backend**: 6 hours
- **Testing**: 2 hours

## Next Steps After Completion
1. Begin SPEC-002: Authentication & Authorization
2. Integrate login flow with App Shell
3. Add role-based navigation filtering
4. Implement user session management

---

**Status**: Ready to Implement  
**Priority**: High (Foundation for all features)  
**Assigned To**: Frontend Team (Web + Mobile)  
**Reviewer**: Architecture Team  
**Created**: 2026-01-13  
**Target Complete**: 2026-01-20
