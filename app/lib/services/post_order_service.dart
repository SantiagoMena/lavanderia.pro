import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/util/constants.dart';

Future<ClientProfile?> postOrder(String token, OrderModel? order) async {
  print(['business-order/${order!.businessId}', token]);
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
  if (response.statusCode == 200) {
    print(response.body);
    // Si la llamada al servidor fue exitosa, analiza el JSON
    return ClientProfile.fromJson(json.decode(response.body));
  } else {
    print(response.body);
  }
  return null;
}

class ClientProfile {
  final String? id;
  final String? name;
  final String? createdAt;
  final String? updatedAt;

  ClientProfile({this.id, this.name, this.createdAt, this.updatedAt});

  factory ClientProfile.fromJson(Map<String, dynamic> json) {
    return ClientProfile(
      id: json['id'],
      name: json['name'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}