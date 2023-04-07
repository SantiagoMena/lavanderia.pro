package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/databases"
)

var productCollection = "product"

type ProductRepository struct {
	database databases.Database
}

func NewProductRepository(database databases.Database) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}

func (productRepository *ProductRepository) FindAllProduct() ([]types.Product, error) {
	productDb, err := productRepository.database.FindAll(productCollection)

	if err != nil {
		return nil, err
	}

	var productMap []types.Product
	if err = productDb.All(context.TODO(), &productMap); err != nil {
		panic(err)
	}

	return productMap, nil
}

func (productRepository *ProductRepository) Create(product *types.Product) (types.Product, error) {
	t := time.Now()
	product.CreatedAt = &t

	businessId, _ := primitive.ObjectIDFromHex(product.Business)

	productDb, err := productRepository.database.Create("product", bson.D{
		{Key: "name", Value: product.Name},
		{Key: "price", Value: product.Price},
		{Key: "created_at", Value: product.CreatedAt},
		{Key: "business", Value: businessId},
	})

	if err != nil {
		return types.Product{}, err
	}

	insertedId := productDb.InsertedID.(primitive.ObjectID).Hex()

	newProduct := types.Product{
		ID:        insertedId,
		Name:      product.Name,
		Price:     product.Price,
		Business:  product.Business,
		CreatedAt: product.CreatedAt,
	}

	return newProduct, nil
}

func (productRepository *ProductRepository) GetAllProductsByBusiness(business string) ([]types.Product, error) {
	fmt.Println("business")
	fmt.Println(business)
	businessId, _ := primitive.ObjectIDFromHex(business)

	fmt.Println("businessId")
	fmt.Println(businessId)
	productsDb, err := productRepository.database.FindAllFilter(productCollection, bson.D{
		{Key: "business", Value: businessId},
		{Key: "deleted_at", Value: nil},
	})

	if err != nil {
		return []types.Product{}, err
	}

	var productsMap []types.Product
	if err = productsDb.All(context.TODO(), &productsMap); err != nil {
		return []types.Product{}, err
	}

	return productsMap, nil
}

func (productRepository *ProductRepository) Delete(product *types.Product) (types.Product, error) {
	t := time.Now()
	product.DeletedAt = &t

	id, _ := primitive.ObjectIDFromHex(product.ID)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"deleted_at", product.DeletedAt}}}}

	objectUpdated, err := productRepository.database.UpdateOne(productCollection, filter, update)
	if err != nil {
		return types.Product{}, err
	}

	var deletedProduct types.Product

	objectUpdt, _ := bson.Marshal(objectUpdated)
	bson.Unmarshal(objectUpdt, &deletedProduct)
	deletedProduct.DeletedAt = product.DeletedAt

	return deletedProduct, nil
}

func (productRepository *ProductRepository) Get(product *types.Product) (types.Product, error) {
	// t := time.Now()
	// product.UpdatedAt = &t

	id, _ := primitive.ObjectIDFromHex(product.ID)

	filter := bson.D{
		{Key: "_id", Value: id},
		{Key: "deleted_at", Value: nil},
	}

	objectProduct, err := productRepository.database.FindOne(productCollection, filter)

	if err != nil {
		return types.Product{}, err
	}

	var foundProduct types.Product

	objectUpdt, _ := bson.Marshal(objectProduct)
	bson.Unmarshal(objectUpdt, &foundProduct)

	return foundProduct, nil
}

func (productRepository *ProductRepository) Update(product *types.Product) (types.Product, error) {
	t := time.Now()
	product.UpdatedAt = &t

	id, _ := primitive.ObjectIDFromHex(product.ID)

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{
		{"name", product.Name},
		{"price", product.Price},
		{"updated_at", product.UpdatedAt},
	}}}

	updatedProduct, err := productRepository.database.UpdateOne(productCollection, filter, update)

	if err != nil {
		return types.Product{}, err
	}

	// Unmarshal
	var updatedProductUnmarshal types.Product
	productUpdatedObj, _ := bson.Marshal(updatedProduct)
	bson.Unmarshal(productUpdatedObj, &updatedProductUnmarshal)
	productId, _ := primitive.ObjectIDFromHex(updatedProductUnmarshal.ID)

	productUpdatedFound, errFind := productRepository.database.FindOne(productCollection, bson.D{
		{"_id", productId},
	})

	if errFind != nil {
		return types.Product{}, errFind
	}

	// Unmarshal
	var productUnmarshal types.Product
	productMarshal, _ := bson.Marshal(productUpdatedFound)
	bson.Unmarshal(productMarshal, &productUnmarshal)

	return productUnmarshal, nil
}
