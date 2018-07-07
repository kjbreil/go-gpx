// Copyright 2013, 2014 Peter Vasil. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gpx

import (
	"log"
	"testing"
	"time"
)

var g *GPX

func init() {
	log.Println("gpx test init")
}

func TestParse(t *testing.T) {
	var err error
	g, err = ParseFile("testdata/file.gpx")

	if err != nil {
		t.Error("Error parsing GPX file: ", err)
	}

	// t.Log("Test parser")
	timestampA := g.Metadata.Timestamp
	timestampE := "2012-03-17T15:44:18Z"
	if timestampA != timestampE {
		t.Errorf("timestamp expected: %s, actual: %s", timestampE, timestampA)
	}

	trknameA := g.Tracks[0].Name
	trknameE := "17-MRZ-12 16:44:12"
	if trknameA != trknameE {
		t.Errorf("Trackname expected: %s, actual: %s", trknameE, trknameA)
	}

	numPointsA := len(g.Tracks[0].Segments[0].Waypoints)
	numPointsE := 4
	if numPointsE != numPointsA {
		t.Errorf("Number of tracks expected: %d, actual: %d", numPointsE, numPointsA)
	}
}

func TestParseNonUtf8(t *testing.T) {
	g2, err := ParseFile("testdata/file2.gpx")

	if err != nil {
		t.Error("Error parsing GPX file: ", err)
	}

	timestampA := g2.Metadata.Timestamp
	timestampE := "2012-03-17T15:44:18Z"
	if timestampA != timestampE {
		t.Errorf("timestamp expected: %s, actual: %s", timestampE, timestampA)
	}
}

func TestLength2DSeg(t *testing.T) {
	lengthA := g.Tracks[0].Segments[0].Length2D()
	lengthE := 56.77577732775905

	if lengthA != lengthE {
		t.Errorf("Length 2d expected: %f, actual %f", lengthE, lengthA)
	}
}

func TestLength3DSeg(t *testing.T) {
	lengthA := g.Tracks[0].Segments[0].Length3D()
	lengthE := 61.76815317436073

	if lengthA != lengthE {
		t.Errorf("Length 3d expected: %f, actual %f", lengthE, lengthA)
	}
}

func TestTimePoint(t *testing.T) {
	timeA := g.Tracks[0].Segments[0].Waypoints[0].Time()
	//2012-03-17T12:46:19Z
	timeE := time.Date(2012, 3, 17, 12, 46, 19, 0, time.UTC)

	if timeA != timeE {
		t.Errorf("Time expected: %s, actual: %s", timeE.String(), timeA.String())
	}
}

func TestTimeBoundsSeg(t *testing.T) {
	startA, endA := g.Tracks[0].Segments[0].TimeBounds()
	startE := time.Date(2012, 3, 17, 12, 46, 19, 0, time.UTC)
	endE := time.Date(2012, 3, 17, 12, 47, 23, 0, time.UTC)

	if startA != startE || endA != endE {
		t.Errorf("TimeBounds expected: %v, %v, actual: %v, %v", startE, endE, startA, endA)
	}
}

func TestBoundsSeg(t *testing.T) {
	boundsA := g.Tracks[0].Segments[0].Bounds()
	boundsE := Bounds{
		MaxLat: 52.5117189623, MinLat: 52.5113534275,
		MaxLon: 13.4571944922, MinLon: 13.4567520116,
	}

	if boundsE != *boundsA {
		t.Errorf("Bounds expected: %s, actual: %s", boundsE.String(), boundsA.String())
	}
}

func TestBoundsGpx(t *testing.T) {
	boundsA := g.Bounds()
	boundsE := Bounds{
		MaxLat: 52.5117189623, MinLat: 52.5113534275,
		MaxLon: 13.4571944922, MinLon: 13.4567520116,
	}

	if boundsE != *boundsA {
		t.Errorf("Bounds expected: %s, actual: %s", boundsE.String(), boundsA.String())
	}
}

