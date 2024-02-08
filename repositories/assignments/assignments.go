package assignments

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"project.com/restful-api/repositories/models"
	"project.com/restful-api/utilities"
)

type Repository interface {
	Get(id string) []models.Assignment
	GetAll() []models.Assignment
	Add(assign models.Assignment, uid string)
	Del(id models.Assignment, uid string)
	Put(assign models.Assignment)
}

type Prodconnection struct {
	Db *gorm.DB
}

func (con Prodconnection) Get(id string) []models.Assignment {
	statserr := utilities.StatsdClient.Inc("repo.assignments.Get", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	var assigns []models.Assignment
	//con.Db.Raw("SELECT * FROM assignments WHERE ID = ?", id).Scan(&assigns)
	con.Db.First(&assigns, "id = ?", id)
	return assigns
}

func (con Prodconnection) GetAll() []models.Assignment {
	statserr := utilities.StatsdClient.Inc("repo.assignments.GetAll", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	var assigns []models.Assignment
	//con.Db.Raw("SELECT * FROM assignments").Scan(&assigns)
	con.Db.Find(&assigns)
	return assigns
}

func (con Prodconnection) Add(assign models.Assignment, uid string) {
	statserr := utilities.StatsdClient.Inc("repo.assignments.Add", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	con.Db.Create(&assign)
	con.Db.Create(&models.Account_Assignment{Account_Id: uid, Assignment_Id: assign.ID})
	//con.Db.Exec("INSERT INTO account_assignments(account_id, assignment_id) VALUES (?, ?)", uid, assign.ID)
}

func (con Prodconnection) Del(assign models.Assignment, uid string) {
	statserr := utilities.StatsdClient.Inc("repo.assignments.Delete", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	tmpAA := models.Account_Assignment{Account_Id: uid, Assignment_Id: assign.ID}
	con.Db.Where("account_id = ? AND assignment_id = ?", tmpAA.Account_Id, tmpAA.Assignment_Id).Delete(&tmpAA)
	con.Db.Delete(&assign)
}

func (con Prodconnection) Put(assign models.Assignment) {
	statserr := utilities.StatsdClient.Inc("repo.assignments.Put", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelDebug, "Statd metric logging failed: "+statserr.Error())
	}

	//err := con.Db.Exec("UPDATE assignments SET name = ?, points = ?, num_of_attemps = ?, assignment_updated = ? WHERE ID = ?", assign.Name, assign.Points, assign.Num_Of_Attemps, assign.Assignment_Updated, assign.ID)
	err := con.Db.Model(&models.Assignment{}).Where("id = ?", assign.ID).Updates(&models.Assignment{Name: assign.Name, Points: assign.Points, Num_Of_Attemps: assign.Num_Of_Attemps, Assignment_Updated: assign.Assignment_Updated, Deadline: assign.Deadline})
	fmt.Println(err)
}
