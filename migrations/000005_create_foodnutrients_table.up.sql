--drop table if EXISTS foodnutrients;
CREATE  TABLE foodnutrients ( 
	id                   uuid  NOT NULL  ,
	prod_id              uuid  NOT NULL  ,
	code                 integer  NOT NULL  ,
	nutrient             varchar(50)  NOT NULL  ,
	percentage           numeric(5,2)  NOT NULL  ,
	"value"              integer  NOT NULL  ,
	energy               numeric(10,2)  NOT NULL  ,
	CONSTRAINT "Pk_foodnutrients" PRIMARY KEY ( id )
 );

CREATE INDEX idx_foodnutrients_id_prod ON foodnutrients ( prod_id );

CREATE INDEX idx_foodnutrients_nutrient ON foodnutrients ( nutrient );

COMMENT ON COLUMN foodnutrients.id IS 'Идентификатор';

COMMENT ON COLUMN foodnutrients.prod_id IS 'Идентификатор продукта/блюда';

COMMENT ON COLUMN foodnutrients.code IS 'цифровой код';

COMMENT ON COLUMN foodnutrients.nutrient IS 'Название параметра';

COMMENT ON COLUMN foodnutrients.percentage IS 'относительное содержание к весу в %';

COMMENT ON COLUMN foodnutrients."value" IS 'абсолютный вес в одной порции в граммах';

COMMENT ON COLUMN foodnutrients.energy IS 'энергетическая ценность одной порции в килокалориях';
