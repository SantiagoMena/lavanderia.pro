import 'package:flutter/material.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';

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
                    builder: (context) => ProcessedOrderClient(orderId: orderItem.id,)
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
                      Expanded(
                          child: Align(
                              alignment: Alignment.topRight,
                              child: OrderStatus(orderItem: orderItem)
                          )
                      ),
                    ]
                ),
                Row(
                    children: [
                      Text('Products: ${orderItem.totalProducts}', style: TextStyle(color: Colors.black),),
                      Expanded(child: Text("")),
                    ]
                ),
                Row(
                    children: [
                      Text('Price: \$${orderItem.totalPrice}', style: TextStyle(color: Colors.black),),
                      Expanded(child: Text("")),
                    ]
                ),
                Row(
                    children: [
                      Text('Address: ${orderItem?.address?.address ?? ""}', style: TextStyle(color: Colors.black),),
                      Expanded(child: Text("")),
                    ]
                ),
                Row(
                    children: [
                      Text("Date: ${orderItem.createdAt}", style: const TextStyle(color: Colors.black),),
                      const Expanded(child: Text("")),
                    ]
                ),
              ],
            ),
          ),
        )
    );
  }
}

class OrderStatus extends StatelessWidget {
  const OrderStatus({
    super.key,
    required this.orderItem,
  });

  final Order orderItem;

  @override
  Widget build(BuildContext context) {


    if(orderItem.getStatus() == 'new') {

    }

    return Text(
      orderItem.getStatus(),
      style: const TextStyle(color: Colors.green),
    );
  }
}
