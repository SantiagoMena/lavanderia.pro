import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';

class HomeClientTab extends StatelessWidget {
  const HomeClientTab({super.key, this.token});

  final String? token;

  @override
  Widget build(BuildContext context) {

    Profile? profile;

    return MaterialApp(
      home: DefaultTabController(
        length: 3,
        child: Scaffold(
          bottomNavigationBar: Container(
            color: Colors.green,
            child: const TabBar(
              tabs: [
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: Tab(icon: Icon(Icons.person))
                  ),
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: Tab(icon: Icon(Icons.add))
                  ),
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: Tab(icon: Icon(Icons.assignment_turned_in_outlined))
                  ),
                ],
            ),
            // title: const Text('Tabs Demo'),
          ),
          body: const TabBarView(
            children: [
              Icon(Icons.person),
              Icon(Icons.add),
              Icon(Icons.assignment_turned_in_outlined),
            ],
          ),
        ),
      ),
    );
  }
}
