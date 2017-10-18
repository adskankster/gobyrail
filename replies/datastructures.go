package replies

// Response Structures

type ServiceLocation struct {
	Location Location `xml:"location"`
}

type Location struct {
	Name             string `xml:"locationName"`
	Crs              string `xml:"crs"`
	Via              string `xml:"via"`
	AssocIsCancelled bool   `xml:"assocIsCancelled"`
	FutureChangeTo   string `xml:"futureChangeto"`
}

type ServiceItem struct {
	Origin                  ServiceLocation   `xml:"origin"`
	Destination             ServiceLocation   `xml:"destination"`
	CurrentOrigins          []ServiceLocation `xml:"currentOrigins>location"`
	CurrentDestinations     []ServiceLocation `xml:"currentDestinations>location`
	STA                     string            `xml:"sta"`
	STD                     string            `xml:"std"`
	ETA                     string            `xml:"eta"`
	ETD                     string            `xml:"etd"`
	Platform                string            `xml:"platform"`
	Operator                string            `xml:"operator"`
	OperatorCode            string            `xml:"operatorCode"`
	ServiceID               string            `xml:"serviceID"`
	AdhocAlerts             []string          `xml:"adhocAlerts>adhocAlertText`
	CancelReason            string            `xml:"cancelReason"`
	DelayReason             string            `xml:"delayReason"`
	DetachFront             bool              `xml:"detachFront"`
	FilterLocationCancelled bool              `xml:"filterLocatioName`
	Crs                     string            `xml:"crs"`
	Via                     string            `xml:"via"`
	AssocIsCancelled        bool              `xml:"assocIsCancelled"`
	FutureChangeTo          string            `xml:"futureChangeto"`
	IsCancelled             bool              `xml:"isCancelled"`
	IsCircularRoute         bool              `xml:"isCircularRoute"`
	IsReverseFormation      bool              `xml:"isReverseFormation"`
	Length                  int               `xml:"length"`
	ServiceType             string            `xml:"serviceType"`
}

type ServiceItemWithCallingPoints struct {
	PreviousCallingPoints   []CallingPoint `xml:"previousCallingPoints>callingPoint`
	SubsequentCallingPoints []CallingPoint `xml:"subsequentCallingPoints>callingPoint`
	Service                 ServiceItem
}

type StationBoard struct {
	GeneratedAt          string        `xml:"generatedAt"`
	LocationName         string        `xml:"locationName"`
	Crs                  string        `xml:"crs"`
	FilterLocationName   string        `xml:"filterLocationName"`
	FilterCrs            string        `xml:"filtercrs"`
	FilterType           string        `xml:"filterType"`
	NrccMessages         []string      `xml:"nrccMessages>message"`
	PlatformAvailable    bool          `xml:"platformAvailable"`
	AreServicesAvailable bool          `xml:"areServicesAvailable"`
	TrainServices        []ServiceItem `xml:"trainServices>service"`
}

type StationBoardWithDetails struct {
	GeneratedAt         string        `xml:"generatedAt"`
	LocationName        string        `xml:"locationName"`
	Crs                 string        `xml:"crs"`
	FilterLocationName  string        `xml:"filterLocationName"`
	FilterCrs           string        `xml:"filtercrs"`
	FilterType          string        `xml:"filterType"`
	NrccMessages        []string      `xml:"nrccMessages>message"`
	PlatformAvailable   bool          `xml:"platformAvailable"`
	AreSrvicesAvailable bool          `xml:"areServicesAvailable"`
	TrainServices       []ServiceItem `xml:"trainServices>service"`
}

type DeparturesBoard struct {
	Departures []DepartureItem `xml:"departures>departureItem`
	Board      StationBoard
}

type DeparturesBoardWithDetails struct {
	Departures []DepartureItem `xml:"departures>departureItem`
	Board      StationBoard
}

type DepartureItem struct {
	Crs     string      `xml:"crs"`
	Service ServiceItem `xml:"service"`
}

type DepartureItemWithCallingPoints struct {
	Crs     string      `xml:"crs"`
	Service ServiceItem `xml:"service"`
}

