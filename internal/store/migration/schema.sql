CREATE TABLE IF NOT EXISTS category (
    id  SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS rule (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    CONSTRAINT fk_category 
        FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE IF NOT EXISTS team (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    school TEXT NOT NULL,
    grade TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    CONSTRAINT fk_category
        FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE IF NOT EXISTS member (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    is_leader BOOLEAN NOT NULL DEFAULT false,
    team_id INTEGER NOT NULL,
    CONSTRAINT fk_team
        FOREIGN KEY (team_id) REFERENCES team(id)
);

CREATE TABLE IF NOT EXISTS robot (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    weight_cm DECIMAL(5,2),
    width_cm DECIMAL(5,2),
    height_cm TEXT,
    length_cm DECIMAL(5,2),
    is_valid BOOLEAN NOT NULL DEFAULT false,
    invalid_reason TEXT,
    autonomous TEXT,
    power_button TEXT,
    internal_power TEXT,
    status TEXT NOT NULL DEFAULT 'pending'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    CONSTRAINT fk_team
        FOREIGN KEY (team_id) REFERENCEs team(id)
    CONSTRAINT fk_category
    FOREIGN KEY (category_id) REFERENCES category(id)
);
