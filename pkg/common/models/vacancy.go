package models

type Vacancy struct {
	Id          string
	IdPlatform  int
	Name        string
	Description string
	Url         string
	Skills      []string
	SalaryFrom  int64
	SalaryTo    int64
	Currency    string
	Address     string
	City        string
	Experience  string
	Schedule    string
	Employment  string
}
