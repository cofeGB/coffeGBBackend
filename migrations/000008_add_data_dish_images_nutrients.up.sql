DO $$
DECLARE DishId uuid;
BEGIN  
          
INSERT INTO dish("id","created_at", "updated_at", "category", "creator_gu_guid", 
				"title", "description", "weight", "volume", "price", "quantity", "availability", "warnings", "is_run")
			VALUES(uuid_generate_v4(),now(),now(),(select id from categories where query = 'starters'),'a7f4485d-ecb7-4d1d-817d-3f5036fae88a',
			        'Закуска «Уральский рулет»',  'Интересный мясной рулет с добавлением сырокопченого окорока и яичного омлета с болгарским перцем и кунжутом.',
					100,220,  84, 1, 0,'no warning', true) RETURNING id INTO DishId;

INSERT INTO images("id", "id_prod", "created_at", "updated_at", "file_name")
            VALUES(uuid_generate_v4(), DishId, now(), now(),'rulet.jpg' );

INSERT INTO foodnutrients("id", "prod_id", "code", "nutrient", "percentage", "value", "energy")
          VALUES (uuid_generate_v4(), DishId, 1, 'proteins', 15.8, 0, 515.2),
            (uuid_generate_v4(), DishId, 2, 'fats', 18.8, 0, 515.2), 
            (uuid_generate_v4(), DishId, 3, 'carbohydrates', 0.5, 0, 515.2);


--2
INSERT INTO dish("id","created_at", "updated_at", "category", "creator_gu_guid", 
				"title", "description", "weight", "volume", "price", "quantity", "availability", "warnings", "is_run")
			VALUES(uuid_generate_v4(),now(),now(),(select id from categories where query = 'sandwich'),'a7f4485d-ecb7-4d1d-817d-3f5036fae88a',
			        'Сэндвич с курицей и беконом',  'Сэндвич с курицей, беконом, зеленым салатом, помидором, огурцом, красным луком, горчицей и сыром.',
					100,250,  78, 1, 1,'no warning', true) RETURNING id INTO DishId;

INSERT INTO images("id", "id_prod", "created_at", "updated_at", "file_name")
            VALUES(uuid_generate_v4(), DishId,now(),now(),'sandwich.jpg' );

INSERT INTO foodnutrients("id", "prod_id", "code", "nutrient", "percentage", "value", "energy")
          VALUES (uuid_generate_v4(), DishId, 1, 'proteins', 0, 36, 537),
            (uuid_generate_v4(), DishId, 2, 'fats', 0, 23, 537), 
           (uuid_generate_v4(), DishId, 3, 'carbohydrates', 0, 44, 537);



--3
INSERT INTO dish("id","created_at", "updated_at", "category", "creator_gu_guid", 
				"title", "description", "weight", "volume", "price", "quantity", "availability", "warnings", "is_run")
			VALUES(uuid_generate_v4(),now(),now(),(select id from categories where query = 'salad'),'a7f4485d-ecb7-4d1d-817d-3f5036fae88a',
			        'Греческий салат',  'Салат из помидоров, сладкого перца, огурцов, феты, шалота и маслин, заправленный оливковым маслом с солью, чёрным перцем, орегано.',
					100,250,  95, 1, 2,'no warning', true) RETURNING id INTO DishId;

INSERT INTO images("id", "id_prod", "created_at", "updated_at", "file_name")
            VALUES(uuid_generate_v4(), DishId,now(),now(),'salad.jpg');

INSERT INTO foodnutrients("id", "prod_id", "code", "nutrient", "percentage", "value", "energy")
          VALUES (uuid_generate_v4(), DishId, 1, 'proteins', 0, 9.5, 482.7),
           (uuid_generate_v4(), DishId, 2, 'fats', 0, 43.7, 482.7), 
            (uuid_generate_v4(), DishId, 3, 'carbohydrates', 0, 12.4, 482.7);


