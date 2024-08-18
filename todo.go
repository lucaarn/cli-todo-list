package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

type todo struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

type todoList []todo

func (t *todoList) SaveTasks() {
	file, err := os.Create("todos.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
		return
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
			return
		}
	}
}
