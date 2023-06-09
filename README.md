# Restaurant API
## Description
The REST API is used to update the food menu of a Restaurant and the Chefs profile of the webpage. You can also perform normal CRUD operations via the endpoints (`/menu/` and `/chef/`).

The data from the API is displayed on the webpage after refreshing the webpage. Refreshing automatically sends a **GET** request to the endpoint to retrieve the data

## Requirement
- Install the lastest Golang (required for first setup)
- Install Docker (required for second setup)
- HTTP Client for interacting with the API: Postman, Curl, Thunder Client, Web Browser or any client of your choice

## Setup
There are 2 ways to run this project: `cloning the project` and `docker build`

### Cloning the Project
- Clone the project
- Run the project: 
    - from the command line, start the web server by running `go run main.go`
    - on your browser or HTTP client, enter `http://localhost:3000` to **load the web page**. The default port is ***3000***
    - to change the port or address, for example to ***7000***, on the commad line, run `go run main.go --listenaddr :7000`
    - on your browser or HTTP client like Postman, enter `http://localhost:7000` 

        [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/25499133-d07cee34-4e11-4ae7-ad33-643cd9a6e3b5?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D25499133-d07cee34-4e11-4ae7-ad33-643cd9a6e3b5%26entityType%3Dcollection%26workspaceId%3D6089f5e4-d1b8-4a4c-ba19-c5bec8f7facc)



### Docker Build
***COMING SOON...***

## Endpoints
### Menu
- **Add a Food Menu**: POST `http://localhost:3000/menu/`
~~~
Example of JSON body to send as part of the POST request

[
    {
        "type" : "Breakfast",
        "meal" : "Potato with Bacon",
        "price" : "$5",
        "desc": "mildly fried potatoes with bacon"
    },  

    {
        "type" : "Lunch",
        "meal" : "Chicken Fried Rice",
        "price" : "$15",
        "desc": "Rice with stirred fried vege with baked chicken"
    }
]
~~~
- **Returns list of Food Menu**: GET `http://localhost:3000/menu/`. **GET** request retrieves the data and it is displayed on the browser
- **Returns list of a meal type in the food menu**: There are 4 acceptable meal types: ***Breakfast, Lunch, Dinner & Drinks***.
GET `http://localhost:3000/menu/drinks` or GET `http://localhost:3000/menu/lunch`
- **Returns Meals that match a search pattern**: GET `http://localhost:3000/menu/breakfast/hotdog` or `http://localhost:3000/menu/dinner/rice`
- **Deletes all Menu**: DELETE `http://localhost:3000/menu/`
- **Deletes a Meal that matches a search pattern**: DELETE `http://localhost:3000/menu/drinks/mangolasse`

### Chef
- **Add a Chef**: POST `http://localhost:3000/chef/`
~~~
Example of JSON body to send as part of the POST request

[
    {
        "Name": "John Doe",
        "About": "John has 10 years experience making delicious meals for 5-star hotels and famous restaurants in North America"
    },

    {
        "Name": "Chocho Okoye",
        "About": "Chocho has 20 years experience cooking for famous restaurants in African and the Caribbean"
    }
]
~~~
- **Returns list of Chefs**: GET `http://localhost:3000/chef/`. **GET** request retrieves the data and it is displayed on the browser
- **Returns a Chef that matches a search pattern**: GET `http://localhost:3000/chef/john`
- **Deletes all Chefs**: DELETE `http://localhost:3000/chef/`
- **Deletes a Chef that matches a search pattern**: DELETE `http://localhost:3000/chef/johndoe`
