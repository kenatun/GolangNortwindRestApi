package product

import "database/sql"

type Repository interface {
	GetProductById(productId int) (*Product, error)
	GetProducts(params *getProductRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
	GetBestSellers() ([]*ProductTop, error)
	GetTotalVentas() (float64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetProductById(productId int) (*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
				standard_cost,list_price,
				category
				FROM products
				WHERE id=?`
	row := repo.db.QueryRow(sql, productId)
	product := &Product{}

	err := row.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description,
		&product.StandardCost, &product.ListPrice, &product.Category)
	if err != nil {
		panic(err)
	}
	return product, err
}

func (repo *repository) GetProducts(params *getProductRequest) ([]*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),
				standard_cost,list_price,
				category
				FROM products
				LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	if err != nil {
		panic(err)
	}
	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost,
			&product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = "SELECT COUNT(*) FROM products"
	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}

func (repo *repository) InsertProduct(params *getAddProductRequest) (int64, error) {
	const sql = `
	INSERT INTO products
	(product_code,product_name,category,description,list_price,standard_cost)
	VALUES(?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName,
		params.Category, params.Description, params.ListPrice, params.StandardCost)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) UpdateProduct(params *updateProductRequest) (int64, error) {
	const sql = `
		UPDATE products
		SET Product_Code = ?,
		Product_Name=?,
		Category=?,
		Description=?,
		List_Price=?,
		Standard_Cost=?
		WHERE id =?`
	_, err := repo.db.Exec(sql, params.ProductCode, params.ProductName,
		params.Category, params.Description, params.ListPrice, params.StandardCost, params.ID)

	if err != nil {
		panic(err)
	}

	return params.ID, nil

}

func (repo *repository) DeleteProduct(params *deleteProductRequest) (int64, error) {
	const sql = `DELETE FROM products WHERE id =?`
	result, err := repo.db.Exec(sql, params.ProductID)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, nil
}

func (repo *repository) GetBestSellers() ([]*ProductTop, error) {
	const sql = `select
					od.product_id,
					p.product_name,
					SUM(od.quantity*od.unit_price) vendido
				from order_details od
				inner join products p on od.product_id = p.id
				GROUP BY od.product_id
				ORDER by vendido desc
				limit 10`

	results, err := repo.db.Query(sql)
	if err != nil {
		panic(err)
	}
	var products []*ProductTop
	for results.Next() {
		product := &ProductTop{}
		err = results.Scan(&product.ID, &product.ProductName, &product.Vendidos)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *repository) GetTotalVentas() (float64, error) {
	const sql = `select SUM(od.quantity*od.unit_price) vendido
				from order_details od`
	var total float64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}
	return total, nil
}
