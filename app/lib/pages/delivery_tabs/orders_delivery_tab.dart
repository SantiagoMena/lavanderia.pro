import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/client_tabs/business_client_view.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';

class OrdersDeliveryTab extends StatefulWidget {
   const OrdersDeliveryTab({super.key, this.token});

  final String? token;

  @override
  State<OrdersDeliveryTab> createState() => _OrdersDeliveryTabState();
}

class _OrdersDeliveryTabState extends State<OrdersDeliveryTab> {
  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
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
                    child: Text("Select a order")
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
                    var orderItem = items[index];
                    return OrderCard(items: items, ordertIndex: index);
                  },
                ),
              ),
              // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
            ],
           ),
        );
      }
    );
  }
}

class OrderCard extends StatelessWidget {
  const OrderCard({
    super.key,
    required this.items,
    required this.ordertIndex,
  });

  final int ordertIndex;
  final List<String> items;

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: EdgeInsets.symmetric(horizontal: 50, vertical: 16),
        child: ElevatedButton(
          onPressed: () {
            Navigator.push(
                context,
                MaterialPageRoute(
                    builder: (context) => ProcessedOrderClient()
                )
            );
          },
          style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
          child: Padding(
            padding: EdgeInsets.symmetric(horizontal: 10, vertical: 10),
            child: Column(
              children: [
                Row(
                    children: [
                      Container(child:
                        Text(items[ordertIndex], style: TextStyle(color: Colors.black)),
                      ),
                      Expanded(child: Text("")),
                      Expanded(
                          child: Align(
                              alignment: Alignment.topRight,
                              child: Text(
                                  "Active",
                                  style: TextStyle(color: Colors.green),
                              )
                          )
                      ),
                    ]
                ),
                Row(
                    children: [
                      Container(child: Text("Date: 30/04/2023", style: TextStyle(color: Colors.black),)),
                      Expanded(child: Text("")),
                    ]
                ),
                Row(
                    children: [
                      Container(child: Text('Price: \$0000', style: TextStyle(color: Colors.black),)),
                      Expanded(child: Text("")),
                    ]
                ),
              ],
            ),
          ),
        )
    );
  }
}
