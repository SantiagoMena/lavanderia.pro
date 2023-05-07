
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/models/address.dart';
import 'package:lavanderiapro/pages/business_tabs/business_view.dart';
import 'package:lavanderiapro/services/post_address_service.dart';
import 'package:lavanderiapro/services/put_address_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AddressUpdateForm extends StatefulWidget {
  const AddressUpdateForm({super.key, required this.address});

  final Address address;

  @override
  State<AddressUpdateForm> createState() => _AddressUpdateFormState();
}

class _AddressUpdateFormState extends State<AddressUpdateForm> {
  final _formKey = GlobalKey<FormState>();
  TextEditingController addressNameController = TextEditingController();
  TextEditingController addressController = TextEditingController();
  TextEditingController phoneController = TextEditingController();
  TextEditingController extraInfoController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    addressNameController.text = widget.address!.name ?? "";
    addressController.text = widget.address!.address ?? "";
    phoneController.text = widget.address!.phone ?? "";
    extraInfoController.text = widget.address!.extra ?? "";

    return Scaffold(
      appBar: AppBar(
        title: Text("Update Address: ${widget.address.name}"),
      ),
      body: Form(
        key: _formKey,
        child: SingleChildScrollView(
          reverse: true,
          child: Column(
            children: [
              const Padding(
                  padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                  child: Text("Update Address")
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: TextFormField(
                  controller: addressNameController,
                  decoration: const InputDecoration(border: OutlineInputBorder(), label: Text("Name Address")),
                  validator: (value) {
                    if(value == null || value.isEmpty){
                      return "Empty address name alert";
                    }
                    return null;
                  },
                ),
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: TextFormField(
                  controller: addressController,
                  decoration: const InputDecoration(border: OutlineInputBorder(), label: Text("Address")),
                  validator: (value) {
                    if(value == null || value.isEmpty){
                      return "Empty address alert";
                    }
                    return null;
                  },
                ),
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: TextFormField(
                  controller: phoneController,
                  decoration: const InputDecoration(border: OutlineInputBorder(), label: Text("Phone")),
                  validator: (value) {
                    if(value == null || value.isEmpty){
                      return "Empty phone alert";
                    }
                    return null;
                  },
                ),
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: TextFormField(
                  controller: extraInfoController,
                  decoration: const InputDecoration(border: OutlineInputBorder(), label: Text("Extra Info")),
                  validator: (value) {
                    return null;
                  },
                ),
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: Center(
                  child: FutureBuilder(
                    future: SharedPreferences.getInstance(),
                    builder: (context, snapshot) {
                      if(snapshot.hasData) {
                        String token = snapshot.data!.getString('token') ?? "";
                        return ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          minimumSize: const Size.fromHeight(50),
                          backgroundColor: Colors.green,
                        ),
                        onPressed: () {
                          Address address = Address.fromJson({
                            'id': widget.address!.id,
                            'address': addressController.text,
                            'name': addressNameController.text,
                            'phone': phoneController.text,
                            'extra': extraInfoController.text,
                          });
                          
                          FocusManager.instance.primaryFocus?.unfocus();
                          if(_formKey.currentState!.validate()){
                            putAddress(token, address).then((addressCreated) {
                              Navigator.pop(context);
                              print(['addressCreated', addressCreated]);
                            });
                          } else {
                            ScaffoldMessenger.of(context).showSnackBar(
                                const SnackBar(content: SnackBarFillInputLabel())
                            );
                          }
                        },
                        child: const Text('Update Address'),
                      );
                      }
                      else {
                        return const CircularProgressIndicator();
                      }
                    }
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

class SnackBarFillInputLabel extends StatelessWidget {
  const SnackBarFillInputLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarFillInput);
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
      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
      child: ElevatedButton(
        style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
        onPressed: () {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) => const BusinessView()
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
