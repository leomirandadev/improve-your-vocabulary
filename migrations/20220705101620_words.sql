-- +goose Up
CREATE TABLE words(
    id INT AUTO_INCREMENT NOT NULL,
    word VARCHAR(200) NOT NULL,
    user_id INT NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE words;