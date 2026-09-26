import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'app_theme.dart';

enum ThemeModeEnum { light, dark, system }

class ThemeProvider extends ChangeNotifier {
  ThemeModeEnum _themeMode = ThemeModeEnum.system;
  bool _isInitialized = false;

  ThemeModeEnum get themeMode => _themeMode;
  bool get isInitialized => _isInitialized;

  ThemeData get currentTheme {
    switch (_themeMode) {
      case ThemeModeEnum.light:
        return AppTheme.lightTheme;
      case ThemeModeEnum.dark:
        return AppTheme.darkTheme;
      case ThemeModeEnum.system:
        return AppTheme.systemTheme;
    }
  }

  ThemeMode get platformThemeMode {
    switch (_themeMode) {
      case ThemeModeEnum.light:
        return ThemeMode.light;
      case ThemeModeEnum.dark:
        return ThemeMode.dark;
      case ThemeModeEnum.system:
        return ThemeMode.system;
    }
  }

  Future<void> init() async {
    if (_isInitialized) return;

    try {
      final prefs = await SharedPreferences.getInstance();
      final storedMode = prefs.getString('theme_mode');
      
      if (storedMode != null) {
        _themeMode = ThemeModeEnum.values.firstWhere(
          (e) => e.name == storedMode,
          orElse: () => ThemeModeEnum.system,
        );
      }
    } catch (e) {
      debugPrint('Error loading theme preference: $e');
    }

    _isInitialized = true;
    notifyListeners();
  }

  Future<void> setThemeMode(ThemeModeEnum mode) async {
    if (_themeMode == mode) return;

    _themeMode = mode;
    
    try {
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('theme_mode', mode.name);
    } catch (e) {
      debugPrint('Error saving theme preference: $e');
    }

    notifyListeners();
  }

  Future<void> toggleTheme() async {
    ThemeModeEnum newMode;
    switch (_themeMode) {
      case ThemeModeEnum.light:
        newMode = ThemeModeEnum.dark;
        break;
      case ThemeModeEnum.dark:
        newMode = ThemeModeEnum.light;
        break;
      case ThemeModeEnum.system:
        newMode = ThemeModeEnum.dark;
        break;
    }
    await setThemeMode(newMode);
  }
}
