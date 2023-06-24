package symptom

// @Summary Create symptom
// @Description Create a new symptom
// @Tags Symptom
// @Accept json
// @Produce json
// @Param name body string true "Name of the symptom"
// @Success 200 {object} entity.Symptom "Symptom created successfully"
// @Failure 400 {object} entity.Symptom "Invalid input"
// @Router /api/v1/symptoms/admin [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get symptoms
// @Description Get all symptoms
// @Tags Symptom
// @Produce json
// @Success 200 {array} entity.Symptom "Symptoms retrieved successfully"
// @Router /api/v1/symptoms/admin [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Add user to symptom
// @Description Add a user to a symptom
// @Tags Symptom
// @Accept json
// @Produce json
// @Param userUUID body string true "User UUID"
// @Param symptomUUID body string true "Symptom UUID"
// @Success 200 {object} entity.Symptom "User added to symptom successfully"
// @Failure 400 {object} entity.Symptom "Invalid user UUID, input, or symptom UUID"
// @Router /api/v1/symptoms/user/add [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Remove user from symptom
// @Description Remove a user from a symptom
// @Tags Symptom
// @Accept json
// @Produce json
// @Param userUUID body string true "User UUID"
// @Param symptomUUID body string true "Symptom UUID"
// @Success 200 {object} entity.Symptom "User removed from symptom successfully"
// @Failure 400 {object} entity.Symptom "Invalid user UUID, input, or symptom UUID"
// @Router /api/v1/symptoms/user/remove [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get symptoms by user
// @Description Get all symptoms related to a user
// @Tags Symptom
// @Produce json
// @Success 200 {array} entity.Symptom "Symptoms retrieved successfully"
// @Router /api/v1/symptoms/user [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