func TestSpeedSeg(t *testing.T) {
	speedA := g.Tracks[0].Segments[0].Speed(2)
	speedE := 1.5386074011963367

	if speedE != speedA {
		t.Errorf("Speed expected: %f, actual: %f", speedE, speedA)
	}
}

func TestDurationSeg(t *testing.T) {
	durA := g.Tracks[0].Segments[0].Duration()
	durE := 64.0

	if durE != durA {
		t.Errorf("Duration expected: %f, actual: %f", durE, durA)
	}
}

func TestUphillDownHillSeg(t *testing.T) {
	upA, downA := g.Tracks[0].Segments[0].UphillDownhill()
	upE := 5.863000000000007
	downE := 1.5430000000000064

	if upA != upE || downA != downE {
		t.Errorf("UphillDownhill expected: %+v, %+v, actual: %+v, %+v", upA, downA, upE, downE)
	}
}

func TestMovingData(t *testing.T) {
	movDataA := g.MovingData()
	movDataE := MovingData{
		MovingTime:      39.0,
		StoppedTime:     25.0,
		MovingDistance:  55.28705571308896,
		StoppedDistance: 6.481097461271765,
		MaxSpeed:        0.0,
	}

	if movDataE != *movDataA {
		t.Errorf("Moving data expected: %+v, actual: %+v", movDataE, movDataA)
	}
}

func TestUphillDownhill(t *testing.T) {
	upA, downA := g.UphillDownhill()
	upE := 5.863000000000007
	downE := 1.5430000000000064

	if upA != upE || downA != downE {
		t.Errorf("UphillDownhill expected: %+v, %+v, actual: %+v, %+v", upA, downA, upE, downE)
	}
}

func TestToXml(t *testing.T) {
	xmlA := string(g.ToXML())
	xmlE := `<?xml version="1.0" encoding="UTF-8"?>
<gpx xmlns="http://www.topografix.com/GPX/1/1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd" version="1.1" creator="eTrex 10">
	<metadata>
		<link href="http://www.garmin.com">
			<text>Garmin International</text>
		</link>
		<time>2012-03-17T15:44:18Z</time>
	</metadata>
	<wpt lat="37.085751" lon="-121.17042">
		<ele>195.440933</ele>
		<time>2012-03-21T21:24:43Z</time>
		<name>001</name>
		<sym>Flag, Blue</sym>
	</wpt>
	<wpt lat="37.085751" lon="-121.17042">
		<ele>195.438324</ele>
		<time>2012-03-21T21:24:44Z</time>
		<name>002</name>
		<sym>Flag, Blue</sym>
	</wpt>
	<trk>
		<name>17-MRZ-12 16:44:12</name>
		<trkseg>
			<trkpt lat="52.5113534275" lon="13.4571944922">
				<ele>59.26</ele>
				<time>2012-03-17T12:46:19Z</time>
			</trkpt>
			<trkpt lat="52.5113568641" lon="13.4571697656">
				<ele>65.51</ele>
				<time>2012-03-17T12:46:44Z</time>
			</trkpt>
			<trkpt lat="52.511710329" lon="13.456941694">
				<ele>65.99</ele>
				<time>2012-03-17T12:47:01Z</time>
			</trkpt>
			<trkpt lat="52.5117189623" lon="13.4567520116">
				<ele>63.58</ele>
				<time>2012-03-17T12:47:23Z</time>
			</trkpt>
		</trkseg>
	</trk>
</gpx>`

	if xmlE != xmlA {
		t.Errorf("XML expected: \n%s, \nactual \n%s", xmlE, xmlA)
	}
}

