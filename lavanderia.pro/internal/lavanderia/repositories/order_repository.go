package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var orderCollection = "order"

type OrderRepository struct {
	database databases.Database
}

func NewOrderRepository(database databases.Database) *OrderRepository {
	return &OrderRepository{
		database: database,
	}
}

func (orderRepository *OrderRepository) FindAllOrder() ([]types.Order, error) {
	orderDb, err := orderRepository.database.FindAll(orderCollection)

	if err != nil {
		return nil, err
	}

	var orderMap []types.Order
	if err = orderDb.All(context.TODO(), &orderMap); err != nil {
		panic(err)
	}

	return orderMap, nil
}

func (orderRepository *OrderRepository) Create(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.CreatedAt = &t

	orderDb, err := orderRepository.database.Create("order", bson.D{
		{Key: "business", Value: order.Business},
		{Key: "products", Value: order.Products},
		{Key: "client", Value: order.Client},
		{Key: "address", Value: order.Address},
		{Key: "created_at", Value: order.CreatedAt},
	})

	if err != nil {
		return types.Order{}, err
	}

	insertedId := orderDb.InsertedID.(primitive.ObjectID).Hex()

	// TODO: find inserted and return
	newOrder := types.Order{
		ID:        insertedId,
		Business:  order.Business,
		Client:    order.Client,
		Address:   order.Address,
		Products:  order.Products,
		CreatedAt: order.CreatedAt,
	}

	return newOrder, nil
}

func (orderRepository *OrderRepository) GetAllOrdersByBusiness(business string) ([]types.Order, error) {
	businessId, errBusinessId := primitive.ObjectIDFromHex(business)

	if errBusinessId != nil {
		return []types.Order{}, errors.New("wrong business id")
	}

	ordersDb, err := orderRepository.database.FindAllFilter(orderCollection, bson.D{
		{Key: "business", Value: businessId},
		{Key: "deleted_at", Value: nil},
	})

	if err != nil {
		return []types.Order{}, err
	}

	var ordersMap []types.Order
	if err = ordersDb.All(context.TODO(), &ordersMap); err != nil {
		return []types.Order{}, err
	}

	return ordersMap, nil
}

func (orderRepository *OrderRepository) Delete(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.DeletedAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "accepted_at", Value: nil},
		{Key: "rejected_at", Value: nil},
		{Key: "assigned_pickup_at", Value: nil},
		{Key: "pickup_client_at", Value: nil},
		{Key: "processing_at", Value: nil},
		{Key: "finished_at", Value: nil},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: order.DeletedAt}}}}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var deletedOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &deletedOrder)
	deletedOrder.DeletedAt = order.DeletedAt

	return deletedOrder, nil
}

func (orderRepository *OrderRepository) Get(order *types.Order) (types.Order, error) {
	id, errorId := primitive.ObjectIDFromHex(order.ID)

	if errorId != nil {
		return types.Order{}, errorId
	}

	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "deleted_at", Value: nil},
	}

	objectOrder, err := orderRepository.database.FindOne(orderCollection, filter)

	if err != nil {
		return types.Order{}, err
	}

	var foundOrder types.Order

	objectUpdt, _ := bson.Marshal(objectOrder)
	bson.Unmarshal(objectUpdt, &foundOrder)

	return foundOrder, nil
}

func (orderRepository *OrderRepository) Accept(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.AcceptedAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "rejected_at", Value: nil},
		{Key: "assigned_pickup_at", Value: nil},
		{Key: "pickup_client_at", Value: nil},
		{Key: "processing_at", Value: nil},
		{Key: "finished_at", Value: nil},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "accepted_at", Value: order.AcceptedAt}}}}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var acceptedOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &acceptedOrder)
	acceptedOrder.AcceptedAt = order.AcceptedAt

	return acceptedOrder, nil
}

func (orderRepository *OrderRepository) Reject(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.RejectedAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "rejected_at", Value: nil},
		{Key: "assigned_pickup_at", Value: nil},
		{Key: "pickup_client_at", Value: nil},
		{Key: "processing_at", Value: nil},
		{Key: "finished_at", Value: nil},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "rejected_at", Value: order.RejectedAt}}}}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var rejectedOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &rejectedOrder)
	rejectedOrder.RejectedAt = order.RejectedAt

	return rejectedOrder, nil
}

