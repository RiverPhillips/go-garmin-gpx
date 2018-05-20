package gpx

import (
	"encoding/xml"
	"io/ioutil"
)

// ParseFile takes a file and parses it
func ParseFile(fileName string) (*GPX, error) {
	g := GPX{}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return &g, err
	}

	err = Parse(bytes, &g)
	if err != nil {
		return &g, err
	}
	return &g, nil
}

// Parse bytes of xml
func Parse(bytes []byte, g *GPX) error {
	err := xml.Unmarshal(bytes, g)
	if err != nil {
		return err
	}
	return nil
}

// Comments from http://www.topografix.com/GPX/1/1/

// GPX is the root element
type GPX struct {
	XMLName   xml.Name   `xml:"gpx"`
	Version   string     `xml:"version,attr"`
	Creator   string     `xml:"creator,attr"`
	Metadata  Metadata   `xml:"metadata,omitempty"`
	Waypoints []WayPoint `xml:"wpt,omitempty"`
	Routes    []Route    `xml:"rte,omitempty"`
	Tracks    Track      `xml:"trk"`
}

// Metadata has information about the GPX file
type Metadata struct {
	XMLName     xml.Name  `xml:"metadata"`
	Name        string    `xml:"name,omitempty"`
	Description string    `xml:"desc,omitempty"`
	Author      Person    `xml:"author,omitempty"`
	Copyright   Copyright `xml:"copyright,omitempty"`
	Links       []Link    `xml:"link,omitempty"`
	Timestamp   string    `xml:"time,omitempty"`
	Keywords    string    `xml:"keywords,omitempty"`
	Bounds      Bounds    `xml:"bounds"`
	Extensions  Extension `xml:"extensions,omitempty"`
}

// WayPoint is a point of interest, or named feature on a map.
type WayPoint struct {
	Latitude                      Latitude    `xml:"lat,attr"`
	Longitude                     Longitude   `xml:"lon,attr"`
	Elevation                     float64     `xml:"ele,omitempty"`
	Timestamp                     string      `xml:"time,omitempty"`
	MagneticVariation             Degrees     `xml:"magvar,omitempty"`
	GeoIDHeight                   string      `xml:"geoidheight,omitempty"`
	Name                          string      `xml:"name,omitempty"`
	Comment                       string      `xml:"cmt,omitempty"`
	Description                   string      `xml:"desc,omitempty"`
	Source                        string      `xml:"src,omitempty"`
	Links                         []Link      `xml:"link"`
	Symbol                        string      `xml:"sym,omitempty"`
	Type                          string      `xml:"type,omitempty"`
	Fix                           Fix         `xml:"fix,omitempty"`
	Sat                           int         `xml:"sat,omitempty"`
	HorizontalDilutionOfPrecision float64     `xml:"hdop,omitempty"`
	VerticalDilutionOfPrecision   float64     `xml:"vdop,omitempty"`
	PositionDilutionOfPrecision   float64     `xml:"pdop,omitempty"`
	AgeOfGpsData                  float64     `xml:"ageofgpsdata,omitempty"`
	DifferentialGPSID             DGPSStation `xml:"dgpsid,omitempty"`
	Extensions                    Extension   `xml:"extensions,omitempty"`
}

// Route is an ordered list of Waypoints representing a series of points leading to a destination.
type Route struct {
	XMLName     xml.Name   `xml:"rte"`
	Name        string     `xml:"name,omitempty"`
	Comment     string     `xml:"cmt,omitempty"`
	Description string     `xml:"desc,omitempty"`
	Source      string     `xml:"src,omitempty"`
	Links       []Link     `xml:"link"`
	Number      int        `xml:"number,omitempty"`
	Type        string     `xml:"type,omitempty"`
	Extensions  Extension  `xml:"extensions,omitempty"`
	RoutePoints []WayPoint `xml:"rtept"`
}

