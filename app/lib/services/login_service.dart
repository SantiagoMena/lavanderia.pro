import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<Token?> emailLogin(String email, String password) async {
  var url = Uri.http(API_HOST, 'auth/login');
  final response = await http.post(
      url,
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'email': email,
        'password': password
      }
  ));
print(response.body);
  if (response.statusCode == 200) {
    // Si la llamada al servidor fue exitosa, analiza el JSON
    return Token.fromJson(json.decode(response.body));
  } else {
    // Si la llamada no fue exitosa, lanza un error.
    // throw Exception('Failed to login');
  }
}

class Token {
  final String? token;
  final String? refresh_token;

  Token({this.token, this.refresh_token});

  factory Token.fromJson(Map<String, dynamic> json) {
    return Token(
      token: json['token'],
      refresh_token: json['refresh_token'],
    );
  }
}