--4
INSERT INTO dish("id","created_at", "updated_at", "category", "creator_gu_guid", 
				"title", "description", "weight", "volume", "price", "quantity", "availability", "warnings", "is_run")
			VALUES(uuid_generate_v4(),now(),now(),(select id from categories where query = 'desserts'),'a7f4485d-ecb7-4d1d-817d-3f5036fae88a',
			        'Торт Прага',  'Состоит из трёх бисквитных коржей с двумя слоями крема «Пражский». Верхняя и боковые поверхности покрыты повидлом и заглазированы помадкой (либо боковые поверхности отделаны кремом и бисквитной крошкой). Сверху украшен рисунком из крема',
					100,200,  83, 1, 1,'no warning', true) RETURNING id INTO DishId;

INSERT INTO images("id", "id_prod", "created_at", "updated_at", "file_name")
            VALUES(uuid_generate_v4(), DishId,now(),now(),'dessert.jpg');

INSERT INTO foodnutrients("id", "prod_id", "code", "nutrient", "percentage", "value", "energy")
          VALUES (uuid_generate_v4(), DishId, 1, 'proteins', 0, 12.7, 720.7),
           (uuid_generate_v4(), DishId, 2, 'fats', 0, 33.8, 720.7), 
             (uuid_generate_v4(), DishId, 3, 'carbohydrates', 0, 85.1, 720.7);


--5
INSERT INTO dish("id","created_at", "updated_at", "category", "creator_gu_guid", 
				"title", "description", "weight", "volume", "price", "quantity", "availability", "warnings", "is_run")
			VALUES(uuid_generate_v4(),now(),now(),(select id from categories where query = 'coffee'),'a7f4485d-ecb7-4d1d-817d-3f5036fae88a',
			        'Капучино',  'Напиток на основе эспрессо со вспененным молоком. Вкус капучино идеально сбалансирован. Благодаря свежесваренному эспрессо он насыщенно кофейный, с ярко выраженным ароматом, а из-за молока - мягкий, с натуральной сладостью.',
					100, 150, 83, 1, 1,'no warning', true) RETURNING id INTO DishId;

INSERT INTO images("id", "id_prod", "created_at", "updated_at", "file_name")
            VALUES(uuid_generate_v4(), DishId,now(),now(),'coffee.jpeg');

INSERT INTO foodnutrients("id", "prod_id", "code", "nutrient", "percentage", "value", "energy")
          VALUES (uuid_generate_v4(), DishId, 1, 'proteins', 0, 3, 59),
           (uuid_generate_v4(), DishId, 2, 'fats', 0, 3, 59), 
             (uuid_generate_v4(), DishId, 3, 'carbohydrates', 0, 5, 59);


--6
INSERT INTO dish("id","created_at", "updated_at", "category", "creator_gu_guid", 
				"title", "description", "weight", "volume", "price", "quantity", "availability", "warnings", "is_run")
			VALUES(uuid_generate_v4(),now(),now(),(select id from categories where query = 'tea'),'a7f4485d-ecb7-4d1d-817d-3f5036fae88a',
			        'Зеленый чай с имбирем лимоном и медом',  'Ароматный, согревающий и полезный напиток в сезон холодов!',
					100, 150, 95, 1, 2,'no warning', true) RETURNING id INTO DishId;

INSERT INTO images("id", "id_prod", "created_at", "updated_at", "file_name")
            VALUES(uuid_generate_v4(), DishId,now(),now(),'tea.jpg');

INSERT INTO foodnutrients("id", "prod_id", "code", "nutrient", "percentage", "value", "energy")
          VALUES (uuid_generate_v4(), DishId, 1, 'proteins', 0, 0, 89),
            (uuid_generate_v4(), DishId, 2, 'fats', 0, 0, 89), 
             (uuid_generate_v4(), DishId, 3, 'carbohydrates', 0, 18, 89);




          
END $$;