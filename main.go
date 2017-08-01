package main

import (
	"farmerCalendar/library"
	"farmerCalendar/service"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var date = "04-17"

	var processor = service.Processor{}

	processor.LoadDays()

	findDayErr := processor.FindDay(date)
	if findDayErr != nil {
		fmt.Println(findDayErr)
	}

	spew.Dump(processor.Today)

	harvestDay, findHarvestErr := processor.FindHarvestDay()
	if findHarvestErr != nil {
		fmt.Println(findHarvestErr)
	}

	spew.Dump(harvestDay.ID, library.YearDay2Date(harvestDay.ID))
}
