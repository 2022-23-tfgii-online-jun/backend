package medicalrecord

// @Summary Create medical record
// @Description Create a new medical record
// @Tags Medical Record
// @Accept json
// @Produce json
// @Param body body entity.MedicalRecord true "Medical record object"
// @Success 200 {object} entity.MedicalRecord "Medical record created successfully"
// @Failure 400 {object} entity.MedicalRecord "Invalid input"
// @Router /api/v1/medicalrecords [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get medical record
// @Description Get the medical record for the authenticated user
// @Tags Medical Record
// @Produce json
// @Success 200 {object} entity.MedicalRecord "Medical record retrieved successfully"
// @Failure 400 {object} entity.MedicalRecord "Invalid user UUID"
// @Router /api/v1/medicalrecords [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Update medical record
// @Description Update an existing medical record
// @Tags Medical Record
// @Accept json
// @Produce json
// @Param uuid path string true "Medical record UUID"
// @Param body body entity.MedicalRecord true "Medical record object"
// @Success 200 {object} entity.MedicalRecord "Medical record updated successfully"
// @Failure 400 {object} entity.MedicalRecord "Invalid input"
// @Router /api/v1/medicalrecords/{uuid} [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
