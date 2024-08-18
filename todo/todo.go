package todo

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

type TodoList []Todo

func (t *TodoList) Save() {
	file, err := os.Create("todos.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, todo := range *t {
		record := []string{
			strconv.Itoa(todo.ID),
			todo.Description,
			todo.CreatedAt.Format(time.RFC3339),
			strconv.FormatBool(todo.IsComplete),
		}

		err := writer.Write(record)
		if err != nil {
			log.Fatalln("failed to write to file", err)
		}
	}
}

func Load() TodoList {
	file, err := os.Open("todos.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()

	var todos TodoList

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("failed to read file", err)
	}

	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalln("failed to convert id to int", err)
		}

		createdAt, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			log.Fatalln("failed to convert createdAt to time", err)
		}

		isComplete, err := strconv.ParseBool(record[3])
		if err != nil {
			log.Fatalln("failed to convert isComplete to bool", err)
		}

		todos = append(todos, Todo{
			ID:          id,
			Description: record[1],
			CreatedAt:   createdAt,
			IsComplete:  isComplete,
		})
	}

	return todos
}

func GetLatestID(list TodoList) int {
	return list[len(list)-1].ID
}
