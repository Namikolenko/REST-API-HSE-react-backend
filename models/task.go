package models

import (
	"api/utils"
	"fmt"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Id          uint   `json:"id"`
	UserId      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (task *Task) Validate() (map[string]interface{}, bool) {

	if task.Name == "" {
		return utils.Message(false, "Contact name should be on the payload"), false
	}

	if task.UserId <= 0 {
		return utils.Message(false, "User is not recognized"), false
	}

	return utils.Message(true, "success"), true
}

func (task *Task) Create() map[string]interface{} {

	if resp, err := task.Validate();
		!err {
		return resp
	}

	GetDB().Create(task)

	resp := utils.Message(true, "success")
	resp["task"] = task
	return resp
}

func GetTasks(user uint) []*Task {

	tasks := make([]*Task, 0)
	err := GetDB().Table("tasks").Where("user_id = ?", user).Find(&tasks).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tasks
}

func ChangeStateById(id, user uint) map[string]interface{} {

	task := &Task{}

	err := GetDB().Table("tasks").Where("user_id = ? AND id = ?", user, id).First(task).Error
	if err != nil {
		return utils.Message(false, "There is no existing Task with this props!")
	}
	completed := task.Completed
	task.Completed = !completed
	GetDB().Save(task)

	return utils.Message(true, "Success!")
}

func DeleteTask(id, user uint) map[string]interface{} {

	task := &Task{}

	err := GetDB().Table("tasks").Where("user_id = ? AND id = ?", user, id).First(task).Error
	if err != nil {
		return utils.Message(false, "There is no existing Task with this props!")
	}
	GetDB().Delete(task)

	return utils.Message(true, "Success!")
}
