package services

import (
	"go-get-stuff-done/models"
)


type NextTask struct {
	rests models.Result
	mins_completed models.Result
}

func ( t NextTask) GetNext() {
	if t.mins_completed.Val > t.rests.Val*120 {
	// create a rest task
	}
	//get oldest task by priority
}
