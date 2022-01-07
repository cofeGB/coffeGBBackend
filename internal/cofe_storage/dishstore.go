package cofe_storage

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services"
)

var _ cofe_services.DishStore = &Dishs{}

type DishNutrientDb struct {
	ID         uuid.UUID `db:"id"`
	DishId     uuid.UUID `db:"dish_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Category   string    `db:"category"`
	Percentage float64   `db:"percentage"`
	Value      int32     `db:"value"`
}

type DishIngredientDB struct {
	ID          uuid.UUID // guGuid - идентификатор ингредиента (Global Unchanged Guidentifier)
	Category    string
	//DishId      uuid.UUID
	Title       string
	Description string
	Weight      float64
	Volume      float64
	Price       float64
	Quantity    float64
	//  Nutrients FoodNutrients
	//	ImgUrls []string
	Brand        string
	Origin       string
	Availability string
	Warnings     string
}

type DishDb struct {
	ID         uuid.UUID `db:"id"` //DishGuid uuid.UUID
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	CategoryId uuid.UUID `db:"category_id"`
	//DishGuid uuid.UUID
	CreatorGuGuid uuid.UUID `db:"creator_gu_guid"`
	Title         string    `db:"title"`
	Description   string    `db:"description"`
	Weight        float64   `db:"weight"`
	Volume        float64   `db:"volume"`
	Price         float64   `db:"price"`
	Quantity      int32     `db:"quantity"`
	//Nutrients     uuid.UUID `db:"nutrients"`
	//	ImgUrls       uuid.UUID `db:"img_urls"`
	//	Ingredients   uuid.UUID `db:"ingredients"`
	Availability int32  `db:"availability"`
	Warnings     string `db:"warnings"`
	IsRun        bool   `db:"is_run"`
}

type Dishs struct {
	PG *sql.DB
}

func NewDishStore(dbConn *CofeDB) *Dishs {
	return &Dishs{
		PG: dbConn.PG,
	}

}

func (dish *Dishs) GetListDish(ctx context.Context) ([]cofe_services.Dish, error) {

	listDishDb, err := dish.GetListDishDb(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}



	cofeDish := cofe_services.Dish{}
	listCofeDish := []cofe_services.Dish{}

	for _, d := range listDishDb {

		dishCategory, err := dish.GetDishCategory(ctx, d.CategoryId)
		if err != nil {
			log.Println(err)

		}
		dishNutrient, err := dish.GetDishNutrient(ctx, d.ID)
		if err != nil {
			log.Println(err)

		}

		dishImage, err := dish.GetDishImageUrls(ctx, d.ID)
		if err != nil {
			log.Println(err)

		}

		dishIngradient, err := dish.GetDishIngredient(ctx, d.ID)
		if err != nil {
			log.Println(err)

		}

		cofeDish = cofe_services.Dish{
			Category:      dishCategory,
			DishGuid:      d.ID,
			CreatorGuGuid: d.CreatorGuGuid,
			Title:         d.Title,
			Description:   d.Description,
			Weight:        d.Weight,
			Volume:        d.Volume,
			Price:         d.Price,
			Quantity:      d.Quantity,
			Nutrients:     dishNutrient,
			ImgUrls:       dishImage,
			Ingredients:   dishIngradient,
			Availability:  d.Availability,
			Warnings:      d.Warnings,
		}

		listCofeDish = append(listCofeDish, cofeDish)
	}

	return listCofeDish, nil
}

func (dish *Dishs) GetListDishDb(ctx context.Context) ([]DishDb, error) {
	rows, err := dish.PG.QueryContext(ctx,
		`SELECT id, created_at, updated_at, category, creator_gu_guid, title, description, weight, volume,
				price, quantity, availability, warnings, is_run
	            FROM dish where is_run = true`)
	if err != nil {
		log.Println(err)
		return []DishDb{}, err
	}

	defer rows.Close()
	d := DishDb{}
	dd := []DishDb{}



	for rows.Next() {
		err := rows.Scan(&d.ID, &d.CreatedAt, &d.UpdatedAt, &d.CategoryId, &d.CreatorGuGuid, &d.Title,
			&d.Description, &d.Weight, &d.Volume, &d.Price, &d.Quantity, &d.Availability, &d.Warnings, &d.IsRun)
		if err != nil {
			log.Println(err)
			continue
		}
		// d = DishDb{
		// 	ID:            d.ID,
		// 	CreatedAt:     d.CreatedAt,
		// 	UpdatedAt:     d.UpdatedAt,
		// 	CategoryId:      d.CategoryId,
		// 	CreatorGuGuid: d.CreatorGuGuid,
		// 	Title:         d.Title,
		// 	Description:   d.Description,
		// 	Weight:        d.Weight,
		// 	Volume:        d.Volume,
		// 	Price:         d.Price,
		// 	Quantity:      d.Quantity,
		// 	Availability:  d.Availability,
		// 	Warnings:      d.Warnings,
		// 	IsRun:         d.IsRun,
		// }

		dd = append(dd, d)

	}

	return dd, nil
}

func (dish *Dishs) GetDishCategory(ctx context.Context, id uuid.UUID) (string, error) {
	row := dish.PG.QueryRowContext(ctx, "select title from categories where id = $1", id)

	var category string
	err := row.Scan(&category)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return category, nil
}

func (dish *Dishs) GetDishNutrient(ctx context.Context, idProd uuid.UUID) (cofe_services.FoodNutrients, error) {

	proteins := cofe_services.Nutrient{}
	fats := cofe_services.Nutrient{}
	carbohydrates := cofe_services.Nutrient{}
	foodNutrients := cofe_services.FoodNutrients{
		Proteins:      proteins,
		Fats:          fats,
		Carbohydrates: carbohydrates,
		Energy:        0,
	}
	

	rows, err := dish.PG.QueryContext(ctx, `select nutrient, percentage, "value", energy from foodnutrients where prod_id = $1 order by code`, idProd)
	if err != nil {
		log.Println(err)
		return foodNutrients, err
	}
	defer rows.Close()
	var nutrient string
	var percentage float64
	var value int32
	var energy float64
	for rows.Next() {

		err := rows.Scan(&nutrient, &percentage, &value, &energy)
		if err != nil {
			log.Println(err)
			continue
		}
		if nutrient == "proteins" {
			proteins = cofe_services.Nutrient{
				Percentage: percentage,
				Value:      value,
			}

		}
		if nutrient == "fats" {
			fats = cofe_services.Nutrient{
				Percentage: percentage,
				Value:      value,
			}
		}
		if nutrient == "carbohydrates" {
			carbohydrates = cofe_services.Nutrient{
				Percentage: percentage,
				Value:      value,
			}
		}

	}
	foodNutrients = cofe_services.FoodNutrients{
		Proteins:      proteins,
		Fats:          fats,
		Carbohydrates: carbohydrates,
		Energy:        energy,
	}

	return foodNutrients, nil
}

func (dish *Dishs) GetDishImageUrls(ctx context.Context, idProd uuid.UUID) ([]string, error) {
	s := ""
	var ss []string
	//ss = append(ss, s)
	rows, err := dish.PG.QueryContext(ctx, "select file_name from images where id_prod = $1", idProd)
	if err != nil {
		log.Println(err)
		return ss, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&s)
		if err != nil {
			log.Println(err)
			continue
		}
		ss = append(ss, s)
	}

	return ss, nil
}

func (dish *Dishs) GetDishIngredient(ctx context.Context, id uuid.UUID) ([]cofe_services.Ingredient, error) {
	listIngradient := []cofe_services.Ingredient{}
	//ingradient := cofe_services.Ingredient{}
	//listIngradient = append(listIngradient, ingradient)
	return listIngradient, nil
}
