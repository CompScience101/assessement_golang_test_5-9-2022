# Assessment Test REST API

> My REST API to implement retailer's rewards program.

## Quick Start for running golang REST API without go docker container

#1) Copy "assessment_test_5_9_2022.exe" to local hard drive.
#2) Open command line prompt and run "assessment_test_5_9_2022.exe" from command line.
    In windows run command "assessment_test_5_9_2022.exe". 
    In Linux run command "./assessment_test_5_9_2022.exe"

## Quick Start for running golang REST API with go in docker container
#1) copy my code(all files including dockerfile) from my repository to your local machine that you have docker installed on.
#2) run command from CLI:   docker build assessment_test_5_9_2022 ./
#3) run command from CLI:   docker run -it -p 8080:8080 --name testapi assessment_test_5_9_2022
#4) then use any client to test my api on port 8080, with the given endpoints below.

## Unit Testing
#1) copy my code(all files including dockerfile )
#2) 



## Using Endpoints
- When using the "Create transaction" endpoint, please make sure to provide a decimal value for the cost field of your json,
  and only use up to two decimal places, because any more will cause an error.
- When using the "Get Health" endpoint you don't need to pass any data, just use the url and path specified and if it succeeds
  you will get a 200 status response along with content body of the running status, otherwise you will get no response from 
  server if it is not available.

## Endpoints

### Get Health
``` bash
# URL: api/health
# Method: GET  
# HTTP/1.1
# Response sample
# code: 200
# content: {running: "true"}
```

### Create Transaction
``` bash
# URL: api/transaction 
# Method:POST 
# HTTP/1.1

# Request sample
# {
#   "cost": 120.00,
# }

# Response sample
# Code: 201
# Content: {
#   reward: 90
# }
```


```


### Writer

Nicholas Donald