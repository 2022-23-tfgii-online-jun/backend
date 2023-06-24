package user

// @Summary User login
// @Description Authenticate user and generate JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param body body entity.DefaultCredentials true "User credentials"
// @Success 200 {object} TokenResponse "Token generated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Router /api/v1/users/login [post]
func _() {
	// Swagger annotations.
}

// @Summary User Sign Up
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param body body entity.SignUp true "User Sign Up data"
// @Success 201 {object} ResponseUserData "User registered successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Router /api/v1/users/signup [post]
func _() {
	// Swagger annotations.
}

// @Summary Update user
// @Description Update existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param body body entity.UpdateUser true "User update data"
// @Success 200 {object} ResponseUserData "User updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid input"
// @Router /api/v1/users [patch]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Get user information
// @Description Get user information
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "User information retrieved successfully"
// @Failure 404 {object} ErrorResponse "User not found"
// @Failure 500 {object} ErrorResponse "An error occurred while getting the user"
// @Router /api/v1/users [get]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Set user active status
// @Description Set user's active status
// @Tags Users
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Param body body SetStatusRequest true "Set user active status"
// @Success 200 {object} ResponseUserData "User active status updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 500 {object} ErrorResponse "An error occurred while updating the user's active status"
// @Router /api/v1/users/active/{uuid} [patch]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// @Summary Set user banned status
// @Description Set user's banned status
// @Tags Users
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Param body body SetStatusRequest true "Set user banned status"
// @Success 200 {object} ResponseUserData "User banned status updated successfully"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 500 {object} ErrorResponse "An error occurred while updating the user's banned status"
// @Router /api/v1/users/banned/{uuid} [patch]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func _() {
	// Swagger annotations.
}

// TokenResponse represents the response structure for the login endpoint.
type TokenResponse struct {
	Token string `json:"token"`
}

// ErrorResponse represents the response structure for error responses.
type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SetStatusRequest represents the request structure for updating user status.
type SetStatusRequest struct {
	Status bool `json:"status"`
}
