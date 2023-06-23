package monitoring

// @Summary Create monitoring
// @Description Create a new monitoring
// @Tags Monitoring
// @Accept json
// @Produce json
// @Param body body entity.RequestCreateMonitoring true "Monitoring object"
// @Success 200 {object} entity.Monitoring "Monitoring created successfully"
// @Failure 400 {object} entity.Monitoring "Invalid input"
// @Router /api/v1/monitorings [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get monitorings
// @Description Get all monitorings for the authenticated user
// @Tags Monitoring
// @Produce json
// @Success 200 {array} entity.Monitoring "Monitorings retrieved successfully"
// @Failure 400 {object} entity.Monitoring "Invalid user UUID"
// @Router /api/v1/monitorings [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
