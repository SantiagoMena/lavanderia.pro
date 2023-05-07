import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<List<Address>> getAddressClient(String token) async {
  var url = Uri.http(API_HOST, 'addresses');
  final response = await http.get(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );

  if (response.statusCode == 200) {
    Iterable l = json.decode(response.body);
    print(response.body);
    return List<Address>.from(l.map((model)=> Address.fromJson(model)));
  } else {
    return List<Address>.empty();
  }
}

class Address {
  String? id;
  String? client;
  String? name;
  String? address;
  String? phone;
  String? extra;
  String? createdAt;
  String? updatedAt;

  Address({
    this.id,
    this.client,
    this.name,
    this.address,
    this.phone,
    this.extra,
    this.createdAt,
    this.updatedAt,
  });

  factory Address.fromJson(Map<String, dynamic> json) {
    return Address(
      id: json['id'],
      client: json['client'],
      name: json['name'],
      address: json['address'],
      phone: json['phone'],
      extra: json['extra'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}