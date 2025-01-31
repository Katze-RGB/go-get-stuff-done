package handlers

import (
	"fmt"
	"go-get-stuff-done/db"
	"go-get-stuff-done/models"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("meow meow meow brrt meow pew pew")
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.TodoTask)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := task.IsValid(); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	db.DB.Db.Create(&task)
	return c.Status(200).JSON(task)

}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.TodoTask
	result := db.DB.Db.Delete(&task, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "task not found. This either never existed or already was deleted."})
	}

	return c.Status(200).JSON(fiber.Map{"message": "task deleted"})
}

func CompleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.TodoTask
	result := db.DB.Db.Model(&task).Where("id=? and completed=?", id, false).Update("Completed", true)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "task not found. This task either never existed or has already been completed"})
	}
	return c.Status(200).JSON("completed task")
}

func ListTasks(c *fiber.Ctx) error {
	tasks := []models.TodoTask{}
	db.DB.Db.Find(&tasks)
	return c.Status(200).JSON(tasks)
}

func GetNextTask(c *fiber.Ctx) error {

	var task models.Task

	var rests models.Result
	var mins_completed models.Result
	//sums up time of tasks completet during the current day and returns an int
	db.DB.Db.Raw(`SELECT sum(Estimated_length) as val
				FROM todo_tasks
				WHERE date_trunc('day',updated_at)>=CURRENT_DATE AND Priority <= 3 AND completed = true AND deleted_at is NULL`).Scan(&mins_completed)
	//checks to see how many rest periods we've had so far today
	db.DB.Db.Raw(`SELECT Count(id) as val
				FROM todo_tasks
				WHERE Priority = 4 and date_trunc('day',created_at)>=CURRENT_DATE and deleted_at is NULL`).Scan(&rests)
	//if we've done 2 hours more work since our last rest period, or if we've got time debt, kick a nap at highest priority into the db
	if (mins_completed.Val - rests.Val*120) >= 20 {
		fmt.Println("look at you, you go-getter. Have a little treat.")
		var rest models.TodoTask
		rest.Priority = 4
		rest.Description = "take a little nap"
		rest.Estimated_length = 15
		db.DB.Db.Create(&rest)
	}
	//find the highest priority task by age.
	db.DB.Db.Raw(`SELECT priority,
						estimated_length,
						description,
						id
						FROM todo_tasks
						WHERE completed=false and deleted_at IS null
						ORDER BY priority DESC, created_at DESC LIMIT 1`).Scan(&task)
	task.EngPriority = task.FriendlyPriority()
	return (c.Status(200).JSON(task))

}

func ProductivityReport(c *fiber.Ctx) error {
	date := c.Params("date")
	var report models.ProductivityReport
	db.DB.Db.Raw(`SELECT date_trunc('day',updated_at) as date,
						count(id) as tasks_completed,
						sum(estimated_length) as mins_spent
						FROM todo_tasks
						WHERE completed=true and date_trunc('day', updated_at)=? and deleted_at is null
						GROUP BY date_trunc('day',updated_at)`, date).Scan(&report)
	return (c.Status(200).JSON(report))
}
