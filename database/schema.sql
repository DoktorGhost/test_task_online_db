CREATE TABLE IF NOT EXISTS Client (
    client_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Stand_name (
    stand_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Stands (
    id SERIAL PRIMARY KEY,
    stand_id INT,
    product_id INT, 
    parent BOOLEAN,
    FOREIGN KEY (product_id) REFERENCES Products(product_id),
    FOREIGN KEY (stand_id) REFERENCES Stand_name(stand_id)
);


CREATE TABLE IF NOT EXISTS Orders (
    order_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    date DATE NOT NULL,
    FOREIGN KEY (client_id) REFERENCES Client(client_id)
);

CREATE TABLE IF NOT EXISTS Order_details (
    order_details_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    count INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (product_id) REFERENCES Products(product_id)
);


-- НАПОЛНЕНИЕ ТАБЛИЦ

INSERT INTO Client (client_id, name) VALUES (1, 'ВАЛЕРА');
INSERT INTO Client (client_id, name) VALUES (2, 'СЕРГЕЙ');

INSERT INTO Products (product_id, name) VALUES (1, 'Ноутбук');
INSERT INTO Products (product_id, name) VALUES (2, 'Монитор');
INSERT INTO Products (product_id, name) VALUES (3, 'Телефон');
INSERT INTO Products (product_id, name) VALUES (4, 'Системный блок');
INSERT INTO Products (product_id, name) VALUES (5, 'Часы');
INSERT INTO Products (product_id, name) VALUES (6, 'Микрофон');

INSERT INTO Orders (order_id, client_id, date) VALUES (10, 1, '2023-02-01');
INSERT INTO Orders (order_id, client_id, date) VALUES (11, 2, '2023-03-15');
INSERT INTO Orders (order_id, client_id, date) VALUES (14, 1, '2023-02-04');
INSERT INTO Orders (order_id, client_id, date) VALUES (15, 2, '2023-06-01');

INSERT INTO Stand_name (stand_id, name) VALUES (1, 'A');
INSERT INTO Stand_name (stand_id, name) VALUES (2, 'Б');
INSERT INTO Stand_name (stand_id, name) VALUES (3, 'В');
INSERT INTO Stand_name (stand_id, name) VALUES (4, 'З');
INSERT INTO Stand_name (stand_id, name) VALUES (5, 'Ж');

INSERT INTO Stands (stand_id, product_id, parent) VALUES (1, 1, TRUE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (1, 2, TRUE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (2, 3, TRUE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (3, 3, FALSE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (4, 3, FALSE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (5, 4, TRUE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (5, 5, TRUE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (5, 6, TRUE);
INSERT INTO Stands (stand_id, product_id, parent) VALUES (1, 5, FALSE);

INSERT INTO Order_details (order_id, product_id, count) VALUES (10, 1, 2);
INSERT INTO Order_details (order_id, product_id, count) VALUES (10, 3, 1);
INSERT INTO Order_details (order_id, product_id, count) VALUES (10, 6, 1);
INSERT INTO Order_details (order_id, product_id, count) VALUES (11, 2, 3);
INSERT INTO Order_details (order_id, product_id, count) VALUES (14, 1, 3);
INSERT INTO Order_details (order_id, product_id, count) VALUES (14, 4, 4);
INSERT INTO Order_details (order_id, product_id, count) VALUES (15, 5, 1);
