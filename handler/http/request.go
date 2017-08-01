package http

import (
	"encoding/json"
	"farmerCalendar/library"
	"farmerCalendar/service"
	"fmt"
	"net/http"
)

type jsonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func calHarvest(w http.ResponseWriter, r *http.Request) {
	var response = jsonResponse{}
	post := r.PostFormValue("date")

	if post == "" {
		response.Code = -1
		response.Msg = "invalid param"
		response.Data = ""

		fmt.Fprintf(w, genJSONStr(response))
		return
	}

	var date = post

	var processor = service.Processor{}

	processor.LoadDays()

	findDayErr := processor.FindDay(date)
	if findDayErr != nil {
		fmt.Println(findDayErr)
	}

	harvestDay, findHarvestErr := processor.FindHarvestDay()
	if findHarvestErr != nil {
		fmt.Println(findHarvestErr)
	}

	harvestDate := library.YearDay2Date(harvestDay.ID)

	response.Code = 0
	response.Msg = "success"
	response.Data = harvestDate

	fmt.Fprintf(w, genJSONStr(response))
}

func genJSONStr(res jsonResponse) string {
	var resBytes []byte
	var jsonErr error

	resBytes, jsonErr = json.Marshal(res)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return string(resBytes)
}
