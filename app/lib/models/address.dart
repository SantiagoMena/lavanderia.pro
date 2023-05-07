class Address {
  final String? id;
  late String? name;
  late String? address;
  late String? client;
  late String? phone;
  late String? extra;
  late String? createdAt;
  late String? updatedAt;

  Address({this.id, this.name, this.address, this.client, this.phone, this.extra, this.createdAt, this.updatedAt});

  factory Address.fromJson(Map<String, dynamic> json) {
    return Address(
      id: json['id'],
      name: json['name'],
      address: json['address'],
      client: json['client'],
      phone: json['phone'],
      extra: json['extra'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
    );
  }

  void setName(String nameNew) {
    name = nameNew;
  }

  void setAddress(String addressNew) {
    address = addressNew;
  }

  void setPhone(String phoneNew) {
    phone = phoneNew;
  }

  void setExtra(String extraNew) {
    extra = extraNew;
  }
}