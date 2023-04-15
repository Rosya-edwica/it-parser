package models

type Vacancy struct {
	Id          string
	IdPlatform  int
	Name        string
	Description string
	Url         string
	Skills      []string
	SalaryTo    int
	SalaryFrom  int
	Currency    string
	Address     string
	City        string
	Experience  string
	Schedule    string
	Employment  string
}
