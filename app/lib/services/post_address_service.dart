import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/util/constants.dart';

import '../models/address.dart';

Future<Address?> postAddress(String token, Address? address) async {
  var url = Uri.http(API_HOST, 'address');
  print(['address!.name', address!.name]);
  print(['address!.address', address!.address]);
  print(['address!.phone', address!.phone]);
  print(['address!.extra', address!.extra]);
  final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      },
      body: json.encode({
        'name': address!.name,
        'address': address!.address,
        'phone': address!.phone,
        'extra': address!.extra,
      }
  ));
  if (response.statusCode == 201) {
    return Address.fromJson(json.decode(response.body));
  } else {
    print(response.body);
  }
  return null;
}
