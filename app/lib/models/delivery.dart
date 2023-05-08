class Delivery {
  final String? id;
  final String? name;
  final String? createdAt;
  final String? updatedAt;

  Delivery({this.id, this.name, this.createdAt, this.updatedAt});

  factory Delivery.fromJson(Map<String, dynamic> json) {
    return Delivery(
      id: json['id'],
      name: json['name'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}