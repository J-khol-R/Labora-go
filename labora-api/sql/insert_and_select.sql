
INSERT INTO items (id, customer_name, order_date, product, quantity, price)
VALUES (1, 'valentina', '2023-05-03', 'naranja',10, 20),
(2, 'juan', '2023-06-12', 'mora', 12, 65),
(3, 'valeria', '2023-04-20', 'mango', 5, 10),
(4, 'rodrigo', '2023-01-04', 'uva', 11, 35),
(5, 'sofia', '2023-12-31', 'fresa', 8,  80);

SELECT * FROM items WHERE quantity > 2 AND price > 50;
