-- name: CreateProduct :exec
INSERT INTO products (name, description, price, stock, category_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, NOW(), NOW());

-- name: GetAllProducts :many
SELECT 
    p.id,
    p.name,
    p.description,
    p.price,
    p.stock,
    c.name AS category
FROM products p
JOIN categories c ON p.category_id = c.id;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = ?;

-- name: GetProductById :one
SELECT 
    id,
    name,
    description,
    price,
    stock,
    category_id,
    created_at,
    updated_at
FROM products
WHERE id = ?;

-- name: GetProductsByCategory :many
SELECT 
    p.id,
    p.name,
    p.description,
    p.price,
    p.stock,
    c.name AS category
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE p.category_id = sqlc.arg(category_id);

-- name: UpdateProduct :exec
UPDATE products
SET 
    name = ?,
    description = ?,
    price = ?,
    stock = ?,
    category_id = ?,
    updated_at = NOW()
WHERE id = ?;