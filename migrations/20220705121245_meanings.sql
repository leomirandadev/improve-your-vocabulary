-- +goose Up
CREATE TABLE meanings(
    id INT AUTO_INCREMENT NOT NULL,
    meaning VARCHAR(200) NOT NULL,
    word_id INT NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id)
    FOREIGN KEY (word_id) REFERENCES words(id)
);

-- +goose Down
DROP TABLE meanings;