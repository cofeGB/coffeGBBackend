package cofe_storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services"
)

var _ cofe_services.CategoryStore = &CategoryStore{}

type Category struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Query     string    `db:"query"`
	Icon      string    `db:"icon"`
	ItemOrder int32     `db:"item_order"`
}

type CategoryStore struct {
	PG *sql.DB
}

func NewCategoryStore(db *CofeDB) *CategoryStore {
	return &CategoryStore{
		PG: db.PG,
	}
}

func (cs *CategoryStore) GetCategories(ctx context.Context) ([]cofe_services.Category, error) {

	rows, err := cs.PG.QueryContext(ctx, "select * from categories order by item_order")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	items := []cofe_services.Category{}
	for rows.Next() {
		p := Category{}
		err := rows.Scan(&p.ID, &p.Title, &p.Query, &p.Icon, &p.ItemOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		item := cofe_services.Category{
			ID:        p.ID,
			Title:     p.Title,
			Query:     p.Query,
			Icon:      p.Icon,
			ItemOrder: p.ItemOrder,
		}
		items = append(items, item)
	}

	return items, nil
}
