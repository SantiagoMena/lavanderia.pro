import 'package:flutter/material.dart';
import 'package:lavanderiapro/auth/login.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Lavanderia.pro',
      theme: ThemeData(
        primarySwatch: Colors.green,
      ),
      home: const LoginPage(title: 'Login'),
    );
  }
}
