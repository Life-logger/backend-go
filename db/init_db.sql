CREATE TABLE Users (
    user_id INT NOT NULL AUTO_INCREMENT, 
    user_name VARCHAR(10) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


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

-- -- INSERT INTO ADMIN (email,passcode,admin_type,brand_id) VALUES
-- -- ('ahnnn000@gmail.com','$2a$12$cqXwxUhyP5bkcdoz7Tge1.6DgDeN9h4mqfTqMycCdm01K8rsKwbVG','SUPER',NULL);


CREATE TABLE Categories (
    id INT NOT NULL AUTO_INCREMENT,
    user_email VARCHAR(40),
    color VARCHAR(7),
    title VARCHAR(20) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE Blocks (
    block_id INT AUTO_INCREMENT,
    start_time INT NOT NULL,
    end_time INT NOT NULL,
    duration INT NOT NULL,
    color VARCHAR(7) REFERENCES Categories(categories_id) ON DELETE CASCADE ON UPDATE NO ACTION,
    background_image_url VARCHAR(255),
    block_memo VARCHAR(255),
    block_pin BOOLEAN NOT NULL,
    PRIMARY KEY (block_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- CREATE TABLE Calendar (
--     date DATE PRIMARY KEY,
--     representative_block_id INT REFERENCES Blocks(block_id)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
