package repositories

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"lavanderia.pro/api/types"
	"lavanderia.pro/internal/lavanderia/config"
	"lavanderia.pro/internal/lavanderia/databases"
)

func TestCreateOrder(t *testing.T) {
	if err := godotenv.Load("../../../.env.test"); err != nil {
		fmt.Println("No .env.test file found")
	}

	authRepository := MakeAuthRepositoryForTestOrder()
	businessRepository := MakeBusinessRepositoryTestOrder()
	deliveryRepository := MakeDeliveryRepositoryTestOrder()
	clientRepository := MakeClientRepositoryTestOrder()
	addressRepository := MakeAddressRepositoryTestOrder()
	productRepository := MakeProductRepositoryTestOrder()
	orderRepository := MakeOrderRepositoryTestOrder()

	// Register Business
	pwd := []byte("PwD")
	password, errPass := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPass, nil, "GenerateFromPassword() returns error")

	email := []string{"new@", MakeRandomString(7), "test.com"}

	authBusiness, errAuthBusiness := authRepository.Create(&types.Auth{
		Email:    strings.Join(email, ""),
		Password: string(password),
	})

	assert.Equal(t, nil, errAuthBusiness, "error on create business auth")
	assert.NotEmpty(t, authBusiness, "error on create business auth")

	business, errBussiness := businessRepository.Create(&types.Business{
		Auth: authBusiness.ID,
		Name: "Test",
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
	})

	assert.Equal(t, nil, errBussiness, "error on create business")
	assert.NotEmpty(t, business, "error on create business")

	// Register Delivery
	passwordDelivery, errPassDelivery := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPassDelivery, nil, "GenerateFromPassword() returns error")

	emailDelivery := []string{"new@", MakeRandomString(7), "test.com"}

	authDelivery, errAuthDelivery := authRepository.Create(&types.Auth{
		Email:    strings.Join(emailDelivery, ""),
		Password: string(passwordDelivery),
	})

	assert.Equal(t, nil, errAuthDelivery, "error on create delivery auth")
	assert.NotEmpty(t, authDelivery, "error on create delivery auth")

	delivery, errDelivery := deliveryRepository.Create(&types.Delivery{
		Name: "Test",
		Auth: authDelivery.ID,
	})

	assert.Equal(t, nil, errDelivery, "error on create delivery")
	assert.NotEmpty(t, delivery, "error on create delivery")
	assert.NotEmpty(t, delivery.CreatedAt, "error on create delivery CreatedAt empty")

	// Register Client
	passwordClient, errPassClient := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	assert.Equal(t, errPassClient, nil, "GenerateFromPassword() returns error")

	emailClient := []string{"new@", MakeRandomString(7), "test.com"}

	authClient, errAuthClient := authRepository.Create(&types.Auth{
		Email:    strings.Join(emailClient, ""),
		Password: string(passwordClient),
	})

	assert.Equal(t, nil, errAuthClient, "error on create client auth")
	assert.NotEmpty(t, authClient, "error on create client auth")

	client, errClient := clientRepository.Create(&types.Client{
		Name: "Test",
		Auth: authClient.ID,
	})

	assert.Equal(t, nil, errClient, "error on create client")
	assert.NotEmpty(t, client, "error on create client")

	// Create Address
	address, errAddress := addressRepository.Create(&types.Address{
		Client: client.ID,
		Position: types.Geometry{
			Type:        "Point",
			Coordinates: []float64{-71.327767, -41.138444},
		},
		Name:    "TEST_ORDER",
		Extra:   "Call me",
		Phone:   "+123123123",
		Address: "Av. Pioneros 201, S.C Bariloche, Argentina",
	})

	assert.Equal(t, nil, errAddress, "error on create address")
	assert.NotEmpty(t, address, "error on create address")
	assert.NotEmpty(t, address.CreatedAt, "error on create address CreatedAt empty")

	// Make Product
	product, errProduct := productRepository.Create(&types.Product{
		Name:     "test_product",
		Price:    0.123,
		Business: business.ID,
	})

	assert.Equal(t, nil, errProduct, "error on create product")
	assert.NotEmpty(t, product, "error on create product")

	// Create product list
	productList := &[]types.OrderProduct{
		{Product: product, Amount: 7},
	}

	// Create Order
	order, errOrder := orderRepository.Create(&types.Order{
		Products: *productList,
		Client:   client,
		Business: business,
		Delivery: delivery,
	})

	assert.Equal(t, nil, errOrder, "error on create order")
	assert.NotEmpty(t, order, "error on create order")
}

func MakeAuthRepositoryForTestOrder() *AuthRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	authRepository := NewAuthRepository(mongo, config)

	return authRepository
}

func MakeBusinessRepositoryTestOrder() *BusinessRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	businessRepository := NewBusinessRepository(mongo)

	return businessRepository
}

func MakeDeliveryRepositoryTestOrder() *DeliveryRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	deliveryRepository := NewDeliveryRepository(mongo)

	return deliveryRepository
}

func MakeClientRepositoryTestOrder() *ClientRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	clientRepository := NewClientRepository(mongo)

	return clientRepository
}

func MakeAddressRepositoryTestOrder() *AddressRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	addressRepository := NewAddressRepository(mongo)

	return addressRepository
}

func MakeProductRepositoryTestOrder() *ProductRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	productRepository := NewProductRepository(mongo)

	return productRepository
}

func MakeOrderRepositoryTestOrder() *OrderRepository {
	config := config.NewConfig()
	mongo := databases.NewMongoDatabase(config)
	orderRepository := NewOrderRepository(mongo)

	return orderRepository
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func MakeRandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