type ServiceDetails struct {
	ATA                     string            `xml:"ata"`
	ATD                     string            `xml:"atd"`
	GeneratedAt             string            `xml:"generatedAt"`
	LocationName            string            `xml:"locationName"`
	OverdueMessage          string            `xml:"overdueMessage"`
	Crs                     ServiceLocation   `xml:"crs"`
	Destination             ServiceLocation   `xml:"destination"`
	CurrentOrigins          []ServiceLocation `xml:"currentOrigins>location"`
	CurrentDestinations     []ServiceLocation `xml:"currentDestinations>location`
	STA                     string            `xml:"sta"`
	STD                     string            `xml:"sta"`
	ETA                     string            `xml:"eta"`
	ETD                     string            `xml:"etd"`
	Platform                string            `xml:"platform"`
	Operator                string            `xml:"operator"`
	OperatorCode            string            `xml:"operatorCode"`
	ServiceID               string            `xml:"serviceID"`
	AdhocAlerts             []string          `xml:"adhocAlerts>adhocAlertText`
	CancelReason            string            `xml:"cancelReason"`
	DelayReason             string            `xml:"delayReason"`
	DetachFront             bool              `xml:"detachFront"`
	FilterLocationCancelled bool              `xml:"filterLocationCancelled"`
	IsCancelled             bool              `xml:"isCancelled"`
	IsCircularRoute         bool              `xml:"isCircularRoute"`
	IsReverseFormation      bool              `xml:"isReverseFormation"`
	Length                  int               `xml:"length"`
	ServiceType             string            `xml:"serviceType"`
	PreviousCallingPoints   []CallingPoint    `xml:"previousCallingPoints>callingPoint`
	SubsequentCallingPoints []CallingPoint    `xml:"subsequentCallingPoints>callingPoint`
}

type CallingPoint struct {
	Name        string   `xml:"locationName"`
	Crs         string   `xml:"crs"`
	At          string   `xml:"at"`
	Et          string   `xml:"et"`
	IsCancelled bool     `xml:"isCancelled"`
	DetachFront bool     `xml:"detachFront"`
	Length      int      `xml:"length"`
	AdhocAlerts []string `xml:"adhocAlerts>adhocAlertText`
}

