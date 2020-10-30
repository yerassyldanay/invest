package model

// this will be used to count the number of rows
// after running raw sql
type Counter struct {
	Number				int					`json:"number"`
	Count				int					`json:"count"`
}
