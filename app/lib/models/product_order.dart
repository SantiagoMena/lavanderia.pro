import 'package:lavanderiapro/models/product.dart';

class ProductOrder {
  final Product? product;
  final num? amount;

  ProductOrder({this.product, this.amount});

  factory ProductOrder.fromJson(Map<String, dynamic> json) {
    return ProductOrder(
      product: json['product'],
      amount: json['amount'],
    );
  }
}