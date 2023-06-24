package treatment

// @Summary Create treatment
// @Description Create a new treatment
// @Tags Treatment
// @Accept json
// @Produce json
// @Param name body string true "Name of the treatment"
// @Param description body string false "Description of the treatment"
// @Success 200 {object} entity.Treatment "Treatment created successfully"
// @Failure 400 {object} entity.Treatment "Invalid input"
// @Router /api/v1/treatments [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get treatments
// @Description Get all treatments of a user
// @Tags Treatment
// @Produce json
// @Success 200 {array} entity.Treatment "Treatments retrieved successfully"
// @Router /api/v1/treatments [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Delete treatment
// @Description Delete a treatment
// @Tags Treatment
// @Produce json
// @Param uuid path string true "UUID of the treatment"
// @Success 200 {object} entity.Treatment "Treatment deleted successfully"
// @Failure 500 {object} entity.Treatment "An error occurred while deleting the treatment"
// @Router /api/v1/treatments/{uuid} [delete]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Update treatment
// @Description Update a treatment
// @Tags Treatment
// @Accept json
// @Produce json
// @Param uuid path string true "UUID of the treatment"
// @Param name body string true "Name of the treatment"
// @Param description body string false "Description of the treatment"
// @Success 200 {object} entity.Treatment "Treatment updated successfully"
// @Failure 400 {object} entity.Treatment "Invalid input"
// @Failure 500 {object} entity.Treatment "An error occurred while updating the treatment"
// @Router /api/v1/treatments/{uuid} [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
