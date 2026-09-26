import 'package:flutter/material.dart';
import 'package:get/get.dart';
import '../theme/theme_provider.dart';
import 'app_drawer.dart';
import 'app_bar.dart';

class AppShell extends StatelessWidget {
  final Widget body;
  final int currentIndex;
  final Function(int) onTabChanged;
  final List<NavigationDestination> destinations;

  const AppShell({
    super.key,
    required this.body,
    required this.currentIndex,
    required this.onTabChanged,
    required this.destinations,
  });

  @override
  Widget build(BuildContext context) {
    final themeProvider = Get.find<ThemeProvider>();

    return Scaffold(
      appBar: const AppShellAppBar(),
      drawer: const AppShellDrawer(),
      body: body,
      bottomNavigationBar: NavigationBar(
        selectedIndex: currentIndex,
        onDestinationSelected: onTabChanged,
        destinations: destinations,
      ),
    );
  }
}
