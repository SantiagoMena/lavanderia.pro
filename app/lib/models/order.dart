import 'dart:collection';
import 'package:lavanderiapro/models/product.dart';

class OrderModel {
  /// Internal, private state of the cart.
  final List<Product> _items = [];

  late String businessId = "";

  late String addressId = "";

  /// An unmodifiable view of the items in the cart.
  UnmodifiableListView<Product> get items => UnmodifiableListView(_items);

  /// The current total price of all items (assuming all items cost $42).
  num get totalPrice => _items.isNotEmpty ? _items.map((e) => e.price).reduce((value, element) => element != null ? value! + element : value) ?? 0 : 0;

  int get count => _items.length;

  void setBusinessId(String id){
    businessId = id;
  }

  void setAddressId(String id){
    addressId = id;
  }

  /// Adds [item] to cart. This and [removeAll] are the only ways to modify the
  void add(Product item) {
    _items.add(item);
  }

  /// Removes all items from the cart.
  void removeAll() {
    _items.clear();
  }

  void remove(Product item) {
    for(final element in _items){
      if(element.id == item.id){
        _items.remove(element);
        break;
      }
    }
  }

  void delete(Product item) {
    for(final element in _items){
      if(element.id == item.id){
        _items.remove(element);
      }
    }
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

  int countProduct(Product item) {
    int counter = 0;

    _items.forEach((element) {
      if(element.id == item.id){
        counter++;
      }
    });

    return counter;
  }
}