class Product {
  final String? id;
  final String? name;
  final num? price;
  final String? createdAt;
  final String? updatedAt;

  Product({this.id, this.name, this.price, this.createdAt, this.updatedAt});

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
      id: json['id'],
      name: json['name'],
      price: json['price'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}