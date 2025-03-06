package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"superindo-test/model"
	"time"
)

type CartRepository interface {
	AddCart(req model.RequestCart) (data model.Cart, err error)
	GetListCart(userId string) (data []model.CartList, err error)
}

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (cr *cartRepository) AddCart(req model.RequestCart) (data model.Cart, err error) {

	productUuid, err := uuid.Parse(req.ProductId)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	userUuid, err := uuid.Parse(req.UserId)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	cart := model.Cart{
		ProductId: productUuid,
		UserId:    userUuid,
		Quantity:  req.Quantity,
		CreatedAt: time.Now(),
	}

	if err = cr.DB.Create(&cart).Error; err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	data = cart
	return
}

func (cr *cartRepository) GetListCart(userId string) (data []model.CartList, err error) {
	var cart []model.CartList
	query := `select 
    			c.product_id, 
    			p.name as name, 
				Sum(c.quantity) as quantity, 
				p.price 
				from carts c
				inner join products p on p.product_id = c.product_id 
				where c.user_id = ?
				group by c.product_id, p."name", p.price`

	if err = cr.DB.Debug().Raw(query, userId).Scan(&cart).Error; err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	for _, v := range cart {
		data = append(data, model.CartList{
			ProductId:       v.ProductId,
			ProductName:     v.ProductName,
			ProductPrice:    v.ProductPrice,
			ProductQuantity: v.ProductQuantity,
			TotalPrice:      v.ProductPrice * float64(v.ProductQuantity),
		})
	}
	return
}
