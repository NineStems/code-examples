package main

import (
	"encoding/json"
	calc "go-rti-testing/src/calculate"
	loggin "go-rti-testing/src/log"
	models "go-rti-testing/src/models"
	"io/ioutil"
	"net/http"
)

type Req struct {
	ProductIn   models.Product     `json:"productIn"`
	ConditionIn []models.Condition `json:"conditionIn"`
}
type ResponsObject struct {
	Succeed  bool         `json:"succeed"`
	Comments string       `json:"comments"`
	Offer    models.Offer `json:"offer"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		loggin.LogWorkApp("start processing request", loggin.UsualLog)
		var request Req
		var response ResponsObject
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			loggin.LogWorkApp("Error with getting req`s body:"+err.Error(), loggin.Problem)
			return
		}
		err = json.Unmarshal(body, &request)
		if err != nil {
			loggin.LogWorkApp("Error with unmarshall req`s body:"+err.Error(), loggin.Problem)
			return
		}
		var p models.Product = request.ProductIn
		var c []models.Condition = request.ConditionIn

		offer, err := calc.Calculate(&p, c)
		/*Необходимо прокидывать из функции расчёта текст произошедшего и передавать его в JSON добавляемый в ответ*/
		if err != nil {
			loggin.LogWorkApp("Error with calculate offer:"+err.Error(), loggin.Problem)
			return
		}
		response.Succeed = true
		response.Comments = ""
		response.Offer = *offer

		resultOffer, err := json.Marshal(response)
		if err != nil {
			loggin.LogWorkApp("Error with marshal respons`s:"+err.Error(), loggin.Problem)
			return
		}
		w.Write(resultOffer)

		loggin.LogWorkApp("end processing request", loggin.UsualLog)
	})
	http.ListenAndServe(":81", nil)

}
