-- 1. Listar los datos de los autores.
SELECT *
FROM autor;

-- 2. Listar nombre y edad de los estudiantes
SELECT Nombre, Edad
FROM estudiante;

-- 3. ¿Qué estudiantes pertenecen a la carrera informática?
SELECT *
FROM estudiante
WHERE Carrera = 'Informática';

-- 4. ¿Qué autores son de nacionalidad francesa o italiana?
SELECT *
FROM autor
WHERE Nacionalidad = 'Francesa' OR Nacionalidad = 'Italiana';

-- 5. ¿Qué libros no son del área de internet?
SELECT *
FROM libro
WHERE Area != 'Internet';

-- 6. Listar los libros de la editorial Salamandra.
SELECT *
FROM libro
WHERE Editorial = 'Salamandra';

-- 7. Listar los datos de los estudiantes cuya edad es mayor al promedio.
SELECT *
FROM estudiante
WHERE Edad > (SELECT AVG(Edad) FROM estudiante);

-- 8. Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
SELECT Nombre
FROM estudiante
WHERE Apellido LIKE 'G%';

-- 9. Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
SELECT autor.Nombre
FROM autor
JOIN libroautor ON autor.IdAutor = libroautor.IdAutor
JOIN libro ON libro.IdLibro = libroautor.IdLibro
WHERE libro.Titulo = 'El Universo: Guía de viaje';

-- 10. ¿Qué libros se prestaron al lector “Filippo Galli”?
SELECT l.Titulo
FROM libro l
JOIN prestamo p ON l.IdLibro = p.IdLibro
JOIN estudiante e ON e.IdLector = p.IdLector
WHERE e.Nombre = 'Filippo' AND e.Apellido = 'Galli';

-- 11. Listar el nombre del estudiante de menor edad.
SELECT Nombre
FROM estudiante
WHERE Edad = (SELECT MIN(Edad) FROM estudiante);

SELECT Nombre
FROM estudiante
ORDER BY Edad
LIMIT 1;

-- 12. Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
SELECT e.Nombre 
FROM estudiante e
JOIN prestamo p ON e.IdLector = p.IdLector
JOIN libro l ON l.IdLibro = p.IdLibro
WHERE l.Area = 'Base de Datos';

-- 13. Listar los libros que pertenecen a la autora J.K. Rowling.
SELECT l.Titulo
FROM libro l
JOIN libroautor la ON l.IdLibro = la.IdLibro
JOIN autor a ON a.IdAutor = la.IdAutor
WHERE a.Nombre = 'J.K. Rowling';

-- 14. Listar títulos de los libros que debían devolverse el 16/07/2021.
SELECT l.Titulo
FROM libro l
JOIN prestamo p ON l.IdLibro = p.IdLibro
WHERE p.FechaDevolucion = '2021-07-16';




