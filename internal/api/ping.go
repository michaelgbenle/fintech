package api

//PingHandler is for testing the connections
func (u *HTTPHandler) PingHandler(c *gin.Context) {
	data := &model.Student{}

	// healthcheck
	helpers.Response(c, "pong", 200, data, nil)
}
