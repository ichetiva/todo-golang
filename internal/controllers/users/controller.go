package users

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ichetiva/todo-golang/config"
	"github.com/ichetiva/todo-golang/internal/use_cases/security"
	"github.com/ichetiva/todo-golang/pkg/postgres"
	"github.com/ichetiva/todo-golang/pkg/postgres/models"
	"gorm.io/gorm"
)

type Controller struct {
	Config *config.Config
}

// @summary     Create user
// @description Create user to database with data which contains username and password
// @tags        users
// @accept      json
// @produce		json
// @router      /user/create [post]
// @param 		username body string true "Your future username"
// @param 		password body string true "Your future password"
// @success 	200 {object} UserResponse
// @failure 	400 {object} ErrorResponse "Non-valid data"
// @failure		400 {object} ErrorResponse "Non-valid password"
// @failure 	502 {object} ErrorResponse "Create user failed"
func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestData UserCreation
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}
	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Create user failed",
		})
		_, _ = w.Write([]byte(data))
	}

	user := models.User{
		Username: requestData.Username,
		Password: security.Hash(requestData.Password),
	}
	db.Create(&user)

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(map[string]interface{}{
		"data": User{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
	_, _ = w.Write([]byte(data))
}

// @summary     Authorization user
// @description Get or create token from/to database
// @tags        users
// @accept      json
// @produce		json
// @router      /user/authorizate [post]
// @param 		username body string true "Your username"
// @param 		password body string true "Your password"
// @success 	200 {object} SessionResponse
// @failure 	400 {object} ErrorResponse "Non-valid data"
// @failure 	400 {object} ErrorResponse "User with this credentials not found"
// @failure 	400 {object} ErrorResponse "Wrong password"
// @failure 	502 {object} ErrorResponse "Create session failed"
func (c *Controller) AuthorizationUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var requestData SessionCreate
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Non-valid data",
		})
		_, _ = w.Write([]byte(data))
	}

	db, err := postgres.GetDatabase(c.Config)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Create session failure",
		})
		_, _ = w.Write([]byte(data))
	}

	var user models.User
	result := db.Where(
		"username = ? and password = ?",
		requestData.Username, security.Hash(requestData.Password),
	).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "User with this credentials not found",
		})
		_, _ = w.Write([]byte(data))
	}

	if !(security.Match(requestData.Password, user.Password)) {
		w.WriteHeader(http.StatusBadRequest)
		data, _ := json.Marshal(map[string]interface{}{
			"message": "Wrong password",
		})
		_, _ = w.Write([]byte(data))
	}

	var session models.Session
	result = db.Where(
		"user_refer = ?", user.ID,
	).First(&session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		session = models.Session{
			UserRefer: user.ID,
			Token:     security.MakeToken(),
		}
		db.Create(&session)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(map[string]interface{}{
		"data": Session{
			Token: session.Token,
		},
	})
	_, _ = w.Write([]byte(data))
}
