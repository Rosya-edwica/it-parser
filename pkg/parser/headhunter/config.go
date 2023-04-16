package headhunter

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

const TOKEN = "QQAVSIBVU4B0JCR296THKB22JP05A92H329U49TDD9CRIS8DT9BRPPT7M9OLQ6HD"

var headers = map[string]string{
	"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.",
	"Authorization": fmt.Sprintf("Bearer %s", TOKEN),
}

func getJson(url string) (json string, err error) {
	// Создаем клиента, который будет подключаться к странице
	client := http.Client{ // Ставим время ожидания на 120 секунд. Если ответа не будет больше 2 минут, будет ошибка
		Timeout: 120 * time.Second,
	}
	// Создаем GET-запрос к url без дополнительных параметров. body = nil
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	// Добавляем заголовок к запросу
	req.Header.Set("User-Agent", headers["User-Agent"])
	req.Header.Set("Authorization", headers["Authorization"])
	response, err := client.Do(req) // Пытаемся подключиться
	if err != nil {
		return
	}
	defer response.Body.Close() // Обязательно закрываем содержимое ответа
	// Сохраняем содержимое в виде списка байтов
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	// приводим список байтов к строке и выводим ее вместе с отсутствием ошибок
	return string(data), nil
}

// Удобная функция для проверки ошибок
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func renameExperience(experience string) string {
	switch experience {
	case "Нет опыта":
		return "Intern"
	case "От 1 года до 3 лет":
		return "Junior"
	case "От 3 до 6 лет":
		return "Middle"
	case "Более 6 лет":
		return "Senior"
	default:
		return "Lead"
	}
}

func getSkills(json string) (skills []string) {
	for _, skill := range gjson.Get(json, "key_skills").Array() {
		skills = append(skills, skill.Get("name").String())
	}
	return
}
