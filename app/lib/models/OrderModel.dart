import 'dart:collection';
import 'package:lavanderiapro/services/get_all_products_business_service.dart';

class OrderModel {
  /// Internal, private state of the cart.
  final List<Product> _items = [];

  /// An unmodifiable view of the items in the cart.
  UnmodifiableListView<Product> get items => UnmodifiableListView(_items);

  /// The current total price of all items (assuming all items cost $42).
  num get totalPrice => _items.isNotEmpty ? _items.map((e) => e.price).reduce((value, element) => element != null ? value! + element : value) ?? 0 : 0;

  int get count => _items.length;

  /// Adds [item] to cart. This and [removeAll] are the only ways to modify the
  void add(Product item) {
    _items.add(item);
  }

  /// Removes all items from the cart.
  void removeAll() {
    _items.clear();
  }

  void remove(Product item) {
    _items.remove(item);
  }

  List<Product> getGrouped() {
    List<Product> group = List<Product>.empty(growable: true);

    _items.forEach((element) {
      Iterable<Product> elementFound = group.where((elementGroup) => elementGroup.id == element.id);
      if(elementFound.isEmpty) {
        group.add(element);
      }
    });

    return group;
  }
}