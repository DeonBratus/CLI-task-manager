package tskMng

type status string

const (
	ToDo       status = "todo"
	InProgress status = "in-progress"
	Done       status = "done"
)

type Task struct {
	ID          int    `json:"ID"`
	TaskName    string `json:"Name"`
	Description string `json:"Desc"`
	Status      status `json:"Status"`
	CreatedBy   User   `json:"CreatedAt"`
	UpdatedBy   User   `json:"UpdatedAt"`
}

func (tsk *Task) UpdateDesc(desc string, user User) {
	tsk.Description = desc
	tsk.UpdatedBy = user
}

func (tsk *Task) UpdateStatus(st status) status {
	tsk.Status = st
	return tsk.Status
}

func (tsk *Task) UpdateName(newName string) {
	tsk.TaskName = newName
}
