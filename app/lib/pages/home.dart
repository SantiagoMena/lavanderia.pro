import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/pages/business_tabs/business_tab.dart';
import 'package:lavanderiapro/pages/business_tabs/home_business.dart';
import 'package:lavanderiapro/pages/client_tabs/home_client.dart';
import 'package:lavanderiapro/pages/delivery_tabs/home_delivery.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    Profile? profile;
    List<Widget> children = [
      Padding(
          padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
          child: Center(
            child: Text('👤️', style: TextStyle(fontSize: 75)),
          )
      )
    ];

    return Scaffold(
      body: FutureBuilder(
        future: SharedPreferences.getInstance(),
        builder: (context, snapshot) {
          if(snapshot.hasData) {
            String? tokenStored = snapshot.data!.getString('token');

            getProfile(tokenStored ?? '').then((profile) {
              if(profile!.business == true) {
                children.add(const LoginAsBusinessButton());
              }
              if(profile.client == true) {
                children.add(const LoginAsClientButton());
              }
              if(profile.delivery == true) {
                children.add(const LoginAsDeliveryButton());
              }

              if(profile.business == true && profile.client == false && profile.delivery == false) {
                Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) =>
                            const HomeBusinessTab()
                    )
                );
              }

              if(profile.business == false && profile.client == true && profile.delivery == false) {
                Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) =>
                            const HomeClientTab()
                    )
                );
                // return HomeClientTab();
              }

              if(profile.business == false && profile.client == false && profile.delivery == true) {
                Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) =>
                            const HomeDeliveryTab()
                    )
                );
                // return HomeDeliveryTab();
              }
            }).catchError((onError) => print(onError));

          } else if (snapshot.connectionState != ConnectionState.done) {
            children = [CircularProgressIndicator()];
          } else {

          }


          return ListView(
            children: children,
          );
        },
      )
    );
  }
}

class LoginAsDeliveryButton extends StatelessWidget {
  const LoginAsDeliveryButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
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
      child: Text(AppLocalizations.of(context)!.loginAsDeliveryButtonLabel, style: TextStyle(fontSize: 18)),
    ),
          ),
        );
  }
}

class LoginAsClientButton extends StatelessWidget {
  const LoginAsClientButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
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
      child: Text(AppLocalizations.of(context)!.loginAsClientButtonLabel, style: TextStyle(fontSize: 18)),
    ),
          ),
        );
  }
}

class LoginAsBusinessButton extends StatelessWidget {
  const LoginAsBusinessButton({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
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
          child: Text(AppLocalizations.of(context)!.loginAsBusinessButtonLabel, style: TextStyle(fontSize: 18)),
        ),
      ),
    );
  }
}