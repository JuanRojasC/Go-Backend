CREATE IF NOT EXISTS DATABASE emple_dep;

DROP TABLE IF EXISTS departamentos;

CREATE TABLE IF NOT EXISTS departamentos(
    'id' INTEGER AUTO_INCREMENT,
    'nombre' VARCHAR(200),
    'direccion' VARCHAR(200),
    PRIMARY KEY ('id')
);

DROP TABLE IF EXISTS empleados;

CREATE TABLE IF NOT EXISTS empleados(
    'legajo' INTEGER INT,
    'nombre' VARCHAR(100),
    'apellido' VARCHAR(100)
    'dni' INTEGER,
    'cargo' VARCHAR(100),
    'sueldo_neto' DECIMAL(12, 2),
    'fecha_incorporacion' DATE,
    'fecha_nacimiento' DATE,
    'dep_id' INTEGER,
    PRIMARY KEY ('legajo')
    FOREIGN KEY ('dep_id') REFERENCES departamentos(id)
);

INSERT INTO departamentos (nombre, direccion) VALUES ('a', 'calle 1');
INSERT INTO departamentos (nombre, direccion) VALUES ('b', 'calle 2');
INSERT INTO departamentos (nombre, direccion) VALUES ('c', 'calle 3');
INSERT INTO departamentos (nombre, direccion) VALUES ('d', 'calle 4');
INSERT INTO departamentos (nombre, direccion) VALUES ('e', 'calle 5');

INSERT INTO empleados VALUES (1, 'juan', 'rojas', 74365495, 'software developer', 1000, STR_TO_DATE('14-03-2022', '%d-%m-%Y'), STR_TO_DATE('2-10-2000', '%d-%m-%Y'), 1);
INSERT INTO empleados VALUES (12, 'pedro', 'alcantaraz', 345345, 'contador', 1000, STR_TO_DATE('8-10-2022', '%d-%m-%Y'), STR_TO_DATE('20-01-2000', '%d-%m-%Y'), 2);
INSERT INTO empleados VALUES (123, 'sofia', 'mendez', 74365495, 'gerente', 1000, STR_TO_DATE('27-03-2022', '%d-%m-%Y'), STR_TO_DATE('7-02-2000', '%d-%m-%Y'), 1);


