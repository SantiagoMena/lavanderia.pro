class Client {
  final String? id;
  final String? name;
  final String? createdAt;
  final String? updatedAt;

  Client({this.id, this.name, this.createdAt, this.updatedAt});

  factory Client.fromJson(Map<String, dynamic> json) {
    return Client(
      id: json['id'],
      name: json['name'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }
}