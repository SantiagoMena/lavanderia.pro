import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/pages/client_tabs/address_create_form.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';
import 'package:lavanderiapro/services/get_address_client_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AddressesClientView extends StatefulWidget {
   const AddressesClientView({super.key, this.token});

  final String? token;

  @override
  State<AddressesClientView> createState() => _AddressesClientViewState();
}

class _AddressesClientViewState extends State<AddressesClientView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
        title: Text("Manage Addresses"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) => AddressCreateForm()
              )
          );
        },
        child: const Icon(Icons.add),
      ),
      body: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints viewportConstraints) {
          var items = List<String>.generate(15, (i) => 'Address Title $i');
          return Align(
            alignment: Alignment.bottomCenter,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.start,
                mainAxisSize: MainAxisSize.min,
                children: [
                  const SizedBox(
                    height: 50,
                    child: Align(
                      alignment: Alignment.center,
                      child: Padding(
                          padding: EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                          child: Text("Select a Address")
                      ),
                    ),
                  ),
                  Expanded(
                    child: FutureBuilder(
                      future: SharedPreferences.getInstance(),
                      builder: (context, snapshot) {
                        if(snapshot.hasData){
                          return FutureBuilder(
                            future: getAddressClient(snapshot.data!.getString('token') ?? ""),
                            builder: (context, snapshot) {
                            if(snapshot.hasData) {
                              return ListView.builder(
                                itemCount: snapshot.data!.length,
                                itemBuilder: (context, index) {
                                  var addressItem = snapshot.data![index];
                                  return AddressCard(addressItem: addressItem);
                                },
                              );
                            } else {
                              return const Text("Empty Addresses");
                            }
                          }
                        );
                        }
                        else {
                          return const CircularProgressIndicator();
                        }
                      }
                    ),
                  ),
                  // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
                ],
               ),
          );
        }
      ),
    );
  }
}

class AddressCard extends StatelessWidget {
  const AddressCard({
    super.key,
    required this.addressItem,
  });

  final Address addressItem;

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: EdgeInsets.symmetric(horizontal: 50, vertical: 16),
        child: ElevatedButton(
          onPressed: () {
            Navigator.push(
                context,
                MaterialPageRoute(
                    builder: (context) => AddressCreateForm()
                )
            );
          },
          style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
          child: Padding(
            padding: EdgeInsets.symmetric(horizontal: 10, vertical: 10),
            child: Column(
              children: [
                Row(
                    children: [
                      Container(child:
                        Text(addressItem.name ?? "", style: TextStyle(color: Colors.black)),
                      ),
                      Expanded(child: Text("")),
                      Expanded(
                          child: Align(
                              alignment: Alignment.topRight,
                              child: Text(
                                "Active",
                                style: TextStyle(color: Colors.green),
                              )
                          )
                      ),
                    ]
                ),
                Row(
                    children: [
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: Text(""),
                        ),
                    ]
                ),
                Row(
                    children: [
                      Expanded(child: Text(addressItem.address ?? "", style: TextStyle(color: Colors.black),)),
                    ]
                ),
              ],
            ),
          ),
        )
    );
  }
}