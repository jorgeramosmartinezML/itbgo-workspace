-- 1. Listar todos los clientes
SELECT * FROM cliente;

-- 2. Mostrar todos los planes de intenet disponibles
SELECT * FROM plan;

-- 3. Obtener el nombre y el plan de internet contratando por cada cliente
SELECT c.firstname, c.lastname, ct.id_plan 
FROM customers c
JOIN contract ct ON c.dni = ct.id_costumer
JOIN plan p ON p.id = ct.id_plan;

-- 4. Contar la cantidad de clientes en cada provincia
SELECT province, COUNT(*) AS cantidad_clientes
FROM customers
GROUP BY province;

-- 5. Mostrar los clientes que tienen un descuento en su plan de internet
SELECT c.firstname, c.lastname, ct.id_plan
FROM customers c
JOIN contract ct ON c.dni = ct.id_costumer
JOIN plan p ON p.id = ct.id_plan
WHERE ct.discount > 0;

-- 6. Calcular el precio total que paga cada cliente por su plan de internet (sin descuento)
SELECT c.firstname, c.lastname, SUM(p.price) AS precio_total
FROM customers c
JOIN contract ct ON c.dni = ct.id_costumer
JOIN plan p ON p.id = ct.id_plan;
GROUP BY c.dni, c.firstname, c.lastname;

-- 7. Encontrar los clientes que tienen más de 30 años.
SELECT firstname, lastname
FROM customers
WHERE birthday <= DATE_SUB(CURDATE(), INTERVAL 30 YEAR);

-- 8. Listar los clientes de una provincia específica, ordenados alfabéticamente por apellidos.
SELECT firstname, lastname
WHERE customers
WHERE province = 'Madrid'
ORDER BY lastname;

-- 9. Mostart los planes de internet que tienen una velocidad mayor a 100 Megas
SELECT id, speed, price
FROM plan
WHERE speed > 100;

-- 10. Contar la cantidad de clientes que tienen contratado cada plan de internet.
SELECT id_plan, COUNT(*) AS cantidad_clientes
FROM plan p
INNER JOIN contract ct ON p.id = ct.id_plan
GROUP BY id_plan;
