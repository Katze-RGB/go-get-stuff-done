package models

import (
	"errors"
	"slices"
	"time"

	"gorm.io/gorm"
)

type TodoTask struct {
	gorm.Model
	ID               uint   `json:"id" gorm:"primaryKey"`
	Description      string `json:"description" gorm:"text;not null;default:null"` // pretty self explanatory
	Priority         int    `json:"priority" gorm:"text;not null;default:1"`       // 1-3 where 3 is the most and 1 is the least. 4 exists but is secret ;)
	Estimated_length int    `json:"estimated_length" gorm:"default:5;not_null"`    // estimated time of tasks in mins
	Completed        bool   `json:"completed" gorm:"default:false"`                // is it done?
}

func (t TodoTask) IsValid() error {
	valid_priorities := []int{1, 2, 3, 4}
	if !slices.Contains(valid_priorities, t.Priority) {
		return errors.New("Priority must be 1,2,3. Please try again")
	}

	if t.Estimated_length <= 0 {
		return errors.New("tasks don't take negative time stop that")
	}

	return nil
}

type ProductivityReport struct {
	Date           time.Time `json:"date"`
	TasksCompleted int       `json:"tasks_completed"`
	MinsSpent      int       `json:"mins_spent"`
}

type Result struct {
	Val int
}

type Task struct {
	ID              int    `json:"id"`
	Priority        int    `json:"priority"`
	EstimatedLength int    `json:"estimated_length"`
	Description     string `json:"description"`
	EngPriority     string `json:"eng_priority"`
}

func (t Task) FriendlyPriority() string {
	switch t.Priority {
	case 1:
		return "low"
	case 2:
		return "medium"
	case 3:
		return "high"
	case 4:
		return "naptime RIGHT MEOW"
	}
	return "unknown"
}