/*

<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xmlns:xsd="http://www.w3.org/2001/XMLSchema">
<soap:Body>
<GetDepartureBoardResponse xmlns="http://thalesgroup.com/RTTI/2016-02-16/ldb/">
	<GetStationBoardResult xmlns:lt3="http://thalesgroup.com/RTTI/2015-05-14/ldb/types"
							xmlns:lt5="http://thalesgroup.com/RTTI/2016-02-16/ldb/types"
							xmlns:lt4="http://thalesgroup.com/RTTI/2015-11-27/ldb/types"
							xmlns:lt="http://thalesgroup.com/RTTI/2012-01-13/ldb/types"
							xmlns:lt2="http://thalesgroup.com/RTTI/2014-02-20/ldb/types">
		<lt4:generatedAt>2016-06-15T10:16:18.0069606+01:00</lt4:generatedAt>
		<lt4:locationName>Sheffield</lt4:locationName>
		<lt4:crs>SHF</lt4:crs>
		<lt4:nrccMessages>
			<lt:message>Trains between Birmingham New Street and Leicester / Derby are being delayed by up to 20 minutes. More information in &lt;A href="http://nationalrail.co.uk/service_disruptions/140323.aspx"&gt;Latest Travel News&lt;/A&gt;.</lt:message>
		</lt4:nrccMessages>
		<lt4:platformAvailable>true</lt4:platformAvailable>
		<lt5:trainServices>
			<lt5:service>
				<lt4:std>09:47</lt4:std>
				<lt4:etd>10:19</lt4:etd>
				<lt4:platform>1B</lt4:platform>
				<lt4:operator>CrossCountry</lt4:operator>
				<lt4:operatorCode>XC</lt4:operatorCode>
				<lt4:serviceType>train</lt4:serviceType>
				<lt4:delayReason>This train has been delayed by a tree blocking the railway</lt4:delayReason>
				<lt4:serviceID>80XwraLkmCdGh2Vmp0u65Q==</lt4:serviceID>
				<lt5:rsid>XC206000</lt5:rsid>
				<lt5:origin>
					<lt4:location>
						<lt4:locationName>Guildford</lt4:locationName>
						<lt4:crs>GLD</lt4:crs>
					</lt4:location>
				</lt5:origin>
				<lt5:destination>
					<lt4:location>
						<lt4:locationName>Newcastle</lt4:locationName>
						<lt4:crs>NCL</lt4:crs>
						<lt4:via>via Doncaster</lt4:via>
					</lt4:location>
				</lt5:destination>
			</lt5:service>
			<lt5:service>
				<lt4:std>10:10</lt4:std>
				<lt4:etd>10:14</lt4:etd>
				<lt4:platform>1B</lt4:platform>
				<lt4:operator>TransPennine Express</lt4:operator>
				<lt4:operatorCode>TP</lt4:operatorCode>
				<lt4:serviceType>train</lt4:serviceType>
				<lt4:serviceID>uaX8CX0JUfV7PxatLHg6OQ==</lt4:serviceID>
				<lt5:rsid>TP601600</lt5:rsid>
				<lt5:origin>
					<lt4:location>
						<lt4:locationName>Manchester Airport</lt4:locationName>
						<lt4:crs>MIA</lt4:crs>
					</lt4:location>
				</lt5:origin>
				<lt5:destination>
					<lt4:location>
						<lt4:locationName>Cleethorpes</lt4:locationName>
						<lt4:crs>CLE</lt4:crs>
					</lt4:location>
				</lt5:destination>
			</lt5:service>
			<lt5:service>
				<lt4:std>10:11</lt4:std><lt4:etd>10:15</lt4:etd><lt4:platform>6A</lt4:platform><lt4:operator>TransPennine Express</lt4:operator><lt4:operatorCode>TP</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:serviceID>xjiYYVIXcFv/4j/47716Ig==</lt4:serviceID><lt5:rsid>TP602500</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Cleethorpes</lt4:locationName><lt4:crs>CLE</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>Manchester Airport</lt4:locationName><lt4:crs>MIA</lt4:crs></lt4:location></lt5:destination></lt5:service><lt5:service><lt4:std>10:14</lt4:std><lt4:etd>10:16</lt4:etd><lt4:platform>1A</lt4:platform><lt4:operator>Northern</lt4:operator><lt4:operatorCode>NT</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:serviceID>l2BBb2eU3RJGF5HxL26w7Q==</lt4:serviceID><lt5:rsid>NT030900</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Sheffield</lt4:locationName><lt4:crs>SHF</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>Leeds</lt4:locationName><lt4:crs>LDS</lt4:crs></lt4:location></lt5:destination></lt5:service><lt5:service><lt4:std>10:18</lt4:std><lt4:etd>On time</lt4:etd><lt4:platform>2B</lt4:platform><lt4:operator>Northern</lt4:operator><lt4:operatorCode>NT</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:serviceID>teLi8G+FiRoZ+bTtwQNaWQ==</lt4:serviceID><lt5:rsid>NT656300</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Nottingham</lt4:locationName><lt4:crs>NOT</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>Leeds</lt4:locationName><lt4:crs>LDS</lt4:crs><lt4:via>via Barnsley</lt4:via></lt4:location></lt5:destination></lt5:service><lt5:service><lt4:std>10:21</lt4:std><lt4:etd>10:33</lt4:etd><lt4:platform>1B</lt4:platform><lt4:operator>CrossCountry</lt4:operator><lt4:operatorCode>XC</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:delayReason>This train has been delayed by an obstruction on the track</lt4:delayReason><lt4:serviceID>8RyniiMeSWbt5DJIPEmDXA==</lt4:serviceID><lt5:rsid>XC106000</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Plymouth</lt4:locationName><lt4:crs>PLY</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>Edinburgh</lt4:locationName><lt4:crs>EDB</lt4:crs><lt4:via>via Leeds</lt4:via></lt4:location></lt5:destination></lt5:service><lt5:service><lt4:std>10:24</lt4:std><lt4:etd>On time</lt4:etd><lt4:platform>6</lt4:platform><lt4:operator>CrossCountry</lt4:operator><lt4:operatorCode>XC</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:serviceID>LHTaewYOCM0oCEUXE7eYDg==</lt4:serviceID><lt5:rsid>XC260000</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Edinburgh</lt4:locationName><lt4:crs>EDB</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>Reading</lt4:locationName><lt4:crs>RDG</lt4:crs></lt4:location></lt5:destination></lt5:service><lt5:service><lt4:std>10:24</lt4:std><lt4:etd>On time</lt4:etd><lt4:platform>3A</lt4:platform><lt4:operator>Northern</lt4:operator><lt4:operatorCode>NT</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:serviceID>qrzMsQUqSGTTqDNKc0zGUg==</lt4:serviceID><lt5:rsid>NT958600</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Sheffield</lt4:locationName><lt4:crs>SHF</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>Scunthorpe</lt4:locationName><lt4:crs>SCU</lt4:crs></lt4:location></lt5:destination></lt5:service><lt5:service><lt4:std>10:29</lt4:std><lt4:etd>On time</lt4:etd><lt4:platform>5</lt4:platform><lt4:operator>East Midlands Trains</lt4:operator><lt4:operatorCode>EM</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:serviceID>aJXkp2Kpdx/zkCssmLeQOA==</lt4:serviceID><lt5:rsid>EM123500</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Sheffield</lt4:locationName><lt4:crs>SHF</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>London St Pancras (Intl)</lt4:locationName><lt4:crs>STP</lt4:crs></lt4:location></lt5:destination></lt5:service><lt5:service><lt4:std>10:36</lt4:std><lt4:etd>On time</lt4:etd><lt4:platform>4B</lt4:platform><lt4:operator>Northern</lt4:operator><lt4:operatorCode>NT</lt4:operatorCode><lt4:serviceType>train</lt4:serviceType><lt4:serviceID>Bc8fL7SMyjy8M+/6BMTJMg==</lt4:serviceID><lt5:rsid>NT961100</lt5:rsid><lt5:origin><lt4:location><lt4:locationName>Sheffield</lt4:locationName><lt4:crs>SHF</lt4:crs></lt4:location></lt5:origin><lt5:destination><lt4:location><lt4:locationName>Huddersfield</lt4:locationName><lt4:crs>HUD</lt4:crs></lt4:location></lt5:destination></lt5:service></lt5:trainServices></GetStationBoardResult></GetDepartureBoardResponse></soap:Body></soap:Envelope>


*/
