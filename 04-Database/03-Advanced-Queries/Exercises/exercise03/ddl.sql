CREATE DATABASE biblioteca;

USE biblioteca;

CREATE TABLE autor (
    idAutor INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(50) NOT NULL,
    nacionalidad VARCHAR(50) NOT NULL
);

CREATE TABLE libro (
    idLibro INT PRIMARY KEY AUTO_INCREMENT,
    titulo VARCHAR(50) NOT NULL,
    area VARCHAR(50) NOT NULL,
    editorial VARCHAR(50) NOT NULL
);

CREATE TABLE libroAutor (
    idLibro INT,
    idAutor INT,
    PRIMARY KEY (idLibro, idAutor),
    FOREIGN KEY (idLibro) REFERENCES libro(idLibro),
    FOREIGN KEY (idAutor) REFERENCES autor(idAutor)
);

CREATE TABLE estudiante (
    idLector INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(50) NOT NULL,
    apellido VARCHAR(50) NOT NULL,
    direccion VARCHAR(50) NOT NULL,
    carrera VARCHAR(50) NOT NULL,
    edad INT NOT NULL
);

CREATE TABLE prestamo (
    idPrestamo INT PRIMARY KEY AUTO_INCREMENT,
    idLibro INT,
    idLector INT,
    fechaPrestamo DATE NOT NULL,
    fechaDevolucion DATE NOT NULL,
    devuelto BOOLEAN NOT NULL,
    FOREIGN KEY (idLibro) REFERENCES libro(idLibro),
    FOREIGN KEY (idLector) REFERENCES estudiante(idLector)
);

INSERT INTO autor (nombre, nacionalidad) VALUES
('J.K. Rowling', 'Británica'),
('Stephen King', 'Estadounidense'),
('George R.R. Martin', 'Estadounidense'),
('J.R.R. Tolkien', 'Británica'),
('Isaac Asimov', 'Rusa');

INSERT INTO libro (titulo, area, editorial) VALUES
('Harry Potter y la piedra filosofal', 'Fantasía', 'Salamandra'),
('It', 'Terror', 'Viking Press'),
('Canción de hielo y fuego', 'Fantasía', 'Bantam Spectra'),
('El Señor de los Anillos', 'Fantasía', 'George Allen & Unwin'),
('Fundación', 'Ciencia Ficción', 'Gnome Press');

INSERT INTO libroAutor (idLibro, idAutor) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5);

INSERT INTO estudiante (nombre, apellido, direccion, carrera, edad) VALUES
('Harry', 'Potter', '4 Privet Drive', 'Magia', 11),
('Bill', 'Denbrough', '29 Neibolt Street', 'Ingeniería', 12),
('Jon', 'Snow', 'Winterfell', 'Administración', 14),
('Frodo', 'Bolsón', 'La Comarca', 'Ciencias de la Computación', 50),
('Hari', 'Seldon', 'Trántor', 'Matemáticas', 60);

INSERT INTO prestamo (idLibro, idLector, fechaPrestamo, fechaDevolucion, devuelto) VALUES
(1, 1, '2020-01-01', '2020-01-15', TRUE),
(2, 2, '2020-01-01', '2020-01-15', TRUE),
(3, 3, '2020-01-01', '2020-01-15', TRUE),
(4, 4, '2020-01-01', '2020-01-15', TRUE),
(5, 5, '2020-01-01', '2020-01-15', TRUE);