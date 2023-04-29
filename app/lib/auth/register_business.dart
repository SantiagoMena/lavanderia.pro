import 'package:flutter/material.dart';
import 'package:lavanderiapro/pages/home.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/services/login_service.dart';
import 'package:lavanderiapro/services/register_business_service.dart';


class RegisterBusinessPage extends StatefulWidget {
  const RegisterBusinessPage({super.key, required this.title});

  final String title;

  @override
  State<RegisterBusinessPage> createState() => _RegisterBusinessPageState();
}

class _RegisterBusinessPageState extends State<RegisterBusinessPage> {
  final _formKey = GlobalKey<FormState>();
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  TextEditingController nameController = TextEditingController();


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(AppLocalizations.of(context)!.registerBusinessPageTitle),
      ),
      body: Form(
          key: _formKey,
          child: SingleChildScrollView(
            reverse: true,
            child: Column(
              children: [
                Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                    child: Center(
                      child: Text('🗝️', style: TextStyle(fontSize: 75)),
                    )
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: TextFormField(
                    controller: nameController,
                    decoration: const InputDecoration(border: OutlineInputBorder(), label: NameLabel()),
                    validator: (value) {
                      if(value == null || value.isEmpty){
                        return AppLocalizations.of(context)!.emptyNameAlert;
                      }
                      return null;
                    },
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: TextFormField(
                    controller: emailController,
                    decoration: const InputDecoration(border: OutlineInputBorder(), label: EmailLabel()),
                    validator: (value) {
                      if(value == null || value.isEmpty){
                        return AppLocalizations.of(context)!.emptyEmailAlert;
                      }
                      return null;
                    },
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: TextFormField(
                    controller: passwordController,
                    obscureText: true,
                    decoration: const InputDecoration(border: OutlineInputBorder(), label: PasswordLabel()),
                    validator: (value) {
                      if(value == null || value.isEmpty){
                        return AppLocalizations.of(context)!.emptyPasswordAlert;
                      }
                      return null;
                    },
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: Center(
                    child: ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        minimumSize: const Size.fromHeight(50), // NEW
                      ),
                      onPressed: () {
                        if(_formKey.currentState!.validate()){
                          // Login User
                          emailBusinessRegister(nameController.text, emailController.text, passwordController.text)
                            .then(
                              (business) => business!.created_at!.length > 0 ?
                              emailLogin(emailController.text, passwordController.text).then(
                                (token) =>
                                  token!.token!.length > 0 ?
                                      Navigator.push(
                                          context,
                                          MaterialPageRoute(
                                              builder: (context) => HomePage(token: token!.token ?? '')
                                          )
                                      ) :
                                  ScaffoldMessenger.of(context).showSnackBar(
                                    const SnackBar(content: SnackBarRegisterError())
                                  )
                              ) :
                              ScaffoldMessenger.of(context).showSnackBar(
                                const SnackBar(content: SnackBarRegisterError())
                              )
                            ).catchError((e) => ScaffoldMessenger.of(context).showSnackBar(
                              const SnackBar(content: SnackBarRegisterError())
                            )).catchError((e) => ScaffoldMessenger.of(context).showSnackBar(
                              const SnackBar(content: SnackBarRegisterError())
                          ));
                        } else {
                          ScaffoldMessenger.of(context).showSnackBar(
                              const SnackBar(content: SnackBarFillInput())
                          );
                        }
                      },
                      child: SubmitLabel(),
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: Center(
                    child: ElevatedButton(
                      style: ElevatedButton.styleFrom(
                        minimumSize: const Size.fromHeight(50), // NEW
                      ),
                      onPressed: () {
                        Navigator.pop(context);
                      },
                      child: Text(AppLocalizations.of(context)!.loginLabel, style: TextStyle(fontSize: 18)),
                    ),
                  ),
                ),
                Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
              ],
            ),
          ),
      ),
    );
  }
}

class NameLabel extends StatelessWidget {
  const NameLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.businessNameLabel);
  }
}

class SnackBarRegisterError extends StatelessWidget {
  const SnackBarRegisterError({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarRegisterError);
  }
}

class SnackBarFillInput extends StatelessWidget {
  const SnackBarFillInput({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarFillInput);
  }
}

class SnackBarInvalidCredentials extends StatelessWidget {
  const SnackBarInvalidCredentials({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarInvalidCredentials);
  }
}

class SubmitLabel extends StatelessWidget {
  const SubmitLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.submitRegisterLabel, style: TextStyle(fontSize: 18));
  }
}

class EmailLabel extends StatelessWidget {
  const EmailLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.emailLabel);
  }
}

class PasswordLabel extends StatelessWidget {
  const PasswordLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.passwordLabel);
  }
}
