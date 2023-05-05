CREATE TABLE oauth_access_tokens  (
    id INT NOT NULL AUTO_INCREMENT,
    oauth_client_id INT NOT NULL,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    scope VARCHAR(255) NOT NULL,
    expired_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id)
)