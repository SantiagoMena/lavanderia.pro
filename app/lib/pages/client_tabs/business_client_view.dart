import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/models/order.dart';
import 'package:lavanderiapro/pages/business_tabs/products_business_view.dart';
import 'package:lavanderiapro/pages/client_tabs/check_order_client_view.dart';
import 'package:lavanderiapro/pages/client_tabs/processed_order_client_view.dart';
import 'package:lavanderiapro/services/get_all_products_business_service.dart';
import 'package:lavanderiapro/services/search_business_service.dart';
import 'package:lavanderiapro/models/product.dart';
import 'package:lavanderiapro/widgets/ProductCardCart.dart';

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
           order.setBusinessId(widget.businessItem!.id ?? "");
           Navigator.push(
               context,
               MaterialPageRoute(
                   builder: (context) => CheckOrderClient(order: order)
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
                                return ProductCardCart(
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
                       title:  Text(AppLocalizations.of(context)!.productsSelectedLabel(order.count)),
                       subtitle: Text(AppLocalizations.of(context)!.showFullOrderLabel),
                       children: getOrderProducts()
                    ),
                  ],
              ),
            ),
        );
      }
    )
   );
  }
}
