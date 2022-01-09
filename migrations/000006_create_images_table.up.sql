CREATE  TABLE images ( 
	id                   uuid  NOT NULL  ,
	id_prod              uuid  NOT NULL  ,
	created_at            date  NOT NULL  ,
	updated_at            date DEFAULT CURRENT_DATE NOT NULL  ,
	file_name            varchar(256)  NOT NULL  ,
	is_run               boolean DEFAULT true NOT NULL  ,
	CONSTRAINT pk_images_id PRIMARY KEY ( id )
 );

CREATE INDEX idx_images_id_prod ON images ( id_prod );

COMMENT ON TABLE images IS 'Таблица изображений';

COMMENT ON COLUMN images.id IS 'идентификатор';

COMMENT ON COLUMN images.id_prod IS 'Идентификатор продукта';

COMMENT ON COLUMN images.created_at IS 'Дата создания';

COMMENT ON COLUMN images.updated_at IS 'Дата изменения';

COMMENT ON COLUMN images.file_name IS 'Путь к файлу';
