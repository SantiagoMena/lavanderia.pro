import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:lavanderiapro/util/constants.dart';

Future<BusinessProfile?> getBusinessProfile(String token) async {
  var urlBusiness = Uri.http(API_HOST, 'business');
  final responseBusiness = await http.get(
      urlBusiness,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${token}'
      }
  );

  if (responseBusiness.statusCode == 200) {
    return BusinessProfile.fromJson(json.decode(responseBusiness.body));
  } else {
    return BusinessProfile();
  }
}

class BusinessProfile {
  String? name;

  BusinessProfile({this.name});

  factory BusinessProfile.fromJson(Map<String, dynamic> json) {
    return BusinessProfile(
      name: json['name'],
    );
  }
}