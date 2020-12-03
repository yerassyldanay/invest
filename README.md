#### Invest project

##### About the roles 

Roles are hard-coded within the project.

```text
>> Roles: <<
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


