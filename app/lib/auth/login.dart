import 'package:flutter/material.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/auth/register_client.dart';
import 'package:lavanderiapro/pages/home.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/services/login_service.dart';
import 'package:shared_preferences/shared_preferences.dart';


class LoginPage extends StatefulWidget {
  const LoginPage({super.key, required this.title});

  final String title;

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _formKey = GlobalKey<FormState>();
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();


  @override
  Widget build(BuildContext context) {
    var icons = ['🧺','🧼','👕','👚','👔','👖','🧦','👙','👗'];
    return Scaffold(
      appBar: AppBar(
        title: Text(AppLocalizations.of(context)!.loginLabel),
      ),
      body: Form(
        key: _formKey,
        child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
        child:
          SingleChildScrollView(
          reverse: true,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: Center(
                    child: Text((icons..shuffle()).first, style: TextStyle(fontSize: 75)),
                  )
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
                        emailLogin(emailController.text, passwordController.text)
                            .then((token) {
                          if(token!.token!.length > 0){
                            try {
                              Future<SharedPreferences> _sprefs = SharedPreferences.getInstance();
                              _sprefs.then((prefs) {
                                prefs.setString('token', token!.token ?? '');
                                prefs.setString('refresh_token', token!.refresh_token ?? '');
                              },
                              onError: (error) {
                                print("SharedPreferences ERROR = $error");
                              });

                            } catch (e) {
                              print(e);
                            }

                            Navigator.push(
                                context,
                                MaterialPageRoute(
                                    builder: (context) =>
                                        const HomePage()
                                )
                            );
                          } else {
                            ScaffoldMessenger.of(context).showSnackBar(
                                const SnackBar(content: SnackBarLoginError())
                            );
                          }
                        }
                        ).catchError((e) => ScaffoldMessenger.of(context).showSnackBar(
                            const SnackBar(content: SnackBarLoginError())
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
                      Navigator.push(
                          context,
                          MaterialPageRoute(
                              builder: (context) => RegisterClientPage(title: 'Register Client')
                          )
                      );
                    },
                    child: Text(AppLocalizations.of(context)!.registerLabel, style: TextStyle(fontSize: 18)),
                  ),
                ),
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: Divider(),
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: Center(
                  child: ElevatedButton(
                    style: ElevatedButton.styleFrom(
                      minimumSize: const Size.fromHeight(50), // NEW
                    ),
                    onPressed: () {
                      Navigator.push(
                          context,
                          MaterialPageRoute(
                              builder: (context) => RegisterBusinessPage(title: 'Register Business')
                          )
                      );
                    },
                    child: Text(AppLocalizations.of(context)!.registerBusinessLabel, style: TextStyle(fontSize: 18)),
                  ),
                ),
              ),
              Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
            ],
          ),
        ),
      ),
      ),
    );
  }
}

class SnackBarLoginError extends StatelessWidget {
  const SnackBarLoginError({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarLoginError);
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
    return Text(AppLocalizations.of(context)!.submitLoginLabel, style: TextStyle(fontSize: 18));
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
