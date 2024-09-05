-- dep.ADMIN definition
CREATE TABLE ADMIN (
    id INT NOT NULL AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    passcode varchar(255) NOT NULL,
    admin_type varchar(255) NOT NULL,
    brand_id INT,
    PRIMARY KEY (id),
    UNIQUE KEY ADMIN_unique (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO ADMIN (email,passcode,admin_type,brand_id) VALUES
('ahnnn000@gmail.com','$2a$12$cqXwxUhyP5bkcdoz7Tge1.6DgDeN9h4mqfTqMycCdm01K8rsKwbVG','SUPER',NULL);

CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY, --SERIAL: 자동으로 증가하는 정수 타입
    --password_hash VARCHAR(255) NOT NULL,
    user_name VARCHAR(10) NOT NULL,
    --preferences JSONB
);

CREATE TABLE Categories (
    categories_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(user_id),
    color VARCHAR(7),
    --start_time TIMESTAMP NOT NULL,
    --end_time TIMESTAMP NOT NULL,
    categories_title VARCHAR(20) NOT NULL,
    categories_total_time INT NOT NULL DEFAULT 0,
);

CREATE TABLE Blocks (
    block_id SERIAL PRIMARY KEY,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    duration time.Duration NOT NULL,
    color VARCHAR(7) REFERENCES Categories(categories_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    background_image_url TEXT,
    block_memo TEXT,
    block_pin BOOLEAN NOT NULL,
);


-- CREATE TABLE CategoriesBlocks (
--     categories_id INT REFERENCES categories(categories_id),
--     block_id INT REFERENCES Blocks(block_id),
--     PRIMARY KEY(categories_id, block_id)
-- );

/* CREATE TABLE Tags (
    tag_id SERIAL PRIMARY KEY,
    tag_name VARCHAR(100) UNIQUE NOT NULL
);--? */

/* CREATE TABLE CategoriesTags (
    categories_id INT REFERENCES categories(categories_id),
    tag_id INT REFERENCES Tags(tag_id),
    PRIMARY KEY(categories_id, tag_id)
); */

CREATE TABLE Calendar (
    date DATE PRIMARY KEY,
    representative_block_id INT REFERENCES Blocks(block_id),
);

-- CREATE TABLE CalendarBlocks (
--     date DATE REFERENCES Calendar(date),
--     block_id INT REFERENCES Blocks(block_id),
--     PRIMARY KEY(date, block_id)
-- );