package services

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func filterOutWhitespaces(arr []string) []string {
	filtered := []string{}
	for _, v := range arr {
		if len(v) > 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

type parsedDate struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
}

func newParsedDate() parsedDate {
	return parsedDate{0, 1, 1, 0, 0, 0}
}

func (pd *parsedDate) SetYear(year int) {
	pd.year = year
}
func (pd *parsedDate) SetMonth(month int) {
	pd.month = month
}
func (pd *parsedDate) SetDay(day int) {
	pd.day = day
}

func parseDateString(ds []int) parsedDate {
	fmt.Println(ds)
	pd := newParsedDate()
	dsLen := len(ds)
	switch dsLen {
	case 1:
		pd.SetYear(ds[0])
	case 2:
		pd.SetYear(ds[0])
		pd.SetMonth(ds[1])
	case 3:
		pd.SetYear(ds[0])
		pd.SetMonth(ds[1])
		pd.SetDay(ds[2])
	}
	return pd
}

// DateError - custom error type for handling invalid dates
type DateError struct{}

func (de DateError) Error() string {
	return fmt.Sprint("Invalid date")
}

//TimestampResponse - struct representing response type for the client
type TimestampResponse struct {
	Unix int64  `json:"unix"`
	Utc  string `json:"utc"`
}

//Timestamp -
type Timestamp struct{}

func (ts Timestamp) isValidUnix(u string) bool {

	if strings.HasSuffix(u, "-") {
		return false
	}
	converted, err := strconv.ParseInt(u, 10, 64)
	if err != nil {
		return false
	}
	if converted < 0 {
		return false
	}
	return true
}

func (ts Timestamp) fromUnix(u int64) TimestampResponse {
	return newTimestampResponse(u)
}

func (ts Timestamp) fromDateString(ds string) (TimestampResponse, error) {
	dsElements := filterOutWhitespaces(strings.Split(ds, "-"))
	convertedElements := []int{}
	for _, element := range dsElements {
		converted, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			return TimestampResponse{}, DateError{}
		}
		convertedElements = append(convertedElements, int(converted))
	}
	pd := parseDateString(convertedElements)
	loc, _ := time.LoadLocation("UTC")
	date := time.Date(pd.year, time.Month(pd.month), pd.day, pd.hour, pd.minute, pd.second, 0, loc)
	return newTimestampResponse(date.Unix() * 1000), nil

}

//Parse - takes date string as an input and decide which constructor to use to create TimestampResponse
func (ts Timestamp) Parse(ds string) (TimestampResponse, error) {
	if len(ds) == 0 {
		return ts.fromUnix(time.Now().Unix()), nil
	}

	dsElements := filterOutWhitespaces(strings.SplitAfter(ds, "-"))

	if len(dsElements) == 1 && ts.isValidUnix(dsElements[0]) {
		converted, err := strconv.ParseInt(dsElements[0], 10, 64)
		if err != nil {
			return TimestampResponse{}, DateError{}
		}
		return ts.fromUnix(converted / 1000), nil
	}
	return ts.fromDateString(ds)

}

func newTimestampResponse(u int64) TimestampResponse {
	miliseconds := u * 1000
	loc, _ := time.LoadLocation("UTC")
	utc := time.Unix(u, 0).In(loc).Format(http.TimeFormat)
	return TimestampResponse{miliseconds, utc}

}

//Run - main function for the service
func (ts Timestamp) Run(w io.Writer) {
	fmt.Fprint(w, "Timestamp service")
}
