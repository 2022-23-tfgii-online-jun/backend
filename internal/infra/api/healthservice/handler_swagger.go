package healthservice

// @Summary Create health service
// @Description Create a new health service
// @Tags Health Services
// @Accept json
// @Produce json
// @Param body body entity.RequestCreateHealthService true "Health service object"
// @Success 200 {object} entity.HealthService "Health service created successfully"
// @Failure 400 {object} entity.HealthService "Invalid request body"
// @Router /api/v1/healthservices [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get all health services
// @Description Get all health services
// @Tags Health Services
// @Produce json
// @Success 200 {array} entity.HealthService "Health services retrieved successfully"
// @Router /api/v1/healthservices [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Add rating to health service
// @Description Add a rating to a health service
// @Tags Health Services
// @Accept json
// @Produce json
// @Param body body entity.HealthServiceRating true "Rating object"
// @Success 200 "Rating added successfully"
// @Failure 400 {object} entity.HealthService "Invalid request body"
// @Failure 404 {object} entity.HealthService "Health service not found"
// @Router /api/v1/healthservices/rating [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
