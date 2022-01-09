CREATE  TABLE ingredient ( 
	id                   uuid  NOT NULL  ,
	category_id          uuid  NOT NULL  ,
	title                varchar(256)  NOT NULL  ,
	description          text DEFAULT 'no comment' NOT NULL  ,
	weight               numeric(10,2) DEFAULT 0 NOT NULL  ,
	"volume "            numeric(10,2) DEFAULT 0 NOT NULL  ,
	price                numeric(10,2) DEFAULT 0 NOT NULL  ,
	quantity             numeric(10,2) DEFAULT 0 NOT NULL  ,
	brand                varchar(256) DEFAULT 'no brand' NOT NULL  ,
	"origin "            varchar(256) DEFAULT 'no origin' NOT NULL  ,
	availability         integer DEFAULT 0 NOT NULL  ,
	warnings             text DEFAULT 'no warnings' NOT NULL  ,
	CONSTRAINT pk_ingredient_id PRIMARY KEY ( id )
 );

CREATE INDEX idx_ingredient_id_prod ON ingredient ( category_id );

COMMENT ON COLUMN ingredient.category_id IS 'категория ингредиента \n @typedef {string} IngredientCategory - категрия ингредиента: ''специи''|''крупы''|...';

COMMENT ON COLUMN ingredient.title IS '- наименование ингредиента';

COMMENT ON COLUMN ingredient.description IS '- подробное описание ингредиента';

COMMENT ON COLUMN ingredient.weight IS '- вес одной порции в граммах';

COMMENT ON COLUMN ingredient."volume " IS '- объем одной порции в граммах';

COMMENT ON COLUMN ingredient.price IS '- цена одной порции в рублях';

COMMENT ON COLUMN ingredient.quantity IS '- количество порций ингредиента, в штуках относительно weight или volume';

COMMENT ON COLUMN ingredient.brand IS '- производитель ингредиента';

COMMENT ON COLUMN ingredient."origin " IS '- страна происхождения ингредиента';

COMMENT ON COLUMN ingredient.availability IS '- наличие ингредиента\n @typedef {(0|1|2)} Availability - статус наличия: нет, мало, достаточно';

COMMENT ON COLUMN ingredient.warnings IS '- предупреждения о свойствах продукта (для аллергиков, диабетиков?)';

