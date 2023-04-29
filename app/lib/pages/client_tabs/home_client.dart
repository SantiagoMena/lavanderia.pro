import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class HomeClientTab extends StatelessWidget {
  const HomeClientTab({super.key, this.token});

  final String? token;

  @override
  Widget build(BuildContext context) {

    Profile? profile;

    return MaterialApp(
      localizationsDelegates: AppLocalizations.localizationsDelegates,
      supportedLocales: AppLocalizations.supportedLocales,
      home: DefaultTabController(
        length: 3,
        child: Scaffold(
          bottomNavigationBar: Container(
            color: Colors.green,
            child: const TabBar(
              indicatorColor: Colors.white,
              tabs: [
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: ProfileTabClient()
                  ),
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: NewOrderTabClient()
                  ),
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: OrdersTabClient()
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

class OrdersTabClient extends StatelessWidget {
  const OrdersTabClient({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Tab(
      text: AppLocalizations.of(context)!.ordersTabClientLabel,
      icon: Icon(Icons.assignment_turned_in_outlined)
    );
  }
}

class NewOrderTabClient extends StatelessWidget {
  const NewOrderTabClient({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Tab(
        text: AppLocalizations.of(context)!.newOrderTabClientLabel,
        icon: Icon(Icons.add)
    );
  }
}

class ProfileTabClient extends StatelessWidget {
  const ProfileTabClient({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Tab(
        text: AppLocalizations.of(context)!.profileTabClientLabel,
        icon: Icon(Icons.person)
    );
  }
}
