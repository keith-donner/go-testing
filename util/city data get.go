package util

import (
	"or/o"
	"strings"
	"time"
)

// send in the city code and it will return the following
//
// metroCode
// cityName
// timeZone.Long
// timeZone.Short
// offsetUtc

type city struct {
	AirportCodes   []string
	MetroCode      string
	CityName       string
	Country        o.Country
	TimeZone       timeZoneValues
	OffsetUtc      int
	TimeStamp      timeStamp
	HemisphereCode int
}

type timeZoneValues struct {
	Long  string
	Short string
}

type timeStamp struct {
	Military timeStampValues
	US       timeStampValues
}

type timeStampValues struct {
	DateTime time.Time
	Text     string
}

// this will return multiple values for the city code
//
// TimeZone (long and short)
//
// # AirportCodes: all the surrounding airport codes []string
//
// # MetroCode
//
// # CityName
//
// # OffSetUtc
//
// TimeStamp.Military | TimeStamp.US (each contain values: DateTime and Text for use in remarks)
//
// HemisphereCode: used for BSP (0-9)
// type segment struct {
// }

// func CityDataFromSegment(segment segment, useDepart bool) city {
// 	return CityData()

// }

func CityData(cityCode string, languageCode o.LanguageCode) city {

	CityValues := new(city)
	CityValues.MetroCode = "YVR"
	// CityValues.Country = o.CountryGet("US")

	CityValues.TimeZone.Long = "12"
	CityValues.TimeZone.Short = "ET"
	CityValues.OffsetUtc = 4
	// CityValues.Country.IsoCode = "FD"
	CityValues.Country.Name = "FD"
	CityValues.CityName = "YYRR"
	strings.Split(CityValues.TimeZone.Long, "")
	CityValues.TimeStamp.Military.DateTime = time.Now()
	CityValues.TimeStamp.Military.Text = "time"
	CityValues.TimeStamp.US.DateTime = time.Now()
	CityValues.TimeStamp.US.Text = "time"
	CityValues.HemisphereCode = 0
	return *CityValues
}

// func CountriesInPnrGet (pnr pnr) []Country{

// }
