import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<DeliveryProfile?> putDeliveryProfile(String token, String name) async {
  var url = Uri.http(API_HOST, 'delivery/profile');
  final response = await http.put(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      },
      body: json.encode({
        'name': name,
      }
  ));
  if (response.statusCode == 200) {
    // Si la llamada al servidor fue exitosa, analiza el JSON
    return DeliveryProfile.fromJson(json.decode(response.body));
  } else {
    // Si la llamada no fue exitosa, lanza un error.
    // throw Exception('Failed to login');
  }
  return null;
}

class DeliveryProfile {
  final String? id;
  final String? name;
  final String? createdAt;
  final String? updatedAt;

  DeliveryProfile({this.id, this.name, this.createdAt, this.updatedAt});

  factory DeliveryProfile.fromJson(Map<String, dynamic> json) {
    return DeliveryProfile(
      id: json['id'],
      name: json['name'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}