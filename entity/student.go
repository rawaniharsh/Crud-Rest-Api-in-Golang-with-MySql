package entity

//Student object for database and rest operations
type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	RollNo string `json:"rollNo"`
	Class  string `json:"class"`
}
