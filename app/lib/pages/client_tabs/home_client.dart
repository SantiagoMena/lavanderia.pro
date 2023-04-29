import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';

class HomeClientTab extends StatelessWidget {
  const HomeClientTab({super.key, this.token});

  final String? token;

  @override
  Widget build(BuildContext context) {


    Profile? profile;
    List<Widget> children = [
      Padding(
          padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
          child: Center(
            child: Text('👕', style: TextStyle(fontSize: 75)),
          )
      )
    ];

      //children.add(const LoginAsBusinessButton());


    return Scaffold(
      /*appBar: AppBar(
        title: const Text('Client Home'),
      ),*/
      body: ListView(
        children:children,
      )
    );
  }
}
