package category

// @Summary Create category
// @Description Create a new category
// @Tags Categories
// @Accept json
// @Produce json
// @Param body body entity.Category true "Category object"
// @Success 200 {object} entity.Category "Category created successfully"
// @Failure 400 {object} entity.Category "Invalid request body"
// @Router /api/v1/categories [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get all categories
// @Description Get all categories
// @Tags Categories
// @Produce json
// @Success 200 {array} entity.Category "Categories retrieved successfully"
// @Router /api/v1/categories [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Update category
// @Description Update an existing category
// @Tags Categories
// @Accept json
// @Produce json
// @Param uuid path string true "UUID of the category"
// @Param body body entity.Category true "Category object"
// @Success 200 {object} entity.Category "Category updated successfully"
// @Failure 400 {object} entity.Category "Invalid request body"
// @Failure 404 {object} entity.Category "Category not found"
// @Router /api/v1/categories/{uuid} [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Delete category
// @Description Delete an existing category
// @Tags Categories
// @Param uuid path string true "UUID of the category"
// @Success 200 "Category deleted successfully"
// @Failure 404 {object} entity.Category "Category not found"
// @Router /api/v1/categories/{uuid} [delete]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
