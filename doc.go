// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package simple_go_http_server

	Live DEMO http://159.203.107.223:30458


	Use

	as utility for logs view:

	General:

	small Docker container image
	can be demployed on
		Amazon EC2 Container Service
		Google Container Engine
		Azure Container Service etc..
		as well can be used as stand alone application

	simple deployment on docker:
		docker run -t -p 9090:8080 remotejob/simple_go_http_server:0.0
		test on http://localhost:9090

	stand alone use simple:
		./server
		test on http://localhost:8080
		binary ./server compiled for Linux ELF 64-bit LSB  executable, x86-64, version 1 (SYSV), statically linked, not stripped
		so to use under Windows an Mac it must be recompiled

	deployment on cluster:
		file deployment.yaml can be useful

	size:
		Doker image after deployment on Registry:
		size 3.963 MB!! It's greate!


	Docker and Golang:
		Golang: definitely winning in relation to Docker compare with others languages

	Files:
		Makefile - compile and deploy Docker image
		Dockerfile - it's Dockerfile for Docker image creation
		deployment.yaml - kubernetes deployment on any Containers
*/
package main
