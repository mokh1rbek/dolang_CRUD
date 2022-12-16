
CREATE TABLE film (
    film_id UUID NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR,
    release_year DATE NOT NULL,
    duration INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE category (
    category_id UUID NOT NULL,
    name VARCHAR(25) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
