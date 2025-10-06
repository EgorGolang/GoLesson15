CREATE TABLE IF NOT EXISTS userstest(
                                    id SERIAL PRIMARY KEY,
                                    name VARCHAR NOT NULL,
                                    email VARCHAR NOT NULL UNIQUE,
                                    age INT NOT NULL
);