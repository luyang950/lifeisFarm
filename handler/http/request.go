package http

import (
	"encoding/json"
	"lifeisFarm/library"
	"lifeisFarm/service"
	"fmt"
	"net/http"
)

type jsonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []string `json:"data"`
}

func calHarvest(w http.ResponseWriter, r *http.Request) {
	var response = jsonResponse{}
	post := r.PostFormValue("date")

	if post == "" {
		fmt.Println("invalid param")

		response.Code = -1
		response.Msg = "invalid param"
		response.Data = []string{}

		fmt.Fprintf(w, genJSONStr(response))
		return
	}

	fmt.Println("plantDate:", post)

	var plantDate = post

	var processor = service.Processor{}

	processor.LoadDays()

	findDayErr := processor.FindDay(plantDate)
	if findDayErr != nil {
		fmt.Println(findDayErr)
	}

	resDays, findHarvestErr := processor.FindHarvestDay()
	if findHarvestErr != nil {
		fmt.Println(findHarvestErr)
	}

	var resDates = make([]string, 0)
	resDates = append(resDates, plantDate)

	for _, v := range resDays {
		resDates = append(resDates, library.YearDay2Date(v.ID))
	}

	response.Code = 0
	response.Msg = "success"
	response.Data = resDates

	fmt.Println("res dates:", resDates)
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
