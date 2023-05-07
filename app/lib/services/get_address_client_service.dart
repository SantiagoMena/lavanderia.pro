import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/models/address.dart';
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