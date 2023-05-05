import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<List<Product>?> getAllProductsBusiness(String token, String businessId) async {
  var urlClient = Uri.http(API_HOST, 'business/$businessId/products');
  final response = await http.get(
      urlClient,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      }
  );

  if (response.statusCode == 200) {
    Iterable l = json.decode(response.body);
    print(response.body);
    return List<Product>.from(l.map((model)=> Product.fromJson(model)));
  } else {
    return List<Product>.empty();
  }
}

class Product {
  final String? id;
  final String? name;
  final num? price;
  final String? createdAt;
  final String? updatedAt;

  Product({this.id, this.name, this.price, this.createdAt, this.updatedAt});

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      id: json['id'],
      name: json['name'],
      price: json['price'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}