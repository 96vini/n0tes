type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	dbUser := &UserModel{
		ID:    user.ID.String(),
		Email: user.Email.String(),
		Name:  user.Name,
	}

	return r.db.WithContext(ctx).Create(dbUser).Error
}