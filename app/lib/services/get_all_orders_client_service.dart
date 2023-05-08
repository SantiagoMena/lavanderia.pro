import 'dart:async';
import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:lavanderiapro/models/address.dart';
import 'package:lavanderiapro/models/business.dart';
import 'package:lavanderiapro/models/client.dart';
import 'package:lavanderiapro/models/delivery.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/models/product_order.dart';
import 'package:lavanderiapro/util/constants.dart';
import 'package:lavanderiapro/models/product.dart';

Future<List<Order>?> getAllOrderClient(String token) async {
  var urlClient = Uri.http(API_HOST, 'business-order/client');
  final response = await http.get(
      urlClient,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      }
  );
  if (response.statusCode == 200) {
    Iterable l = json.decode(response.body);

    return List<Order>.from(l.map((model)  {
      List<ProductOrder> products = List<ProductOrder>.empty(growable: true);
      for(final modelProduct in model['products']) {
        products.add(ProductOrder.fromJson({
          'product': Product.fromJson(modelProduct['product']),
          'amount': modelProduct['amount']
        }));
      }

      return Order.fromJson({
        'id': model['id'],
        'address': Address.fromJson({
          'id': model['address']['id'],
          'name': model['address']['name'],
          'address': model['address']['address'],
          'phone': model['address']['phone'],
          'extra': model['address']['extra'],
        }),
        'business': Business.fromJson({
          'id': model['business']['id'],
          'name': model['business']['name'],
        }),
        'client': Client.fromJson({
          'id': model['client']['id'],
          'name': model['client']['name'],
        }),
        'pickup': Delivery.fromJson({
          'id': model['pickup']['id'],
          'name': model['pickup']['name'],
        }),
        'delivery': Delivery.fromJson({
          'id': model['delivery']['id'],
          'name': model['delivery']['name'],
        }),
        'products': products,
        'created_at': model['created_at'],
        'accepted_at': model['accepted_at'],
        'rejected_at': model['rejected_at'],
        'assigned_pickup_at': model['assigned_pickup_at'],
        'pickup_client_at': model['pickup_client_at'],
        'processing_at': model['processing_at'],
        'finished_at': model['finished_at'],
        'assigned_delivery_at': model['assigned_delivery_at'],
        'pickup_business_at': model['pickup_business_at'],
        'delivered_client_at': model['delivered_client_at'],
        'deleted_at': model['deleted_at'],
      });
    }));
  } else {
    return List<Order>.empty();
  }
}
