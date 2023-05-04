import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<Profile?> getProfile(String token) async {
  var profile = Profile();

  var urlDelivery = Uri.http(API_HOST, 'delivery/profile');
  final responseDelivery = await http.get(
      urlDelivery,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );
  profile.delivery = responseDelivery.statusCode == 200;

  var urlClient = Uri.http(API_HOST, 'client/profile');
  final responseClient = await http.get(
      urlClient,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );
  profile.client = responseClient.statusCode == 200;

  var urlBusiness = Uri.http(API_HOST, 'business');
  final responseBusiness = await http.get(
      urlBusiness,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );

  profile.business = responseBusiness.statusCode == 200;

  return profile;
}

class Profile {
  bool? business;
  bool? delivery;
  bool? client;

  Profile({this.business, this.delivery, this.client});

  factory Profile.fromJson(Map<String, dynamic> json) {
    return Profile(
      business: json['business'],
      delivery: json['delivery'],
      client: json['client'],
    );
  }
}