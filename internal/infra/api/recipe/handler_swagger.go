package recipe

// @Summary Create recipe
// @Description Create a new recipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param body body entity.RequestCreateRecipe true "Recipe object"
// @Success 200 {object} entity.Recipe "Recipe created successfully"
// @Failure 400 {object} entity.Recipe "Invalid input"
// @Router /api/v1/recipes [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get recipes
// @Description Get all recipes
// @Tags Recipe
// @Produce json
// @Success 200 {array} entity.Recipe "Recipes retrieved successfully"
// @Router /api/v1/recipes [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Update recipe
// @Description Update an existing recipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param uuid path string true "Recipe UUID"
// @Param body body entity.RequestUpdateRecipe true "Recipe object"
// @Success 200 {object} entity.Recipe "Recipe updated successfully"
// @Failure 400 {object} entity.Recipe "Invalid input or UUID format"
// @Router /api/v1/recipes/{uuid} [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Delete recipe
// @Description Delete an existing recipe
// @Tags Recipe
// @Produce json
// @Param uuid path string true "Recipe UUID"
// @Success 200 {object} entity.Recipe "Recipe deleted successfully"
// @Failure 400 {object} entity.Recipe "Invalid UUID format"
// @Router /api/v1/recipes/{uuid} [delete]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Vote recipe
// @Description Vote a recipe
// @Tags Recipe
// @Accept json
// @Produce json
// @Param uuid path string true "Recipe UUID"
// @Param body body entity.RequestVoteRecipe true "Vote object"
// @Success 200 {object} entity.Recipe "Recipe voted successfully"
// @Failure 400 {object} entity.Recipe "Invalid UUID format or input"
// @Router /api/v1/recipes/{uuid}/vote [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
