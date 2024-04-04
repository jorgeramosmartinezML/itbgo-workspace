-- 1. Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
SELECT d.nombre_depto, e.puesto, d.localidad
FROM empleado e
JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE e.puesto = 'Vendedor';

-- 2. Visualizar los departamentos con más de cinco empleados.
SELECT d.nombre_depto, COUNT(e.cod_emp) AS cantidad_empleados
FROM empleado e
JOIN departamento d ON e.depto_nro = d.depto_nro
GROUP BY d.nombre_depto
HAVING COUNT(e.cod_emp) > 5;

-- 3. Mostrar el nombre, salario y nombre del departamento de los empleados que
-- tengan el mismo puesto que ‘Mito Barchuk’.
SELECT e.nombre, e.salario, d.nombre_depto
FROM empleado e
JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE e.puesto = (SELECT e2.puesto
                  FROM empleado e2
                  WHERE e2.nombre = 'Mito' AND e2.apellido = 'Barchuk');

-- 4. Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, 
-- ordenados por nombre.
SELECT e.*
FROM empleado e
JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Contabilidad'
ORDER BY e.nombre;

SELECT e.*
FROM empleado e
WHERE e.depto_nro = (SELECT d.depto_nro
                     FROM departamento d
                     WHERE d.nombre_depto = 'Contabilidad')
ORDER BY e.nombre;


-- 5. Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT e.nombre
FROM empleado e
WHERE e.salario = (SELECT MIN(e2.salario)
                   FROM empleado e2);

SELECT e.nombre
FROM empleado e
ORDER BY e.salario
LIMIT 1;

-- 6. Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e.*
FROM empleado e
JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Ventas'
ORDER BY e.salario DESC
LIMIT 1;

