package requests

import (
	"encoding/xml"
)

// Specific xml definitions

type GetDepartureBoardRequest struct {
	XMLName    xml.Name `xml:"GetDepartureBoardRequest"`
	Sch        string   `xml:"xmlns,attr"`
	NumRows    int      `xml:"numRows"`
	Crs        string   `xml:"crs"`
	FilterCrs  string   `xml:"filterCrs"`
	FilterType string   `xml:"filterType"`
	TimeOffset int      `xml:"timeOffset"`
	TimeWindow int      `xml:"timeWindow"`
}

// Generic xml - used in all requests

type Body struct {
	Sch1        string      `xml:"xmlns:xsi,attr"`
	Sch2        string      `xml:"xmlns:xsd,attr"`
	RequestBody interface{} // body type goes in here
}

type AccessToken struct {
	Sch1       string `xml:"xmlns:h,attr"`
	Sch2       string `xml:"xmlns,attr"`
	Sch3       string `xml:"xmlns:xsi,attr"`
	Sch4       string `xml:"xmlns:xsd,attr"`
	TokenValue string `xml:"TokenValue"`
}

type Header struct {
	Token AccessToken `xml:"h:AccessToken"`
}

type Request struct {
	XMLName   xml.Name `xml:"s:Envelope"`
	SE        string   `xml:"xmlns:s,attr"`
	ReqHeader Header   `xml:"s:Header"`
	ReqBody   Body     `xml:"s:Body"`
}

type FilterList struct {
	Crs []string `xml:">crs"`
}

/*

<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
	<s:Header>
		<h:AccessToken xmlns:h="http://thalesgroup.com/RTTI/2013-11-28/Token/types/"
					   xmlns="http://thalesgroup.com/RTTI/2013-11-28/Token/types/"
					   xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
					   xmlns:xsd="http://www.w3.org/2001/XMLSchema">
			<TokenValue>9bb0b4c2-6126-42c0-9c53-b6aa050cc7bd</TokenValue>
		</h:AccessToken>
	</s:Header>
	<s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xmlns:xsd="http://www.w3.org/2001/XMLSchema">
		<GetDepartureBoardRequest xmlns="http://thalesgroup.com/RTTI/2016-02-16/ldb/">
			<numRows>10</numRows>
			<crs>SHF</crs>
		</GetDepartureBoardRequest>
	</s:Body>
</s:Envelope>


*/
