type UserHandler struct {
	createUserUC *usecases.CreateUserUseCase
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	err := h.createUserUC.Execute(c.Request.Context(), req)
	if err != nil {
		switch err {
		case domain.ErrUserAlreadyExists:
			c.JSON(409, gin.H{"error": "user already exists"})
		default:
			c.JSON(500, gin.H{"error": "internal error"})
		}
		return
	}

	c.JSON(201, gin.H{"message": "user created"})
}