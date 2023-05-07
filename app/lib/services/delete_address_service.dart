import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/util/constants.dart';

import '../models/address.dart';

Future<Address?> deleteAddress(String token, Address? address) async {
  var url = Uri.http(API_HOST, 'address/${address!.id}');
  final response = await http.delete(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      }
  );

  if (response.statusCode == 200) {
    print(response.body);
    return Address.fromJson(json.decode(response.body));
  } else {
    print(response.body);
  }
  return null;
}
