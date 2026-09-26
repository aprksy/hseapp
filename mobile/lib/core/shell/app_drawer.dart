import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class AppShellDrawer extends StatelessWidget {
  const AppShellDrawer({super.key});

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: [
          DrawerHeader(
            decoration: BoxDecoration(
              color: Theme.of(context).colorScheme.primaryContainer,
            ),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                const CircleAvatar(
                  radius: 32,
                  child: Icon(Icons.person, size: 32),
                ),
                const SizedBox(height: 8),
                Text(
                  'User Name', // TODO: Replace with actual user name
                  style: Theme.of(context).textTheme.titleMedium?.copyWith(
                    color: Theme.of(context).colorScheme.onPrimaryContainer,
                  ),
                ),
                Text(
                  'user@example.com', // TODO: Replace with actual email
                  style: Theme.of(context).textTheme.bodySmall?.copyWith(
                    color: Theme.of(context).colorScheme.onPrimaryContainer,
                  ),
                ),
              ],
            ),
          ),
          ListTile(
            leading: const Icon(Icons.dashboard_outlined),
            title: Text(l10n.dashboard),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to dashboard
            },
          ),
          ListTile(
            leading: const Icon(Icons.checklist_outlined),
            title: Text(l10n.compliance),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to compliance
            },
          ),
          ListTile(
            leading: const Icon(Icons.warning_amber_outlined),
            title: Text(l10n.incidents),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to incidents
            },
          ),
          ListTile(
            leading: const Icon(Icons.calendar_today_outlined),
            title: Text(l10n.planning),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to planning
            },
          ),
          ListTile(
            leading: const Icon(Icons.monitor_heart_outlined),
            title: Text(l10n.monitoring),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to monitoring
            },
          ),
          ListTile(
            leading: const Icon(Icons.analytics_outlined),
            title: Text(l10n.evaluation),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to evaluation
            },
          ),
          ListTile(
            leading: const Icon(Icons.description_outlined),
            title: Text(l10n.reports),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to reports
            },
          ),
          ListTile(
            leading: const Icon(Icons.project_outlined),
            title: Text(l10n.projects),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to projects
            },
          ),
          const Divider(),
          ListTile(
            leading: const Icon(Icons.settings_outlined),
            title: Text(l10n.settings),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to settings
            },
          ),
          ListTile(
            leading: const Icon(Icons.help_outline),
            title: Text(l10n.help),
            onTap: () {
              Navigator.pop(context);
              // TODO: Navigate to help
            },
          ),
          const Divider(),
          ListTile(
            leading: const Icon(Icons.logout, color: Colors.red),
            title: Text(l10n.logout, style: const TextStyle(color: Colors.red)),
            onTap: () {
              Navigator.pop(context);
              // TODO: Handle logout
            },
          ),
          const Spacer(),
          Padding(
            padding: const EdgeInsets.all(16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  l10n.copyright,
                  style: Theme.of(context).textTheme.bodySmall,
                ),
                const SizedBox(height: 4),
                Text(
                  l10n.compliance,
                  style: Theme.of(context).textTheme.bodySmall?.copyWith(
                    color: Theme.of(context).colorScheme.primary,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
