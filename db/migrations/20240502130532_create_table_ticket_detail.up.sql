CREATE TABLE ticket_detail (
    ticket_detail_id INT AUTO_INCREMENT PRIMARY KEY,
    type varchar(50),
    price int,
    continent_name varchar(100),
    stock_ticket int,
    stock int,
    stock_ordered int,
    country_name VARCHAR(100),
    country_city VARCHAR(100),
    country_place VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
