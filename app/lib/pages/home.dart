import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/login.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/pages/business_tabs/home_business.dart';
import 'package:lavanderiapro/pages/client_tabs/home_client.dart';
import 'package:lavanderiapro/pages/delivery_tabs/delivery_client.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key, this.token});

  final String? token;

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

      //children.add(const LoginAsBusinessButton());


    return Scaffold(
      /*appBar: AppBar(
        title: const Text('Select Profile'),
      ),*/
      body: FutureBuilder(
        future: getProfile(token ?? ''),
        builder: (context, snapshot) {


          if(snapshot.hasData) {
            if(snapshot.data!.business == true) {
              children.add(const LoginAsBusinessButton());
            }
            if(snapshot.data!.client == true) {
              children.add(const LoginAsClientButton());
            }
            if(snapshot.data!.delivery == true) {
              children.add(const LoginAsDeliveryButton());
            }

            print([snapshot.data!.business, snapshot.data!.client, snapshot.data!.delivery]);
            if(snapshot.data!.business == true && snapshot.data!.client == false && snapshot.data!.delivery == false) {
              return HomeBusinessTab();
            }

            if(snapshot.data!.business == false && snapshot.data!.client == true && snapshot.data!.delivery == false) {
              return HomeClientTab();
            }

            if(snapshot.data!.business == false && snapshot.data!.client == false && snapshot.data!.delivery == true) {
              return HomeDeliveryTab();
            }

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