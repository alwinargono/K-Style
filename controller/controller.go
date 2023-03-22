package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test/dbconnection"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	var newMember dbconnection.Member
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter JSON Data for insert to database")
	}

	json.Unmarshal(reqBody, &newMember)
	dbconnection.InsertToMemberTable(newMember)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newMember)
	json.NewEncoder(w).Encode("Insert Successful")
}

func Update(w http.ResponseWriter, r *http.Request) {
	var newMember dbconnection.Member
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter memberID and data that wants to be updated")
	}

	json.Unmarshal(reqBody, &newMember)
	dbconnection.UpdateToMemberTable(newMember)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newMember)
	json.NewEncoder(w).Encode("Update Successful")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var newMember dbconnection.Member
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter memberID that wants to be deleted")
	}

	json.Unmarshal(reqBody, &newMember)
	dbconnection.DeleteFromMemberTable(newMember)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newMember)
}

func ViewAll(w http.ResponseWriter, r *http.Request) {
	var newMember dbconnection.Member
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "JSON input error")
	}

	json.Unmarshal(reqBody, &newMember)
	temp := dbconnection.ViewAllMemberTable(newMember)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(temp)
}

func FindProduct(w http.ResponseWriter, r *http.Request) {
	var productID dbconnection.Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please input product_id")
	}

	json.Unmarshal(reqBody, &productID)
	temp := dbconnection.ViewProduct(productID)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(temp)
}

func LikeOrDislike(w http.ResponseWriter, r *http.Request) {
	var option dbconnection.LikeDislike
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please input option for Like or Dislike, reviewID and add memberID for Like")
	}

	json.Unmarshal(reqBody, &option)

	temp := dbconnection.InsertLikeorDislike(option)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(temp)
}
