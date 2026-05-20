-- name: AddToCart :exec
INSERT INTO carrito_items (carrito_id, producto_id, cantidad)
VALUES (?, ?, ?);

-- name: CreateCarrito :exec
INSERT INTO carrito (usuario_id)
VALUES (?);

-- name: GetCartItems :many
SELECT 
    ci.carrito_id,
    p.id,
    p.name,
    CAST(p.price AS DECIMAL(10,2)) AS price,
    ci.cantidad
FROM carrito_items ci
JOIN products p ON ci.producto_id = p.id
WHERE ci.carrito_id = ?;

-- name: GetAllCarritos :many
SELECT 
    c.id AS carrito_id,
    c.usuario_id AS usuario_id
FROM carrito c;

-- name: GetCarritosByProducto :many
SELECT 
    ci.carrito_id,
    p.id,
    p.name,
    ci.cantidad
FROM carrito_items ci
JOIN products p ON ci.producto_id = p.id
WHERE p.id = ?;

-- name: GetCarritosByUserV3 :many
SELECT 
    c.id,
    c.usuario_id
FROM carrito c
WHERE c.usuario_id = sqlc.arg(usuario_id);

-- name: DeleteCartItem :exec
DELETE FROM carrito_items
WHERE carrito_id = ? AND producto_id = ?