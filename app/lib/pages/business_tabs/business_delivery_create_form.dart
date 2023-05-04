import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/business_tabs/business_view.dart';
import 'package:lavanderiapro/pages/business_tabs/products_business_view.dart';
import 'package:lavanderiapro/pages/client_tabs/business_client_view.dart';

class BusinessDeliveryCreateForm extends StatefulWidget {
   const BusinessDeliveryCreateForm({super.key, this.token});

  final String? token;

  @override
  State<BusinessDeliveryCreateForm> createState() => _BusinessDeliveryCreateFormState();
}

class _BusinessDeliveryCreateFormState extends State<BusinessDeliveryCreateForm> {
  final _formKey = GlobalKey<FormState>();
  TextEditingController nameController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Add your onPressed code here!
        },
        child: const Icon(Icons.add),
      ),
      appBar: AppBar(
        title: Text("Create Delivery Account"),
      ),
      body: Form(
        key: _formKey,
        child: SingleChildScrollView(
          reverse: true,
          child: Column(
            children: [
                const Padding(
                    padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                    child: Text("Create New Delivery")
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: TextFormField(
                    controller: nameController,
                    decoration: const InputDecoration(border: OutlineInputBorder(), label: BusinessNameLabel()),
                    validator: (value) {
                      if(value == null || value.isEmpty){
                        return AppLocalizations.of(context)!.emptyNameAlert;
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
                        if(_formKey.currentState!.validate()){
                          // Change Name
                          Navigator.push(
                              context,
                              MaterialPageRoute(
                                  builder: (context) => ProductBusinessView()
                              )
                          );
                        } else {
                          ScaffoldMessenger.of(context).showSnackBar(
                              const SnackBar(content: Text("FillInputSnackBar()"))
                          );
                        }
                      },
                      child: Text(AppLocalizations.of(context)!.createBusinessButtonLabel),
                    ),
                  ),
                ),
                Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
              ],
            ),
        ),
        ),
    );
  }
}

class BusinessNameLabel extends StatelessWidget {
  const BusinessNameLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.businessNameLabel);
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
                  builder: (context) => BusinessView()
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
