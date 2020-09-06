#### Invest project

##### About the roles 
Within the scope of the project, every role is considered as a user

```text
>> Roles (users): <<
admin
investor
manager
lawyer
financier
```

##### Database

A meaning of some column names

```json
{
  "fname": "first name",
  "sname": "second name",
  "mname": "middle name - in our case, it is father's name of a person",

  "country_code": "+7",
  "phone": "xxxyyyaabb"
}
```

#####Warning

The warning levels inside the code (It is created because of incomtability with the 'logrus' package)

```go
package nameit

const (
	PanicLevel int = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)
```

##### Permissions
```go

```

#### Statuses & their meaning
```text
200 - ok. everything is fine & as expected
201 - ok. created
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

#### Wrapper Statuses
```text
This is a list of statuses that might be returned to any request that go though
wrapper that check session token, language & permission and whether email address 
is confir03med

421 - invalid parameters (e.g. session token is invalid) 
423 - email address is not confirmed 
424 - has not such a permission 
500 - other internal error or path is not correct
```

###### Supported Languages
```text
kk;q=*.*
en;
en-*
```