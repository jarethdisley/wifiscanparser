package wifiscanparser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	// AirportRE is a regex for parsing Airport scan output
	AirportRE = "\\s*([a-zA-Z0-9-_\\s]*)\\s*([a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2}:[a-fA-F0-9]{2})\\s*([-|+]{1}[0-9]*)\\s*([0-9]*,*[-|+]*[0-9]*)\\s*([Y|N]{1})\\s*([A-Z-]*)\\s*(.*)"
)

// WifiScanParser parses the output from an airport WIFI scan and generates
// a WifiInfo struct for each network found.
type WifiScanParser struct {
	path string
	typ  string
}

// WifiInfo represents meta data about a WIFI network
type WifiInfo struct {
	SSID, BSSID string
	RSSI        int
	Channel     string
	HT          bool
	CountryCode string
	Security    string
}

func (wfi *WifiInfo) toString() string {
	return fmt.Sprintf("WifiInfo:\n\tSSID: %s\n\tBSSID: %s\n\tRSSI: %d\n\tChannel: %s\n\tHT: %t\n\tContryCode: %s\n\tSecurity: %s", wfi.SSID, wfi.BSSID, wfi.RSSI, wfi.Channel, wfi.HT, wfi.CountryCode, wfi.Security)
}

// NewWifiScanParser initialises an instance of WifiScanParser
func NewWifiScanParser(path, typ string) (*WifiScanParser, error) {

	// Check we are parsing a supported type (Only support airport for now)
	switch typ {
	case "airport":
		break
	default:
		fmt.Printf("Unsupported parse type '%s'\n", typ)
		return nil, fmt.Errorf("Unsupported parse type '%s'", typ)
	}

	// Check that the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", path)
		return nil, err
	}

	// Finally initialise an instance of WifiScanParser
	parser := &WifiScanParser{
		path: path,
		typ:  typ,
	}

	return parser, nil
}

// Parse opens / reads the file initially specified and parses each of the lines
func (wsp *WifiScanParser) Parse() []*WifiInfo {

	// Compile the regex
	re := regexp.MustCompile(AirportRE)

	// Read in the scan file (defer the close)
	inFile, err := os.Open(wsp.path)
	if err != nil {
		fmt.Printf("Error reading " + wsp.path + "\n")
	}
	defer inFile.Close()

	// Read each of the lines
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	// Parse each of the lines
	networks := []*WifiInfo{}
	for scanner.Scan() {
		network := processSubmatch(re.FindStringSubmatch(scanner.Text()))
		if network != nil {
			fmt.Println(network.toString())
			networks = append(networks, network)
		}
	}
	return networks
}

func processSubmatch(matches []string) *WifiInfo {

	// Ignore non matches
	if len(matches) == 0 {
		return nil
	}

	// First match of regex is always the full string
	matches = matches[1:]

	// Parse the RSSI int
	rssi, rerr := strconv.ParseInt(matches[2], 10, 32)
	if rerr != nil {
		return nil
	}

	// Parse the HT boolean
	var ht bool
	if matches[4] == "Y" {
		ht = true
	} else {
		ht = false
	}

	// Trim each of the strings and build the struct
	return &WifiInfo{
		SSID:        strings.Trim(matches[0], " "),
		BSSID:       strings.Trim(matches[1], " "),
		RSSI:        int(rssi),
		Channel:     matches[3],
		HT:          ht,
		CountryCode: strings.Trim(matches[5], " "),
		Security:    strings.Trim(matches[6], " "),
	}
}
