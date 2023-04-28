import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<Business?> emailBusinessRegister(String name, String email, String password) async {
  var url = Uri.http(API_HOST, 'business/register');
  final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'name': name,
        'email': email,
        'password': password
      }
  ));
print(response.body);
  if (response.statusCode == 201) {
    // Si la llamada al servidor fue exitosa, analiza el JSON
    return Business.fromJson(json.decode(response.body));
  } else {
    // Si la llamada no fue exitosa, lanza un error.
    // throw Exception('Failed to login');
  }
}

class Business {
  final String? id;
  final String? name;
  final String? created_at;

  Business({this.id, this.name, this.created_at});

  factory Business.fromJson(Map<String, dynamic> json) {
    return Business(
      id: json['id'],
      name: json['name'],
      created_at: json['created_at'],
    );
  }
}