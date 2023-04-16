package database

import (
	"fmt"
	"strings"

	"github.com/Rosya-edwica/it-parser/pkg/common/models"
)

// FIXME: Dynamic TableName and DBUrl

const TableName = "vacancy"

func AddVacancy(v models.Vacancy) {
	db := Connect()
	defer db.Close()

	query := fmt.Sprintf(`INSERT INTO %s(id, id_platform, name, description, url, skills, salary_to, salary_from, currency, city, address, experience, schedule, employment)
	VALUES('%s', %d, '%s', '%s', '%s', '%s', %d, %d, '%s', '%s', '%s', '%s', '%s', '%s')`,
		TableName, v.Id, v.IdPlatform, v.Name, v.Description, v.Url, arrayToPostgres(v.Skills), v.SalaryTo, v.SalaryFrom, v.Currency, v.City, v.Address, v.Experience, v.Schedule, v.Employment)

	tx, _ := db.Begin()
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tx.Commit()
	checkErr(err)
	fmt.Println("Вакансия сохранена! ", v.Id)
	return
}

func arrayToPostgres(arrray []string) string {
	if len(arrray) > 0 {
		return "{" + strings.Join(arrray, ",") + "}"
	}
	return "{}"
}
