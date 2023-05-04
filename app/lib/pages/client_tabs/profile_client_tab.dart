import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/pages/client_tabs/addresses_client_view.dart';
import 'package:lavanderiapro/services/change_password_service.dart';
import 'package:lavanderiapro/services/get_business_profile_service.dart';
import 'package:lavanderiapro/services/get_client_profile_service.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ProfileClientTab extends StatefulWidget {
   const ProfileClientTab({super.key});

  @override
  State<ProfileClientTab> createState() => _ProfileClientTabState();
}

class _ProfileClientTabState extends State<ProfileClientTab> {
  final _formNameKey = GlobalKey<FormState>();
  final _formPasswordKey = GlobalKey<FormState>();

  TextEditingController nameController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  TextEditingController newPasswordController = TextEditingController();

  @override
  Widget build(BuildContext context) {

    Profile? profile;

    return FutureBuilder(
      future: SharedPreferences.getInstance(),
      builder: (context, snapshot) {
        if (snapshot.connectionState != ConnectionState.done) {
          return CircularProgressIndicator();
        }

        if(snapshot.hasData) {
          String token = snapshot.data!.getString('token') ?? '';

          return FutureBuilder(
            future: getClientProfile(token),
            builder: (context, snapshot) {
            return LayoutBuilder(builder: (BuildContext context, BoxConstraints viewportConstraints) {
              if(snapshot.hasData){
                print(['nameController.text', nameController.text]);
                nameController.text = snapshot.data!.name ?? '';
                return Align(
                alignment: Alignment.topCenter,
                child: SingleChildScrollView(
                  reverse: true,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    mainAxisAlignment: MainAxisAlignment.start,
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      Padding(
                        padding: const EdgeInsets.symmetric(
                            horizontal: 8, vertical: 16),
                        child: Center(
                          child: ElevatedButton(
                            style: ElevatedButton.styleFrom(
                              minimumSize: const Size.fromHeight(50),
                              backgroundColor: Colors.green,
                            ),
                            onPressed: () {
                              Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                      builder: (context) =>
                                          AddressesClientView()));
                            },
                            child: Text(AppLocalizations.of(context)!
                                .manageAddressesButtonLabel),
                          ),
                        ),
                      ),
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
                            children: [
                              Padding(
                                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                                child: TextFormField(
                                  controller: passwordController,
                                  obscureText: true,
                                  decoration: const InputDecoration(border: OutlineInputBorder(), label: CurrentPasswordLabel()),
                                  validator: (value) {
                                    if (value == null || value.isEmpty) {
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
                                    if (value == null || value.isEmpty) {
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
                                                    content: SnackBarErrorOnChangePasswordClient()
                                                ));

                                                return;
                                              }
                                              if(auth!.id!.length > 0) {
                                                ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                                                    content: SnackBarPasswordChangedSuccessfullyClient()
                                                ));

                                                passwordController.text = "";
                                                newPasswordController.text = "";
                                              } else {
                                                ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                                                    content: SnackBarErrorOnChangePasswordClient()
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
                        padding: const EdgeInsets.symmetric(
                            horizontal: 8, vertical: 16),
                        child: Center(
                          child: ElevatedButton(
                            style: ElevatedButton.styleFrom(
                              minimumSize: const Size.fromHeight(50),
                              backgroundColor: Colors.green,
                            ),
                            onPressed: () {
                              Navigator.pushNamedAndRemoveUntil(
                                  context, '/login', (_) => false);
                            },
                            child: Text("Logout"),
                          ),
                        ),
                      ),
                      Padding(
                          padding: EdgeInsets.only(
                              bottom: MediaQuery.of(context)
                                  .viewInsets
                                  .bottom))
                    ],
                  ),
                ),
              );
              } else {
                return CircularProgressIndicator();
              }
            });
          }
          );
        }

        return CircularProgressIndicator();

    });
  }
}

class SnackBarPasswordChangedSuccessfullyClient extends StatelessWidget {
  const SnackBarPasswordChangedSuccessfullyClient({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarPasswordChangedSuccessfully);
  }
}

class SnackBarErrorOnChangePasswordClient extends StatelessWidget {
  const SnackBarErrorOnChangePasswordClient({
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

