import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/services/change_password_service.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ProfileDeliveryTab extends StatefulWidget {
   const ProfileDeliveryTab({super.key, this.token});

  final String? token;

  @override
  State<ProfileDeliveryTab> createState() => _ProfileDeliveryTabState();
}

class _ProfileDeliveryTabState extends State<ProfileDeliveryTab> {
  final _formPasswordKey = GlobalKey<FormState>();
  final _formNameKey = GlobalKey<FormState>();

  TextEditingController nameController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  TextEditingController newPasswordController = TextEditingController();

  @override
  Widget build(BuildContext context) {

    Profile? profile;

    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints viewportConstraints) {
      return Align(
        alignment: Alignment.topCenter,
        child: SingleChildScrollView(
              reverse: true,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.start,
                mainAxisSize: MainAxisSize.min,
                children: [
                  Form(
                    key: _formNameKey,
                    child: Column(
                      children: [
                        Padding(
                          padding: const EdgeInsets.symmetric(
                              horizontal: 8, vertical: 16),
                          child: TextFormField(
                            controller: nameController,
                            decoration: const InputDecoration(
                                border: OutlineInputBorder(),
                                label: ClientNameLabel()),
                            validator: (value) {
                              if (value == null || value.isEmpty) {
                                return AppLocalizations.of(context)!
                                    .emptyNameAlert;
                              }
                              return null;
                            },
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.symmetric(
                              horizontal: 8, vertical: 16),
                          child: Center(
                            child: ElevatedButton(
                              style: ElevatedButton.styleFrom(
                                minimumSize:
                                const Size.fromHeight(50),
                                backgroundColor: Colors.green,
                              ),
                              onPressed: () {
                                if (_formNameKey.currentState!
                                    .validate()) {
                                  // Change Name
                                } else {
                                  ScaffoldMessenger.of(context)
                                      .showSnackBar(const SnackBar(
                                      content:
                                      FillInputSnackBar()));
                                }
                              },
                              child: Text(
                                  AppLocalizations.of(context)!
                                      .changeNameButtonLabel),
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                  FutureBuilder(
                    future: SharedPreferences.getInstance(),
                    builder: (context, snapshot) {
                      if(snapshot.hasData) {
                        return Form(
                              key: _formPasswordKey,
                               child: Column(
                                   children:[
                                 Padding(
                                   padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                                   child: TextFormField(
                                     controller: passwordController,
                                     obscureText: true,
                                     decoration: const InputDecoration(border: OutlineInputBorder(), label: CurrentPasswordLabel()),
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
                                   child: TextFormField(
                                     controller: newPasswordController,
                                     obscureText: true,
                                     decoration: const InputDecoration(border: OutlineInputBorder(), label: NewPasswordLabel()),
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
                                         minimumSize: const Size.fromHeight(50),
                                         backgroundColor: Colors.green,
                                       ),
                                       onPressed: () {
                                         FocusManager.instance.primaryFocus?.unfocus();
                                         if(_formPasswordKey.currentState!.validate()){
                                           String? token = snapshot.data!.getString('token') ?? '';
                                           changePassword(token, passwordController.text, newPasswordController.text)
                                               .then((auth) {
                                             if(auth == null){
                                               ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                                                   content: SnackBarErrorOnChangePasswordDelivery()
                                               ));

                                               return;
                                             }
                                             if(auth!.id!.length > 0) {
                                               ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                                                   content: SnackBarPasswordChangedSuccessfullyDelivery()
                                               ));

                                               passwordController.text = "";
                                               newPasswordController.text = "";
                                             } else {
                                               ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                                                   content: SnackBarErrorOnChangePasswordDelivery()
                                               ));
                                             }
                                           })
                                           .catchError((error) => error);
                                         } else {
                                           ScaffoldMessenger.of(context).showSnackBar(
                                               const SnackBar(content: FillInputSnackBar())
                                           );
                                         }
                                       },
                                       child: Text(AppLocalizations.of(context)!.changePasswordButtonLabel),
                                     ),
                                   ),
                                 ),
                                 ],
                               ),
                             );
                      } else {
                        return CircularProgressIndicator();
                      }
                    },
                  ),
                  Padding(
                   padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                   child: Center(
                     child: ElevatedButton(
                       style: ElevatedButton.styleFrom(
                         minimumSize: const Size.fromHeight(50),
                         backgroundColor: Colors.green,
                       ),
                       onPressed: () {
                         Navigator.pushNamedAndRemoveUntil(context, '/login', (_) => false);
                       },
                       child: Text("Logout"),
                     ),
                   ),
                 ),
                  Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
                ],
               ),
              ),
        );
      }
    );
  }
}

class SnackBarPasswordChangedSuccessfullyDelivery extends StatelessWidget {
  const SnackBarPasswordChangedSuccessfullyDelivery({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarPasswordChangedSuccessfully);
  }
}

class SnackBarErrorOnChangePasswordDelivery extends StatelessWidget {
  const SnackBarErrorOnChangePasswordDelivery({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarErrorOnChangePassword);
  }
}

class NewPasswordLabel extends StatelessWidget {
  const NewPasswordLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.newPasswordLabel);
  }
}

class CurrentPasswordLabel extends StatelessWidget {
  const CurrentPasswordLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.currentPasswordLabel);
  }
}

class ClientNameLabel extends StatelessWidget {
  const ClientNameLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.clientNameLabel);
  }
}

class FillInputSnackBar extends StatelessWidget {
  const FillInputSnackBar({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarFillInput);
  }
}

