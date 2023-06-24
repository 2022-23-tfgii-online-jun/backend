package reminder

// @Summary Create reminder
// @Description Create a new reminder
// @Tags Reminder
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name of the reminder"
// @Param type formData string true "Type of the reminder"
// @Param date formData string true "Date of the reminder (format: dd/MM/yyyy)"
// @Param note formData string false "Additional note for the reminder"
// @Param notification formData string true "Notification details (JSON array)"
// @Param task formData string true "Task details (JSON array)"
// @Success 200 {object} entity.Reminder "Reminder created successfully"
// @Failure 400 {object} entity.Reminder "Invalid input or date format"
// @Router /api/v1/reminders [post]
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get reminders
// @Description Get all reminders
// @Tags Reminder
// @Produce json
// @Success 200 {array} entity.Reminder "Reminders fetched successfully"
// @Router /api/v1/reminders [get]
// @Security Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Update reminder
// @Description Update an existing reminder
// @Tags Reminder
// @Accept multipart/form-data
// @Produce json
// @Param uuid query string true "Reminder UUID"
// @Param name formData string true "Name of the reminder"
// @Param type formData string true "Type of the reminder"
// @Param date formData string true "Date of the reminder (format: dd/MM/yyyy)"
// @Param note formData string false "Additional note for the reminder"
// @Param notification formData string true "Notification details (JSON array)"
// @Param task formData string true "Task details (JSON array)"
// @Success 200 {object} entity.Reminder "Reminder updated successfully"
// @Failure 400 {object} entity.Reminder "Invalid UUID format, input, or date format"
// @Router /api/v1/reminders [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Delete reminder
// @Description Delete an existing reminder
// @Tags Reminder
// @Produce json
// @Param uuid query string true "Reminder UUID"
// @Success 200 {object} entity.Reminder "Reminder deleted successfully"
// @Failure 400 {object} entity.Reminder "Invalid UUID format"
// @Router /api/v1/reminders [delete]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
