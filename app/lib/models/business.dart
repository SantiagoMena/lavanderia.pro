class Business {
  final String? id;
  final String? name;
  final String? createdAt;
  final String? updatedAt;

  Business({this.id, this.name, this.createdAt, this.updatedAt});

  factory Business.fromJson(Map<String, dynamic> json) {
    return Business(
      id: json['id'],
      name: json['name'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}