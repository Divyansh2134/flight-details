# flight-details

# Frontend

## working locally

- Download dependencies
```npm ci```

- Starting server locally
```npm run dev```

- Access using url ```http://localhost:5173/```

## working with docker

- Build Docker file using docker command for example:- ```docker build -t  frontend:1.0 .``` 

- Run Docker Image:- ```docker run -p 80:80 frontend:1.0```

- Access using url ```http://localhost:80/```

# Backend 

## Working Locally

- Download dependencies ```go mod download```

- For running Locally ```go run .```

- It has two post urls to access 

- Creating flight data
```http://localhost:8000/create-flight```
 for creating flights database in mongodb(make sure your local mongodb is running)
with data: 
```
{
	"flight": "AA126",
	"status": "On Time",
	"gate": "A2",
	"remarks": "Boarding"
}
```

- Updating the existing data ```http://localhost:8000/update-flight-details```

with data :
```
{
	"flight": "AA126",
	"status": "On Time",
	"gate": "A3",
	"remarks": "Gate Changed to A3"
}
```

## Docker build

- Build Docker file using docker command for example:- ```docker build -t  backend:1.0 .``` 

- Run Docker Image:- ```docker run -p 8000:8000 backend:1.0```

- Access using url Ports create: ```http://localhost:8000/create-flight```
update: ```http://localhost:8000/update-flight-details```

# Database

- I am using MongoDB as a Database and have choose my mock data myself as it is not provide.

- First start MongoDb locally.

- create flight data POST request on ```http://localhost:8000/create-flight``` 
with data structued as
```
{
	"flight": "AA126",
	"status": "On Time",
	"gate": "A2",
	"remarks": "Boarding"
}
```

- Update Flight data with POST request on ```http://localhost:8000/update-flight-details```
with data structured as 
```
{
	"flight": "AA126",
	"status": "On Time",
	"gate": "A3",
	"remarks": "Gate Changed to A3"
}
```



