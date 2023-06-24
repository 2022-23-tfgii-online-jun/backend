package answer

// @Summary Create answer
// @Description Create a answer for a specific question
// @Tags Answers
// @Accept json
// @Produce json
// @Param question_uuid path string true "UUID of the question"
// @Param userUUID query string true "UUID of the user"
// @Param body body entity.RequestCreateAnswer true "Body of the answer"
// @Success 200 {object} entity.Answer "Answer created successfully"
// @Failure 400 {object} entity.Answer "Invalid request body"
// @Router /api/v1/questions/{question_uuid}/answer [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
