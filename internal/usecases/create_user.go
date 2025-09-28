type CreateUserUseCase struct {
	userRepo domain.UserRepository
	emailSvc EmailService
	eventBus EventBus
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, req CreateUserRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	user, err := domain.NewUser(req.Email, req.Name)
	if err != nil {
		return err
	}

	exists, err := uc.userRepo.ExistsByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return domain.ErrUserAlreadyExists
	}

	if err := uc.userRepo.Save(ctx, user); err != nil {
		return err
	}

	uc.emailSvc.SendWelcomeEmail(user.Email)
	uc.eventBus.Publish(UserCreatedEvent{UserID: user.ID})

	return nil
}