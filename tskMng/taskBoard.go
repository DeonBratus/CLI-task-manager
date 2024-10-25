package tskMng

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TaskBoard struct {
	TaskList []Task `json:"TaskList"`
}

func searchTaskWithName(name string, tb *TaskBoard) *Task {
	for _, task := range tb.TaskList {
		if task.TaskName == name {
			return &task
		}
	}
	return nil
}

func searchTasksWithStatus(status string, tb *TaskBoard) *[]Task {
	var tasks []Task
	for _, task := range tb.TaskList {
		if string(task.Status) == status {
			tasks = append(tasks, task)
		}
	}

	return &tasks
}

func (tb *TaskBoard) CreateTask(name string, user User) {
	for _, task := range tb.TaskList {
		if task.TaskName == name {
			fmt.Print("You can't create task with same names. Rename your task, please!")
		}
	}
	tb.TaskList = append(tb.TaskList, Task{ID: len(tb.TaskList) + 1, TaskName: name, CreatedBy: user})
	tb.TaskList[len(tb.TaskList)-1].Status = ToDo

}

func (tb *TaskBoard) UpdateTask(name string, user User) {
	changingType := ""
	changingStr := ""

	for idx := range tb.TaskList {
		if tb.TaskList[idx].TaskName == name {
			fmt.Println("What do you want to edit? name, status, desc")
			fmt.Scan(&changingType)
			fmt.Print("You edit ", changingType, "\n>")

			switch changingType {
			case "name":
				fmt.Scan(&changingStr)
				tb.TaskList[idx].UpdateName(changingStr)
			case "desc":
				fmt.Scan(&changingStr)
				tb.TaskList[idx].UpdateDesc(changingStr, user)
			case "status":
				fmt.Scan(&changingStr)
				tb.TaskList[idx].UpdateStatus(status(changingStr))
			default:
				fmt.Println("Incorrect type of changes. Try again with correct type.")
			}
			return
		}
	}
	fmt.Print("Task doesn't exist. Add task for edit, please.\n")

}

func (tb *TaskBoard) DeleteTask(name string) {
	for idx, task := range tb.TaskList {

		if task.TaskName == name {

			tb.TaskList = append(tb.TaskList[:idx], tb.TaskList[idx+1:]...)

			for i := idx; i <= len(tb.TaskList)-1; i++ {
				tb.TaskList[i].ID--
			}
			return
		}
	}
	fmt.Print("Task doesn't exist. Add task before deleting this, please\n")

}

func (tb *TaskBoard) ShowTaskInfo(name string) {
	task := searchTaskWithName(name, tb)
	fmt.Printf("ID%d | %s | %s |\n\"%s\"\n", task.ID, task.TaskName, task.Status, task.Description)

}

func (tb *TaskBoard) ShowListDone() {
	tasks := searchTasksWithStatus("done", tb)
	for _, task := range *tasks {
		fmt.Printf("ID%d | %s | %s |\n", task.ID, task.TaskName, task.Status)
	}
}

func (tb *TaskBoard) ShowListTodo() {
	tasks := searchTasksWithStatus("todo", tb)
	for _, task := range *tasks {
		fmt.Printf("ID%d | %s | %s |\n", task.ID, task.TaskName, task.Status)
	}
}

func (tb *TaskBoard) ShowListInProgress() {
	tasks := searchTasksWithStatus("in-progress", tb)
	for _, task := range *tasks {
		fmt.Printf("ID%d | %s | %s |\n", task.ID, task.TaskName, task.Status)
	}
}

func (tb *TaskBoard) ShowALlTasksList() {
	for _, task := range tb.TaskList {
		fmt.Printf("ID%d | %s | %s |\n", task.ID, task.TaskName, task.Status)
	}
}

func (tb *TaskBoard) HandlerCLI(taskcmdstr string, user User) {
	splitstr := strings.Split(taskcmdstr, " ")
	taskName := ""
	cmd := splitstr[0]
	if len(splitstr) == 0 {
		fmt.Println("Invalid command. Please provide both command and task name.")
		return
	} else if len(splitstr) == 1 {

	} else {
		taskName = splitstr[1]

	}

	switch strings.ToLower(cmd) {
	case "add":
		tb.CreateTask(taskName, user)
	case "del":
		tb.DeleteTask(taskName)
	case "edit":
		tb.UpdateTask(taskName, user)
	case "show":
		tb.ShowTaskInfo(taskName)
	case "ldn":
		tb.ShowListDone()
	case "ltd":
		tb.ShowListTodo()
	case "lip":
		tb.ShowListInProgress()
	case "lall":
		tb.ShowALlTasksList()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func (tb *TaskBoard) TaskManageCLI(user User) {
	var taskcmdstr string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		taskcmdstr = scanner.Text()

		if strings.ToLower(taskcmdstr) == "quit" {
			break
		}
		if taskcmdstr != "" {
			tb.HandlerCLI(taskcmdstr, user)
		}
	}
}
