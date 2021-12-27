CREATE TABLE categories(  
    id uuid NOT NULL primary key,
    title varchar NOT null UNIQUE,
    query varchar(255) NOT NULL UNIQUE,
    icon varchar,
    item_order integer not null
);

INSERT INTO categories
values 
    (uuid_generate_v4(), 'Закуски','starters', 'starters',0),
    (uuid_generate_v4(), 'Сендвичи','sandwich', 'sandwich',1),
    (uuid_generate_v4(), 'Салаты',  'salad', 'salad',  2),
    (uuid_generate_v4(), 'Десерты','desserts', 'desserts', 3),
    (uuid_generate_v4(), 'Кофе','coffee', 'coffee', 4),
    (uuid_generate_v4(), 'Чай','tea', 'tea', 5);