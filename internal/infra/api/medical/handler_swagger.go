package medical

// @Summary Get all medical records
// @Description Get all medical records
// @Tags Medical
// @Produce json
// @Success 200 {array} entity.Medical "Medical records retrieved successfully"
// @Router /api/v1/medical [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Add rating to medical record
// @Description Add a rating to a medical record
// @Tags Medical
// @Accept json
// @Produce json
// @Param body body entity.MedicalRating true "Rating object"
// @Success 200 "Rating added successfully"
// @Failure 400 {object} entity.Medical "Invalid input"
// @Failure 404 {object} entity.Medical "Medical record not found"
// @Router /api/v1/medical/rating [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
