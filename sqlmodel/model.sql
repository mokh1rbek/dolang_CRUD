

CREATE TABLE films (
    film_id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(24),
    description TEXT,
    release_year DATE,
    rental_duration NUMERIC,
    
)






CREATE TABLE actor (
    actor_id UUID DEFAULT gen_random_uuid(),
    first_name VARCHAR(24),
    last_name VARCHAR(24),
    last_update TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE category (
    category_id UUID DEFAULT gen_random_uuid(),
    name VARCHAR(255),
    last_update TIMESTAMP NOT NULL DEFAULT NOW()
);