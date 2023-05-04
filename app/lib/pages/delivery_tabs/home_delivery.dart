import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/pages/client_tabs/new_order_client_tab.dart';
import 'package:lavanderiapro/pages/client_tabs/orders_client_tab.dart';
import 'package:lavanderiapro/pages/client_tabs/profile_client_tab.dart';
import 'package:lavanderiapro/pages/delivery_tabs/orders_delivery_tab.dart';
import 'package:lavanderiapro/pages/delivery_tabs/orders_history_delivery_tab.dart';
import 'package:lavanderiapro/pages/delivery_tabs/profile_delivery_tab.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class HomeDeliveryTab extends StatelessWidget {
  const HomeDeliveryTab({super.key, this.token});

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
                  child: ProfileTabDeliveryLabel()
              ),
              Padding(
                  padding: const EdgeInsets.symmetric(vertical: 4),
                  child: OrdersDeliveryTabLabel()
              ),
              Padding(
                  padding: const EdgeInsets.symmetric(vertical: 4),
                  child: OrdersTabDeliveryLabel()
              ),
            ],
          ),
          // title: const Text('Tabs Demo'),
        ),
        body: const TabBarView(
          children: [
            ProfileDeliveryTab(),
            OrdersDeliveryTab(),
            OrdersHistoryDeliveryTab(),
          ],
        ),
      ),
    );
  }
}

class OrdersTabDeliveryLabel extends StatelessWidget {
  const OrdersTabDeliveryLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Tab(
        text: "History",
        icon: Icon(Icons.assignment_turned_in_outlined)
    );
  }
}

class OrdersDeliveryTabLabel extends StatelessWidget {
  const OrdersDeliveryTabLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Tab(
        text: "Orders",
        icon: Icon(Icons.delivery_dining)
    );
  }
}

class ProfileTabDeliveryLabel extends StatelessWidget {
  const ProfileTabDeliveryLabel({
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
