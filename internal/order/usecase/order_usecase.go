package usecase

import (
	"errors"
	"github.com/google/uuid"
	"go_online_course/internal/cart/usecase"
	entity2 "go_online_course/internal/discount/entity"
	usecase2 "go_online_course/internal/discount/usecase"
	"go_online_course/internal/order/dto"
	"go_online_course/internal/order/entity"
	"go_online_course/internal/order/repository"
	entity4 "go_online_course/internal/order_detail/entity"
	usecase4 "go_online_course/internal/order_detail/usecase"
	dto2 "go_online_course/internal/payment/dto"
	usecase5 "go_online_course/internal/payment/usecase"
	entity3 "go_online_course/internal/product/entity"
	usecase3 "go_online_course/internal/product/usecase"
	"strconv"
)

type OrderUseCase interface {
	FindAll(offset int, limit int) []entity.Order
	FindAllByUserID(offset int, limit int, userID int) []entity.Order
	FindByID(id int) (*entity.Order, error)
	FindByExternalID(externalID string) (*entity.Order, error)
	Create(dto dto.OrderRequestBody) (*entity.Order, error)
	Update(id int, dto dto.OrderRequestBody) (*entity.Order, error)
}

type OrderUseCaseImpl struct {
	repository         repository.OrderRepository
	cartUseCase        usecase.CartUseCase
	discountUseCase    usecase2.DiscountUseCase
	productUseCase     usecase3.ProductUseCase
	orderDetailUseCase usecase4.OrderDetailUseCase
	paymentUseCase     usecase5.PaymentUseCase
}

func (useCase *OrderUseCaseImpl) FindAllByUserID(offset int, limit int, userID int) []entity.Order {
	return useCase.repository.FindAllByUserID(offset, limit, userID)
}

func (useCase *OrderUseCaseImpl) Update(id int, dto dto.OrderRequestBody) (*entity.Order, error) {
	order, err := useCase.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	order.Status = dto.Status

	updateOrder, err := useCase.repository.Update(*order)
	if err != nil {
		return nil, err
	}
	return updateOrder, nil
}

func (useCase *OrderUseCaseImpl) FindByExternalID(externalID string) (*entity.Order, error) {
	return useCase.repository.FindOneByExternalID(externalID)
}

func (useCase *OrderUseCaseImpl) FindAll(offset int, limit int) []entity.Order {
	return useCase.repository.FindAll(offset, limit)
}

func (useCase *OrderUseCaseImpl) FindByID(id int) (*entity.Order, error) {
	return useCase.repository.FindByID(id)
}

func (useCase *OrderUseCaseImpl) Create(dto dto.OrderRequestBody) (*entity.Order, error) {
	//define first
	price := 0
	totalPrice := 0
	description := ""
	var products []entity3.Product

	order := entity.Order{
		UserID: dto.UserID,
		Status: "pending",
	}

	var dataDiscount *entity2.Discount
	//	find data by user id
	carts := useCase.cartUseCase.FindByUserID(int(dto.UserID), 0, 99)

	//check the carts is empty or not
	if len(carts) == 0 {
		//	if cart isEmpty
		if dto.ProductID == nil {
			return nil, errors.New("your carts is empty")
		}
	}

	//	check discount
	if dto.DiscountCode != nil {
		discount, err := useCase.discountUseCase.FindByCode(*dto.DiscountCode)
		if err != nil {
			return nil, errors.New("discount is invalid")
		}
		if discount.RemainingQuantity == 0 {
			return nil, errors.New("discount quota has run out")
		}
		dataDiscount = discount
		//	other validation here
	}
	if len(carts) > 0 {
		//if used cart
		for _, cart := range carts {
			product, err := useCase.productUseCase.FindById(int(cart.ProductID))
			if err != nil {
				return nil, err
			}
			products = append(products, *product)
		}
	} else if dto.ProductID != nil {
		//if user directly checkout
		product, err := useCase.productUseCase.FindById(int(*dto.ProductID))
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	//calculate price & create description to xendit payment gateway
	for index, product := range products {
		price += int(product.Price)
		i := strconv.Itoa(index + 1)
		description += i + ". Product : " + product.Title + "<br/>"
	}

	totalPrice = price
	if dataDiscount != nil {
		//	calculate discount logic
		if dataDiscount.Type == "rebate" {
			totalPrice = price - int(dataDiscount.Value)
		} else if dataDiscount.Type == "percent" {
			totalPrice = price - (price / 100 * int(dataDiscount.Value))
		}
		order.DiscountID = &dataDiscount.ID
	}

	order.Price = int64(price)
	order.TotalPrice = int64(totalPrice) //price after discounted
	order.CreatedByID = &dto.UserID

	externalId := uuid.New().String()
	order.ExternalID = externalId

	//	insert data to order table
	data, err := useCase.repository.Create(order)
	if err != nil {
		return nil, err
	}
	//	insert data to order_detail table
	for _, product := range products {
		orderDetail := entity4.OrderDetail{
			OrderID:     data.ID,
			ProductID:   product.ID,
			CreatedByID: &order.UserID,
			Price:       product.Price,
		}
		useCase.orderDetailUseCase.Create(orderDetail)
	}
	//	hit to xendit payment gateway
	dataPayment := dto2.PaymentRequestBody{
		ExternalID:  externalId,
		Amount:      data.TotalPrice,
		PayerEmail:  dto.Email,
		Description: description,
	}

	payment, err := useCase.paymentUseCase.Create(dataPayment)
	if err != nil {
		return nil, err
	}
	data.CheckoutLink = payment.InvoiceURL
	useCase.repository.Update(*data)
	//	Update RemainingQuantityDiscount
	if dto.DiscountCode != nil {
		_, err := useCase.discountUseCase.UpdateRemainingQuantity(int(dataDiscount.ID), 1, "-")
		if err != nil {
			return nil, err
		}
	}

	//	delete carts while checkout success
	err = useCase.cartUseCase.DeleteByUserID(int(dto.UserID))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewOrderUseCase(
	repository repository.OrderRepository,
	cartUseCase usecase.CartUseCase,
	discountUseCase usecase2.DiscountUseCase,
	productUseCase usecase3.ProductUseCase,
	orderDetailUseCase usecase4.OrderDetailUseCase,
	paymentUseCase usecase5.PaymentUseCase,
) OrderUseCase {
	return &OrderUseCaseImpl{
		repository:         repository,
		cartUseCase:        cartUseCase,
		discountUseCase:    discountUseCase,
		productUseCase:     productUseCase,
		orderDetailUseCase: orderDetailUseCase,
		paymentUseCase:     paymentUseCase,
	}
}
