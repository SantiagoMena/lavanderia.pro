import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/business_tabs/business_delivery_create_form.dart';
import 'package:lavanderiapro/pages/business_tabs/business_delivery_update_form.dart';
import 'package:lavanderiapro/pages/business_tabs/business_update_form.dart';
import 'package:lavanderiapro/pages/business_tabs/business_view.dart';
import 'package:lavanderiapro/pages/business_tabs/product_create_form.dart';
import 'package:lavanderiapro/pages/client_tabs/business_client_view.dart';

class DeliveryBusinessView extends StatefulWidget {
   const DeliveryBusinessView({super.key, this.token});

  final String? token;

  @override
  State<DeliveryBusinessView> createState() => _DeliveryBusinessViewState();
}

class _DeliveryBusinessViewState extends State<DeliveryBusinessView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Add your onPressed code here!
          Navigator.push(context, MaterialPageRoute(builder: (builder) => BusinessDeliveryCreateForm()));
        },
        child: const Icon(Icons.add),
      ),
      appBar: AppBar(
        title: const Text("Manage Business Delivery"),
      ),
      body: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints viewportConstraints) {
          var items = List<String>.generate(15, (i) => 'Delivery $i');
          return Align(
          alignment: Alignment.topCenter,
          child: Column(
            children:[
                const Padding(
                  padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: Align(
                    alignment: Alignment.center,
                    child: Padding(
                        padding: EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                        child: Text("Manage Users")
                    ),
                  ),
                ),
                Expanded(
                  child: ListView.builder(
                    itemCount: items.length,
                    itemBuilder: (context, index) {
                      var businessItem = items[index];
                      return DeliveryCard(businessItem: businessItem);
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

class DeliveryCard extends StatelessWidget {
  const DeliveryCard({
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
                  builder: (context) => const BusinessDeliveryUpdateForm()
              )
          );
        },
        child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 50),
            child: Text(businessItem, style: const TextStyle(color: Colors.black),),
        )
      ),
    );
  }
}
