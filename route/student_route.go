package route

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jrione/gin-crud/config"
	"github.com/jrione/gin-crud/domain"
	_studentRepo "github.com/jrione/gin-crud/repository/postgres"
	_studentUseCase "github.com/jrione/gin-crud/usecase"
)

type StudentHandler struct {
	SusCase domain.StudentUseCase
}

func NewStudentRoute(env *config.Config, timeout time.Duration, db *sql.DB, gr *gin.RouterGroup) {
	_student := _studentRepo.NewStudentRepository(db)
	studentUseCase := _studentUseCase.NewStudentUseCase(_student, timeout)

	handler := &StudentHandler{
		SusCase: studentUseCase,
	}

	gr.GET("/student", handler.FetchStudent)
}

func (s StudentHandler) FetchStudent(gin *gin.Context) {
	ctx := gin
	listStudent, err := s.SusCase.Fetch(ctx)
	if err != nil {
		log.Fatal(err)
	}
	gin.JSON(http.StatusOK, listStudent)
}
