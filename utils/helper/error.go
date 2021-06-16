package helper

func IfErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
