import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/business_tabs/business_view.dart';
import 'package:lavanderiapro/pages/business_tabs/product_create_form.dart';
import 'package:lavanderiapro/pages/client_tabs/business_client_view.dart';

class ProductBusinessView extends StatefulWidget {
   const ProductBusinessView({super.key, this.token});

  final String? token;

  @override
  State<ProductBusinessView> createState() => _ProductBusinessViewState();
}

class _ProductBusinessViewState extends State<ProductBusinessView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Add your onPressed code here!
          Navigator.push(context, MaterialPageRoute(builder: (builder) => ProductCreateForm()));
        },
        child: const Icon(Icons.add),
      ),
      appBar: AppBar(
        title: Text("Manage Products"),
      ),
      body: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints viewportConstraints) {
          var items = List<String>.generate(15, (i) => 'Products $i');
          return Align(
          alignment: Alignment.topCenter,
          child: Column(
            children:[
                const SizedBox(
                  height: 50,
                  child: Align(
                    alignment: Alignment.center,
                    child: Padding(
                      padding: EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                      child: Text("Select Product")
                    ),
                  ),
                ),
                Expanded(
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
                ), // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
              ],
             ),
          );
        }
      ),
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
        style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
        onPressed: () {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) => ProductCreateForm()
              )
          );
        },
        child: Padding(
            child: Text(businessItem, style: TextStyle(color: Colors.black),),
            padding: EdgeInsets.symmetric(horizontal: 8, vertical: 50),
        )
      ),
    );
  }
}
