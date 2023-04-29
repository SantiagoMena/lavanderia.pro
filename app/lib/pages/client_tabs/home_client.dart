import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/pages/client_tabs/profile_client_tab.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class HomeClientTab extends StatelessWidget {
  const HomeClientTab({super.key, this.token});

  final String? token;

  @override
  Widget build(BuildContext context) {

    Profile? profile;

    return  DefaultTabController(
        length: 3,
        child: Scaffold(
          resizeToAvoidBottomInset: false,
          appBar:  AppBar(
            backgroundColor: Colors.white,
          ),
          bottomNavigationBar: Container(
            color: Colors.green,
            child: const TabBar(
              indicatorColor: Colors.white,
              tabs: [
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: ProfileTabClientLabel()
                  ),
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: NewOrderTabClientLabel()
                  ),
                  Padding(
                      padding: const EdgeInsets.symmetric(vertical: 4),
                      child: OrdersTabClientLabel()
                  ),
                ],
            ),
            // title: const Text('Tabs Demo'),
          ),
          body: const TabBarView(
            children: [
              ProfileClientTab(),
              Icon(Icons.add),
              Icon(Icons.assignment_turned_in_outlined),
            ],
          ),
        ),
    );
  }
}

class OrdersTabClientLabel extends StatelessWidget {
  const OrdersTabClientLabel({
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

class NewOrderTabClientLabel extends StatelessWidget {
  const NewOrderTabClientLabel({
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

class ProfileTabClientLabel extends StatelessWidget {
  const ProfileTabClientLabel({
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
