package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"test/models"
	"time"
)

type CommController struct {
	DB *gorm.DB
}

func NewCommController(DB *gorm.DB) CommController {
	return CommController{DB}
}
func (pc *CommController) CreateComm(ctx *gin.Context) {

	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.CreateCommRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var postID = uuid.Must(uuid.Parse(payload.Post))
	now := time.Now()
	newComm := models.Comm{
		User:      currentUser.ID,
		Post:      postID,
		Content:   payload.Content,
		CreatedAt: now,
	}

	result := pc.DB.Create(&newComm)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Comm with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newComm})
}
func (pc *CommController) DeleteComm(ctx *gin.Context) {
	commId := ctx.Param("commId")

	result := pc.DB.Delete(&models.Comm{}, "id = ?", commId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No comm with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (pc *CommController) FindComms(ctx *gin.Context) {
	postId := ctx.Param("postId")
	//var page = ctx.DefaultQuery("page", "1")
	//var limit = ctx.DefaultQuery("limit", "10")

	//intPage, _ := strconv.Atoi(page)
	//intLimit, _ := strconv.Atoi(limit)
	//offset := (intPage - 1) * intLimit

	var comms []models.Comm
	results := pc.DB.Find(&comms, "post = ?", postId)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(comms), "data": comms})
}
