package registration

type RegistrationUsecase interface {
	RegisterSession(userId, sessionId uint) error
}
