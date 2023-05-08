import 'package:lavanderiapro/models/address.dart';
import 'package:lavanderiapro/models/business.dart';
import 'package:lavanderiapro/models/client.dart';
import 'package:lavanderiapro/models/delivery.dart';
import 'package:lavanderiapro/models/product_order.dart';

class Order {
  final String? id;
  final Business? business;
  final Client? client;
  final Address? address;
  final Delivery? pickup;
  final Delivery? delivery;
  final List<ProductOrder>? products;
  final String? createdAt;
  final String? updatedAt;
  final String? acceptedAt;
  final String? rejectedAt;
  final String? assignedPickupAt;
  final String? pickupClientAt;
  final String? assignedDeliveryClientAt;
  final String? processingAt;
  final String? finishedAt;
  final String? deliveredClientAt;

  Order({
    this.id,
    this.business,
    this.client,
    this.address,
    this.pickup,
    this.delivery,
    this.products,
    this.createdAt,
    this.updatedAt,
    this.acceptedAt,
    this.rejectedAt,
    this.assignedPickupAt,
    this.pickupClientAt,
    this.assignedDeliveryClientAt,
    this.processingAt,
    this.finishedAt,
    this.deliveredClientAt,
  });

  factory Order.fromJson(Map<String, dynamic> json) {
    return Order(
      id: json['id'],
      business: json['business'],
      client: json['client'],
      address: json['address'],
      pickup: json['pickup'],
      delivery: json['delivery'],
      products: json['products'],
      createdAt: json['created_at'],
      updatedAt: json['updated_at'],
      acceptedAt: json['accepted_at'],
      rejectedAt: json['rejected_at'],
      assignedPickupAt: json['assigned_pickup_at'],
      pickupClientAt: json['pickup_client_at'],
      assignedDeliveryClientAt: json['assigned_delivery_client_at'],
      processingAt: json['processing_at'],
      finishedAt: json['finished_at'],
      deliveredClientAt: json['delivered_client_at'],
    );
  }
}