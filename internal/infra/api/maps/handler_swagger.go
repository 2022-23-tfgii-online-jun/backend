package maps

// @Summary Create map
// @Description Create a new map
// @Tags Maps
// @Accept json
// @Produce json
// @Param body body entity.RequestCreateUpdateMap true "Map object"
// @Success 200 {object} entity.Map "Map created successfully"
// @Failure 400 {object} entity.Map "Invalid request body"
// @Router /api/v1/maps [post]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get all maps
// @Description Get all maps
// @Tags Maps
// @Produce json
// @Success 200 {array} entity.Map "Maps retrieved successfully"
// @Router /api/v1/maps [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Update map
// @Description Update an existing map
// @Tags Maps
// @Accept json
// @Produce json
// @Param uuid path string true "UUID of the map"
// @Param body body entity.RequestCreateUpdateMap true "Map object"
// @Success 200 {object} entity.Map "Map updated successfully"
// @Failure 400 {object} entity.Map "Invalid request body"
// @Failure 404 {object} entity.Map "Map not found"
// @Router /api/v1/maps/{uuid} [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Delete map
// @Description Delete an existing map
// @Tags Maps
// @Param uuid path string true "UUID of the map"
// @Success 200 "Map deleted successfully"
// @Failure 404 {object} entity.Map "Map not found"
// @Router /api/v1/maps/{uuid} [delete]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}
