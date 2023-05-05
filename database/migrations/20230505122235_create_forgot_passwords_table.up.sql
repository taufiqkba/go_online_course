CREATE TABLE forgot_passwords (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    valid BOOLEAN NOT NULL DEFAULT 1,
    expired_at TIMESTAMP NULL,
    code INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id),
    INDEX idx_forgot_passwords_users_id(id)
    CONSTRAINT FK_forgot_passwords_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL
)