CREATE TABLE categories (
                            id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                            name VARCHAR(50) NOT NULL,
                            created_at TIMESTAMP NOT NULL
)ENGINE=InnoDB;

CREATE TABLE tasks (
                       id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       category_id INT,
                       title VARCHAR(100) NOT NULL,
                       description TEXT,
                       priority ENUM('low', 'medium', 'high') NOT NULL,
                       status ENUM('pending', 'in_progress') NOT NULL,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
)ENGINE=InnoDB;

CREATE TABLE completed_tasks (
                                 id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                 category_id INT,
                                 title VARCHAR(100) NOT NULL,
                                 description TEXT,
                                 priority ENUM('low', 'medium', 'high') NOT NULL,
                                 status VARCHAR(10) NOT NULL,
                                 created_at TIMESTAMP NOT NULL,
                                 completed_at TIMESTAMP NOT NULL,
                                 FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
)ENGINE=InnoDB;