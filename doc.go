// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package weathermap_go simple web site

	Live DEMO http://104.155.23.78:8080

	Angular2 code https://github.com/remotejob/weathermap_ang2_webpack.


	Use

	The points --> smallest Docker image, test web performatce (Angular 2):

	General:

	small Docker container image
	can be demployed on
		Amazon EC2 Container Service
		Google Container Engine
		Azure Container Service etc..

	Geo Location:
		get coordinate by
		free GeoLite2-City DB
		https://dev.maxmind.com/geoip/geoip2/geolite2/

	Weather data:
		by simple https://openweathermap.org/  API service

	Site created by Angular 2 framework:
		code splited and organized by Webpack

	Probably mostly intresting part:
		all static web files incorporate inside single file ./server
		how it possible --> https://github.com/rakyll/statik/

	Size:
		Doker image after deployment on Google Container Registry:
			size without GeoLite2-City.mmdn 3 MB!!
			total size 35 MB included GeoLite2-City.mmdn.

	Docker and Golang:
		Golang: definitely winning in relation to Docker compare with others languages

	Files:
		Makefile - compile and deploy Docker image
		Dockerfile - it's Dockerfile for Docker image creation
		deployment.yaml - kubernetes deployment on Google Container Engine
*/
package main
