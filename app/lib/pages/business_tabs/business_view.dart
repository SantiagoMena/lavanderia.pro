import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';

class BusinessView extends StatefulWidget {
   const BusinessView({super.key, this.token});

  final String? token;

  @override
  State<BusinessView> createState() => _BusinessViewState();
}

class _BusinessViewState extends State<BusinessView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Business Name"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Add your onPressed code here!
        },
        child: const Icon(Icons.add),
      ),
      body: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints viewportConstraints) {
          var items = List<String>.generate(15, (i) => 'Product $i');
          return Align(
            alignment: Alignment.bottomCenter,
            child: SingleChildScrollView(
              reverse: true,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.start,
                mainAxisSize: MainAxisSize.min,
                children: [
                  SizedBox(
                    height: 50,
                    child: Text("Manage Products"),
                  ),
                  Expanded(
                    child: ListView.builder(
                      itemCount: items.length,
                      /*prototypeItem: ListTile(
                        title: Text(items.first),
                      ),*/
                      itemBuilder: (context, index) {
                        return ManageProductCard(items: items, productIndex: index);
                      },
                    ),
                  ),
                  // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
                ],
               ),
          ),
          );
        }
      ),
    );
  }
}


class ManageProductCard extends StatelessWidget {
  const ManageProductCard({
    super.key,
    required this.items,
    required this.productIndex,
  });

  final int productIndex;
  final List<String> items;

  @override
  Widget build(BuildContext context) {
    return Text(items[productIndex]);
  }
}
