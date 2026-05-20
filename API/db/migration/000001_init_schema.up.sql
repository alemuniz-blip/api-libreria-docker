CREATE TABLE users(
	id int AUTO_INCREMENT NOT NULL,
	name varchar(100) NOT NULL,
	email varchar(150),
	password varchar(255),
	role varchar(30),
	created_at datetime DEFAULT NULL,
	updated_at datetime DEFAULT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE categories(
	id int AUTO_INCREMENT NOT NULL,
	name varchar(100) NOT NULL,
	created_at datetime DEFAULT NULL,
	updated_at datetime DEFAULT NULL,
	PRIMARY KEY(id)
);

CREATE TABLE products(
	id int AUTO_INCREMENT NOT NULL,
	name varchar(150) NOT NULL,
	description text,
	price decimal(10,2),
	stock int,
	category_id int NOT NULL,
	created_at datetime DEFAULT NULL,
	updated_at datetime DEFAULT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE carrito (
    id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id INT NOT NULL
);

CREATE TABLE carrito_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    carrito_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL,

    FOREIGN KEY (carrito_id) REFERENCES carrito(id),
    FOREIGN KEY (producto_id) REFERENCES products(id)
);

CREATE TABLE compra (
    id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id INT,
    total DECIMAL(10,2),
    estado VARCHAR(50),
    fecha DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE detalle_compra (
    id INT AUTO_INCREMENT PRIMARY KEY,
    compra_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL,
    precio_unitario DECIMAL(10,2),
    subtotal DECIMAL(10,2),

    FOREIGN KEY (compra_id) REFERENCES compra(id),
    FOREIGN KEY (producto_id) REFERENCES products(id)
);