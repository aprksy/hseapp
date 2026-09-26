import 'package:flutter/material.dart';
import 'package:get/get.dart';
import '../theme/theme_provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class AppShellAppBar extends StatelessWidget implements PreferredSizeWidget {
  const AppShellAppBar({super.key});

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);

  @override
  Widget build(BuildContext context) {
    final themeProvider = Get.find<ThemeProvider>();
    final l10n = AppLocalizations.of(context)!;

    return AppBar(
      title: Text(l10n.appName),
      actions: [
        // Notification button
        IconButton(
          icon: const Badge(
            child: Icon(Icons.notifications_outlined),
          ),
          onPressed: () {
            // TODO: Navigate to notifications
          },
          tooltip: l10n.notifications,
        ),
        // Theme toggle
        PopupMenuButton<ThemeModeEnum>(
          icon: Icon(
            themeProvider.themeMode == ThemeModeEnum.dark
                ? Icons.dark_mode
                : themeProvider.themeMode == ThemeModeEnum.light
                    ? Icons.light_mode
                    : Icons.brightness_auto,
          ),
          onSelected: (mode) => themeProvider.setThemeMode(mode),
          itemBuilder: (context) => [
            PopupMenuItem(
              value: ThemeModeEnum.light,
              child: Row(
                children: [
                  const Icon(Icons.light_mode, size: 20),
                  const SizedBox(width: 8),
                  Text(l10n.themeLight),
                ],
              ),
            ),
            PopupMenuItem(
              value: ThemeModeEnum.dark,
              child: Row(
                children: [
                  const Icon(Icons.dark_mode, size: 20),
                  const SizedBox(width: 8),
                  Text(l10n.themeDark),
                ],
              ),
            ),
            PopupMenuItem(
              value: ThemeModeEnum.system,
              child: Row(
                children: [
                  const Icon(Icons.brightness_auto, size: 20),
                  const SizedBox(width: 8),
                  Text(l10n.themeSystem),
                ],
              ),
            ),
          ],
        ),
        // Language toggle
        PopupMenuButton<String>(
          icon: const Icon(Icons.language),
          onSelected: (locale) {
            // TODO: Implement locale change
            Get.updateLocale(Locale(locale.split('_')[0], locale.split('_')[1]));
          },
          itemBuilder: (context) => [
            const PopupMenuItem(
              value: 'id_ID',
              child: Row(
                children: [
                  Icon(Icons.flag, size: 20),
                  SizedBox(width: 8),
                  Text('Bahasa Indonesia'),
                ],
              ),
            ),
            const PopupMenuItem(
              value: 'en_ID',
              child: Row(
                children: [
                  Icon(Icons.flag, size: 20),
                  SizedBox(width: 8),
                  Text('English'),
                ],
              ),
            ),
          ],
        ),
        // User menu
        PopupMenuButton<String>(
          icon: const CircleAvatar(
            radius: 16,
            child: Icon(Icons.person, size: 20),
          ),
          onSelected: (value) {
            switch (value) {
              case 'profile':
                // TODO: Navigate to profile
                break;
              case 'settings':
                // TODO: Navigate to settings
                break;
              case 'logout':
                // TODO: Handle logout
                break;
            }
          },
          itemBuilder: (context) => [
            PopupMenuItem(
              value: 'profile',
              child: Row(
                children: [
                  const Icon(Icons.person_outline, size: 20),
                  const SizedBox(width: 8),
                  Text(l10n.profile),
                ],
              ),
            ),
            PopupMenuItem(
              value: 'settings',
              child: Row(
                children: [
                  const Icon(Icons.settings, size: 20),
                  const SizedBox(width: 8),
                  Text(l10n.preferences),
                ],
              ),
            ),
            const PopupMenuDivider(),
            PopupMenuItem(
              value: 'logout',
              child: Row(
                children: [
                  const Icon(Icons.logout, size: 20, color: Colors.red),
                  const SizedBox(width: 8),
                  Text(l10n.logout, style: const TextStyle(color: Colors.red)),
                ],
              ),
            ),
          ],
        ),
      ],
    );
  }
}
