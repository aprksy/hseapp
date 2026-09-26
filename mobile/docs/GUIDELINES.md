# Mobile Client Development Guidelines

## Overview
This document provides guidelines for developing the mobile client of the HSE App using Flutter, following clean architecture principles and ensuring feature parity with the web client.

## Project Structure

```
mobile/
├── lib/
│   ├── core/
│   │   ├── constants/          # App-wide constants
│   │   ├── errors/             # Error handling classes
│   │   ├── theme/              # Theme configuration
│   │   ├── utils/              # Utility functions
│   │   └── network/            # Network layer
│   │       ├── api_client.dart    # HTTP client
│   │       ├── interceptors.dart  # Auth, logging interceptors
│   │       └── network_info.dart  # Connectivity checker
│   ├── features/               # Feature modules (DDD-style)
│   │   ├── auth/
│   │   │   ├── data/
│   │   │   │   ├── datasources/   # Remote & local data sources
│   │   │   │   ├── models/        # DTOs
│   │   │   │   └── repositories/  # Repository implementations
│   │   │   ├── domain/
│   │   │   │   ├── entities/      # Business objects
│   │   │   │   ├── repositories/  # Repository interfaces
│   │   │   │   └── usecases/      # Business logic
│   │   │   └── presentation/
│   │   │       ├── pages/         # Screens
│   │   │       ├── widgets/       # Feature widgets
│   │   │       ├── controllers/   # State management (GetX/Bloc)
│   │   │       └── bindings/      # Dependency injection
│   │   ├── dashboard/
│   │   ├── compliance/
│   │   ├── incidents/
│   │   ├── notifications/
│   │   └── profile/
│   ├── routes/                 # Navigation routing
│   │   ├── app_pages.dart
│   │   └── app_routes.dart
│   ├── l10n/                   # Localization files
│   │   ├── messages_id.arb
│   │   └── messages_en.arb
│   └── main.dart               # App entry point
├── assets/                     # Static assets
│   ├── images/
│   ├── icons/
│   ├── fonts/
│   └── animations/             # Lottie animations
├── test/                       # Test files
│   ├── unit/
│   ├── widget/
│   └── integration/
├── android/                    # Android-specific files
├── ios/                        # iOS-specific files
├── pubspec.yaml                # Dependencies
├── analysis_options.yaml       # Linter rules
└── Makefile                    # Build automation
```

## Clean Architecture Layers

### Data Layer
Handles data retrieval from remote (API) and local (SQLite, Hive) sources:

```dart
// features/auth/data/datasources/auth_remote_datasource.dart
abstract class AuthRemoteDataSource {
  Future<LoginResponseModel> login(LoginRequestModel request);
  Future<UserModel> getCurrentUser();
  Future<void> logout();
}

class AuthRemoteDataSourceImpl implements AuthRemoteDataSource {
  final ApiClient apiClient;
  
  AuthRemoteDataSourceImpl(this.apiClient);
  
  @override
  Future<LoginResponseModel> login(LoginRequestModel request) async {
    final response = await apiClient.post('/api/v1/auth/login', data: request.toJson());
    return LoginResponseModel.fromJson(response.data);
  }
}
```

### Domain Layer
Contains business logic, entities, and repository interfaces:

```dart
// features/auth/domain/entities/user.dart
class User extends Equatable {
  final String id;
  final String email;
  final List<String> roles;
  final String tenantId;
  final List<String> sites;
  
  const User({
    required this.id,
    required this.email,
    required this.roles,
    required this.tenantId,
    required this.sites,
  });
  
  bool hasRole(String role) => roles.contains(role);
  bool canAccessSite(String siteId) => sites.contains(siteId) || sites.contains('*');
  
  @override
  List<Object?> get props => [id, email, roles, tenantId, sites];
}

// features/auth/domain/repositories/auth_repository.dart
abstract class AuthRepository {
  Future<Either<AppError, LoginResponse>> login(LoginRequest request);
  Future<Either<AppError, User>> getCurrentUser();
  Future<Either<AppError, void>> logout();
}
```

### Presentation Layer
Manages UI and state using GetX or BLoC:

