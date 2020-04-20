package employee

type Employee struct {
	ID            int    `json:"id"`
	Lastname      string `json:"lastname"`
	FirstName     string `json:"firstname"`
	Company       string `json:"company"`
	EmailAddress  string `json:"emailAddress"`
	JobTitle      string `json:"jobTitle"`
	BusinessPhone string `json:"businessPhone"`
	HomePhone     string `json:"homePhone"`
	MobilePhone   string `json:"mobilePhone"`
	FaxNumber     string `json:"faxNumber"`
	Address       string `json:"address"`
}

type EmployeeList struct {
	Data         []*Employee `json:"data"`
	TotalRecords int64       `json:"totalRecords"`
}

type BestEmployee struct {
	ID          int    `json:"id"`
	TotalVentas int    `json:"TotalVentas"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
}
