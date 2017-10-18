// webservice
package network

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adskankster/gobyrail/requests"
)

const url = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb9.asmx"
const actionBase = "http://thalesgroup.com/RTTI/2012-01-13/ldb/"

func MakeRequest(request string, reqType interface{}) ([]byte, error) {
	buffer, err := requests.BuildReq(reqType)
	if err != nil {
		return nil, fmt.Errorf("Could not build request: %v\r", err)
	}
	// fmt.Printf("%s\r\r", string(req))

	client := http.Client{}
	httpReq, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return nil, fmt.Errorf("Could not create HTTP POST: %v\r", err)
	}

	httpReq.Header.Add("SOAPAction", actionBase+request)
	httpReq.Header.Add("Content-Type", "text/xml")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("Could not send request: %v\r", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Request failed: %s\r", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Could not read response: %v\r", err)
	}

	return body, nil
}
