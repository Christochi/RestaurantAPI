# Restaurant API
## Description
JSON API allows one to update the food menu of a Restaurant and the Chefs. You can also perform normal CRUD operations via the endpoints

## Requirement
- Install the lastest Golang (required for first setup)
- Install Docker (required for second setup)
- HTTP Client for interacting with the API: Postman, Curl, Thunder Client, Web Browser or any client of your choice

## Setup
There are 2 ways to run this project: `cloning the project` and `docker build`

### Cloning the Project
- Clone the project
- Run the project: 
    - from the command line, run `go run main.go`
    - on your browser or HTTP client, enter `http://localhost:3000`. The default port is ***3000***
    - to change the port or address, for example to ***7000***, on the commad line, run `go run main.go --listenaddr :7000`
    - on your browser or HTTP client, enter `http://localhost:7000` 

### Docker Build
***COMING SOON...***

## Endpoints
### Menu
- **Add a Food Menu**: POST `http://localhost/menu/`
~~~
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
- **Returns list of Food Menu**: GET `http://localhost/menu/`
- **Returns list of a meal type in the food menu**: There are 4 acceptable meal type: ***Breakfast, Lunch, Dinner & Drinks***.
GET `http://localhost/drinks` or GET `http://localhost/lunch`
- **Returns Meals that match a search pattern**: GET `http://localhost/menu/breakfast/hotdog` or `http://localhost/menu/dinner/rice`
- **Deletes all Menu**: DELETE `http://localhost/menu/`
- **Deletes a Meal that matches a search pattern**: DELETE `http://localhost/menu/drinks/mangolasse`

### Chef
- **Add a Chef**: POST `http://localhost/chef/`
~~~
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
- **Returns list of Chefs**: GET `http://localhost/chef/`
- **Returns a Chef that matches a search pattern**: GET `http://localhost/chef/John`
- **Deletes all Chefs**: DELETE `http://localhost/chef/`
- **Deletes a Chef that matches a search pattern**: DELETE `http://localhost/chef/johndoe`
