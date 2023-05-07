import 'package:flutter/material.dart';
import 'package:lavanderiapro/models/product.dart';

class ProductCardCart extends StatelessWidget {
  const ProductCardCart({
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
