package wifiscanparser

import (
	. "launchpad.net/gocheck"
	"testing"
)

type WifiScanParserSuite struct{}

func Test(t *testing.T)                          { TestingT(t) }
func (t *WifiScanParserSuite) TearDownTest(c *C) {}
func (t *WifiScanParserSuite) SetupTest(c *C)    {}

var _ = Suite(&WifiScanParserSuite{})

func (t *WifiScanParserSuite) TestInitialisation(c *C) {
	wsp, err := NewWifiScanParser("data/test_scan", "airport")
	c.Check(err, IsNil)
	c.Check(wsp, NotNil)
}

func (t *WifiScanParserSuite) TestUnsupportedType(c *C) {
	wsp, err := NewWifiScanParser("data/test_scan_missing", "windows")
	c.Check(err, NotNil)
	c.Check(wsp, IsNil)
}

func (t *WifiScanParserSuite) TestNoFile(c *C) {
	wsp, err := NewWifiScanParser("data/test_scan_missing", "airport")
	c.Check(err, NotNil)
	c.Check(wsp, IsNil)
}

func (t *WifiScanParserSuite) TestParseSingleError(c *C) {

	// Initialise
	wsp, err := NewWifiScanParser("data/test_scan_single_error", "airport")

	// Pre-check
	c.Check(err, IsNil)
	c.Check(wsp, NotNil)

	// Build expected results
	expected := []*WifiInfo{}

	// Act
	actual := wsp.Parse()

	//Check
	c.Check(actual, DeepEquals, expected)
}

func (t *WifiScanParserSuite) TestParseSingle(c *C) {
	wsp, _ := NewWifiScanParser("data/test_scan_single", "airport")

	// Build expected results
	expected := []*WifiInfo{
		&WifiInfo{
			"Park Villas",
			"e0:3f:49:50:67:44",
			-89,
			"13,-1",
			true,
			"DE",
			"WPA2(PSK/AES/AES)",
		},
	}

	// Act
	actual := wsp.Parse()

	//Check
	c.Check(actual, DeepEquals, expected)
}

func (t *WifiScanParserSuite) TestParseFull(c *C) {
	wsp, _ := NewWifiScanParser("data/test_scan", "airport")

	// Build expected results
	expected := []*WifiInfo{
		&WifiInfo{
			"Park Villas",
			"e0:3f:49:50:67:44",
			-89,
			"13,-1",
			true,
			"DE",
			"WPA2(PSK/AES/AES)",
		},
		&WifiInfo{
			"Top Floor Cavendish",
			"30:91:8f:37:6e:2f",
			-86,
			"11",
			true,
			"--",
			"WPA(PSK/TKIP/TKIP) WPA2(PSK/AES/TKIP)",
		},
		&WifiInfo{
			"SKY67411",
			"34:08:04:2e:65:89",
			-84,
			"11",
			false,
			"--",
			"WPA(PSK/AES,TKIP/TKIP) WPA2(PSK/AES,TKIP/TKIP)",
		},
		&WifiInfo{
			"SKY08081",
			"c0:3e:0f:89:41:ed",
			-79,
			"11",
			true,
			"--",
			"WPA2(PSK/AES/AES)",
		},
		&WifiInfo{
			"BTWiFi-with-FON",
			"4a:55:9c:0d:91:fe",
			-72,
			"6",
			true,
			"--",
			"NONE",
		},
		&WifiInfo{
			"BTHub3-GHNF",
			"f4:55:9c:0d:91:fd",
			-71,
			"6",
			true,
			"--",
			"WPA(PSK/AES,TKIP/TKIP) WPA2(PSK/AES,TKIP/TKIP)",
		},
		&WifiInfo{
			"VM273083-2G",
			"e4:f4:c6:82:7b:c0",
			-55,
			"6",
			true,
			"GB",
			"WPA(PSK/AES,TKIP/TKIP) WPA2(PSK/AES,TKIP/TKIP)",
		},
		&WifiInfo{
			"VM273083-5G",
			"e4:f4:c6:6b:43:c0",
			-67,
			"44,+1",
			true,
			"GB",
			"WPA2(PSK/AES/AES)",
		},
	}

	// Act
	actual := wsp.Parse()

	// Check
	for i, exp := range expected {
		act := actual[i]
		c.Check(act, DeepEquals, exp)

	}
}
