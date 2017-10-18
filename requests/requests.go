// Requests
// Functions perform validation. See https://lite.realtime.nationalrail.co.uk/OpenLDBWS/
// for the definitions and validation criteria.
package requests

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/adskankster/gobyrail/helpers"
	"github.com/spf13/viper"
)

// adskankster
// awellings@kapin.co.uk
// Actual token - do not alter.
//const token = "9bb0b4c2-6126-42c0-9c53-b6aa050cc7bd"

// xml version string
const xmlVer = `<?xml version="1.0" encoding="utf-8"?>`

// xml namespace urls. These may/will change so they are defined up here.
const ldbSchema = "http://thalesgroup.com/RTTI/2016-02-16/ldb/"
const tokenSchema = "http://thalesgroup.com/RTTI/2013-11-28/Token/types"
const w3Schema = "http://www.w3.org/2001/XMLSchema"
const w3SchemaInst = "http://www.w3.org/2001/XMLSchema-instance"
const envSchema = "http://schemas.xmlsoap.org/soap/envelope/"

// Builds and returns the xml required for all requests
func BuildReq(reqType interface{}) (bytes.Buffer, error) {
	req := &Request{}

	req.SE = envSchema

	req.ReqHeader.Token.Sch1 = tokenSchema
	req.ReqHeader.Token.Sch2 = tokenSchema
	req.ReqHeader.Token.Sch3 = w3Schema
	req.ReqHeader.Token.Sch4 = w3SchemaInst

	req.ReqHeader.Token.TokenValue = viper.GetString("authtoken")

	req.ReqBody.Sch1 = w3SchemaInst
	req.ReqBody.Sch2 = w3Schema

	req.ReqBody.RequestBody = reqType

	var buffer bytes.Buffer
	buffer.WriteString(xmlVer)

	b, err := xml.MarshalIndent(req, "", " ")
	buffer.Write(b)

	helpers.DebugMsg("***Request: \r" + string(b))

	return buffer, err
}

// Specific request builders

func BuildGetDepartureBoardRequest(numRows int, crs string, filterCrs string, filterType string, timeOffset int, timeWindow int) (*GetDepartureBoardRequest, error) {

	// Validate
	var errmsg = "%s has invalid value: %v"

	if numRows < 0 || numRows > 150 {
		return nil, fmt.Errorf(errmsg, "numRows", numRows)
	}

	// Check crs value?
	if len(crs) != 3 {
		return nil, fmt.Errorf(errmsg, "crs", crs)
	}
	// Check filterCrs?
	if len(filterCrs) != 3 && len(filterCrs) != 0 {
		return nil, fmt.Errorf(errmsg, "filterCrs", filterCrs)
	}

	if filterType != "from" && filterType != "to" {
		return nil, fmt.Errorf(errmsg, "filterType", filterType)
	}

	if timeOffset < -120 && timeOffset > 120 {
		return nil, fmt.Errorf(errmsg, "timeOffset", timeOffset)
	}

	if timeWindow < -120 && timeOffset > 120 {
		return nil, fmt.Errorf(errmsg, "timeWindow", timeWindow)
	}

	// Build request
	req := &GetDepartureBoardRequest{
		Sch:        ldbSchema,
		NumRows:    numRows,
		Crs:        crs,
		FilterCrs:  filterCrs,
		FilterType: filterType,
		TimeOffset: timeOffset,
		TimeWindow: timeWindow,
	}

	return req, nil
}
