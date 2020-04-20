package order

import (
	"database/sql"

	"github.com/GolangNortwindRestApi/helper"
)

type Repository interface {
	GetOrderById(param *getOrderByIdRequest) (*OrderItem, error)
}
type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetOrderById(param *getOrderByIdRequest) (*OrderItem, error) {
	var sql = `select o.id, o.customer_id,o.order_date,o.status_id,os.status_name,
	concat(c.first_name,'',c.last_name) as customer_name,
    c.company,
    c.address,
    c.business_phone,
    c.city
    from orders o
	inner join orders_status os on o.status_id = os.id
    inner join customers c on o.customer_id = c.id
	where o.id = ?`

	order := &OrderItem{}

	row := repo.db.QueryRow(sql, param.orderId)
	err := row.Scan(&order.ID, &order.CustomerID, &order.OrderDate,
		&order.StatusId, &order.StatusName, &order.Customer,
		&order.Company, &order.Address, &order.Phone, &order.City)
	helper.Catch(err)
	orderDetail, err := GetOrderDetail(repo, &param.orderId)
	helper.Catch(err)
	order.Data = orderDetail
	return order, nil
}

func GetOrderDetail(repo *repository, orderId *int64) ([]*OrderDetailItem, error) {
	const sql = `select order_id,od.id,quantity,unit_price,p.product_name,product_id
    from order_details od
    inner join products p on od.product_id = p.id
	where od.order_id = ?`

	results, err := repo.db.Query(sql, orderId)
	helper.Catch(err)

	var orders []*OrderDetailItem
	for results.Next() {
		order := &OrderDetailItem{}
		err = results.Scan(&order.OrderId, &order.ID, &order.Quantity,
			&order.UnitPrice, &order.ProductName, &order.ProductId)
		helper.Catch(err)
		orders = append(orders, order)
	}
	return orders, nil

}
