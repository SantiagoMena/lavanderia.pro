import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/util/constants.dart';

Future<PostOrderResponse?> postOrder(String token, OrderModel? order) async {
  var url = Uri.http(API_HOST, 'business-order/${order!.businessId}');
  var products = [];

  order.getGrouped().forEach((element) {
    products.add({
      'product': {
        'id': element.id
      },
      'amount': order.countProduct(element)
    });
  });

  final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      },
      body: json.encode({
        'products': products,
        'address': {
          'id': order.addressId
        }
      }
  ));
  if (response.statusCode == 201) {
    return PostOrderResponse.fromJson(json.decode(response.body));
  } else {
    print(response.body);
  }
  return null;
}

class PostOrderResponse {
  final String? id;

  PostOrderResponse({this.id});

  factory PostOrderResponse.fromJson(Map<String, dynamic> json) {
    return PostOrderResponse(
      id: json['id'],
    );
  }
}