package handlers

import (
	"fmt"
	"log"

	"testing"

	"online_electronic_store/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// creds := fmt.Sprintf("user=%s password=%s dbname=oes_test sslmode=disable",
	// 	os.Getenv("USER"), os.Getenv("PASSWORD"))
	creds := fmt.Sprintf("host=localhost user=postgres password=Avi@2004 dbname=oes sslmode=disable")
	db, err := gorm.Open(postgres.Open(creds), &gorm.Config{})
	if err != nil {
		log.Fatal("Error occured: ", err.Error())
	}
	fmt.Println("db connection established")
	return db
}

func TestInsertDeliveryIntoDB(t *testing.T) {
	db := setupTestDB()
	deliveryTable := &DeliveryTable{dB: db}

	delivery := models.Delivery{
		DeliveryStatus: "Cancelled",
		OrderID:        5,
	}

	err := deliveryTable.InsertDeliveryIntoDB(delivery)
	assert.NoError(t, err, "Received an error while inserting into the DB")

	var newDelivery models.Delivery
	result := db.Where("order_id = ?", 5).First(&newDelivery)
	assert.NoError(t, result.Error, "Didnt get the record")
	assert.Equal(t, "Cancelled", newDelivery.DeliveryStatus, "Delivery status not matching with the inserted value")
	assert.Equal(t, uint(5), newDelivery.OrderID, "order ID should match inserted value")
	assert.NotZero(t, newDelivery.CreatedAt, "Not updating the createdAt field when inserting new record")
}

func TestOrderDeliveryDetailsFromDB(t *testing.T) {
	db := setupTestDB()
	deliveryTable := &DeliveryTable{dB: db}
	db.Create(&models.Delivery{DeliveryStatus: "Delivered", OrderID: 1})
	db.Create(&models.Delivery{DeliveryStatus: "Shipped", OrderID: 2})
	result, err := deliveryTable.OrderDeliveryDetailsFromDB(1)
	assert.NoError(t, err, "Error occured while fetching the record")
	assert.Equal(t, uint(1), result.OrderID, "Wrong order_id value inserted into DB")
	assert.Equal(t, "Delivered", result.DeliveryStatus, "Inserted value and the expected value for delivery status is not same")
	assert.False(t, result.CreatedAt.IsZero(), "Not updating the createdAt field when inserting new record")

	// checking for order that doesn't exist
	_, err = deliveryTable.OrderDeliveryDetailsFromDB(123)
	assert.Error(t, err, "should return error for non-existing record")
}

func TestGetAllDeliveriesFromDB(t *testing.T) {
	db := setupTestDB()
	deliveryTable := &DeliveryTable{dB: db}
	db.Create(&models.Delivery{DeliveryStatus: "Delivered", OrderID: 10})
	db.Create(&models.Delivery{DeliveryStatus: "Shipped", OrderID: 5})
	db.Create(&models.Delivery{DeliveryStatus: "Pending", OrderID: 7})

	deliveries, err := deliveryTable.GetAllDeliveriesFromDB(3)
	assert.NoError(t, err, "Error while getting the valid user details")
	assert.Len(t, deliveries, 2, "UserId 3 has exactly 2 deliveries(in the database)")
	assert.ElementsMatch(t,
		[]string{"Dispatched", "Delivered"},
		[]string{deliveries[0].DeliveryStatus, deliveries[1].DeliveryStatus},
		"the delivery status of the user doesnt match")
	assert.Equal(t, "Dispatched", deliveries[0].DeliveryStatus, "Delivery status value mismatch")
	assert.Equal(t, uint(3), deliveries[0].OrderID, "OrderID value mismatch")

	deliveries, err = deliveryTable.GetAllDeliveriesFromDB(123)
	assert.NoError(t, err, "Didnt receive error when fething with the wrong usrId")
	assert.Len(t, deliveries, 0, "should get an empty slice for non-existing userID")
}
