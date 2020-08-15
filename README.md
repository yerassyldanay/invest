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