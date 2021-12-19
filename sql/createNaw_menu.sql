CREATE TABLE naw_menu(  
    id uuid NOT NULL primary key,
    created_at DATE,                                                
    updated_at DATE,
    title varchar(50) NOT null UNIQUE ,
    path varchar(255) NOT NULL UNIQUE,
    item_order integer not null ,
    is_run boolean default true
);
