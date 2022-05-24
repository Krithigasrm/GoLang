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
type brand struct {
	BrandId          int    `json:"bid"`
	CategoryId       int    `json:"cid"`
	SubcategoryId    int    `json:"scid"`
	BrandName        string `json:"bname"`
	BrandDescription string `json:"bdes"`
	BrandStatus      bool   `json:"bstatus"`
	BrandCreatedBy   string `json:"bcb"`
	BrandModifiedBy  string `json:"bmb"`
}

type ResponseError struct {
	ErrorMessage  string `json:"errormessage"`
	StatusCode    int    `json:"statuscode"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"customessage"`
}
type ResponseErrors struct {
	ErrorMessage  string              `json:"errormessage"`
	StatusCode    int                 `json:"statuscode"`
	Status        bool                `json:"status"`
	CustomMessage map[string][]string `json:"customessage"`
}

type Response struct {
	//ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"statuscode"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"customessage"`
	Response      []primitive.M
}

type Responsedata struct {
	//ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"statuscode"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"customessage"`
	Response      primitive.M
}

type Responseupdate struct {
	//ErrorMessage  string `json:"error message"`
	StatusCode    int    `json:"statuscode"`
	Status        bool   `json:"status"`
	CustomMessage string `json:"custommessage"`
}

var brandCollection = db().Database("useCase").Collection("brand") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var brand brand
	err := json.NewDecoder(r.Body).Decode(&brand) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err1 := brandCollection.FindOne(context.TODO(), bson.D{{"brandid", brand.BrandId}}).Decode(&result)
	fmt.Println(result, err1)

	if err1 != nil {
		var result1 primitive.M //  an unordered representation of a BSON document which is a Map
		err := userCollection.FindOne(context.TODO(), bson.D{{"categoryid", brand.CategoryId}}).Decode(&result1)
		fmt.Println(result1, err)

		if err != nil {
			json.NewEncoder(w).Encode("category id not available")

			if err1 != nil {
				var result2 primitive.M
				err2 := subcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", brand.SubcategoryId}, {"subctegorystatus", true}}).Decode(&result2)
				fmt.Println("errror", err2)

				if err2 != nil {
					json.NewEncoder(w).Encode("subcategory id not available")
				} else {
					fmt.Println(err2)
					insertResult, err := brandCollection.InsertOne(context.TODO(), brand)
					if err != nil {
						log.Fatal(err)
					}

					fmt.Println("Inserted a single document: ", insertResult)
					// json.NewEncoder(w).Encode(insertResult.InsertedID)
					json.NewEncoder(w).Encode("Brand added sucessfully")

				}
			} else {
				json.NewEncoder(w).Encode("Record cannot be inserted, Brand Id already exits")
			}
		}

	}

}

// Get Profile of a particular User by Name

func GetUserBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body brand
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := brandCollection.FindOne(context.TODO(), bson.D{{"bname", body.BrandName}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

func GetBrandId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string
	s, _ := strconv.Atoi(params)
	var result primitive.M
	// var results primitive.M //  an unordered representation of a BSON document which is a Map
	_ = brandCollection.FindOne(context.TODO(), bson.D{{"brandid", s}}).Decode(&result)
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

func GetBrandName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"] //get Parameter value as string
	// s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	// err1 := brandCollection.FindOne(context.TODO(), bson.D{{"brandname", params}}).Decode(&result)
	_ = brandCollection.FindOne(context.TODO(), bson.D{{"brandname", params}}).Decode(&result)
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

func UpdateBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		BrandName        string `json:"bname"` //value that has to be matched
		BrandDescription string `json:"bdes"`  // value that has to be modified
		BrandId          int    `json:"bid"`
		CategoryId       int    `json:"cid"`
		BrandStatus      bool   `json:"bstatus"`
		BrandModifiedBy  string `json:"bmb"`
		SubCategoryId    int    `json:"scid"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"brandid", body.BrandId}} // converting value to BSON type
	after := options.After                      // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	// var brand brand
	var result1 primitive.M //  an unordered representation of a BSON document which is a Map
	err1 := brandCollection.FindOne(context.TODO(), bson.D{{"brandid", body.BrandId}}).Decode(&result1)
	fmt.Println(result1, err1)
	if err1 == nil {
		var result primitive.M
		err2 := subcategoryCollection.FindOne(context.TODO(), bson.D{{"subcategoryid", body.SubCategoryId}, {"subctegorystatus", true}}).Decode(&result)
		fmt.Println("result", err2, result)
		err3 := userCollection.FindOne(context.TODO(), bson.D{{"categoryid", body.CategoryId}, {"categorystatus", true}}).Decode(&result)
		fmt.Println("result", err3, result)
		var result1 primitive.M
		if err3 != nil {
			msg := ResponseError{
				ErrorMessage:  "nill",
				StatusCode:    200,
				Status:        false,
				CustomMessage: "category id not available",
			}
			json.NewEncoder(w).Encode(msg)
		} else {
			if err2 != nil {
				msg := ResponseError{
					ErrorMessage:  "nill",
					StatusCode:    200,
					Status:        false,
					CustomMessage: "subcategory id not available",
				}
				json.NewEncoder(w).Encode(msg)
			} else {
				update := bson.D{{"$set", bson.D{{"branddescription", body.BrandDescription}, {"brandname", body.BrandName}, {"brandstatus", body.BrandStatus}, {"brandmodifiedby", body.BrandModifiedBy}}}}
				updateResult := brandCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

				_ = updateResult.Decode(&result1)
				if result1 == nil {
					msg := ResponseError{
						ErrorMessage:  "nill",
						StatusCode:    200,
						Status:        false,
						CustomMessage: "No Record exits, brand id not exits to update",
					}
					json.NewEncoder(w).Encode(msg)
				} else {

					msg := Responseupdate{
						StatusCode:    200,
						Status:        true,
						CustomMessage: "Updated successfully",
					}
					json.NewEncoder(w).Encode(msg)

				}
			}
		}
		// if result != nil {
		// 	json.NewEncoder(w).Encode("Brand updated Successfully")
		// } else {
		// 	json.NewEncoder(w).Encode("no record found")
		// }

	} else {

		msg := ResponseError{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: "id not found",
		}
		json.NewEncoder(w).Encode(msg)

	}

}

//Delete Profile of User

func DeleteBrand(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := brandCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                    //slice for multiple documents
	cur, err := brandCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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
