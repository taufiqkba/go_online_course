CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    product_category_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    description TEXT NOT NULL,
    image VARCHAR(255) NOT NULL,
    video VARCHAR(255) NOT NULL,
    updated_by TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id)
)