package corm

import (
	"testing"
	"fmt"
)

func TestInsertLBS(t *testing.T) {
	lp := new(LbsPosition)
	lp.ApplicationId = "sunlichuan"
	lp.UserKey = "a"
	lp.Lng = 139.7454149999999
	lp.Lat = 35.658650898203035
	lp.Status = 0
	lp.FormattedAddress = "東京都港区芝公園4-2-8, Minato, Tokyo, Japan"
	//lp.Confidence = ""
	lp.Business = ""
	lp.Country = "Japan"
	lp.Province = "Tokyo"
	lp.City = "Minato"
	lp.District = ""
	lp.Town = ""
	lp.Street = "東京都港区芝公園4-2-8"
	lp.StreetNumber = ""
	lp.Adcode = "0"
	lp.CountryCode = 26000
	//lp.RegionsName = ""
	//lp.RegionsTag = ""
	//lp.RegionsDirection = ""
	lp.SematicDescription = ""
	lp.CityCode = 26041
	lp.WholeMsg = `{"status":0,"result":{"location":{"lng":139.7454149999999,"lat":35.658650898203038},"formatted_address":"東京都港区芝公園4-2-8, Minato, Tokyo, Japan","business":"","addressComponent":{"country":"Japan","country_code":26000,"country_code_iso":"JPN","country_code_iso2":"JP","province":"Tokyo","city":"Minato","city_level":1,"district":"","town":"","adcode":"0","street":"東京都港区芝公園4-2-8","street_number":"","direction":"附近","distance":"40"},"pois":[],"roads":[],"poiRegions":[],"sematic_description":"","cityCode":26041}}`
	result, err := InsertLBS(lp)
	if err != nil {
		fmt.Errorf("%s\n", err)
	} else {
		fmt.Println(result)
	}
}
