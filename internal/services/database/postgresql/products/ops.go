package products

import (
	"database/sql"

	"github.com/Chandra5468/cfp-Products-Service/internal/types"
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

func (s *Store) GetProduct(id string) (pds *types.Product, errs error) {
	query := `select id,name,description,price,quantity,created_at,updated_at from products where name= $1`
	pd := types.Product{}
	err := s.db.QueryRow(query, id).Scan(&pd.Id, &pd.Name, &pd.Description, &pd.Price, &pd.Quantity, &pd.CreatedAt, &pd.UpdatedAt)

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

func (s *Store) UpdateProducts() {

}

func (s *Store) DeleteProduct() {

}
