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
400 - bad request. your data did not go through validation 
404 - info not found on database
405 - method is not allowed. you are trying something that you cannot access
406 - email is not confirmed thus an account is not activated
409 - already in use
417 - internal database error
422 - could not sent message & not stored on db
503 - service is unavailable (e.g. autocomplete for bin)
```
