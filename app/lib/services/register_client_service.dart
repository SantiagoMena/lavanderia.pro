import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<Client?> emailClientRegister(String name, String email, String password) async {
  var url = Uri.http(API_HOST, 'client/register');
  final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'name': name,
        'email': email,
        'password': password
      }
  ));
  if (response.statusCode == 201) {
    // Si la llamada al servidor fue exitosa, analiza el JSON
    return Client.fromJson(json.decode(response.body));
  } else {
    // Si la llamada no fue exitosa, lanza un error.
    // throw Exception('Failed to login');
  }
}

class Client {
  final String? id;
  final String? name;
  final String? created_at;

  Client({this.id, this.name, this.created_at});

  factory Client.fromJson(Map<String, dynamic> json) {
    return Client(
      id: json['id'],
      name: json['name'],
      created_at: json['created_at'],
    );
  }
}