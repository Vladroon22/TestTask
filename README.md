# Test Task REST API

<h2> API Endpoints </h2>

```
POST /users create new user
```
Body request
```
{
    "name" : ..., (string)
    "surname: ..., (string)
    "age": ..., (int)
    "email: ..., (string)
    "phone":..., (string)
}
```

```
GET /users/{id:[0-9]+} - Get the user data
```

```
PUT /users/{id:[0-9]+} - Update the user data
```

<p>Body request</p>
```
{
    "name" : ..., (string)
    "surname: ..., (string)
    "age": ..., (int)
    "email: ..., (string)
    "phone":..., (string)
}
```


<h2> Configuration </h2>

<h4> Set enviroment variable </h4>

```
addr = ":your free port"
DB = "postgres://postgres:55555@localhost:5432/postgres?sslmode=disable"
```

<h4>How to run</h4>

```
make run
```

```
make compose
```
