import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<ClientProfile?> getClientProfile(String token) async {
  var profile = ClientProfile();

  var urlClient = Uri.http(API_HOST, 'client/profile');
  final responseClient = await http.get(
      urlClient,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );

  if (responseClient.statusCode == 200) {
    return ClientProfile.fromJson(json.decode(responseClient.body));
  } else {
    return ClientProfile();
  }
}

class ClientProfile {
  String? name;

  ClientProfile({this.name});

  factory ClientProfile.fromJson(Map<String, dynamic> json) {
    return ClientProfile(
      name: json['name'],
    );
  }
}