func TestNewXml(t *testing.T) {
	gpx := NewGpx()
	gpxTrack := Trk{}

	gpxSegment := Trkseg{}
	gpxSegment.Waypoints = append(gpxSegment.Waypoints, Waypoint{Lat: 2.1234, Lon: 5.1234, Ele: 1234})
	gpxSegment.Waypoints = append(gpxSegment.Waypoints, Waypoint{Lat: 2.1233, Lon: 5.1235, Ele: 1235})
	gpxSegment.Waypoints = append(gpxSegment.Waypoints, Waypoint{Lat: 2.1235, Lon: 5.1236, Ele: 1236})

	gpxTrack.Segments = append(gpxTrack.Segments, gpxSegment)
	gpx.Tracks = append(gpx.Tracks, gpxTrack)

	actualXML := string(toXML(gpx))
	expectedXML := `<gpx xmlns="http://www.topografix.com/GPX/1/1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.topografix.com/GPX/1/1 http://www.topografix.com/GPX/1/1/gpx.xsd" version="1.1" creator="https://github.com/ptrv/go-gpx">
	<trk>
		<trkseg>
			<trkpt lat="2.1234" lon="5.1234">
				<ele>1234</ele>
			</trkpt>
			<trkpt lat="2.1233" lon="5.1235">
				<ele>1235</ele>
			</trkpt>
			<trkpt lat="2.1235" lon="5.1236">
				<ele>1236</ele>
			</trkpt>
		</trkseg>
	</trk>
</gpx>`

	if expectedXML != actualXML {
		t.Errorf("XML expected:\n%s\n, actual: \n%s", expectedXML, actualXML)
	}
}

func TestSplitGpx(t *testing.T) {
	gTmp := g.Clone()
	gTmp.Split(0, 0, 2)

	gpxSegs := gTmp.Tracks[0].Segments

	segsLenA := len(gpxSegs)
	segsLenE := 2
	if segsLenE != segsLenA {
		t.Errorf("GPX segments count expected: %d, actual %d", segsLenE, segsLenA)
	}

	segsA := string(toXML(gpxSegs))
	segsE := `<trkseg>
	<trkpt lat="52.5113534275" lon="13.4571944922">
		<ele>59.26</ele>
		<time>2012-03-17T12:46:19Z</time>
	</trkpt>
	<trkpt lat="52.5113568641" lon="13.4571697656">
		<ele>65.51</ele>
		<time>2012-03-17T12:46:44Z</time>
	</trkpt>
	<trkpt lat="52.511710329" lon="13.456941694">
		<ele>65.99</ele>
		<time>2012-03-17T12:47:01Z</time>
	</trkpt>
</trkseg>
<trkseg>
	<trkpt lat="52.5117189623" lon="13.4567520116">
		<ele>63.58</ele>
		<time>2012-03-17T12:47:23Z</time>
	</trkpt>
</trkseg>`

	if segsE != segsA {
		t.Errorf("XML expected:\n%s\n, actual: \n%s", segsE, segsA)
	}
}

func TestJoin(t *testing.T) {
	gTmp := g.Clone()
	track := gTmp.Tracks[0]

	track.Split(0, 2)
	track.Join(0, 1)

	segsA := string(toXML(track.Segments))
	segsE := `<trkseg>
	<trkpt lat="52.5113534275" lon="13.4571944922">
		<ele>59.26</ele>
		<time>2012-03-17T12:46:19Z</time>
	</trkpt>
	<trkpt lat="52.5113568641" lon="13.4571697656">
		<ele>65.51</ele>
		<time>2012-03-17T12:46:44Z</time>
	</trkpt>
	<trkpt lat="52.511710329" lon="13.456941694">
		<ele>65.99</ele>
		<time>2012-03-17T12:47:01Z</time>
	</trkpt>
	<trkpt lat="52.5117189623" lon="13.4567520116">
		<ele>63.58</ele>
		<time>2012-03-17T12:47:23Z</time>
	</trkpt>
</trkseg>`

	if segsE != segsA {
		t.Errorf("XML expected:\n%s\n, actual: \n%s", segsE, segsA)
	}
}

func TestClone(t *testing.T) {
	gTmp := g.Clone()
	gTmp.Metadata.Timestamp = "2012-03-17T15:44:19Z"

	if gTmp.Metadata.Timestamp == g.Metadata.Timestamp {
		t.Error("Clone failed")
	}
}
