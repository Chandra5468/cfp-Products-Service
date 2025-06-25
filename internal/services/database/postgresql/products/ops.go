package products

import (
	"database/sql"
	"log"

	"github.com/Chandra5468/cfp-Products-Service/internal/types"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllProducts(pincode int16) {
	// rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	// if err != nil {
	// 	return nil, err
	// }
}

func (s *Store) GetProduct(name string) (pds *types.Product, errs error) {
	query := `select id,name,description,price,quantity,created_at,updated_at from products where name= $1`
	pd := types.Product{}
	err := s.db.QueryRow(query, name).Scan(&pd.Id, &pd.Name, &pd.Description, &pd.Price, &pd.Quantity, &pd.CreatedAt, &pd.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pd, nil
		}
		return nil, err
	}

	return &pd, nil
}

func (s *Store) ProductsCount() {

}

func (s *Store) AddProduct() {

}

func (s *Store) UpdateProductsQuantity(pid *uuid.UUID, quantity *int16) error {
	query := `update products set quantity=$1 where id=$2`

	res, err := s.db.Exec(query, quantity, pid)

	if err != nil {
		return err
	}
	rws, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("Updated quantity for product ", pid, rws)
	return nil
}

func (s *Store) DeleteProduct() {

}

func (s *Store) GetProductByID(id *uuid.UUID) (*types.Product, error) {
	// Check the productId and its quantities and more
	query := `select id, name, description, price, quantity, created_at, updated_at from products where id = $1`

	productDetails := &types.Product{}
	err := s.db.QueryRow(query, id).Scan(&productDetails.Id, &productDetails.Name, &productDetails.Description, &productDetails.Price, &productDetails.Quantity, &productDetails.CreatedAt, &productDetails.UpdatedAt)

	if err != nil {
		return nil, err
	} else {
		return productDetails, nil
	}
}
