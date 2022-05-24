package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for storing data
type product struct {
	ProductId          int            `json:"pid"`
	ProductName        string         `json:"pname"`
	ProductDescription string         `json:"pdes"`
	ProductCreatedBy   string         `json:"pcb"`
	ProductModifiedBy  string         `json:"pmb"`
	ProductQuantity    int            `json:"pq"`
	ProductMRP         float32        `json:"pmrp"`
	ProductStatus      bool           `json:"pstatus`
	ProductFixedPrice  float32        `json:"pfp"`
	Product            map[string]int `json:"product"`
}

var productCollection = db().Database("useCase").Collection("product") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var product product
	err := json.NewDecoder(r.Body).Decode(&product) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	err1 := productCollection.FindOne(context.TODO(), bson.D{{"productid", product.ProductId}}).Decode(&result)
	fmt.Println("err", err1, "result", result)
	if result == nil {
		insertResult, err := userCollection.InsertOne(context.TODO(), product)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document: ", insertResult)
		json.NewEncoder(w).Encode("Product Added uccessfully")
	} else {
		json.NewEncoder(w).Encode("Record cannot be inserted, Product Id already exits")
	}

}

// Get Profile of a particular User by Name

func GetUserProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body product
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := productCollection.FindOne(context.TODO(), bson.D{{"productname", body.ProductName}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

func GetProductId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string
	s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err1 := productCollection.FindOne(context.TODO(), bson.D{{"productid", s}}).Decode(&result)
	fmt.Println(result)
	fmt.Println(s)
	if err1 == nil {
		json.NewEncoder(w).Encode(result) // returns a Map containing document
	} else {
		json.NewEncoder(w).Encode("No record found")
	}

}

func GetProductName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"] //get Parameter value as string
	// s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err1 := productCollection.FindOne(context.TODO(), bson.D{{"productname", params}}).Decode(&result)
	fmt.Println(result)
	fmt.Println(params)
	if err1 == nil {
		json.NewEncoder(w).Encode(result) // returns a Map containing document
	} else {
		json.NewEncoder(w).Encode("No record found")
	}
}

//Update Profile of User

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		ProductName        string  `json:"pname"` //value that has to be matched
		ProductDescription string  `json:"pdes"`  // value that has to be modified
		ProductId          int     `json:"pid"`
		ProductModifiedBy  string  `json:"pmb"`
		ProductQuantity    int     `json:"pq"`
		ProductStatus      bool    `json:"pstatus"`
		ProductMRP         float32 `json:"pmrp`
		ProductFixedPrice  float32 `json:"pfp"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"productname", body.ProductName}} // converting value to BSON type
	after := options.After                              // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"productdescription", body.ProductDescription}, {"productid", body.ProductId}, {"productmodifiedby", body.ProductModifiedBy}, {"productfixedprice", body.ProductFixedPrice}, {"productstatus", body.ProductStatus}}}}
	updateResult := productCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode("updated sucessfully")
}

//Delete Profile of User

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := productCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                      //slice for multiple documents
	cur, err := productCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	// cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	// json.NewEncoder(w).Encode(results)
	if results == nil {
		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "Empty Collection",
		}
		json.NewEncoder(w).Encode(msg)
	} else {

		msg := Response{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "success",
			Response:      results}
		json.NewEncoder(w).Encode(msg)
		cur.Close(context.TODO())
	}
}