func (orderRepository *OrderRepository) AssignPickUp(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.AssignedPickUpAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "accepted_at", Value: bson.M{"$ne": nil}},
		{Key: "rejected_at", Value: nil},
		{Key: "assigned_pickup_at", Value: nil},
		{Key: "pickup_client_at", Value: nil},
		{Key: "processing_at", Value: nil},
		{Key: "finished_at", Value: nil},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "assigned_pickup_at", Value: order.AssignedPickUpAt},
				{Key: "pickup", Value: order.PickUp},
			},
		},
	}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var assignedPickUpOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &assignedPickUpOrder)
	assignedPickUpOrder.AssignedPickUpAt = order.AssignedPickUpAt
	assignedPickUpOrder.PickUp = order.PickUp

	return assignedPickUpOrder, nil
}

func (orderRepository *OrderRepository) PickUpClient(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.PickUpClientAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "accepted_at", Value: bson.M{"$ne": nil}},
		{Key: "assigned_pickup_at", Value: bson.M{"$ne": nil}},
		{Key: "rejected_at", Value: nil},
		{Key: "pickup_client_at", Value: nil},
		{Key: "processing_at", Value: nil},
		{Key: "finished_at", Value: nil},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "pickup_client_at", Value: order.PickUpClientAt}}}}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var rejectedOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &rejectedOrder)
	rejectedOrder.PickUpClientAt = order.PickUpClientAt

	return rejectedOrder, nil
}

func (orderRepository *OrderRepository) Process(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.ProcessingAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "accepted_at", Value: bson.M{"$ne": nil}},
		{Key: "assigned_pickup_at", Value: bson.M{"$ne": nil}},
		{Key: "rejected_at", Value: nil},
		{Key: "pickup_client_at", Value: bson.M{"$ne": nil}},
		{Key: "processing_at", Value: nil},
		{Key: "finished_at", Value: nil},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "processing_at", Value: order.ProcessingAt}}}}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var processedOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &processedOrder)
	processedOrder.ProcessingAt = order.ProcessingAt

	return processedOrder, nil
}

func (orderRepository *OrderRepository) Finish(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.FinishedAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "accepted_at", Value: bson.M{"$ne": nil}},
		{Key: "assigned_pickup_at", Value: bson.M{"$ne": nil}},
		{Key: "rejected_at", Value: nil},
		{Key: "pickup_client_at", Value: bson.M{"$ne": nil}},
		{Key: "processing_at", Value: bson.M{"$ne": nil}},
		{Key: "finished_at", Value: nil},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "finished_at", Value: order.FinishedAt}}}}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var finishedOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &finishedOrder)
	finishedOrder.FinishedAt = order.FinishedAt

	return finishedOrder, nil
}

func (orderRepository *OrderRepository) AssignDelivery(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.AssignedDeliveryAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "accepted_at", Value: bson.M{"$ne": nil}},
		{Key: "rejected_at", Value: nil},
		{Key: "assigned_pickup_at", Value: bson.M{"$ne": nil}},
		{Key: "pickup_client_at", Value: bson.M{"$ne": nil}},
		{Key: "processing_at", Value: bson.M{"$ne": nil}},
		{Key: "finished_at", Value: bson.M{"$ne": nil}},
		{Key: "assigned_delivery_at", Value: nil},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "assigned_delivery_at", Value: order.AssignedDeliveryAt},
				{Key: "delivery", Value: order.Delivery},
			},
		},
	}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var assignedDeliveryOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &assignedDeliveryOrder)
	assignedDeliveryOrder.AssignedDeliveryAt = order.AssignedDeliveryAt
	assignedDeliveryOrder.Delivery = order.Delivery

	return assignedDeliveryOrder, nil
}

func (orderRepository *OrderRepository) PickUpBusiness(order *types.Order) (types.Order, error) {
	t := time.Now()
	order.PickUpBusinessAt = &t

	id, _ := primitive.ObjectIDFromHex(order.ID)

	// delete only if all other status are empties
	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "accepted_at", Value: bson.M{"$ne": nil}},
		{Key: "assigned_pickup_at", Value: bson.M{"$ne": nil}},
		{Key: "rejected_at", Value: nil},
		{Key: "pickup_client_at", Value: bson.M{"$ne": nil}},
		{Key: "processing_at", Value: bson.M{"$ne": nil}},
		{Key: "finished_at", Value: bson.M{"$ne": nil}},
		{Key: "assigned_delivery_at", Value: bson.M{"$ne": nil}},
		{Key: "pickup_business_at", Value: nil},
		{Key: "delivered_client_at", Value: nil},
		{Key: "deleted_at", Value: nil},
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "pickup_business_at", Value: order.PickUpBusinessAt}}}}

	objectUpdated, err := orderRepository.database.UpdateOne(orderCollection, filter, update)
	if err != nil {
		return types.Order{}, err
	}

	var assignedOrder types.Order

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &assignedOrder)
	assignedOrder.PickUpBusinessAt = order.PickUpBusinessAt

	return assignedOrder, nil
}
