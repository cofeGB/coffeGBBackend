package nawmenustore

import (
	"context"
	"fmt"
	"log"

	//"database/sql"
	//"log"
	"time"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services/nawmenu"
	"github.com/cofeGB/coffeGBBackend/internal/cofe_storage/storage"

	//"github.com/lib/pq"

	"github.com/google/uuid"
)

var _ nawmenu.NawMenuStore = &NawMenuStor{}

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

type NawMenuStor struct {
	Store *storage.CofeDB
}

func NewNawMenuStore(db *storage.CofeDB) *NawMenuStor {

	st := &NawMenuStor{
		Store: db,
	}
	// err := migrateNawMenu(st)
	// if err != nil {
	// 	log.Fatalf("cannot initialize migration NawMenu struct: %s", err.Error())
	// }
	return st
}

func (nm *NawMenuStor) GetListNawMenu(ctx context.Context) ([]nawmenu.NawMenu, error) {

	//result := nm.store.DB.Order("item_order").Find(&NawMenu)
	// result := nm.store.DB.Find(&NawMenuList)
	// if result.Error != nil {
	// 	log.Println(result.Error)
	// 	return nil, result.Error
	// }
	// rows, err := nm.store.PG.Query("select id, created_at,
	//                                 updated_at,title,
	// 								path,item_order,is_run
	// 								 from naw_menu order by is_run",)

	rows, err := nm.Store.PG.Query("select * from naw_menu order by is_run")


	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer rows.Close()
	//NawMenuList := []NawMenu{}
	mm := []nawmenu.NawMenu{}
	m := nawmenu.NawMenu{}
	for rows.Next() {
		p := NawMenu{}
		err := rows.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Title, &p.Path, &p.ItemOrder, &p.Is_run)
		if err != nil {
			fmt.Println(err)
			continue
		}
		m = nawmenu.NawMenu{
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

func migrateNawMenu(nm *NawMenuStor) error {
	// if err := nm.store.DB.AutoMigrate(&NawMenu{}); err != nil {
	// 	return err
	// }
	return nil
}
