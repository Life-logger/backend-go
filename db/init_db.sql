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
('annie.chang@gmail.com','$2a$12$cqXwxUhyP5bkcdoz7Tge1.6DgDeN9h4mqfTqMycCdm01K8rsKwbVG','SUPER',NULL),
('tom.cruise@gmail.com','$2a$12$cqXwxUhyP5bkcdoz7Tge1.6DgDeN9h4mqfTqMycCdm01K8rsKwbVG','BRAND',2),
('brad.pitt@gmail.com','$2a$12$cqXwxUhyP5bkcdoz7Tge1.6DgDeN9h4mqfTqMycCdm01K8rsKwbVG','BRAND',1),
('ob_manager@gmail.com','$2a$12$cqXwxUhyP5bkcdoz7Tge1.6DgDeN9h4mqfTqMycCdm01K8rsKwbVG','BRAND',3);
