CREATE TABLE orders (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    discount_id INT NOT NULL,
    checkout_link VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    total_price INT NOT NULL,
    external_id VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_by TIMESTAMP NULL,
    updated_by TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id)
)