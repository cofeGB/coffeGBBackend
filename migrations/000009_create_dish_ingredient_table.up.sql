CREATE TABLE dish_ingredient(
    id uuid NOT NULL,
    dish_id uuid NOT NULL,
    ingedient_id uuid NOT NULL
)

CREATE INDEX idx_dish_id ON ingredient ( dish_id );

COMMENT ON COLUMN dish_ingredient.id IS 'идентификатор';
COMMENT ON COLUMN dish_ingredient.dish_id IS 'идентификатор блюда';
COMMENT ON COLUMN dish_ingredient.ingedient_id IS 'идентификатор ингредиента';