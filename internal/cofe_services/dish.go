package cofe_services

import (
	"context"
	//"database/sql"
//	"log"

	"github.com/google/uuid"
)

/**
 * @typedef {Object} Nutrient - питательное вещество
 * @property {number} percentage - относительное содержание к весу в %
 * @property {number} value - абсолютный вес в одной порции в граммах
 */

/**
 * @typedef {Object} FoodNutrients - пищевая ценность
 * @property {Nutrient} proteins - белки
 * @property {Nutrient} fats - жиры
 * @property {Nutrient} carbohydrates - углеводы
 * @property {number} energy - энергетическая ценность одной порции в килокалориях
 */


type Nutrient struct {
	Percentage float64	`json:"percentage"`
	Value      int32	`json:"value"`
}


type FoodNutrients struct {
	Proteins      Nutrient	`json:"proteins"`
	Fats          Nutrient	`json:"fats"`
	Carbohydrates Nutrient	`json:"carbohydrates"`
	Energy float64			`json:"energy"`
}


/**
 * @typedef {(0|1|2)} Availability - статус наличия: нет, мало, достаточно
 */

/**
 * @typedef {string} IngredientGuid - уникальный идентификатор ингредиента
 */

/**
 * @typedef {string} IngredientCategory - категрия ингредиента: 'специи'|'крупы'|...
 */

/**
 * @typedef {Object} Ingredient - публичное представление ингредиента
 * @property {IngredientCategory} category - категория ингредиента
 * @property {IngredientGuid} guGuid - идентификатор ингредиента (Global Unchanged Guidentifier)
 * @property {string} title - наименование ингредиента
 * @property {string} description - подробное описание ингредиента
 * @property {?number} weight - вес одной порции в граммах
 * @property {?number} volume - объем одной порции в граммах
 * @property {number} price - цена одной порции в рублях
 * @property {number} quantity - количество порций ингредиента, в штуках относительно weight или volume
 * @property {FoodNutrients} nutrients - пищевая ценность одной порции ингредиента
 * @property {string[]} imgUrls - ссылки на изображения
 * @property {string} brand - производитель ингредиента
 * @property {string} origin - страна происхождения ингредиента
 * @property {Availability} availability - наличие ингредиента
 * @property {string} warnings - предупреждения о свойствах продукта (для аллергиков, диабетиков?)
 */


 

type Ingredient struct {
	Category string			`json:"Category"`
	ID uuid.UUID			`json:"guGuid"`             
	Title string			`json:"title"` 
	Description string		`json:"description"` 
	Weight float64			`json:"weight"`
	Volume float64			`json:"volume"`
	Price float64			`json:"price"`
    Quantity float64		`json:"quantity"`
    Nutrients FoodNutrients	`json:"nutrients"`
	ImgUrls []string		`json:"imgUrls"`  
	Brand string			`json:"brand"`
	Origin string			`json:"origin"`
	Availability string		`json:"availability"`
	Warnings string			`json:"warnings"`
}


/**
 * @typedef {string} DishCategory - категрия блюда: 'закуски'|'сэндвичи'|'салаты'|'напитки'...
 */

/**
 * @typedef {string} DishGuid - уникальный идентификатор блюда
 */
/**
 * @typedef {Object} Dish - блюдо
 * @property {DishCategory} category - категория блюда
 * @property {DishGuid} dishGuid - идентификатор блюда
 * @property {?UserGuid} creatorGuGuid - автор блюда
 * @property {string} title - наименование блюда
 * @property {string} description - подробное описание блюда
 * @property {?number} weight - вес одной порции в граммах
 * @property {?number} volume - объем одной порции в граммах
 * @property {number} price - цена одной порции в рублях
 * @property {number} quantity - количество порций блюда, в штуках
 * @property {FoodNutrients} nutrients - пищевая ценность одной порции блюда
 * @property {string[]} imgUrls - ссылки на изображения
 * @property {Ingredient[]} ingredients - ингредиенты блюда
 * @property {Availability} availability - доступность блюда
 * @property {string} warnings - предупреждения о свойствах продукта (для аллергиков, диабетиков?)
 */






type Dish struct {
	Category      string		`json:"category"`
	DishGuid      uuid.UUID		`json:"dishGuid"`
	CreatorGuGuid uuid.UUID		`json:"creatorGuGuid"`
	Title         string		`json:"title"`
	Description   string		`json:"description"`
	Weight        float64		`json:"weight"`
	Volume        float64		`json:"volume"`
	Price         float64		`json:"price"`
	Quantity      int32			`json:"quantity"`
	Nutrients     FoodNutrients	`json:"nutrients"`
	ImgUrls       []string		`json:"imgUrls"`
	Ingredients   []Ingredient	`json:"ingredients"`
	Availability  int32			`json:"availability"`
	Warnings      string		`json:"warnings"`
}

type DishStore interface {
	GetListDish(ctx context.Context) ([]Dish, error)
}

type Dishs struct {
	Store DishStore
}

func NewDishs(dishStore DishStore) *Dishs {
	return &Dishs{
		Store: dishStore,
	}
}

func (d *Dishs) GetListDIsh(ctx context.Context) ([]Dish, error) {
	dish, err := d.Store.GetListDish(ctx)
	if err != nil {
		return nil, err
	}


	return dish, nil
}
