package handlers

import (
  "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/task_iman/pkg/logger"
)

type handlerV1 struct {
  log logger.Logger
}

type HandlerV1Options struct {
  Log logger.Logger
}

func New(options *HandlerV1Options) *handlerV1 {
  return &handlerV1{
    log: options.Log,
  }
}
// @Security BearerAuth
// @Router /v1/days [get]
func (h *handlerV1) Days(ctx *gin.Context){
    

    
    targetYear := 2025
    targetMonth := time.January
    targetDay := 1

    currentTime := time.Now()

    targetDate := time.Date(targetYear, targetMonth, targetDay, 0,0,0,0, time.UTC)

    daysLeft := int(targetDate.Sub(currentTime).Hours() /24)

    fmt.Printf("Days until January 1, 2025 %d\n:", daysLeft)

    ctx.JSON(http.StatusOK, gin.H{
      "Days until January 1, 2025: " : daysLeft,
    })
}

