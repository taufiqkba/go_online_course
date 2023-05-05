CREATE TABLE order_details (
    id INT NOT NULL AUTO_INCREMENT,
    product_id INT NOT NULL,
    order_id INT NOT NULL,
    price INT NOT NULL,
    created_by TIMESTAMP NULL,
    updated_by TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id)
)