```dart
// features/auth/presentation/controllers/auth_controller.dart
class AuthController extends GetxController {
  final LoginUseCase loginUseCase;
  final GetCurrentUserUseCase getCurrentUserUseCase;
  
  final Rx<User?> user = Rx<User?>(null);
  final RxBool isLoading = false.obs;
  final RxString errorMessage = ''.obs;
  
  AuthController({
    required this.loginUseCase,
    required this.getCurrentUserUseCase,
  });
  
  Future<void> login(String email, String password) async {
    isLoading.value = true;
    errorMessage.value = '';
    
    final result = await loginUseCase(LoginRequest(email: email, password: password));
    
    result.fold(
      (error) => errorMessage.value = error.message,
      (response) {
        user.value = response.user;
        // Save tokens to local storage
      },
    );
    
    isLoading.value = false;
  }
}
```

## State Management

### Using GetX (Recommended)

```dart
// Simple reactive state
final count = 0.obs;

// Reactive model
class UserState extends GetxController {
  final user = Rxn<User>();
  final isLoading = false.obs;
  
  void updateUser(User newUser) => user.value = newUser;
  void clearUser() => user.value = null;
}
```

### Using BLoC (Alternative)

```dart
// Event-Driven approach
abstract class AuthEvent extends Equatable {}

class LoginRequested extends AuthEvent {
  final String email;
  final String password;
  
  LoginRequested({required this.email, required this.password});
  
  @override
  List<Object?> get props => [email, password];
}

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  // Implementation
}
```

## API Integration

### HTTP Client Setup

```dart
// core/network/api_client.dart
import 'package:dio/dio.dart';

class ApiClient {
  late final Dio _dio;
  
  ApiClient({required String baseURL, required String? token}) {
    _dio = Dio(BaseOptions(
      baseUrl: baseURL,
      connectTimeout: const Duration(seconds: 30),
      receiveTimeout: const Duration(seconds: 30),
      headers: {
        'Content-Type': 'application/json',
        if (token != null) 'Authorization': 'Bearer $token',
      },
    ));
    
    _dio.interceptors.addAll([
      LogInterceptor(requestBody: true, responseBody: true),
      AuthInterceptor(),
    ]);
  }
  
  Future<Response<T>> get<T>(String path, {Map<String, dynamic>? queryParameters}) {
    return _dio.get<T>(path, queryParameters: queryParameters);
  }
  
  Future<Response<T>> post<T>(String path, {dynamic data}) {
    return _dio.post<T>(path, data: data);
  }
  
  Future<Response<T>> put<T>(String path, {dynamic data}) {
    return _dio.put<T>(path, data: data);
  }
  
  Future<Response<T>> delete<T>(String path) {
    return _dio.delete<T>(path);
  }
}
```

## Localization (i18n)

Using Flutter's built-in localization:

```yaml
# pubspec.yaml
flutter:
  generate: true
```

```dart
// l10n/messages_id.arb
{
  "@@locale": "id",
  "welcomeMessage": "Selamat Datang di HSE App",
  "@welcomeMessage": {
    "description": "Welcome message on home screen"
  },
  "loginTitle": "Masuk",
  "incidentReported": "Insiden berhasil dilaporkan"
}
```

Usage:
```dart
Text(AppLocalizations.of(context)!.welcomeMessage)
```

## Theming

### Light/Dark/System Theme

```dart
// core/theme/app_theme.dart
class AppTheme {
  static ThemeData lightTheme = ThemeData(
    brightness: Brightness.light,
    colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
    useMaterial3: true,
  );
  
  static ThemeData darkTheme = ThemeData(
    brightness: Brightness.dark,
    colorScheme: ColorScheme.fromSeed(
      seedColor: Colors.blue,
      brightness: Brightness.dark,
    ),
    useMaterial3: true,
  );
}

// main.dart
GetMaterialApp(
  theme: AppTheme.lightTheme,
  darkTheme: AppTheme.darkTheme,
  themeMode: ThemeMode.system, // or .light / .dark
);
```

## Forms & Validation

Using `form_builder_validators`:

```dart
import 'package:flutter_form_builder/flutter_form_builder.dart';
import 'package:form_builder_validators/form_builder_validators.dart';

FormBuilder(
  child: Column(
    children: [
      FormBuilderTextField(
        name: 'title',
        decoration: const InputDecoration(labelText: 'Judul Insiden'),
        validator: FormBuilderValidators.compose([
          FormBuilderValidators.required(),
          FormBuilderValidators.minLength(5),
          FormBuilderValidators.maxLength(100),
        ]),
      ),
      FormBuilderDropdown<String>(
        name: 'severity',
        decoration: const InputDecoration(labelText: 'Tingkat Keparahan'),
        items: ['LOW', 'MEDIUM', 'HIGH']
            .map((e) => DropdownMenuItem(value: e, child: Text(e)))
            .toList(),
        validator: FormBuilderValidators.required(),
      ),
    ],
  ),
)
```

