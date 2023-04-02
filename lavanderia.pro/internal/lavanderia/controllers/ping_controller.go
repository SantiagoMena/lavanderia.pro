package controllers

type status struct {
	Status string `json:"status"`
}

type PingController struct {
	// c *gin.Context
}

func NewPingController() *PingController {
	return &PingController{}
}

func (controller PingController) Ping() (status, error) {
	statusObj := status{
		Status: "ok",
	}
	// laundries, err := controller.LaundryRepository.FindAllLaundries()
	return statusObj, nil
}

// package controllers

// import (
// 	"github.com/gin-gonic/gin"
// )

// // import (
// // 	"github.com/gin-gonic/gin"
// // )

// // func Ping(c *gin.Context) {

// // 	c.JSON(200, statusObj)
// // }
// // package controllers

// type PingController struct {
// 	c *gin.Context
// }

// func NewPingController() *PingController {
// 	return &PingController{
// 		PingController: Ping(),
// 	}
// }

// func (controller PingController) Ping() (status, error) {
// 	return status{
// 		Status: "ok",
// 	}, nil
// }
