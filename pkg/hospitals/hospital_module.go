package hospitals

// Hospital represents a hospital
type Hospital struct {
	ID          int32        `json:"id"`
	Name        string       `json:"name"`
	City        string       `json:"city"`
	Departments []Department `json:"departments"`
}

// Department represents the department of a hospital
type Department struct {
	ID       int32    `json:"id"`
	Name     string   `json:"name"`
	Hospital Hospital `json:"-"`
	Doctors  []Doctor `json:"doctors"`
}
