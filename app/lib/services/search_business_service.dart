import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<List<Business>?> searchBusiness(String token) async {
  var urlClient = Uri.http(API_HOST, 'business/search');
  final response = await http.get(
      urlClient,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );

  if (response.statusCode == 200) {
    Iterable l = json.decode(response.body);

    return List<Business>.from(l.map((model)=> Business.fromJson(model)));
  } else {
    return List<Business>.empty();
  }
}

class Business {
  final String? id;
  final String? name;
  final String? createdAt;
  final String? updatedAt;

  Business({this.id, this.name, this.createdAt, this.updatedAt});

  factory Business.fromJson(Map<String, dynamic> json) {
    return Business(
      id: json['id'],
      name: json['name'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}