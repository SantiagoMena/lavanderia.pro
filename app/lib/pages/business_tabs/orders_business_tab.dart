import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/business_tabs/business_create_form.dart';
import 'package:lavanderiapro/pages/business_tabs/business_view.dart';
import 'package:lavanderiapro/pages/business_tabs/products_business_view.dart';
import 'package:lavanderiapro/pages/client_tabs/business_client_view.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';

class OrdersBusinessTab extends StatefulWidget {
   const OrdersBusinessTab({super.key, this.token});

  final String? token;

  @override
  State<OrdersBusinessTab> createState() => _OrdersBusinessTabState();
}

class _OrdersBusinessTabState extends State<OrdersBusinessTab> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints viewportConstraints) {
          var items = List<String>.generate(15, (i) => 'Order $i');
          return Align(
          alignment: Alignment.topCenter,
          child: Column(
            children:[
                const SizedBox(
                  height: 50,
                  child: Align(
                    alignment: Alignment.center,
                    child: Padding(
                      padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                      child: Text("Manage Orders")
                    ),
                  ),
                ),
                Expanded(
                  child: ListView.builder(
                    itemCount: items.length,
                    /*prototypeItem: ListTile(
                      title: Text(items.first),
                    ),*/
                    itemBuilder: (context, index) {
                      var businessItem = items[index];
                      return BusinessCard(businessItem: businessItem);
                    },
                  ),
                ), // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
              ],
             ),
          );
        }
      ),
    );
  }
}

class BusinessCard extends StatelessWidget {
  const BusinessCard({
    super.key,
    required this.businessItem,
  });

  final String businessItem;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
        onPressed: () {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) => ProcessedOrderClient()
              )
          );
        },
        child: Padding(
            child: Text(businessItem, style: TextStyle(color: Colors.black),),
            padding: EdgeInsets.symmetric(horizontal: 8, vertical: 50),
        )
      ),
    );
  }
}
