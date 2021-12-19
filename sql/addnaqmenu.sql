
INSERT INTO naw_menu("id", "created_at", "updated_at","title", "path", "item_order","is_run")
values (uuid_generate_v4(), now(), now(),'Закуски','/menu/starters',0, true), 
        (uuid_generate_v4(), now(), now(),'Сендвичи','/menu/sandwich',1, true), 
        (uuid_generate_v4(), now(), now(),'Салаты',  '/menu/salad',  2, true), 
        (uuid_generate_v4(), now(), now(),'Десерты','/menu/desserts', 3, true) ,
        (uuid_generate_v4(), now(), now(),'Кофе','/menu/coffee', 4, true) ,                
        (uuid_generate_v4(), now(), now(),'Чай','/menu/tea', 5, true) 
;               
                        
                        
                        
                        
 