package service

import (
	"encoding/xml"
	"farmerCalendar/library"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	growthToHarvest = 8

)

type Processor struct {
	Calendar Calendar
	Today    Day
}

type Calendar struct {
	Name xml.Name `xml:"calendar"`
	Days []Day    `xml:"day"`
}

type Day struct {
	ID      int    `xml:"id,attr"`
	Season  string `xml:"season,attr"`
	Weather string `xml:",chardata"`
}

func (p *Processor) LoadDays() {
	content, err := ioutil.ReadFile("./datasource/cm_weather1.xml")

	if err != nil {
		fmt.Println(err)
	}

	var calendar Calendar
	err = xml.Unmarshal(content, &calendar)
	if err != nil {
		fmt.Println(err)
	}

	p.Calendar = calendar
}

func (p *Processor) FindDay(date string) error {
	if len(p.Calendar.Days) != 365 {
		return errors.New("calendar not loaded")
	}

	dateSli := strings.Split(date, "-")

	if len(dateSli) != 2 {
		return errors.New("invalid date format")
	}

	var month, day int

	month, _ = strconv.Atoi(dateSli[0])
	day, _ = strconv.Atoi(dateSli[1])

	if month == 0 || day == 0 {
		return errors.New("invalid date")
	}

	yearDay := library.YearDay(month, day)

	if yearDay < 0 {
		return errors.New("invalid date")
	}

	fmt.Println("post:", yearDay)

	p.Today = p.Calendar.Days[yearDay]

	return nil
}

func (p *Processor) FindHarvestDay() ([]Day, error) {
	if len(p.Calendar.Days) != 365 {
		return []Day{}, errors.New("calendar not loaded")
	}
	if p.Today.Season == "" {
		return []Day{}, errors.New("today not loaded")
	}

	var resDays = make([]Day, 0)

	var lastWeather = p.Today.Weather
	var lastSeason = p.Today.Season
	var growth = 0

	for i := p.Today.ID; ; i++ {
		var targetDay Day


		if i >= 365 {
			i = 0
		}
		if i == p.Today.ID {
			growth++
		}
		// wont grow in winter
		if p.Calendar.Days[i].Season == "Winter" {
			continue
		}

		if p.Calendar.Days[i].Weather != lastWeather && (p.Calendar.Days[i].Weather == "Shower" || p.Calendar.Days[i].Weather == "Fair") {
			if lastSeason == "Winter" || (lastWeather != "Shower" && lastWeather != "Fair") {
				lastWeather = p.Calendar.Days[i].Weather
				lastSeason = p.Calendar.Days[i].Season
				continue
			}
			growth++
			lastWeather = p.Calendar.Days[i].Weather
		}

		if growth == growthToHarvest && len(resDays) == 0 {
			targetDay = p.Calendar.Days[i]
			fmt.Println("harvest:", i, targetDay)
			resDays = append(resDays, targetDay)
			continue
		}

		if growth == growthToHarvest + 1&& len(resDays) == 1 {
			targetDay = p.Calendar.Days[i]
			fmt.Println("expire1:", i, targetDay)
			resDays = append(resDays, targetDay)
			continue
		}

		if growth == growthToHarvest + 2 && len(resDays) == 2{
			targetDay = p.Calendar.Days[i]
			fmt.Println("expire2:", i, targetDay)
			resDays = append(resDays, targetDay)
			continue
		}

		if growth == growthToHarvest + 3 && len(resDays) == 3{
			targetDay = p.Calendar.Days[i]
			fmt.Println("expire3:", i, targetDay)
			resDays = append(resDays, targetDay)
			break
		}
	}

	return resDays, nil
}
