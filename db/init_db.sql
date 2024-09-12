CREATE TABLE Users (
    user_id INT NOT NULL AUTO_INCREMENT, 
    user_name VARCHAR(10) NOT NULL,
    user_email VARCHAR(40) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


CREATE TABLE ADMIN (
    id INT NOT NULL AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    passcode varchar(255) NOT NULL,
    admin_type varchar(255) NOT NULL,
    brand_id INT,
    PRIMARY KEY (id),
    UNIQUE KEY ADMIN_unique (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- INSERT INTO ADMIN (email,passcode,admin_type,brand_id) VALUES
-- ('ahnnn000@gmail.com','$2a$12$cqXwxUhyP5bkcdoz7Tge1.6DgDeN9h4mqfTqMycCdm01K8rsKwbVG','SUPER',NULL);


-- 카테고리 별 누적 시간 미포함
CREATE TABLE Categories (
    id INT NOT NULL AUTO_INCREMENT,
    category_user_id INT,
    color VARCHAR(7),
    title VARCHAR(20) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (category_user_id) REFERENCES Users(user_id) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE Blocks (
    block_id INT NOT NULL AUTO_INCREMENT,
    category_id INT,
    start_time INT NOT NULL,
    end_time INT NOT NULL,
    duration INT NOT NULL,
    color VARCHAR(7),
    background_image_url VARCHAR(50),
    block_memo VARCHAR(100),
    PRIMARY KEY(block_id),
    FOREIGN KEY (category_id) REFERENCES Categories(id) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE Calendar (
    date_id INT NOT NULL AUTO_INCREMENT,
    representative_block_id INT,
    PRIMARY KEY(date_id),
    FOREIGN KEY (representative_block_id) REFERENCES Blocks(block_id) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;