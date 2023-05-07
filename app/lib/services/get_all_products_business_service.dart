import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';
import 'package:lavanderiapro/models/product.dart';

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
