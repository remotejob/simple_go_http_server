# simple_go_http_server
--
Package simple_go_http_server

Live DEMO http://159.203.134.218:32175

Use:

as utility for logs view:

General:

small Docker container image can be demployed on

    Amazon EC2 Container Service
    Google Container Engine
    Azure Container Service etc..
    as well can be used as stand alone application

simple deployment on docker:

    docker run -t -p 9090:8080 remotejob/simple_go_http_server:0.3
    test on http://localhost:9090

stand alone use:

    ./server
    test on http://localhost:8080
    binary ./server compiled for Linux ELF 64-bit LSB  executable, x86-64, version 1 (SYSV), statically linked, not stripped
    so to use under Windows an Mac it must be recompiled

database:

    simple .csv file (logs.csv) it can be something more serious. I thinks Redis will be best in that case.

deployment on cluster:

    file deployment.yaml can be useful
    file intend use kubernetes for cluster orchestration.

size:

    Doker image after deployment on Registry:
    size 3.963 MB!! It's greate!

Docker and Golang:

    Golang: definitely winning in relation to Docker compare with others languages

Files:

    Makefile - compile and deploy Docker image
    Dockerfile - it's Dockerfile for Docker image creation
    deployment.yaml - kubernetes deployment on any Cluster Containers
