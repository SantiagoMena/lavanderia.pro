import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';
import 'package:lavanderiapro/services/get_address_client_service.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:lavanderiapro/widgets/ProductCardCart.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:lavanderiapro/models/product.dart';

class CheckOrderClient extends StatefulWidget {
   const CheckOrderClient({super.key, this.order});

   final OrderModel? order;

  @override
  State<CheckOrderClient> createState() => _CheckOrderClientState();
}

class _CheckOrderClientState extends State<CheckOrderClient> {
  final _formKey = GlobalKey<FormState>();
  TextEditingController nameController = TextEditingController();

  void pushProduct(Product product) {
    setState(() {
      widget.order!.add(product);
    });
  }

  void popProduct(Product product) {
    setState(() {
      widget.order!.remove(product);
    });
  }
  @override
  Widget build(BuildContext context) {

    return Scaffold(
      appBar: AppBar(
        title: Text("Check Order \$${widget.order!.totalPrice}"),
      ),
      body: LayoutBuilder(
      builder: (BuildContext context, BoxConstraints viewportConstraints) {
      return Align(
        alignment: Alignment.topCenter,
        child: SingleChildScrollView(
              reverse: true,
              child:  Form(
                key: _formKey,
                child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    mainAxisAlignment: MainAxisAlignment.start,
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      const SizedBox(
                        height: 50,
                        child: Align(
                            alignment: Alignment.center,
                            child: DropdownAddress()
                        ),
                      ),
                      SizedBox(
                        height: viewportConstraints.maxHeight - AppBar().preferredSize.height - 50,
                        child: ListView.builder(
                          itemCount: widget.order!.getGrouped().length ?? 0,
                          itemBuilder: (context, index) {
                            var productItem = widget.order!.getGrouped()[index];
                            if(productItem == null) {
                              return Text("");
                            }

                            return ProductCardCart(
                                productItem: productItem,
                                pushProductCallback: pushProduct,
                                popProductCallback: popProduct,
                                countProductCallback: widget.order!.countProduct,
                            );
                          },
                        ),
                      ),
                      SizedBox(
                        height: 50,
                        child: Padding(
                          padding: const EdgeInsets.symmetric(horizontal: 8),
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
                                          builder: (context) => const ProcessedOrderClient()
                                      )
                                  );
                                } else {
                                  ScaffoldMessenger.of(context).showSnackBar(
                                      const SnackBar(content: Text("FillInputSnackBar()"))
                                  );
                                }
                              },
                              child: Text("Make Order"),
                            ),
                          ),
                        ),
                      ),
                    ],
                   ),
              ),
              ),
        );
      }
      ),
    );
  }
}

const List<String> list = <String>['One', 'Two', 'Three', 'Four'];
class DropdownAddress extends StatefulWidget {
  const DropdownAddress({super.key});

  @override
  State<DropdownAddress> createState() => _DropdownAddressState();
}

class _DropdownAddressState extends State<DropdownAddress> {

  String dropdownValue = list.first;

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: SharedPreferences.getInstance(),
      builder: (context, snapshot) {
        if(snapshot.hasData){
          return FutureBuilder(
            future: getAddressClient(snapshot.data!.getString('token') ?? ""),
            builder: (contextAddress, snapshotAddress) {
              if(snapshotAddress.hasData) {
                List<Address>? addresses = snapshotAddress.data;

                return DropdownButton<String>(
                  value: addresses!.first.id ?? "",
                  icon: const Icon(Icons.arrow_downward),
                  elevation: 16,
                  style: const TextStyle(color: Colors.deepPurple),
                  underline: Container(
                    height: 2,
                    color: Colors.deepPurpleAccent,
                  ),
                  onChanged: (String? value) {
                    // This is called when the user selects an item.
                    setState(() {
                      dropdownValue = value!;
                    });
                  },
                  items: addresses?.map<DropdownMenuItem<String>>((Address address) {
                    return DropdownMenuItem<String>(
                      value: address.id,
                      child: Text(address.address ?? ""),
                    );
                  }).toList(),
                );
              }
              else {
                return const CircularProgressIndicator();
              }
            }
          );
        } else {
          return const CircularProgressIndicator();
        }
      }
    );
  }
}

