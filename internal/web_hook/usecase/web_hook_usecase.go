package usecase

import (
	"errors"
	"fmt"
	"github.com/xendit/xendit-go"
	"go_online_course/internal/class_room/dto"
	usecase2 "go_online_course/internal/class_room/usecase"
	dto2 "go_online_course/internal/order/dto"
	"go_online_course/internal/order/usecase"
	"os"
	"strings"

	"github.com/xendit/xendit-go/invoice"
)

type WebHookUseCase interface {
	UpdatePayment(id string) error
}

type WebHookUseCaseImpl struct {
	orderUseCase     usecase.OrderUseCase
	classRoomUseCase usecase2.ClassRoomUseCase
}

// UpdatePayment implements WebHookUseCase
func (useCase *WebHookUseCaseImpl) UpdatePayment(id string) error {
	//	Check to xendit payment gateway
	params := invoice.GetParams{ID: id}
	dataXendit, err := invoice.Get(&params)
	if err != nil {
		return err
	}

	dataOrder, errorOrderUseCase := useCase.orderUseCase.FindByExternalID(dataXendit.ExternalID)
	if errorOrderUseCase != nil {
		return errorOrderUseCase
	}

	if dataOrder == nil {
		return errors.New("order not found")
	}

	if dataOrder.Status == "settled" {
		return errors.New("payment already processed")
	}

	if dataOrder.Status != "paid" {
		if dataXendit.Status == "PAID" || dataXendit.Status == "SETTLED" {
			// Add to class_room
			for _, orderDetail := range dataOrder.OrderDetails {
				dataClassRoom := dto.ClassRoom{
					UserID:    dataOrder.UserID,
					ProductID: orderDetail.ProductID,
				}
				fmt.Println(dataClassRoom)

				_, err := useCase.classRoomUseCase.Create(dataClassRoom)

				if err != nil {
					fmt.Println(err)
				}
			}

			// Mengirimkan notif melalui WA / Email ?
		}
	}

	//	update data order
	orderDto := dto2.OrderRequestBody{
		Status: strings.ToLower(dataXendit.Status),
	}
	useCase.orderUseCase.Update(int(dataOrder.ID), orderDto)
	return nil
}

func NewWebHookUseCase(
	orderUseCase usecase.OrderUseCase,
	classRoomUseCase usecase2.ClassRoomUseCase,
) WebHookUseCase {
	// Setup Xendit
	xendit.Opt.SecretKey = os.Getenv("XENDIT_KEY")
	return &WebHookUseCaseImpl{
		orderUseCase:     orderUseCase,
		classRoomUseCase: classRoomUseCase,
	}
}

//TODO web_hook use case
