import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<AuthChangePasswordResponse?> changePassword(String token, String password, String newPassword) async {
  print([token, password, newPassword]);
  var url = Uri.http(API_HOST, 'auth/password/change');
  final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token'
      },
      body: json.encode({
        'password': password,
        'new_password': newPassword,
      }
  ));
  if (response.statusCode == 200) {
    // Si la llamada al servidor fue exitosa, analiza el JSON
    return AuthChangePasswordResponse.fromJson(json.decode(response.body));
  } else {
    // Si la llamada no fue exitosa, lanza un error.
    // throw Exception('Failed to login');
  }
}

class AuthChangePasswordResponse {
  final String? id;

  AuthChangePasswordResponse({this.id});

  factory AuthChangePasswordResponse.fromJson(Map<String, dynamic> json) {
    return AuthChangePasswordResponse(
      id: json['id'],
    );
  }
}