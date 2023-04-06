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

func (businessRepository *ProductRepository) GetAllProductsByBusiness(business string) ([]types.Product, error) {
	// businessMap := []types.Business{}

	fmt.Println("business")
	fmt.Println(business)
	businessId, _ := primitive.ObjectIDFromHex(business)

	fmt.Println("businessId")
	fmt.Println(businessId)
	businessDb, err := businessRepository.database.FindAllFilter(productCollection, bson.D{
		{Key: "business", Value: businessId},
	})

	if err != nil {
		return nil, err
	}

	var businessMap []types.Product
	if err = businessDb.All(context.TODO(), &businessMap); err != nil {
		panic(err)
	}

	return businessMap, nil
}
