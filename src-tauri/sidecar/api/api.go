package api

import (
	"automatic_api/hvc"

	"github.com/gin-gonic/gin"
)

type Device struct {
	hvc.Device
	Status bool `json:"status"`
}

func Run() error {
	e := gin.Default()

	return e.Run("0.0.0.0:8000")

}

func GetDevices(c *gin.Context) {
	sql := "select * from device"
	devices := []*hvc.Device{}

}
