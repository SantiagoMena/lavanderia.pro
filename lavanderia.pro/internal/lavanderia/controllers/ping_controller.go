package controllers

type status struct {
	Status string `json:"status"`
}

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (controller PingController) Ping() (status, error) {
	statusObj := status{
		Status: "ok",
	}

	return statusObj, nil
}
