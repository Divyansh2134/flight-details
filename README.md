# flight-details

# Frontend

- building frontend
```npm ci```

- starting server locally
```npm run dev```

- for building ```npm run build```

- access using url ```http://localhost:5173/```

# Backend 

- build ```go mod download```

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




