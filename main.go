package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofrs/uuid/v5"
)

type Trainer struct {
	UUIDv7       string
	UUIDv4       string
	Name         string
	Age          int
	City         string
	NumericID    int
	NumericIDStr string
	Id           primitive.ObjectID `bson:"_id,omitempty"`
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var totalNoOfRecordsToRead = 100000
var maxRecords = 100000000
var oneMillion = 1000000

func generateRandomTrainer(numericID int) *Trainer {
	uuidV7, _ := uuid.NewV7()
	uuidV4, _ := uuid.NewV4()
	name := RandomString(20)
	age := seededRand.Intn(100)
	city := RandomString(20)
	return &Trainer{uuidV7.String(), uuidV4.String(), name, age, city, numericID, strconv.FormatInt(int64(numericID), 10), primitive.NilObjectID}
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func generateBulkAndInsert(collection *mongo.Collection, count int, startNumericID int) {
	trainers := []interface{}{}
	for i := 0; i < count; i++ {
		trainers = append(trainers, generateRandomTrainer(startNumericID+i))
	}
	_, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteAllRecords(collection *mongo.Collection) {
	// Delete all the documents in the collection
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	}
}

func FindByUUIDv7(collection *mongo.Collection, uuidv7 string) Trainer {
	var result Trainer
	filter := bson.D{{"uuidv7", uuidv7}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("Record with uuidv7 %v not found", uuidv7)
		log.Fatal(err)
	}
	return result
}

func FindByMongoID(collection *mongo.Collection, id primitive.ObjectID) Trainer {
	var result Trainer
	filter := bson.D{{"_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("Record with mongoID %v not found", id)
		log.Fatal(err)
	}
	return result
}

func FindByUUIDv4(collection *mongo.Collection, uuidv4 string) Trainer {
	var result Trainer
	filter := bson.D{{"uuidv4", uuidv4}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("Record with uuidV4 %v not found", uuidv4)
		log.Fatal(err)
	}
	return result
}

func FindByNumID(collection *mongo.Collection, numericID int) Trainer {
	// Find a single document
	var result Trainer
	filter := bson.D{{"numericid", numericID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("Record with numericID %v not found", numericID)
		log.Fatal(err)
	}
	return result
}
func FindByNumIDString(collection *mongo.Collection, numIDString string) Trainer {
	// Find a single document
	var result Trainer
	filter := bson.D{{"numericidstr", numIDString}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Printf("Record with numIDString %v not found", numIDString)
		log.Fatal(err)
	}
	return result
}
func setupCollection() (*mongo.Collection, *mongo.Client) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Connected to MongoDB!")
	collection := client.Database("test").Collection("trainers")
	return collection, client
}

func insert100MillionRecords(collection *mongo.Collection) {
	count := 1000
	startNumericID := 0
	for i := 0; i < 100000; i++ {
		generateBulkAndInsert(collection, count, startNumericID)
		startNumericID += count
	}
}

func generateRandomFile() {
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())

	randomRecords := make([]Trainer, totalNoOfRecordsToRead, totalNoOfRecordsToRead)

	for i := 0; i < totalNoOfRecordsToRead; i++ {
		numericID := seededRand.Intn(maxRecords - 1)
		randomRecords[i] = FindByNumID(collection, numericID)
	}
	writeToFile("randomFile.json", randomRecords)
}

func generateNewerRecordsFile() {
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())

	randomRecords := make([]Trainer, totalNoOfRecordsToRead, totalNoOfRecordsToRead)

	for i := 0; i < totalNoOfRecordsToRead; i++ {
		numericID := maxRecords - seededRand.Intn(oneMillion-1)
		randomRecords[i] = FindByNumID(collection, numericID)
	}
	writeToFile("newerRecordsFile.json", randomRecords)
}

func generateOlderRecordsFile() {
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())

	randomRecords := make([]Trainer, totalNoOfRecordsToRead, totalNoOfRecordsToRead)

	for i := 0; i < totalNoOfRecordsToRead; i++ {
		numericID := seededRand.Intn(oneMillion - 1)
		randomRecords[i] = FindByNumID(collection, numericID)
	}
	writeToFile("olderRecordsFile.json", randomRecords)
}

func writeToFile(filename string, records []Trainer) {

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	defer w.Flush()
	defer f.Close()
	for _, trainer := range records {
		t, _ := json.Marshal(trainer)
		w.Write(t)
		w.WriteString("\n")
	}

}

func readFile(filename string) []Trainer {
	records := make([]Trainer, totalNoOfRecordsToRead, totalNoOfRecordsToRead)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	i := 0
	for scanner.Scan() {
		var record Trainer
		json.Unmarshal([]byte(scanner.Text()), &record)
		records[i] = record
		i++
	}
	return records
}

func readRandomFile() []Trainer {
	return readFile("randomFile.json")
}

func readNewerRecordsfile() []Trainer {
	return readFile("newerRecordsFile.json")
}

func readOlderRecordsFile() []Trainer {
	return readFile("olderRecordsFile.json")
}

func main() {
	collection, client := setupCollection()
	defer client.Disconnect(context.TODO())

	// This step takes time!!
	insert100MillionRecords(collection)

	// Following lines are more for testing!

	// element := FindByNumIDString(collection, "2")
	// fmt.Printf("FindByNumIDString : %+v\n", element)

	// element = FindByNumID(collection, 2)
	// fmt.Printf("FindByNumID : %+v\n", element)

	// element = FindByUUIDv4(collection, element.UUIDv4)
	// fmt.Printf("FindByUUIDv4 : %+v\n", element)

	// element = FindByUUIDv7(collection, element.UUIDv7)
	// fmt.Printf("FindByUUIDv7 : %+v\n", element)

	// element = FindByMongoID(collection, element.Id)
	// fmt.Printf("FindByMongoID : %+v\n", element)

	// Generate files required for benchmarking!
	generateOlderRecordsFile()
	generateNewerRecordsFile()
	generateRandomFile()

}
