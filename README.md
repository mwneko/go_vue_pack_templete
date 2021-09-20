# go_vue_pack
This is a SPA based on Vue.js and Golang with sample data

# Introduction

This is a SPA web application

- Front-end: Vue.js
- Back-end: Golang

This text you see here is *actually- written in Markdown! To get a feel
for Markdown's syntax, type some text into the left window and
watch the results in the right.

## CheckList

The functions implemented are listed here:

- Searchable forms based on Vue. Js and vuetify style
- MongoDB connection and query
- Golang models and api
- Axios


## Installation

Requires [Node.js](https://nodejs.org/) v10+ to run.
npm v7.21.1
node.js v16.9.1
go v1.17.1
vue2 @vue/cli 4.5.3

For Vue.js environments
```sh
git clone https://github.com/mwneko/go_vue_pack.git
cd frontend
npm i
npm run build
```

For golang environments
mention: you need clone the project in $GOPATH 
or set current scr to $GOPATH

```sh
export GOPATH=`pwd`
```
Instructions that may be required
```sh
cd main
go get -u github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver
dep init
```
Start the program

one terminal
```sh
cd main
go run main.go
```
second terminal
```sh
cd frontend
npm run serve
```
address: http://localhost:8081/ 
mentions: due to the program not completed, you need use Moesif Origin & CORS Changer for show the sampleData

## Screenshot
![image](https://github.com/mwneko/go_vue_pack/master/screenshot.jpg)


## Question
It uses for mentions me update the project

- Import .csv file script
- Change mongoDB data to .Json
- Merge front end and back end
- Connect to PostgreSQL
- Golang import install
- Datepicker

