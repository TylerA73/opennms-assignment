### Go Version
go 1.18

### Running
I've provided a makefile to simplify the build and start up process of the application

```
make run
```

This will tidy the `go.mod` and pull the dependencies for the project. It will then build the project, and run it for you.

### Endpoints

**API Documentation**

[API Docs can be found here](https://documenter.getpostman.com/view/3676546/2s8YCYoc1A)

`POST`
`http://localhost:4000/stats`

**Example Request Body**
```
{
    "machineId": 12345,
    "stats": {
        "cpuTemp": 90,
        "fanSpeed": 400,
        "HDDSpace": 800
    },
    "lastLoggedIn": "admin/Paul",
    "sysTime": "2022-04-23T18:25:43.511Z"
}
```
**Example Response Body**
```
{"message":"Machine stats successfully created"}
```
Status Code: 201

`GET`
`http://localhost:4000/stats`

**Example Response Body**
```
[
    {
        "id": "5a0ff41d-fe3d-4e54-ad0e-e67a8401adb3",
        "machineId": 12345,
        "stats": {
            "cpuTemp": 90,
            "fanSpeed": 400,
            "HDDSpace": 800
        },
        "lastLoggedIn": "admin/Paul",
        "sysTime": "2022-04-23T18:25:43.511Z"
    }
]
```
```
{"message":"Could not create the machine stats: machineID, stats, lastLoggedIn, and sysTime cannot be empty"}
```
Status Code: 200, 400


### Tests
```
make test
```

This command will run the tests.