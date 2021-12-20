package cofe_storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services"
)

var _ cofe_services.NawMenuStore = &NawMenus{}

type NawMenu struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	//DeletedAt time.Time `db:"deleted_at"`
	Title     string `db:"title"`
	Path      string `db:"path"`
	ItemOrder int32  `db:"item_order"`
	Is_run    bool   `db:"is_run"`
}

type NawMenus struct {
	PG *sql.DB
}

func NewNawMenuStore(db *sql.DB) *NawMenus {

	return &NawMenus{
		PG: db,
	}

}

func (cs *NawMenus) GetListNawMenu(ctx context.Context) ([]cofe_services.NawMenu, error) {

	rows, err := cs.PG.QueryContext(ctx, "select * from naw_menu order by is_run")

	if err != nil {
		log.Println(err)
		//panic(err)
	}
	defer rows.Close()
	mm := []cofe_services.NawMenu{}
	m := cofe_services.NawMenu{}
	for rows.Next() {
		p := NawMenu{}
		err := rows.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Title, &p.Path, &p.ItemOrder, &p.Is_run)
		if err != nil {
			fmt.Println(err)
			continue
		}
		m = cofe_services.NawMenu{
			ID:        p.ID,
			Title:     p.Title,
			Path:      p.Path,
			ItemOrder: p.ItemOrder,
			Is_run:    p.Is_run,
		}
		mm = append(mm, m)
	}

	return mm, nil
}
