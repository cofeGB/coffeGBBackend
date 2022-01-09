/**
 * Форматы объектов для сериализации (для API):
 * a string
 * a number
 * an object (JSON object)
 * an array
 * a boolean
 * null
 */

/**
 * @typedef {string} UserGuid - уникальный идентификатор пользователя
 */

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

/**
 * @typedef {('pending'|'cooking'|'ready'|'cancelled'|'paused'|'closed')} OrderStatus - статус заказа
 */

/**
 * @typedef {string} OrderGuid - уникальный идентификатор заказа
 */

/**
 * @typedef {Object} Order - публичный объект заказа
 * @property {string} creationDate - дата оформления заказа, формат UTC 'Mon, 13 Dec 2021 19:25:53 GMT'
 * @property {string} processingStartDate - дата начала работы с заказом, формат UTC 'Mon, 13 Dec 2021 19:25:53 GMT'
 * @property {string} processingEndDate - дата окончания работы с заказом, формат UTC 'Mon, 13 Dec 2021 19:25:53 GMT'
 * @property {OrderGuid} orderGuid - уникальный идентификатор заказа
 * @property {UserGuid} userGuid - уникальный идентификатор заказа
 * @property {Dish[]} dishes - блюда в заказе
 * @property {OrderStatus} status - статус заказа
 */

/**
 * @typedef {Object} PrivOrder - объект заказа для заведения
 * @property {Order} order - публичный объект заказа
 * ? что нужно заведению ? Исолнитель? Статус склада? Статус готовности отдельных блюд?
 */

/**
 * @typedef {Object} NavMenuItem - элемент навигационного меню приложения
 * @property {string} title - название пункта навигационного меню, как его видит пользователь
 * @property {string} path - полный путь роутера '/menu/sandwiches'|'/menu/salads'...
 * @property {number} itemOrder - порядковый номер данного элемента в меню, целые числа 0,1,2...
 */

/**
 * @typedef {NavMenuItem[]} NavMenuItemList - элемент навигационного меню приложения
 */

/**
 * @typedef {Object} Product - Продукт (для приготовления или хронения на складе) универсальные данные их можно юзать и для подсчета параметров и для показа остатка и для показа в описание блюд
 * @property {string} title - название пункта навигационного меню, как его видит пользователь
 * @property {string} productGuid - Guid продукта
 * @property {?string} supplier - Поставщик продукта
 * @property {?string} storageTime - Время хранения продукта
 * @property {?string} bestBeforeDate - Срок годности продукта
 * @property {?string} warnings - предупреждения о свойствах продукта (для аллергиков, диабетиков?)
 * @property {?number} weight - Вес продукта
 * @property {?number} weightPack - Вес продукта в упаковке
 * @property {?number} price - Цена продукта
 * @property {?number} pricePack - Цена продукта в упаковке
 * @property {?number} quantity - Количество продукта
 * @property {?number} quantityPack - Количество упаковок продукта
 * @property {FoodNutrients} foodNutrients - КБЖУ продукта
 * @property {Nutrient} nutrient - Питательное вещество продукта
 */