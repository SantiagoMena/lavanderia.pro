import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';

class BusinessClientView extends StatefulWidget {
   const BusinessClientView({super.key, this.token});

  final String? token;

  @override
  State<BusinessClientView> createState() => _BusinessClientViewState();
}

class _BusinessClientViewState extends State<BusinessClientView> {
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
            alignment: Alignment.topCenter,
            child: SingleChildScrollView(
              reverse: true,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.start,
                mainAxisSize: MainAxisSize.min,
                children: [
                  SizedBox(
                    height: 600,
                    child: ListView.builder(
                      itemCount: items.length,
                      /*prototypeItem: ListTile(
                        title: Text(items.first),
                      ),*/
                      itemBuilder: (context, index) {
                        return ProductCard(items: items, productIndex: index);
                      },
                    ),
                  ),


                  ExpansionTile(
                    title: Text('Business Name'),
                    subtitle: Text('11 Products selected'),
                    children: <Widget>[
                      ListTile(title: Text('This is tile number 1')),
                      ListTile(title: Text('This is tile number 3')),
                      ListTile(title: Text('This is tile number 4')),
                      ListTile(title: Text('This is tile number 5')),
                      ListTile(title: Text('This is tile number 6')),
                      ListTile(title: Text('This is tile number 7')),
                      ListTile(title: Text('This is tile number 8')),
                      ListTile(title: Text('This is tile number 9')),
                      ListTile(title: Text('This is tile number 10')),
                      ListTile(title: Text('This is tile number 11')),
                      ListTile(
                        title: Row(
                          children: const [
                            Align(
                              alignment: Alignment.topLeft,
                              child: Text(
                                '☝️ Click To Process Order',
                                style: TextStyle(color: Colors.white),
                              ),
                            ),
                            Expanded(child: Text("")),
                            Align(
                              alignment: Alignment.topRight,
                              child: Text(
                                '\$123,4',
                                style: TextStyle(color: Colors.white),
                              ),
                            ),
                          ],
                        ),
                        tileColor: Colors.green,
                        onTap: () {
                          Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => ProcessedOrderClient()
                              )
                          );
                        },
                      ),
                    ],
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

class ProductCard extends StatelessWidget {
  const ProductCard({
    super.key,
    required this.items,
    required this.productIndex,
  });

  final int productIndex;
  final List<String> items;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: 50, vertical: 16),
      child: Card(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 10, vertical: 10),
          child: Column(
            children: [
              Row(
                children: [
                  Container(child: Text(items[productIndex])),
                  Expanded(child: Text("")),
                  Expanded(child:
                    Align(
                        alignment: Alignment.topRight,
                        child: Padding(
                          padding: EdgeInsets.symmetric(horizontal: 5),
                            child: Text("+1")
                        )
                    )
                  ),
                  Container(child: Text("➕")),
                ]
              ),
              Row(
                  children: [
                    Container(child: Text("Desc ...")),
                    Expanded(child: Text("")),
                    Container(child: Text("➖")),
                  ]
              ),
              Row(
                  children: [
                    Container(child: Text('Price: \$0000')),
                    Expanded(child: Text("")),
                    Container(child: Text("🗑️")),
                  ]
              ),
            ],
          ),
        ),
      )
    );
  }
}
