package headhunter

import (
	"fmt"
	"time"

	"github.com/Rosya-edwica/it-parser/pkg/common/database"
	"github.com/Rosya-edwica/it-parser/pkg/common/models"
	"github.com/tidwall/gjson"
)

func ScrapePage() {
	start := time.Now().Unix()

	main_url := "https://api.hh.ru/vacancies?industry=7.538&indystry=539&indystry=7.540&indystry=7.541"
	json, err := getJson(main_url)
	checkErr(err)

	vacancies := gjson.Get(json, "items").Array()
	for _, vacancy := range vacancies {
		vacancy := ScrapeVacancy(vacancy.Get("id").Int())
		database.AddVacancy(vacancy)
	}

	fmt.Printf("Time: %d\n", time.Now().Unix()-start)
}

// FIXME: idPlatform dynamic

func ScrapeVacancy(id int64) (vacancy models.Vacancy) {
	json, err := getJson(fmt.Sprintf("https://api.hh.ru/vacancies/%d", id))
	checkErr(err)

	vacancy.IdPlatform = 1
	vacancy.Description = "Description"
	vacancy.Name = gjson.Get(json, "name").String()
	vacancy.Id = gjson.Get(json, "id").String()
	vacancy.City = gjson.Get(json, "area.name").String()
	vacancy.Employment = gjson.Get(json, "employment.name").String()
	vacancy.Schedule = gjson.Get(json, "schedule.name").String()
	vacancy.Experience = renameExperience(gjson.Get(json, "experience.name").String())
	vacancy.SalaryFrom = gjson.Get(json, "salary.from").Int()
	vacancy.SalaryTo = gjson.Get(json, "salary.to").Int()
	vacancy.Currency = gjson.Get(json, "salary.currency").String()
	vacancy.Skills = getSkills(json)
	vacancy.Url = gjson.Get(json, "alternate_url").String()
	return
}
