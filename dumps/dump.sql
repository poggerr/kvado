CREATE DATABASE IF NOT EXISTS kvado;

CREATE USER "kvado"@"localhost" IDENTIFIED BY "root";

GRANT ALL PRIVILEGES ON * . * TO 'kvado'@'localhost';

FLUSH PRIVILEGES;

CREATE TABLE IF NOT EXISTS Author (
    Id INT PRIMARY KEY,
    Name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Book (
    Id INT PRIMARY KEY,
    AuthorId INT,
    Name VARCHAR(255),
    FOREIGN KEY (AuthorId) REFERENCES Author(Id)

);

INSERT INTO Author (Id, Name)
VALUES
    (1, 'Fyodor Dostoevsky'),
    (2, 'Alexander Volkov'),
    (3, 'Boris Pasternak'),
    (4, 'George Orwell'),
    (5, 'Alexander Belyaev'),
    (6, 'Nikolai Gogol'),
    (7, 'Fyodor Dostoevsky'),
    (8, 'Alexander Pushkin'),
    (9, 'Alexei Tolstoy'),
    (10, 'Mikhail Bulgakov'),
    (11, 'Alexander Pushkin'),
    (12, 'Anton Chekhov'),
    (13, 'Arkady and Boris Strugatsky'),
    (14, 'Mikhail Bulgakov'),
    (15, 'Mikhail Bulgakov'),
    (16, 'Vladimir Bogomolov'),
    (17, 'Mikhail Bulgakov'),
    (18, 'Konstantin Simonov'),
    (19, 'Alexandre Dumas'),
    (20, 'Valentin Kataev'),
    (21, 'Leo Tolstoy'),
    (22, 'Alexander Ostrovsky'),
    (23, 'Ivan Turgenev'),
    (24, 'Mark Twain'),
    (25, 'Anton Chekhov'),
    (26, 'Leo Tolstoy'),
    (27, 'Mark Twain'),
    (28, 'Jules Verne'),
    (29, 'Nikolai Gogol'),
    (30, 'Johann Wolfgang von Goethe'),
    (31, 'Gabriel Garcia Marquez'),
    (32, 'Jack London'),
    (33, 'Victor Hugo');


INSERT INTO Book (Id, AuthorId, Name)
VALUES
    (3691, 1, 'Poor Folk'),
    (5521, 2, 'The Wizard of the Emerald City'),
    (4547, 3, 'Doctor Zhivago'),
    (6130, 4, '1984'),
    (3956, 5, 'The Amphibian Man'),
    (6236, 6, 'Taras Bulba'),
    (2698, 7, 'The House of the Dead'),
    (2722, 8, 'The Bronze Horseman'),
    (3948, 9, 'Walking the Painful Path'),
    (3351, 10, 'Notes of a Young Doctor'),
    (2288, 11, 'Boris Godunov'),
    (2920, 12, 'The Seagull'),
    (4657, 13, 'Hard to Be a God'),
    (3065, 14, 'Morphine'),
    (2503, 15, 'Ivan Vasilyevich'),
    (2248, 16, 'Moment of Truth (In August 44th)'),
    (2848, 17, 'Run'),
    (2869, 18, 'The Living and the Dead'),
    (1731, 19, 'Forty-Five'),
    (2653, 20, 'The Lonely Sail Whitens'),
    (3484, 21, 'The Prisoner of the Caucasus'),
    (2408, 22, 'A Woman without a Dowry'),
    (2912, 23, 'A Sportsman''s Sketches'),
    (5172, 24, 'The Adventures of Tom Sawyer'),
    (2880, 25, 'Three Sisters'),
    (2612, 26, 'Hadji Murat'),
    (4155, 27, 'Adventures of Huckleberry Finn'),
    (3900, 28, 'The Mysterious Island'),
    (2963, 29, 'The Overcoat'),
    (4391, 30, 'Faust'),
    (5434, 31, 'One Hundred Years of Solitude'),
    (3160, 32, 'Martin Eden'),
    (1732, 33, 'The Hunchback of Notre-Dame');
