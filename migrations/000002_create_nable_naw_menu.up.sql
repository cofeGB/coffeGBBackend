CREATE  TABLE naw_menu ( 
	id                   uuid  NOT NULL  ,
	created_at           date DEFAULT CURRENT_DATE NOT NULL  ,
	updated_at           date DEFAULT CURRENT_DATE NOT NULL  ,
	title                varchar(50)  NOT NULL  ,
	"path"               varchar(256)  NOT NULL  ,
	item_order           integer  NOT NULL  ,
	is_run               boolean DEFAULT true NOT NULL  ,
	CONSTRAINT pk_naw_menu_id PRIMARY KEY ( id )
 );

CREATE INDEX idx_naw_menu_title ON naw_menu ( title );

COMMENT ON COLUMN naw_menu.id IS 'Идентификатор';

COMMENT ON COLUMN naw_menu.created_at IS 'Дата создания';

COMMENT ON COLUMN naw_menu.updated_at IS 'Дата изменения';

COMMENT ON COLUMN naw_menu.title IS 'название пункта навигационного меню, как его видит пользователь';

COMMENT ON COLUMN naw_menu."path" IS 'полный путь роутера ''/menu/sandwiches''|''/menu/salads''...';

COMMENT ON COLUMN naw_menu.item_order IS 'Порядок сортировки';

COMMENT ON COLUMN naw_menu.is_run IS 'Признак рабочей записи';




INSERT INTO naw_menu("id", "created_at", "updated_at","title", "path", "item_order","is_run")
values (uuid_generate_v4(), now(), now(),'Закуски','/menu/starters',0, true), 
        (uuid_generate_v4(), now(), now(),'Сендвичи','/menu/sandwich',1, true), 
        (uuid_generate_v4(), now(), now(),'Салаты',  '/menu/salad',  2, true), 
        (uuid_generate_v4(), now(), now(),'Десерты','/menu/desserts', 3, true) ,
        (uuid_generate_v4(), now(), now(),'Кофе','/menu/coffee', 4, true) ,                
        (uuid_generate_v4(), now(), now(),'Чай','/menu/tea', 5, true) 
;    



