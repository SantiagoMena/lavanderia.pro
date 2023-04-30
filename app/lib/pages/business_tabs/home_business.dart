import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/pages/business_tabs/profile_business_tab.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';

class HomeBusinessTab extends StatelessWidget {
  const HomeBusinessTab({super.key, this.token});

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
                  child: ProfileBusinessTabLabel()
              ),
              Padding(
                  padding: const EdgeInsets.symmetric(vertical: 4),
                  child: HomeBusinessTabLabel()
              ),
              Padding(
                  padding: const EdgeInsets.symmetric(vertical: 4),
                  child: OrdersBusinessLabel()
              ),
            ],
          ),
          // title: const Text('Tabs Demo'),
        ),
        body: const TabBarView(
          children: [
            ProfileBusinessTab(),
            Text("NewOrderClientTab()"),
            Text("OrdersClientTab()"),
          ],
        ),
      ),
    );
  }
}

class OrdersBusinessLabel extends StatelessWidget {
  const OrdersBusinessLabel({
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

class HomeBusinessTabLabel extends StatelessWidget {
  const HomeBusinessTabLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Tab(
        text: AppLocalizations.of(context)!.businessTabLabel,
        icon: Icon(Icons.home)
    );
  }
}

class ProfileBusinessTabLabel extends StatelessWidget {
  const ProfileBusinessTabLabel({
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
