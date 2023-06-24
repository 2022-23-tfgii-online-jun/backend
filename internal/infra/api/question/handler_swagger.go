package question

// @Summary Create question
// @Description Create a new question
// @Tags Question
// @Accept json
// @Produce json
// @Param body body entity.RequestCreateQuestion true "Question object"
// @Success 200 {object} entity.Question "Question created successfully"
// @Failure 400 {object} entity.Question "Invalid input"
// @Router /api/v1/questions [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get questions
// @Description Get all questions
// @Tags Question
// @Produce json
// @Success 200 {array} entity.Question "Questions retrieved successfully"
// @Router /api/v1/questions [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get questions and answers
// @Description Get all questions and their answers
// @Tags Question
// @Produce json
// @Success 200 {array} entity.Question "Questions and answers retrieved successfully"
// @Router /api/v1/questions/{uuid} [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
