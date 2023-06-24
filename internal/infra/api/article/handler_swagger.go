package article

// @Summary Create article
// @Description Create an article
// @Tags Articles
// @Accept json
// @Produce json
// @Param userUUID query string true "UUID of the user"
// @Param body body entity.RequestCreateArticle true "Body of the article"
// @Success 200 {object} entity.Article "Article created successfully"
// @Failure 400 {object} entity.Article "Invalid input"
// @Router /api/v1/articles [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get all articles
// @Description Get all articles
// @Tags Articles
// @Accept json
// @Produce json
// @Success 200 {array} entity.Article "Articles retrieved successfully"
// @Failure 500 {object} entity.Article "An error occurred while getting the articles"
// @Router /api/v1/articles [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Update article
// @Description Update an article
// @Tags Articles
// @Accept json
// @Produce json
// @Param uuid path string true "UUID of the article"
// @Param body body entity.RequestUpdateArticle true "Body of the article"
// @Success 200 {object} entity.Article "Article updated successfully"
// @Failure 400 {object} entity.Article "Invalid input"
// @Router /api/v1/articles/{uuid} [patch]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Delete article
// @Description Delete an article
// @Tags Articles
// @Accept json
// @Produce json
// @Param uuid path string true "UUID of the article"
// @Success 200 {object} entity.Article "Article deleted successfully"
// @Failure 500 {object} entity.Article "An error occurred while deleting the article"
// @Router /api/v1/articles/{uuid} [delete]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Add article to category
// @Description Add an article to a category
// @Tags Articles
// @Accept json
// @Produce json
// @Param body body entity.AddArticleToCategoryRequest true "Request to add article to category"
// @Success 200 {object} entity.Article "Article added to Category successfully"
// @Failure 400 {object} entity.Article "Invalid input"
// @Router /api/v1/articles/add-to-category [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
