// Ларионова Арина 363
package main

import (
	"fmt"
)

type Task struct {
	ID          string
	Description string
	Completed   bool
}

type TaskManager struct {
	task   []Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		task:   make([]Task, 0),
		nextID: 1,
	}
}

func (tm *TaskManager) AddTask(description string) {
	tasc := Task{
		ID:          fmt.Sprintf("%d", tm.nextID),
		Description: description,
		Completed:   false,
	}
	tm.task = append(tm.task, tasc)
	tm.nextID++
}

func (tm *TaskManager) RemoveTask(id string) error {
	for i, task := range tm.task {
		if task.ID == id {
			tm.task = append(tm.task[:i], tm.task[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("задача с ID %s не найдена", id)
}

func (tm *TaskManager) CompleteTask(id string) error {
	for i, task := range tm.task {
		if task.ID == id {
			tm.task[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("задача с ID %s не найдена", id)
}

func (tm *TaskManager) FilterCompleted() []Task {
	completedTasks := []Task{}
	for _, task := range tm.task {
		if task.Completed {
			completedTasks = append(completedTasks, task)
		}
	}
	return completedTasks
}

func (tm *TaskManager) FilterPending() []Task {
	pendingTasks := []Task{}
	for _, task := range tm.task {
		if !task.Completed {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

func (tm *TaskManager) PrintTasks() {
	for _, task := range tm.task {
		status := "не выполнена"
		if task.Completed {
			status = "выполнена"
		}
		fmt.Printf("ID: %s, Описание: %s, Статус: %s\n", task.ID, task.Description, status)
	}
}

func main() {
	tm := NewTaskManager()
	tm.AddTask("бла бла")
	tm.AddTask("Купить продукты")
	fmt.Println("Все задачи:")
	tm.PrintTasks()
	tm.CompleteTask("1")
	fmt.Println("\nЗадачи после выполнения одной из них:")
	tm.PrintTasks()
	err := tm.RemoveTask("2")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("\nЗадачи после удаления одной из них:")
	tm.PrintTasks()
}
