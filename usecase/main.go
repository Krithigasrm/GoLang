package main

import (
	"log"
	"net/http"
	"usecase/app"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createcategory", app.CreateCategory).Methods("POST")
	s.HandleFunc("/getAllcategory", app.GetAllCategory).Methods("GET")
	s.HandleFunc("/getcategoryid/{id}", app.GetCategoryId).Methods("GET")
	s.HandleFunc("/getcategoryname/{name}", app.GetCategoryName).Methods("GET")
	s.HandleFunc("/updatecategory", app.UpdateCategory).Methods("PUT")
	s.HandleFunc("/deletecategory/{id}", app.DeleteCategory).Methods("DELETE")

	s.HandleFunc("/createsubcategory", app.CreateSubCategory).Methods("POST")
	s.HandleFunc("/getAllsubcategory", app.GetAllSubCategory).Methods("GET")
	s.HandleFunc("/getUsersubcategory", app.GetUserSubCategory).Methods("POST")
	s.HandleFunc("/getsubcategoryid/{id}", app.GetSubCategoryId).Methods("GET")
	s.HandleFunc("/getsubcategoryname/{name}", app.GetSubCategoryName).Methods("GET")
	s.HandleFunc("/updatesubcategory", app.UpdateSubCategory).Methods("PUT")
	s.HandleFunc("/deletesubcategory/{id}", app.DeleteSubCategory).Methods("DELETE")

	s.HandleFunc("/createbrand", app.CreateBrand).Methods("POST")
	s.HandleFunc("/getbrandid/{id}", app.GetBrandId).Methods("GET")
	s.HandleFunc("/getbrandname/{name}", app.GetBrandName).Methods("GET")
	s.HandleFunc("/getAllbrand", app.GetAllBrand).Methods("GET")
	s.HandleFunc("/getUserbrand", app.GetUserBrand).Methods("POST")
	s.HandleFunc("/updatebrand", app.UpdateBrand).Methods("PUT")
	s.HandleFunc("/deletebrand/{id}", app.DeleteBrand).Methods("DELETE")

	s.HandleFunc("/createvariant", app.CreateVariant).Methods("POST")
	s.HandleFunc("/getAllvariant", app.GetAllVariant).Methods("GET")
	s.HandleFunc("/getvariantid/{id}", app.GetVariantId).Methods("GET")
	s.HandleFunc("/getvariantname/{name}", app.GetVariantName).Methods("GET")
	s.HandleFunc("/getUservariant", app.GetUserVariant).Methods("POST")
	s.HandleFunc("/updatevariant", app.UpdateVariant).Methods("PUT")
	s.HandleFunc("/deletevariant/{id}", app.DeleteVariant).Methods("DELETE")

	s.HandleFunc("/createproduct", app.CreateProduct).Methods("POST")
	s.HandleFunc("/getAllproduct", app.GetAllProduct).Methods("GET")
	s.HandleFunc("/getproductid/{id}", app.GetProductId).Methods("GET")
	s.HandleFunc("/getproductname/{name}", app.GetProductName).Methods("GET")
	s.HandleFunc("/getuserproduct", app.GetUserProduct).Methods("POST")
	s.HandleFunc("/updateproduct", app.UpdateProduct).Methods("PUT")
	s.HandleFunc("/deleteproduct/{id}", app.DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
