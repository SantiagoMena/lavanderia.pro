import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/models/OrderModel.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';
import 'package:lavanderiapro/services/get_all_products_business_service.dart';
import 'package:lavanderiapro/services/search_business_service.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:provider/provider.dart';

class BusinessClientView extends StatefulWidget {
   BusinessClientView({super.key, this.token, this.businessItem});
   final Business? businessItem;
   final String? token;

  @override
  State<BusinessClientView> createState() => _BusinessClientViewState();
}

class _BusinessClientViewState extends State<BusinessClientView> {
   OrderModel order = OrderModel();

   void pushProduct(Product product) {
     setState(() {
       order.add(product);
     });
     print(order.totalPrice);
   }

   void popProduct(Product product) {
     setState(() {
       order.remove(product);
     });
     print(order.totalPrice);
   }

   List<Widget> getOrderProducts() {
     List<Widget> widgets = List<Widget>.empty(growable: true);

     order.getGrouped().forEach((element) {
       widgets.add(
         ListTile(
             title: Row(
               children: [
                 Expanded(child: Text(element!.name ?? "")),
                 Text("x ${order.countProduct(element)}" ?? "x 0"),
                 Text(" "),
                 Text("\$ ${(element.price ?? 0) * order.countProduct(element)}" ?? "0"),
               ],
             )
         )
       );
     });

     widgets.add(
       ListTile(
         title: Row(
           children: [
             const Align(
               alignment: Alignment.topLeft,
               child: Text(
                 '☝️ Process Order',
                 style: TextStyle(color: Colors.white),
               ),
             ),
             const Expanded(child: Text("")),
             Align(
               alignment: Alignment.topRight,
               child: Text(
                 '\$${order.totalPrice}',
                 style: const TextStyle(color: Colors.white),
               ),
             ),
           ],
         ),
         tileColor: Colors.green,
         onTap: () {
           Navigator.push(
               context,
               MaterialPageRoute(
                   builder: (context) =>
                       ProcessedOrderClient()
               )
           );
         },
       )
     );

     return widgets;
   }

   @override
  Widget build(BuildContext context) {
     return Scaffold(
          appBar: AppBar(
          title: Text(widget.businessItem!.name ?? ""),
        ),
        body: LayoutBuilder(
          builder: (BuildContext context, BoxConstraints viewportConstraints) {
            return
              Align(
              alignment: Alignment.bottomCenter,
              child: SingleChildScrollView(
                reverse: true,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  mainAxisAlignment: MainAxisAlignment.start,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    SizedBox(
                      height: viewportConstraints.maxHeight - AppBar().preferredSize.height,
                      child: FutureBuilder(
                        future: getAllProductsBusiness(widget.token ?? "", widget.businessItem!.id ?? ""),
                        builder: (context, snapshot) {
                          if(snapshot.hasData) {
                            return ListView.builder(
                              itemCount: snapshot.data!.length ?? 0,
                              itemBuilder: (context, index) {
                                var productItem = snapshot.data![index];
                                return ProductCard(
                                    productItem: productItem,
                                    pushProductCallback: pushProduct,
                                    popProductCallback: popProduct,
                                    countProductCallback: order.countProduct,
                                );
                              },
                            );
                          } else {
                            return const CircularProgressIndicator();
                          }
                        },
                      ),
                    ),
                     ExpansionTile(
                       title:  Text('${order.count} Product Selected'),
                       subtitle: Text('Show full order'),
                       children: getOrderProducts()
                    ),
                    // Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
                  ],
              ),
            ),
        );
      }
    )
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
    required this.productItem,
    required this.pushProductCallback,
    required this.popProductCallback,
    required this.countProductCallback,
  });

  final Product productItem;
  final Function pushProductCallback;
  final Function popProductCallback;
  final Function countProductCallback;

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
                  Container(child: Text(productItem.name ?? "")),
                  Expanded(child: Text("")),
                ]
              ),
              Row(
                  children: [
                    Container(child: Text("Desc ...")),
                    Expanded(child: Text("")),
                  ]
              ),
              Row(
                  children: [
                    Container(child: Text('Price: \$0000')),
                    Expanded(child: Text("")),
                    Container(child: Row(
                      children: [
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: ElevatedButton(
                            onPressed: () {
                              popProductCallback(productItem);
                            },
                            child: Text("➖")
                          ),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: Text(countProductCallback(productItem).toString()),
                        ),
                        Padding(
                          padding: const EdgeInsets.all(8.0),
                          child: ElevatedButton(
                              onPressed: () {
                                pushProductCallback(productItem);
                              },
                              child: Text("➕")
                          ),
                        ),
                      ],
                    )),
                  ]
              ),
            ],
          ),
        ),
      )
    );
  }
}
