#### Invest project

##### Needed
```text
* docker (used 18.09.7)
* docker-compose (used 1.8.0)
* **optional** goose (to run migrations)
```

##### To run
**Note:** before running backend refer to ./env/.env file & make sure all variables are correct 
```text
To run PostgreSQL:
* env HOST=0.0.0.0 docker-compose -f docker-services.yml up --bu -d
To run backend:
* docker-compose up --bu -d
```

##### Makefile
Inside the 'Makefile', you can find short-cut commands.  
**Note:** 0.0.0.0 makes your database open for outside

```makefile
run_services:
	env HOST=0.0.0.0 docker-compose -f docker-services.yml up --bu -d

run_backend:
	docker-compose up --bu -d

test_database_run:
	docker pull postgres:11-alpine && docker run --name database_test -p 0.0.0.0:7010:5432 -e POSTGRES_USER=spkuser -e POSTGRES_PASSWORD=spkpassword -e POSTGRES_DB=spkdb -d postgres:11-alpine

test_database_container_remove:
	docker rm database_test

test_database_remove:
	docker kill database_test && docker rm database_test

.PHONY: run_services run_backend test_database_run test_database_remove
```

##### About the roles 

Roles are hard-coded within the project.

```text
>> roles <<
admin
investor
manager
expert
```

##### Packages
* **/app** - API path & mapping to functions  
* **/auth** - all pre-checks & pre-operations before letting the user to access resources (e.g. authorization token is handled in this package)  
* **/control** - accept request, parse request data, security check (e.g. has user a right to access this?) & call service function  
* **/db** - .sql files, which are needed to run migration  
* **/documents** - store documents (all uploaded files) & files left after analysis  
* **/env** - Dockerfile for PostgreSQL & environmental variables  
* **/logdir** - all log files (by date)  
* **/model** - all structs (for database scheme) & data access objects  
* **/service** - logic is here (e.g. running several data access objects)  
* **/test** - you can find all test within this package  
* **/utils** - helper functions (e.g. generate random string), constants (e.g. roles) & error messages  

##### Statuses & their meaning
```text
200 - ok. everything is fine & as expected
201 - ok. created
204 - ok. but could not send email message
400 - bad request. your data did not go through validation 
404 - info not found on database
405 - method is not allowed. you are trying something that you cannot access
406 - email is not confirmed thus an account is not activated
409 - already in use
417 - internal database error
422 - could not sent message & not stored on db
500 - internal server error
503 - service is unavailable (e.g. autocomplete for bin)
```

##### Wrapper Statuses
```text
This is a list of statuses that might be returned to any request that go though
wrapper that check session token, language & permission and whether email address 
is confir03med

421 - invalid parameters (e.g. session token is invalid) 
423 - email address is not confirmed 
424 - has not such a permission 
500 - other internal error or path is not correct
```

