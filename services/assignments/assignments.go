package assignments

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"project.com/restful-api/repositories"
	"project.com/restful-api/repositories/models"
	"project.com/restful-api/services/accounts"
	"project.com/restful-api/services/models/request"
	"project.com/restful-api/utilities"
)

func Patch(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.Patch")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.Patch", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	c.AbortWithStatus(http.StatusMethodNotAllowed)
	slog.LogAttrs(context.Background(), slog.LevelWarn, "MethodNotAllowed Assignments.Patch")
	return
}

func Del(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.Delete")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.Delete", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	params := c.Request.URL.Query()
	bbody, rerr := io.ReadAll(c.Request.Body)
	if len(params) > 0 || rerr != nil || len(bbody) != 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uid, valid := accounts.IsBasicAuthorized(c.GetHeader("Authorization"))
	if valid == false {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Unauthorized Assignments.Delete")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	/** Todo: Handle required fields for all */
	/** Todo: Check update time stamp */
	/** Todo: Handle min and max requirements */
	/** Todo: HTTP methods not supported part */
	repo := repositories.GetAssignmentsRepo()
	assign := repo.Get(c.Param("id"))
	if len(assign) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "NotFound Assignments.Delete")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	repoAA := repositories.GetAccountAssignmentsRepo()
	aa := repoAA.GetByAssignment(assign[0].ID)
	if uid != aa.Account_Id {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Forbidden Assignments.Delete")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	repo.Del(assign[0], uid)
	c.Status(http.StatusNoContent)

}

func Put(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.Put")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.Put", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Accept", "application/json")
	ct := c.Request.Header.Get("Content-Type")
	params := c.Request.URL.Query()
	bbody, rerr := io.ReadAll(c.Request.Body)
	if len(params) > 0 || ct != "application/json" || rerr != nil || len(bbody) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var rassign request.Assignment
	jer := json.Unmarshal(bbody, &rassign)
	if jer != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	assign := models.Assignment{
		ID:                 c.Param("id"),
		Name:               rassign.Name,
		Points:             rassign.Points,
		Num_Of_Attemps:     rassign.Num_Of_Attemps,
		Deadline:           rassign.Deadline,
		Assignment_Updated: time.Now(),
	}

	if assign.Points < 1 ||
		assign.Points > 100 ||
		assign.Num_Of_Attemps < 1 ||
		assign.Num_Of_Attemps > 100 ||
		strings.TrimSpace(assign.Name) == "" ||
		assign.Deadline.IsZero() == true {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uid, valid := accounts.IsBasicAuthorized(c.GetHeader("Authorization"))
	if valid == false {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Unauthorized Assignments.Put")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	repoAA := repositories.GetAccountAssignmentsRepo()
	aa := repoAA.GetByAssignment(assign.ID)
	if uid != aa.Account_Id {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Forbidden Assignments.Put")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	repo := repositories.GetAssignmentsRepo()
	assgn := repo.Get(assign.ID)
	if len(assgn) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "NotFound Assignments.Put")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	repo.Put(assign)
	c.Status(http.StatusNoContent)
	return
}

func Post(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.Post")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.Post", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Accept", "application/json")
	ct := c.Request.Header.Get("Content-Type")
	params := c.Request.URL.Query()
	bbody, rerr := io.ReadAll(c.Request.Body)
	if len(params) > 0 || ct != "application/json" || rerr != nil || len(bbody) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var rassign request.Assignment
	jer := json.Unmarshal(bbody, &rassign)
	if jer != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uuid, _ := uuid2.NewUUID()
	assign := models.Assignment{
		ID:             uuid.String(),
		Name:           rassign.Name,
		Points:         rassign.Points,
		Num_Of_Attemps: rassign.Num_Of_Attemps,
		Deadline:       rassign.Deadline,
	}

	if assign.Points < 1 ||
		assign.Points > 100 ||
		assign.Num_Of_Attemps < 1 ||
		assign.Num_Of_Attemps > 100 ||
		strings.TrimSpace(assign.Name) == "" ||
		assign.Deadline.IsZero() == true {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uid, valid := accounts.IsBasicAuthorized(c.GetHeader("Authorization"))
	if valid == false {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Unauthorized Assignments.Post")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	repo := repositories.GetAssignmentsRepo()
	repo.Add(assign, uid)

	res := repo.Get(assign.ID)

	c.JSON(http.StatusCreated, res[0])
}

func GetAll(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.GetAll")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.GetAll", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Accept", "application/json")
	params := c.Request.URL.Query()
	bbody, err := io.ReadAll(c.Request.Body)
	if len(params) > 0 || len(bbody) > 0 || err != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.GetAll")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, valid := accounts.IsBasicAuthorized(c.GetHeader("Authorization"))
	if valid == false {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Unauthorized Assignments.GetAll")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	repo := repositories.GetAssignmentsRepo()
	assigns := repo.GetAll()
	if len(assigns) == 0 {
		c.JSON(http.StatusOK, []models.Assignment{})
		return
	}

	c.JSON(http.StatusOK, assigns)
}

func Get(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.Get")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.Get", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Accept", "application/json")
	params := c.Request.URL.Query()
	bbody, err := io.ReadAll(c.Request.Body)
	if len(params) > 0 || len(bbody) > 0 || err != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.Get")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, valid := accounts.IsBasicAuthorized(c.GetHeader("Authorization"))
	if valid == false {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Unauthorized Assignments.Get")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	repo := repositories.GetAssignmentsRepo()
	assigns := repo.Get(c.Param("id"))
	if len(assigns) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "NotFound Assignments.Get")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, assigns[0])
}

func PostSubmission(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.PostSubmission")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.PostSubmission", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Accept", "application/json")
	ct := c.Request.Header.Get("Content-Type")
	params := c.Request.URL.Query()
	bbody, rerr := io.ReadAll(c.Request.Body)
	if len(params) > 0 || ct != "application/json" || rerr != nil || len(bbody) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid request metadata")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var submissionDetails struct {
		SubmissionUrl string `json:"submission_url"`
	}
	jer := json.Unmarshal(bbody, &submissionDetails)
	if jer != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Failed to unmarshall data")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(submissionDetails.SubmissionUrl) == "" {
		slog.LogAttrs(context.Background(), slog.LevelError, "Invalid submission URL")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	uid, valid := accounts.IsBasicAuthorized(c.GetHeader("Authorization"))
	if valid == false {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Unauthorized Assignments.Post")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	repo := repositories.GetAssignmentsRepo()
	assigns := repo.Get(c.Param("id"))
	if len(assigns) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "NotFound Assignments.Get")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	subRepo := repositories.GetSubmissionsRepo()
	subRec := subRepo.GetByFilter(uid, assigns[0].ID)
	if len(subRec) >= assigns[0].Num_Of_Attemps || time.Now().UTC().After(assigns[0].Deadline) {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "SubmissionLimitation Account_Assignment.GetByFiler")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	uuid, _ := uuid2.NewUUID()
	submission := models.Submission{
		ID:             uuid.String(),
		AccountId:      uid,
		AssignmentId:   assigns[0].ID,
		SubmissionUrl:  submissionDetails.SubmissionUrl,
		SubmissionDate: time.Time{},
	}
	subRepo.Add(submission)

	accountEmail := accounts.Get(uid)

	sub, err := json.Marshal(struct {
		SubmissionEmail string `json:"SubmissionEmail"`
		SubmissionUrl   string `json:"SubmissionUrl"`
		SubmissionId    string `json:"SubmissionId"`
		AssignmentId    string `json:"AssignmentId"`
		UserId          string `json:"UserId"`
	}{
		accountEmail,
		submission.SubmissionUrl,
		submission.ID,
		submission.AssignmentId,
		uid,
	})

	utilities.SendMessage(string(sub))

	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Failed to parse struct at Assignments.PostSubmission"+err.Error())
	}
}

func GetSubmissions(c *gin.Context) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Assignments.GetSubmission")
	statserr := utilities.StatsdClient.Inc("endpoint.assignments.Get", 1, 1.0)
	if statserr != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Statd metric logging failed: "+statserr.Error())
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Accept", "application/json")
	params := c.Request.URL.Query()
	bbody, err := io.ReadAll(c.Request.Body)
	if len(params) > 0 || len(bbody) > 0 || err != nil {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Invalid Assignments.GetSubmission")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, valid := accounts.IsBasicAuthorized(c.GetHeader("Authorization"))
	if valid == false {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "Unauthorized Assignments.GetSubmission")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	repo := repositories.GetSubmissionsRepo()
	assigns := repo.GetByAssignmentId(c.Param("id"))
	if len(assigns) == 0 {
		slog.LogAttrs(context.Background(), slog.LevelWarn, "NotFound Assignments.GetSubmission")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, assigns[0])
}