// Track represents a track - an ordered list of points describing a path
type Track struct {
	XMLName       xml.Name       `xml:"trk"`
	Name          string         `xml:"name,omitempty"`
	Comment       string         `xml:"cmt,omitempty"`
	Description   string         `xml:"desc,omitempty"`
	Source        string         `xml:"src,omitempty"`
	Links         []Link         `xml:"link"`
	Number        int            `xml:"number,omitempty"`
	Type          string         `xml:"type,omitempty"`
	Extensions    Extension      `xml:"extensions,omitempty"`
	TrackSegments []TrackSegment `xml:"trkseg"`
}

// Extension extend GPX by adding your own elements from another schema
type Extension struct {
	XMLName              xml.Name            `xml:"extensions"`
	TrackPointExtensions TrackPointExtension `xml:"TrackPointExtension,omitempty"`
}

// TrackPointExtension tracks temperature, heart rate and cadence specific to garmin devices
type TrackPointExtension struct {
	XMLName     xml.Name `xml:"TrackPointExtension"`
	Temperature int      `xml:"atemp,omitempty"`
	HeartRate   int      `xml:"hr,omitempty"`
	Cadence     int      `xml:"cad,omitempty"`
}

// TrackSegment has a list of continious span of TrackPoints
type TrackSegment struct {
	XMLName    xml.Name   `xml:"trkseg"`
	TrackPoint []WayPoint `xml:"trkpt"`
	Extensions Extension  `xml:"extensions,omitempty"`
}

// Copyright has information about holder and license
type Copyright struct {
	XMLName xml.Name `xml:"copyright"`
	Author  string   `xml:"author,attr"`
	Year    string   `xml:"year,omitempty"`
	License string   `xml:"license,omitempty"`
}

// Link is for an external resource with additional information.
type Link struct {
	XMLName xml.Name `xml:"link"`
	URL     string   `xml:"href,attr,omitempty"`
	Text    string   `xml:"text,omitempty"`
	Type    string   `xml:"type,omitempty"`
}

// Email address which is broken into two parts (id and domain)
type Email struct {
	XMLName xml.Name `xml:"email"`
	ID      string   `xml:"id,attr,omitempty"`
	Domain  string   `xml:"domain,attr,omitempty"`
}

// Person is a person or an organisation
type Person struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name,omitempty"`
	Email   Email    `xml:"email,omitempty"`
	Link    Link     `xml:"link,omitempty"`
}

// Point with optional elevation and time
type Point struct {
	XMLName   xml.Name  `xml:"pt"`
	Latitude  Latitude  `xml:"lat,attr"`
	Longitude Longitude `xml:"lon,attr"`
	Elevation float64   `xml:"ele,omitempty"`
	Timestamp string    `xml:"time,omitempty"`
}

// PointSegment is a sequence of Points
type PointSegment struct {
	XMLName xml.Name `xml:"ptseg"`
	Points  []Point  `xml:"pt"`
}

// Bounds are two latitude longitude pairs defining the extent of an element.
type Bounds struct {
	XMLName xml.Name `xml:"bounds"`
	MinLat  float64  `xml:"minlat,attr"`
	MaxLat  float64  `xml:"maxlat,attr"`
	MinLon  float64  `xml:"minlon,attr"`
	MaxLon  float64  `xml:"maxlon,attr"`
}

// Latitude is the latitude of the point. Decimal degrees, WGS84 datum. The value varies between -90.0 to 90.0
type Latitude float64

// Longitude is the longitude of the point. Decimal degrees, WGS84 datum. The value varies between -180.0 and 180.0
type Longitude float64

// Degrees is used for bearing, heading, course. Units are decimal degrees, true (not magnetic). The value varies between 0.0 and 360.0
type Degrees float64

// Fix represents type of GPS fix
type Fix string

const (
	// None means we didn't get a fix
	None Fix = "none"
	// TwoDimensional fix
	TwoDimensional Fix = "2d"
	// ThreeDimensional fix
	ThreeDimensional Fix = "3d"
	// DGPS means a digital GPS fix
	DGPS Fix = "dgps"
	// PPS means that a military signal was used
	PPS Fix = "pps"
)

// DGPSStation represents a differential GPS station and varies between 0 to 1023
type DGPSStation int
