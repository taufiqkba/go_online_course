CREATE TABLE oauth_refresh_tokens  (
    id INT NOT NULL AUTO_INCREMENT,
    oauth_access_token_id INT NOT NULL,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    expired_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id)
)