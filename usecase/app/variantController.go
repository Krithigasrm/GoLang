package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for storing data
type variant struct {
	VariantId          int    `json:"vid" validate:"required,numeric"`
	Varianname         string `json:"vname" validate:"required,alpha"`
	VariantDescription string `json:"vdes" validate:"required,alpha"`
	VariantStatus      bool   `json:"vstatus" validate:"required"`
	VariantCreatedBy   string `json:"vcb" validate:"required,alpha"`
	VariantModifiedBy  string `json:"vmb" validate:"required,alpha"`
}

var variantCollection = db().Database("useCase").Collection("variant") // get collection "users" from db() which returns *mongo.Client

// Create Profile or Signup

func CreateVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type
	validate = validator.New()
	var variant variant

	err := json.NewDecoder(r.Body).Decode(&variant) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}

	err2 := validate.Struct(variant)
	if err2 != nil {
		errors := make(map[string][]string)

		for _, err2 := range err2.(validator.ValidationErrors) {

			var name string
			name = err2.StructField()

			switch err2.Tag() {
			case "required":
				errors[name] = append(errors[name], "The "+name+" is required")
				break
			case "alpha":
				errors[name] = append(errors[name], "The"+name+" should be characters")
				break
			case "numeric":
				errors[name] = append(errors[name], "The "+name+" should be numeric")
				break
			default:
				errors[name] = append(errors[name], "The "+name+" is invalid")
				break

			}

		}

		msg := ResponseErrors{
			ErrorMessage:  "nill",
			StatusCode:    200,
			Status:        false,
			CustomMessage: errors,
		}

		json.NewEncoder(w).Encode(msg)

		return

	}

	var result primitive.M
	err1 := variantCollection.FindOne(context.TODO(), bson.D{{"variantid", variant.VariantId}}).Decode(&result)
	fmt.Println("err", err1, "result", result)
	if result == nil {
		insertResult, err := variantCollection.InsertOne(context.TODO(), variant)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Inserted a single document: ", insertResult)
		json.NewEncoder(w).Encode("Variant Added Successfully")
	} else {
		json.NewEncoder(w).Encode("Record cannot be inserted,variant id already exits")
	}

}

// Get Profile of a particular User by Name

func GetUserVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var body variant
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	err := variantCollection.FindOne(context.TODO(), bson.D{{"vid", body.VariantId}}).Decode(&result)
	if err != nil {

		fmt.Println(err)

	}
	json.NewEncoder(w).Encode(result) // returns a Map containing document

}

func GetVariantId(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"] //get Parameter value as string
	s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	_ = variantCollection.FindOne(context.TODO(), bson.D{{"variantid", s}}).Decode(&result)
	// if result != nil {

	// 	if err != nil {

	// 		fmt.Println(err)

	// 	}

	// 	json.NewEncoder(w).Encode(result) // returns a Map containing document

	// } else {
	// 	json.NewEncoder(w).Encode("no record found")
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

func GetVariantName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["name"] //get Parameter value as string
	// s, _ := strconv.Atoi(params)
	var result primitive.M //  an unordered representation of a BSON document which is a Map
	_ = variantCollection.FindOne(context.TODO(), bson.D{{"variantname", params}}).Decode(&result)
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

//Update Profile of User

func UpdateVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		// VariantName        string `json:"vname"` //value that has to be matched
		VariantId          int    `json:"vid"`
		VariantDescription string `json:"vdes"` // value that has to be modified
		VariantModifiedBy  string `json:"vmb"`
		VariantStatus      bool   `json:"vstatus"`
		VariantName        string `json:"vname"`
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"variantid", body.VariantId}} // converting value to BSON type
	after := options.After                          // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"variantdescription", body.VariantDescription}, {"variantid", body.VariantId}, {"variantstatus", body.VariantStatus}, {"variantmodifiedby", body.VariantStatus}, {"variantname", body.VariantName}}}}
	updateResult := variantCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

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
			CustomMessage: "Variant Updated successfully",
		}
		json.NewEncoder(w).Encode(msg)

	}

	// json.NewEncoder(w).Encode("Updated Sucessfully")
}

//Delete Profile of User

func DeleteVariant(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Printf(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := variantCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func GetAllVariant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                      //slice for multiple documents
	cur, err := variantCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
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
