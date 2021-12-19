CREATE TABLE naw_menu(  
    id uuid NOT NULL primary key,
    created_at DATE,                                                
    updated_at DATE,
    title varchar(50) NOT null UNIQUE ,
    path varchar(255) NOT NULL UNIQUE,
    item_order integer not null ,
    is_run boolean default true
);

INSERT INTO naw_menu("id", "created_at", "updated_at","title", "path", "item_order","is_run")
values (uuid_generate_v4(), now(), now(),'Закуски','/menu/starters',0, true), 
        (uuid_generate_v4(), now(), now(),'Сендвичи','/menu/sandwich',1, true), 
        (uuid_generate_v4(), now(), now(),'Салаты',  '/menu/salad',  2, true), 
        (uuid_generate_v4(), now(), now(),'Десерты','/menu/desserts', 3, true) ,
        (uuid_generate_v4(), now(), now(),'Кофе','/menu/coffee', 4, true) ,                
        (uuid_generate_v4(), now(), now(),'Чай','/menu/tea', 5, true) 
;    