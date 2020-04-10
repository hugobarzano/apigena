package model

type STUDENTS struct {
	Address STUDENTS_sub1 `json:"address" yml:"address"`
	Age     int64         `json:"age" yml:"age"`
	Car     bool          `json:"car" yml:"car"`
	Hobbies []string      `json:"hobbies" yml:"hobbies"`
	Name    string        `json:"name" yml:"name"`
	ID string      `json:"id" yml:"id"`
}

type STUDENTS_sub1 struct {
	City   string `json:"city" yml:"city"`
	Number int64  `json:"number" yml:"number"`
	Street string `json:"street" yml:"street"`
}
