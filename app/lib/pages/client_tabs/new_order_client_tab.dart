import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/client_tabs/business_client_view.dart';

class NewOrderClientTab extends StatefulWidget {
   const NewOrderClientTab({super.key, this.token});

  final String? token;

  @override
  State<NewOrderClientTab> createState() => _NewOrderClientTabState();
}

class _NewOrderClientTabState extends State<NewOrderClientTab> {
  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (BuildContext context, BoxConstraints viewportConstraints) {
        var items = List<String>.generate(15, (i) => 'Business $i');
        return Align(
        alignment: Alignment.topCenter,
        child: Column(
          children:[
              const Expanded(
                child: Align(
                  alignment: Alignment.center,
                  child: Padding(
                    padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                    child: SelectBusinessLabel()
                  ),
                ),
              ),
              SizedBox(
                height: 550,
                child: ListView.builder(
                  itemCount: items.length,
                  /*prototypeItem: ListTile(
                    title: Text(items.first),
                  ),*/
                  itemBuilder: (context, index) {
                    var businessItem = items[index];
                    return BusinessCard(businessItem: businessItem);
                  },
                ),
              ),
              // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
            ],
           ),
        );
      }
    );
  }
}

class BusinessCard extends StatelessWidget {
  const BusinessCard({
    super.key,
    required this.businessItem,
  });

  final String businessItem;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
      child: ElevatedButton(
        onPressed: () => {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) => BusinessClientView()
              )
          )
        },
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 8, vertical: 50),
            child: Text(businessItem),
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
