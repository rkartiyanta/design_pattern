package structural

var CacheEmployee []Employee

type Employee struct {
	Name string
	Job  string
}

type EmployeeList struct {
	Employees []Employee
}

// employee list only generated once and cached
// the other classes are using the shared employee list
func (c *EmployeeList) GetEmployeeList() []Employee {
	employees := CacheEmployee
	if len(employees) > 0 {
		return employees
	}

	return c.GenerateList()
}

func (c *EmployeeList) GenerateList() []Employee {
	employees := []Employee{
		Employee{
			Name: "Richard",
			Job:  "Backend",
		},
		Employee{
			Name: "Kara",
			Job:  "BA",
		},
		Employee{
			Name: "Sukma",
			Job:  "QA",
		},
	}

	CacheEmployee = employees
	return employees
}
