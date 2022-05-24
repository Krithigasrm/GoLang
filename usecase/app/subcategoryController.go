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
type subcategory struct {
	CategoryId             int    `json:"cid"`
	SubCategoryId          int    `json:"scid"`
	SubCategoryName        string `json:"scname"`
	SubCategoryDescription string `json:"scdes"`
	SubCtegoryStatus       bool   `json:"scstatus"`
	SubCategoryCreatedBy   string `json:"sccb"`
	SubCategoryModifiedBy  string `json:"scmb"`
}

var subcategoryCollection = db().Database("useCase").Collection("subcategory") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var subcategory subcategory
	err := json.NewDecoder(r.Body).Decode(&subcategory) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var result1 primitive.M //  an unordered representation of a BSON document which is a Map
	err1 := subcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", subcategory.SubCategoryId}}).Decode(&result1)
	if err1 != nil {
		var result primitive.M
		err2 := userCollection.FindOne(context.TODO(), bson.D{{"categoryid", subcategory.CategoryId}, {"categorystatus", true}}).Decode(&result)
		fmt.Println("errror", err2)
		if err2 != nil {
			json.NewEncoder(w).Encode("category not available")
		} else {
			fmt.Println(err2)
			insertResult, err := subcategoryCollection.InsertOne(context.TODO(), subcategory)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Inserted a single document: ", insertResult)
			json.NewEncoder(w).Encode("subcategory added successfully")

		}
	} else {
		json.NewEncoder(w).Encode("Record cannot be inserted, Id already exits")
	}

}

// Get Profile of a particular User by Name

func GetUserSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body subcategory
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := subcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", body.SubCategoryId}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

func GetSubCategoryId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string
	s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	_ = subcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", s}}).Decode(&result)
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

func GetSubCategoryName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"] //get Parameter value as string
	// s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	_ = subcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryname", params}}).Decode(&result)
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

//Update Profile of User

func UpdateSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		SubCategoryId          int    `json:"scid"`
		SubCategoryName        string `json:"scname"` //value that has to be matched
		SubCategoryDescription string `json:"scdes"`  // value that has to be modified
		SubCategoryStatus      bool   `json:"scstatus"`
		SubCategoryModifiedBy  string `json:"scmb"`
		CategoryId             int    `json:"cid"`
	}

	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"subcategoryid", body.SubCategoryId}} // converting value to BSON type
	after := options.After                                  // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}

	var result1 primitive.M //  an unordered representation of a BSON document which is a Map
	err1 := subcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", body.SubCategoryId}}).Decode(&result1)
	fmt.Println(result1, err1)
	if err1 == nil {
		var result primitive.M
		err3 := userCollection.FindOne(context.TODO(), bson.D{{"categoryid", body.CategoryId}, {"categorystatus", true}}).Decode(&result)
		fmt.Println("result", result)
		if err3 != nil {
			msg := ResponseError{
				ErrorMessage:  "nill",
				StatusCode:    200,
				Status:        false,
				CustomMessage: "category id not available",
			}
			json.NewEncoder(w).Encode(msg)
		} else {
			fmt.Println(result1)
			update := bson.D{{"$set", bson.D{{"subcategorydescription", body.SubCategoryDescription}, {"subcategoryname", body.SubCategoryName}, {"subcategorymodifiedby", body.SubCategoryModifiedBy}, {"subcategorystatus", body.SubCategoryStatus}}}}
			updateResult := subcategoryCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

			var result1 primitive.M
			_ = updateResult.Decode(&result1)
			if result1 == nil {
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
					CustomMessage: "SubCategory Updated successfully",
				}
				json.NewEncoder(w).Encode(msg)
			}
		}
	} else {
		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "No Record exits",
		}
		json.NewEncoder(w).Encode(msg)
	}
}

//Delete Profile of User

func DeleteSubCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := subcategoryCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllSubCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                          //slice for multiple documents
	cur, err := subcategoryCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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
