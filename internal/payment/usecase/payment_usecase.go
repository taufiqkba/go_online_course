package usecase

type PaymentUseCase interface {
}

type PaymentUseCaseImpl struct {
}

func NewPaymentUseCase() PaymentUseCase {
	return &PaymentUseCaseImpl{}
}
