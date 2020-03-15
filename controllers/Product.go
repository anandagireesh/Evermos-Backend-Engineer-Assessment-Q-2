package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-2/models"
)

func AddProducts(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.Product

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.Product{
		ProductName:     msg.ProductName,
		ProductQuantity: msg.ProductQuantity,
		ProductPrice:    msg.ProductPrice,
	}

	message := models.AddProduct(data)

	log.Info(message)

	var dataresponse []models.StatusMessage

	dataresponse = append(dataresponse, models.StatusMessage{Message: message})

	response := &models.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   dataresponse,
	}

	urlsJson, _ := json.Marshal(response)
	log.Println(urlsJson)
	//log.Println(Bloguser)

	w.Write(urlsJson)

}

func BuyProduct(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.Product

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	data := models.Product{
		ProductName:     msg.ProductName,
		ProductQuantity: msg.ProductQuantity,
		ProductPrice:    msg.ProductPrice,
	}

	message := models.ButProduct(data)
	log.Info(message)

	var dataresponse []models.StatusMessage

	dataresponse = append(dataresponse, models.StatusMessage{Message: message})

	response := &models.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   dataresponse,
	}

	urlsJson, _ := json.Marshal(response)
	log.Println(urlsJson)
	//log.Println(Bloguser)

	w.Write(urlsJson)

}