## Offline Support

### Local Database (Hive)

```dart
// features/incidents/data/datasources/incident_local_datasource.dart
import 'package:hive/hive.dart';

class IncidentLocalDataSourceImpl implements IncidentLocalDataSource {
  final Box<IncidentModel> incidentBox;
  
  IncidentLocalDataSourceImpl(this.incidentBox);
  
  @override
  Future<List<IncidentModel>> getAllIncidents() async {
    return incidentBox.values.toList();
  }
  
  @override
  Future<void> saveIncident(IncidentModel incident) async {
    await incidentBox.put(incident.id, incident);
  }
  
  @override
  Future<void> syncWithServer(List<IncidentModel> serverData) async {
    await incidentBox.clear();
    for (var incident in serverData) {
      await incidentBox.put(incident.id, incident);
    }
  }
}
```

### Connectivity Check

```dart
// core/network/network_info.dart
import 'package:connectivity_plus/connectivity_plus.dart';

abstract class NetworkInfo {
  Future<bool> get isConnected;
}

class NetworkInfoImpl implements NetworkInfo {
  final Connectivity connectivity;
  
  NetworkInfoImpl(this.connectivity);
  
  @override
  Future<bool> get isConnected async {
    final results = await connectivity.checkConnectivity();
    return !results.contains(ConnectivityResult.none);
  }
}
```

## Camera & File Capture

For mobile-specific features:

```dart
// features/incidents/presentation/widgets/camera_capture.dart
import 'package:image_picker/image_picker.dart';

class CameraCaptureWidget extends StatelessWidget {
  final Function(File image) onImageCaptured;
  
  const CameraCaptureWidget({required this.onImageCaptured});
  
  Future<void> _pickImage(BuildContext context) async {
    final picker = ImagePicker();
    final image = await picker.pickImage(source: ImageSource.camera);
    
    if (image != null) {
      onImageCaptured(File(image.path));
    }
  }
  
  @override
  Widget build(BuildContext context) {
    return IconButton(
      icon: const Icon(Icons.camera_alt),
      onPressed: () => _pickImage(context),
    );
  }
}
```

## Push Notifications

Using Firebase Cloud Messaging (FCM):

```dart
// core/services/notification_service.dart
import 'package:firebase_messaging/firebase_messaging.dart';

class NotificationService {
  final FirebaseMessaging _messaging = FirebaseMessaging.instance;
  
  Future<void> initialize() async {
    await _messaging.requestPermission();
    
    final token = await _messaging.getToken();
    print('FCM Token: $token');
    
    FirebaseMessaging.onMessage.listen((RemoteMessage message) {
      print('Received message: ${message.notification?.title}');
      // Handle foreground notification
    });
    
    FirebaseMessaging.onBackgroundMessage(_handleBackgroundMessage);
  }
  
  static Future<void> _handleBackgroundMessage(RemoteMessage message) async {
    print('Handling background message: ${message.messageId}');
  }
}
```

## Testing Strategy

### Unit Tests
```bash
make test-unit
```
Test business logic and use cases.

### Widget Tests
```bash
make test-widget
```
Test individual widgets in isolation.

### Integration Tests
```bash
make test-integration
```
Test complete user flows with mocked APIs.

## Development Workflow

1. Create feature directory structure
2. Define domain entities and use cases
3. Implement data layer (models, datasources, repositories)
4. Build presentation layer (pages, widgets, controllers)
5. Add navigation routes
6. Write tests
7. Update localization files

## Adding New Features

1. Update specification in `/specs`
2. Create domain entities
3. Implement data layer
4. Build UI screens and widgets
5. Add navigation
6. Write tests
7. Update i18n translations (ID & EN)

## Performance Optimization

- Use `const` constructors
- Lazy load images with `cached_network_image`
- Minimize rebuilds with `GetBuilder` or `BlocBuilder`
- Use pagination for large lists
- Profile with Flutter DevTools

## Accessibility

- Semantic labels for all interactive elements
- Sufficient color contrast
- Support for screen readers
- Scalable text sizes

## References

- [Flutter Documentation](https://docs.flutter.dev/)
- [Clean Architecture in Flutter](https://resocoder.com/flutter-clean-architecture-tdd/)
- [GetX Documentation](https://pub.dev/packages/get)
- [Flutter Internationalization](https://docs.flutter.dev/ui/accessibility-and-internationalization/internationalization)
