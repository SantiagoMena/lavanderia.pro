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
                  Center(
                    child: Text("Manage Products"),
                  ),
                  SizedBox(
                    height: 650,
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

class SelectedProducts extends StatelessWidget {
  const SelectedProducts({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text("Products Selected +1");
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
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Add your onPressed code here!
        },
        child: const Icon(Icons.add),
      ),
      /*body: Padding(
        padding: EdgeInsets.symmetric(horizontal: 50, vertical: 16),
        child: ElevatedButton(
          style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
          onPressed: () {

          },
            child: Padding(
              padding: EdgeInsets.symmetric(horizontal: 10, vertical: 10),
              child: Column(
                children: [
                  Row(
                    children: [
                      Container(child: Text(items[productIndex], style: TextStyle(color: Colors.black),)),
                      Expanded(child: Text("")),
                      Expanded(child:
                        Align(
                            alignment: Alignment.topRight,
                            child: Padding(
                              padding: EdgeInsets.symmetric(horizontal: 5),
                                child: Text("Stock: 112", style: TextStyle(color: Colors.black),)
                            )
                        )
                      ),
                    ]
                  ),
                  Row(
                      children: [
                        Container(child: Text("Desc ...", style: TextStyle(color: Colors.black),)),
                        Expanded(child: Text("")),
                      ]
                  ),
                  Row(
                      children: [
                        Container(child: Text('Price: \$0000', style: TextStyle(color: Colors.black),)),
                        Expanded(child: Text("")),
                      ]
                  ),
                ],
              ),
          ),
        )
      ),*/
    );
  }
}
