import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/models/address.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';
import 'package:lavanderiapro/services/get_address_client_service.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:lavanderiapro/services/post_order_service.dart';
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

  void setAddressId(String addressId) {
    setState(() {
      widget.order!.setAddressId(addressId);
    });
  }

  @override
  Widget build(BuildContext context) {

    return Scaffold(
      appBar: AppBar(
        title: Text(AppLocalizations.of(context)!.appBarCheckOrder("\$${widget.order!.totalPrice}")),
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
                      SizedBox(
                        height: 50,
                        child: Align(
                            alignment: Alignment.center,
                            child: DropdownAddress(setAddress: setAddressId)
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
                            child: FutureBuilder(
                              future: SharedPreferences.getInstance(),
                              builder: (context, snapshot) {
                                return ElevatedButton(
                                  style: ElevatedButton.styleFrom(
                                    minimumSize: const Size.fromHeight(50),
                                    backgroundColor: Colors.green,
                                  ),
                                  onPressed: () {
                                    if(_formKey.currentState!.validate() && widget.order!.count > 0 && snapshot.hasData){
                                      String token = snapshot.data!.getString('token') ?? "";
                                      if(widget.order!.addressId.isEmpty || widget.order!.addressId == ""){
                                        getAddressClient(token).then((addresses) {
                                          // TODO: Set default address
                                          widget.order!.setAddressId(addresses.first.id ?? "");
                                            postOrder(token, widget.order).then((order) {
                                              if(order != null) {
                                                Navigator.push(
                                                    context,
                                                    MaterialPageRoute(
                                                        builder: (context) => const ProcessedOrderClient()
                                                    )
                                                );
                                              } else {
                                                ScaffoldMessenger.of(context).showSnackBar(
                                                    const SnackBar(content: SnackBarErrorOnMakeOrderLabel())
                                                );
                                              }
                                            }).catchError((onError) {
                                              ScaffoldMessenger.of(context).showSnackBar(
                                                  const SnackBar(content: SnackBarErrorOnMakeOrderLabel())
                                              );
                                              return onError;
                                            });
                                        })
                                            .catchError((onError) => onError);
                                      } else {

                                        postOrder(token, widget.order).then((order) {
                                          if(order != null) {
                                            Navigator.push(
                                                context,
                                                MaterialPageRoute(
                                                    builder: (context) => const ProcessedOrderClient()
                                                )
                                            );
                                          } else {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                                const SnackBar(content: SnackBarErrorOnMakeOrderLabel())
                                            );
                                          }
                                        }).catchError((onError) {
                                          ScaffoldMessenger.of(context).showSnackBar(
                                              const SnackBar(content: SnackBarErrorOnMakeOrderLabel())
                                          );
                                          return onError;
                                        });
                                      }
                                    } else {
                                      ScaffoldMessenger.of(context).showSnackBar(
                                          const SnackBar(content: SnackBarErrorOnMakeOrderLabel())
                                      );
                                    }
                                  },
                                  child: Text(AppLocalizations.of(context)!.makeOrderLabel),
                                );
                              }
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

class SnackBarErrorOnMakeOrderLabel extends StatelessWidget {
  const SnackBarErrorOnMakeOrderLabel({
    super.key,
  });

  @override
  Widget build(BuildContext context) {
    return Text(AppLocalizations.of(context)!.snackBarErrorOnMakeOrder);
  }
}

class DropdownAddress extends StatefulWidget {
  const DropdownAddress({super.key,  this.setAddress});

  final Function? setAddress;

  @override
  State<DropdownAddress> createState() => _DropdownAddressState();
}

class _DropdownAddressState extends State<DropdownAddress> {
  String? dropdownValue;

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: SharedPreferences.getInstance(),
      builder: (context, snapshot) {
        if(snapshot.hasData) {
          return FutureBuilder(
            future: getAddressClient(snapshot.data!.getString('token') ?? ""),
            builder: (contextAddress, snapshotAddress) {
              if(snapshotAddress.hasData){
                return DropdownButton<String>(
                  autofocus: true,
                  value: dropdownValue ?? snapshotAddress.data!.first.id!,
                  icon: const Icon(Icons.arrow_downward),
                  elevation: 16,
                  style: const TextStyle(color: Colors.black),
                  underline: Container(
                    height: 2,
                    color: Colors.green,
                  ),
                  onChanged: (String? value) {
                    print(value);
                    // This is called when the user selects an item.
                    setState(() {
                      dropdownValue = value!;
                      widget.setAddress!(value);
                    });
                  },
                  items: snapshotAddress.data!.map<DropdownMenuItem<String>>((Address value) {
                    return DropdownMenuItem<String>(
                      value: value.id,
                      child: Text(value.address ?? ""),
                    );
                  }).toList(),
              );
              }
              else {
                return Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 8),
                  child: ElevatedButton(
                    style: ElevatedButton.styleFrom(
                      minimumSize: const Size.fromHeight(50),
                      backgroundColor: Colors.green,
                    ),
                    onPressed: () {
                      // TODO: Redirect to add address view
                      Navigator.pop(context);
                    },
                      child: Text(AppLocalizations.of(context)!.addAddressLabel)
                  ),
                );
              }
            }
          );
        }
        else {
          return CircularProgressIndicator();
        }
      }
    );
  }
}
