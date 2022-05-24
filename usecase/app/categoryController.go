package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"

	// "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for storing data
type category struct {
	CategoryId          int    `json:"cid" validate:"numeric"`
	CategoryName        string `json:"cname" validate:"required,alpha"`
	CategoryDescription string `json:"cdes" validate:"required,alpha"`
	CategoryStatus      bool   `json:"cstatus" validate:"required"`
	CategoryCreatedBy   string `json:"ccb" validate:"required`
	CategoryModifiedBy  string `json:"cmb" validate:"required"`
}

var userCollection = db().Database("useCase").Collection("category") // get collection "users" from db() which returns *mongo.Client
var validate *validator.Validate

// Create Profile or Signup

func CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type
	validate = validator.New()
	var category category
	err := json.NewDecoder(r.Body).Decode(&category) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M
	err1 := userCollection.FindOne(context.TODO(), bson.D{{"categoryid", category.CategoryId}}).Decode(&result)
	fmt.Println("err", err1, "result", result)
	if result == nil {
		insertResult, err := userCollection.InsertOne(context.TODO(), category)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document: ", insertResult)
		json.NewEncoder(w).Encode("Category Added successfully")
	} else {
		json.NewEncoder(w).Encode("category id already exits, record cannot be inserted")
	}

	err2 := validate.Struct(category)
	if err2 != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		// fmt.Println("------ List of tag fields with error ---------")

		// for _, err := range err.(validator.ValidationErrors) {
		// 	fmt.Println(err.StructField())
		// 	fmt.Println(err.ActualTag())
		// 	fmt.Println(err.Kind())
		// 	fmt.Println(err.Value())
		// 	fmt.Println(err.Param())
		// 	fmt.Println("---------------")
		// }
		json.NewEncoder(w).Encode("validation field error")

		var errormessage = ""

		for _, err := range err2.(validator.ValidationErrors) {

			errormessage = errormessage + err.StructField() + " should be " + err.ActualTag() + ", "

		}
		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: errormessage,
		}

		json.NewEncoder(w).Encode(msg)
		return

	}
}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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

func GetCategoryId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string
	s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	_ = userCollection.FindOne(context.TODO(), bson.D{{"categoryid", s}}).Decode(&result)
	fmt.Println(result)
	fmt.Println(s)
	// if err1 == nil {
	// 	json.NewEncoder(w).Encode(result) // returns a Map containing document
	// } else {
	// 	json.NewEncoder(w).Encode("No record found")
	// }
	if result == nil {
		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "No Record Exits",
		}
		json.NewEncoder(w).Encode(msg)
	} else {

		msg := Responsedata{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "success",
			Response:      result}
		json.NewEncoder(w).Encode(msg)

	}
}

func GetCategoryName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"] //get Parameter value as string
	// s, _ := strconv.Atoi(params)
	fmt.Println(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	_ = userCollection.FindOne(context.TODO(), bson.D{{"categoryname", params}}).Decode(&result)
	fmt.Println(result)
	fmt.Println(params)
	// if err1 == nil {
	// 	json.NewEncoder(w).Encode(result) // returns a Map containing document
	// } else {
	// 	json.NewEncoder(w).Encode("No record found")
	// }

	if result == nil {
		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "No Record exits",
		}
		json.NewEncoder(w).Encode(msg)
	} else {

		msg := Responsedata{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "success",
			Response:      result}
		json.NewEncoder(w).Encode(msg)

	}
}

// Get Profile of a particular User by Name

// func GetUserCategory(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)["id"] //get Parameter value as string
// 	s, _ := strconv.Atoi(params)
// 	var result primitive.M //  an unordered representation of a BSON document which is a Map
// 	err := userCollection.FindOne(context.TODO(), bson.D{{"categoryid", s}}).Decode(&result)
// 	if result != nil {

// 		if err != nil {

// 			fmt.Println(err)

// 		}

// 		json.NewEncoder(w).Encode(result) // returns a Map containing document

// 	} else {
// 		json.NewEncoder(w).Encode("no record found")
// 	}

// }

//Update Profile of User

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		CategoryId          int    `json:"cid"`
		CategoryName        string `json:"cname"` //value that has to be matched
		CategoryDescription string `json:"cdes"`  // value that has to be modified
		CategoryStatus      bool   `json:"cstatus"`
		CategoryModifiedby  string `json:"cmb"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"categoryid", body.CategoryId}} // converting value to BSON type
	after := options.After                            // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"categorydescription", body.CategoryDescription}, {"categorystatus", body.CategoryStatus}, {"categorymodifiedby", body.CategoryModifiedby}, {"categoryname", body.CategoryName}}}}
	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	// if result != nil {
	// 	json.NewEncoder(w).Encode("Category updated Successfully")
	// } else {
	// 	json.NewEncoder(w).Encode("no record found")
	// }

	// json.NewEncoder(w).Encode("Updated Sucessfully")

	if result == nil {
		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "No Record exits",
		}
		json.NewEncoder(w).Encode(msg)
	} else {

		msg := Responseupdate{
			StatusCode:    200,
			Status:        true,
			CustomMessage: "Category Updated successfully",
		}
		json.NewEncoder(w).Encode(msg)

	}
}

//Delete Profile of User

func DeleteCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}
