# Mobile Client Makefile

Makefile for mobile client (Flutter) development operations.

## Prerequisites

- Flutter 3.16+
- Dart 3.2+
- Android SDK (for Android builds)
- Xcode (for iOS builds, macOS only)
- Make 4.0+

## Common Commands

### Development

```bash
make dev          # Start app in development mode
make run          # Run on connected device/emulator
make build        # Build production APK/IPA
make clean        # Clean build artifacts
```

### Testing

```bash
make test         # Run all tests
make test-unit    # Run unit tests
make test-widget  # Run widget tests
make test-integration  # Run integration tests
make test-coverage    # Run tests with coverage report
```

### Code Quality

```bash
make lint         # Run Dart analyzer
make format       # Format code (dart format)
make check        # Run all quality checks
make pub-get      # Get dependencies
make pub-upgrade  # Upgrade dependencies
```

### Generation

```bash
make gen-models   # Generate data models from JSON
make gen-icons    # Generate icon files
make gen-assets   # Generate asset bindings
make gen-l10n     # Generate localization files
```

### Build Variants

```bash
make apk          # Build Android APK (debug)
make apk-release  # Build Android APK (release)
make ipa          # Build iOS IPA (requires macOS)
make appbundle    # Build Android App Bundle
```

## Variables

Override these variables as needed:

```bash
make dev DEVICE_ID=emulator-5554
make build FLAVOR=production
make test PLATFORM=android
```

Common variables:
- `DEVICE_ID`: Target device/emulator ID
- `FLAVOR`: Build flavor (development, staging, production)
- `PLATFORM`: Target platform (android, ios, web)
- `DART_FLAGS`: Additional Dart flags

## Examples

### Run on specific emulator
```bash
make dev DEVICE_ID=emulator-5554
```

### Build production APK
```bash
make apk-release FLAVOR=production
```

### Run tests with coverage
```bash
make test-coverage
lcov --list coverage/lcov.info
```

### Generate models from API types
```bash
make gen-models INPUT=../proto/api.proto
```

## Device Management

```bash
make devices      # List connected devices
make emulator     # Start Android emulator
make ios-sim      # Start iOS simulator
make screenshot   # Capture screenshots for store
```

## Internal Targets

These targets are used by other targets and typically not called directly:

- `check-flutter`: Verify Flutter is installed
- `check-deps`: Verify dependencies are installed
- `install`: Install dependencies (flutter pub get)
- `deps`: Sync dependencies

## CI/CD Integration

Used in GitHub Actions workflows:

```yaml
- name: Install Dependencies
  run: make install
  
- name: Analyze
  run: make lint
  
- name: Test
  run: make test-unit
  
- name: Build APK
  run: make apk-release
```

## Project Structure

```
mobile/
├── lib/
│   ├── main.dart       # App entry point
│   ├── app/            # App-level components
│   ├── features/       # Feature modules
│   │   ├── auth/       # Authentication feature
│   │   ├── dashboard/  # Dashboard feature
│   │   └── ...
│   ├── core/           # Core utilities
│   │   ├── factories/  # Factory functions
│   │   ├── services/   # Services (API, storage)
│   │   └── widgets/    # Reusable widgets
│   └── l10n/           # Localization
├── test/
│   ├── unit/           # Unit tests
│   ├── widget/         # Widget tests
│   └── integration/    # Integration tests
├── android/            # Android-specific files
├── ios/                # iOS-specific files
├── assets/             # Images, fonts, etc.
└── pubspec.yaml        # Dart dependencies
```

## Platform-Specific Commands

### Android

```bash
make android-dev      # Run on Android (debug)
make android-build    # Build Android APK
make android-release  # Build release APK
make android-bundle   # Build App Bundle
```

### iOS

```bash
make ios-dev          # Run on iOS (debug)
make ios-build        # Build iOS app
make ios-release      # Build release IPA (requires signing)
```

### Web

```bash
make web-dev          # Run on web (debug)
make web-build        # Build web app
make web-deploy       # Deploy to Firebase Hosting
```

## Related Documentation

- [Root Makefile](../../Makefile)
- [Backend Makefile](../backend/Makefile)
- [Mobile Architecture](../../docs/02-ARCHITECTURE.md#mobile-client)
- [Specification-007: Mobile Client](../../specs/00-SPECIFICATIONS.md#spec-007-mobile-client)
