package model

type ResponseStudents struct {
	Students []Student `json:"students"`
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}
