CREATE TABLE oauth_refresh_tokens  (
    id INT NOT NULL AUTO_INCREMENT,
    oauth_access_token_id INT NULL,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    expired_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id),
    UNIQUE KEY oauth_refresh_tokens_token_unique(token),
    INDEX idx_oauth_refresh_tokens_oauth_access_token_id(oauth_access_token_id),
    INDEX idx_oauth_refresh_tokens_token(token),
    CONSTRAINT FK_oauth_refresh_tokens_oauth_access_token_id FOREIGN KEY (oauth_access_token_id) REFERENCES oauth_access_tokens(id) ON DELETE SET NULL
)