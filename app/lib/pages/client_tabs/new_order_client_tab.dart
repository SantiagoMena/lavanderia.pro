import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/client_tabs/business_client_view.dart';
import 'package:lavanderiapro/services/search_business_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class NewOrderClientTab extends StatefulWidget {
   const NewOrderClientTab({super.key, this.token});

  final String? token;

  @override
  State<NewOrderClientTab> createState() => _NewOrderClientTabState();
}

class _NewOrderClientTabState extends State<NewOrderClientTab> {
  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: SharedPreferences.getInstance(),
      builder: (context, snapshot) {
        if(snapshot.hasData) {
          String? token = snapshot.data!.getString('token');

          return LayoutBuilder(
        builder: (BuildContext context, BoxConstraints viewportConstraints) {
          // var items = List<String>.generate(1, (i) => 'Business $i');
          return Align(
          alignment: Alignment.topCenter,
          child: Column(
            children:[
                const SizedBox(
                  height: 50,
                  child: Align(
                    alignment: Alignment.center,
                    child: Padding(
                      padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                      child: SelectBusinessLabel()
                    ),
                  ),
                ),
                Expanded(
                  child: FutureBuilder(
                    future: searchBusiness(token ?? ""),
                    builder: (context, snapshot) {
                      if(snapshot.hasData) {
                        return ListView.builder(
                        itemCount: snapshot.data!.length ?? 1,
                        itemBuilder: (context, index) {
                          var businessItem = snapshot.data![index];
                          return BusinessCard(businessItem: businessItem, token: token ?? "");
                        },
                      );
                      } else {
                       return CircularProgressIndicator();
                      }
                    },
                  ),
                ),
                // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
              ],
             ),
          );
        }
      );
        } else {
          return CircularProgressIndicator();
        }
      },
    );
  }
}

class BusinessCard extends StatelessWidget {
  const BusinessCard({
    super.key,
    required this.businessItem,
    required this.token,
  });

  final Business businessItem;
  final String token;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
      child: ElevatedButton(
        onPressed: () => {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) => BusinessClientView(businessItem: businessItem, token: token)
              )
          )
        },
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 8, vertical: 50),
            child: Text(businessItem.name ?? ""),
          )
      ),
    );
  }
}

class SelectBusinessLabel extends StatelessWidget {
  const SelectBusinessLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.selectBusinessLabel);
  }
}
