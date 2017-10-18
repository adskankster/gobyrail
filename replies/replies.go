// replies.go
package replies

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/adskankster/gobyrail/helpers"

	"github.com/spf13/viper"
)

var (
	templates template.Template
)

type WriteFormatter interface {
	WriteFormat() error
}

// initialise the templates
func init() {
	pattern := filepath.Join(helpers.CurrentWorkingDir, "templates", "*.tmpl")
	templates = *template.Must(template.ParseGlob(pattern))
}

// Wrappers as results are named after requests, eg StationBoard is never labelled
// as one, but as the result of the request the asked for it ie GetStationBoardResult

type GetStationBoardResult struct {
	XMLName  xml.Name     `xml:"Envelope"`
	Result   StationBoard `xml:"Body>GetDepartureBoardResponse>GetStationBoardResult"`
	Template string
}

func (s *GetStationBoardResult) WriteFormat() error {
	return templates.ExecuteTemplate(os.Stdout, s.Template, s.Result)
}

// Generic method to convert all replies into the relevant, passed in structure.
func HandleReply(result []byte, v WriteFormatter) error {
	//helpers.DebugMsg("***Request: \r" + string(result))

	// The xml is in one, single line and we may wish to reformat, so unmarshal
	// first, even if we're outputting as xml. If there seems to be a performance
	// hit, we can revisit.
	err := xml.Unmarshal(result, &v)
	if err != nil {
		return fmt.Errorf("Error '%v', while unmarshalling result.", err)
	}

	oformat := viper.Get("format")
	helpers.DebugMsg(fmt.Sprintf("output format: %v", oformat))

	// Will/may add new formats
	switch oformat {
	case "xml":
		if result, err = xml.MarshalIndent(v, "", "\t"); err != nil {
			return fmt.Errorf("Error '%v', while marshalling xml for output.", err)
		}

		fmt.Print(string(result))

		return nil

	case "json":
		if result, err = json.MarshalIndent(v, "", "\t"); err != nil {
			return fmt.Errorf("Error '%v', while marshalling json for output.", err)
		}

		fmt.Println(string(result))

		return nil

	default: // formatted
		if oformat != "fmt" && oformat != "" {
			fmt.Printf("Incorrect output format (%s), using default:", oformat)
		}

		return v.WriteFormat()
	}

	// Should never get here - bad form?
	return fmt.Errorf("Unknown Error in Reply handling.")
}
