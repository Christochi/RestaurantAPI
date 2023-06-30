# Restaurant API
## Description
JSON API allows one to update the food menu of a Restaurant and the Chefs. You can also perform normal CRUD operations via the endpoints

## Requirement
- Install the lastest Golang
- Install Docker (only necessary for using the build in a docker container)
- HTTP Client for interacting with the API: Postman, Curl, Thunder Client, Web Browser or any client of your choice

## Setup
There are 2 ways to run this project: `cloning the project` and `docker build`

#### Cloning the Project
- Clone the project
- Run the project: 
    - from the command line, run `go run main.go`
    - on your browser or HTTP client, enter `http://localhost:3000`. The default port is ***3000***
    - on change the port or address, for example, ***7000***, on the commad line, run `go run main.go --listenaddr :7000`
    - on your browser or HTTP client, enter `http://localhost:7000` 

#### Docker Build
***COMING SOON...***

## Endpoints
#### Menu
- **Add a Food Menu**

