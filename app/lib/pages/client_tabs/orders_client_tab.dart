import 'package:flutter/material.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';
import 'package:lavanderiapro/services/get_all_orders_client_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class OrdersClientTab extends StatefulWidget {
   const OrdersClientTab({super.key, this.token});

  final String? token;

  @override
  State<OrdersClientTab> createState() => _OrdersClientTabState();
}

class _OrdersClientTabState extends State<OrdersClientTab> {
  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints viewportConstraints) {
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
                  child: FutureBuilder(
                  future: SharedPreferences.getInstance(),
                  builder: (contexSharedPreferences, snapshot) {
                    if(snapshot.hasData) {
                      String token = snapshot.data!.getString('token') ?? "";
                      return FutureBuilder(
                        future: getAllOrderClient(token),
                        builder: (contextOrders, snapshotOrders) {
                          if(snapshotOrders.hasData) {
                            return ListView.builder(
                              itemCount: snapshotOrders.data!.length,
                              /*prototypeItem: ListTile(
                                title: Text(items.first),
                              ),*/
                              itemBuilder: (context, index) {
                                var orderItem = snapshotOrders.data![index];
                                return OrderCard(orderItem: orderItem);
                              },
                            );
                          }
                          else {
                              return const Text("Not found orders");
                          }
                        }
                      );
                    } else {
                      return const CircularProgressIndicator();
                    }
                  }
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
    required this.orderItem,
  });

  final Order orderItem;

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.symmetric(horizontal: 50, vertical: 16),
        child: ElevatedButton(
          onPressed: () {
            Navigator.push(
                context,
                MaterialPageRoute(
                    builder: (context) => const ProcessedOrderClient()
                )
            );
          },
          style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 10),
            child: Column(
              children: [
                Row(
                    children: [
                      Text(orderItem.business!.name ?? "", style: const TextStyle(color: Colors.black)),
                      const Expanded(child: Text("")),
                      const Expanded(
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
                      Text("Date: ${orderItem.createdAt}", style: const TextStyle(color: Colors.black),),
                      const Expanded(child: Text("")),
                    ]
                ),
                Row(
                    children: const [
                      Text('Price: \$0000', style: TextStyle(color: Colors.black),),
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
