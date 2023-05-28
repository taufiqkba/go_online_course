package usecase

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"go_online_course/internal/payment/dto"
	"os"
)

type PaymentUseCase interface {
	Create(dto dto.PaymentRequestBody) (*xendit.Invoice, error)
}

type PaymentUseCaseImpl struct {
}

func (useCase *PaymentUseCaseImpl) Create(dto dto.PaymentRequestBody) (*xendit.Invoice, error) {
	data := invoice.CreateParams{
		ExternalID:  dto.ExternalID,
		Amount:      float64(dto.Amount),
		Description: dto.Description,
		PayerEmail:  dto.PayerEmail,
		Customer: xendit.InvoiceCustomer{
			Email: dto.PayerEmail,
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"email"},
			InvoiceReminder: []string{"email"},
			InvoicePaid:     []string{"email"},
			InvoiceExpired:  []string{"email"},
		},
		InvoiceDuration:    86400,
		SuccessRedirectURL: os.Getenv("XENDIT_SUCCESS_URL"),
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewPaymentUseCase() PaymentUseCase {
	//Setup Xendit PaymentGateway
	xendit.Opt.SecretKey = os.Getenv("XENDIT_KEY")
	return &PaymentUseCaseImpl{}
}
