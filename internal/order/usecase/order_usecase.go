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
	entity3 "go_online_course/internal/product/entity"
	usecase3 "go_online_course/internal/product/usecase"
	"strconv"
)

type OrderUseCase interface {
	FindAll(offset int, limit int) []entity.Order
	FindByID(id int) (*entity.Order, error)
	Create(dto dto.OrderRequestBody) (*entity.Order, error)
}

type OrderUseCaseImpl struct {
	repository         repository.OrderRepository
	cartUseCase        usecase.CartUseCase
	discountUseCase    usecase2.DiscountUseCase
	productUseCase     usecase3.ProductUseCase
	orderDetailUseCase usecase4.OrderDetailUseCase
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
		dataDiscount, err := useCase.discountUseCase.FindByCode(*dto.DiscountCode)
		if err != nil {
			return nil, errors.New("discount is invalid")
		}
		if dataDiscount.RemainingQuantity == 0 {
			return nil, errors.New("discount quota has run out")
		}

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

	if dataDiscount != nil {
		//	calculate discount logic
		if dataDiscount.Type == "rebate" {
			totalPrice = price - int(dataDiscount.Value)
		} else if dataDiscount.Type == "percent" {
			totalPrice = price - (price / 100 * int(dataDiscount.Value))
		}
		order.DiscountID = dataDiscount.ID
	}

	order.Price = int64(price)
	order.TotalPrice = int64(totalPrice) //price after discounted

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
}

func NewOrderUseCase(
	repository repository.OrderRepository,
	cartUseCase usecase.CartUseCase,
	discountUseCase usecase2.DiscountUseCase,
	productUseCase usecase3.ProductUseCase,
	orderDetailUseCase usecase4.OrderDetailUseCase,
) OrderUseCase {
	return &OrderUseCaseImpl{
		repository:         repository,
		cartUseCase:        cartUseCase,
		discountUseCase:    discountUseCase,
		productUseCase:     productUseCase,
		orderDetailUseCase: orderDetailUseCase,
	}
}