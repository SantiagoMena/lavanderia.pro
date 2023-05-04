import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<DeliveryProfile?> getDeliveryProfile(String token) async {
  var urlClient = Uri.http(API_HOST, 'delivery/profile');
  final responseClient = await http.get(
      urlClient,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );

  if (responseClient.statusCode == 200) {
    return DeliveryProfile.fromJson(json.decode(responseClient.body));
  } else {
    return DeliveryProfile();
  }
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