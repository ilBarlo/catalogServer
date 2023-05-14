package catalog

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://ilBarlo:FlavourGenerator@clusterbarlo.qnlqwmd.mongodb.net/?retryWrites=true&w=majority"
const dbName = "flavours"
const colName = "resources"

var collection *mongo.Collection

// init used for connecting with mongoDB
func init() {
	// Client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to mongodb
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection istance is ready")
}

func getAllFlavours() ([]Flavour, error) {

	// Define the filter for the query
	filter := bson.M{}

	// Perform the query
	var flavours []Flavour
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// Iterate through the query results
	for cur.Next(context.Background()) {
		var f Flavour
		if err := cur.Decode(&f); err != nil {
			return nil, err
		}
		flavours = append(flavours, f)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return flavours, nil
}

func getFlavoursByArchitecture(architecture string) ([]Flavour, error) {

	// Define the filter for the query
	filter := bson.M{"architecture": architecture}

	// Perform the query
	var flavours []Flavour
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// Iterate through the query results
	for cur.Next(context.Background()) {
		var f Flavour
		if err := cur.Decode(&f); err != nil {
			return nil, err
		}
		flavours = append(flavours, f)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return flavours, nil
}

func getFlavoursByOS(os string) ([]Flavour, error) {

	// Define the filter for the query
	filter := bson.M{"operatingsystem": os}

	// Perform the query
	var flavours []Flavour
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// Iterate through the query results
	for cur.Next(context.Background()) {
		var f Flavour
		if err := cur.Decode(&f); err != nil {
			return nil, err
		}
		flavours = append(flavours, f)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return flavours, nil
}

func findMatchingResources(request Request) ([]Flavour, error) {
	flavours, err := getAllFlavours()
	if err != nil {
		return nil, err
	}

	matchingFlavours := []Flavour{}

	for _, flavour := range flavours {
		if request.Architecture != "" && request.Architecture != flavour.Architecture {
			continue
		}
		if request.OS != "" && request.OS != flavour.OperatingSystem {
			continue
		}
		if request.CPURequest != "" && request.CPURequest > flavour.CPUOffer {
			continue
		}
		if request.MemoryRequest != "" && request.MemoryRequest > flavour.MemoryOffer {
			continue
		}
		if request.PodsPlan != "" {
			foundPlan := false
			for _, plan := range flavour.PodsOffer {
				if plan.Name == request.PodsPlan && plan.Available {
					foundPlan = true
					break
				}
			}
			if !foundPlan {
				continue
			}
		}

		matchingFlavours = append(matchingFlavours, flavour)
	}

	return matchingFlavours, nil
}
