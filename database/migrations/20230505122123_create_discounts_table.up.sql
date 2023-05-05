CREATE TABLE discounts (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    remaining_quantity INT NOT NULL,
    type VARCHAR(255) NOT NULL,
    value INT NOT NULL,
    start_date TIMESTAMP NULL,
    end_date TIMESTAMP NULL,
    created_by TIMESTAMP NULL,
    updated_by TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id)
)