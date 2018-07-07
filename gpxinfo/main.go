// Copyright 2013, 2014 Peter Vasil. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/kjbreil/go-gpx"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please provide a GPX file path!")
		return
	}

	gpxFileArg := args[0]
	gpxFile, err := gpx.ParseFile(gpxFileArg)

	if err != nil {
		fmt.Println("Error opening gpx file: ", err)
		return
	}

	gpxPath, _ := filepath.Abs(gpxFileArg)
	fmt.Println("File: ", gpxPath)

	if gpxFile.Metadata != nil {
		fmt.Println("\tGPX name: ", gpxFile.Metadata.Name)
		fmt.Println("\tGPX desctiption: ", gpxFile.Metadata.Desc)
		if gpxFile.Metadata.Author != nil {
			fmt.Println("\tAuthor: ", gpxFile.Metadata.Author.Name)
			if gpxFile.Metadata.Author.Email != nil {
				fmt.Println("\tEmail: ", gpxFile.Metadata.Author.Email)
			}
		}
	}

	len2d := gpxFile.Length2D()
	len3d := gpxFile.Length3D()
	fmt.Println("\tLength 2D: ", len2d/1000.0)
	fmt.Println("\tLength 3D: ", len3d/1000.0)

	fmt.Printf("\tBounds: %+v\n", gpxFile.Bounds())

	md := gpxFile.MovingData()
	fmt.Println("\tMoving time: ", md.MovingTime)
	fmt.Println("\tStopped time: ", md.StoppedTime)

	fmt.Printf("\tMax speed: %fm/s = %fkm/h\n", md.MaxSpeed, md.MaxSpeed*60*60/1000.0)

	uphill, downhill := gpxFile.UphillDownhill()
	fmt.Println("\tTotal uphill: ", uphill)
	fmt.Println("\tTotal downhill: ", downhill)

	start, end := gpxFile.TimeBounds()
	fmt.Println("\tStarted: ", start)
	fmt.Println("\tEnded: ", end)

	fmt.Println()
}
