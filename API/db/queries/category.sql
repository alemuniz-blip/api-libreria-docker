-- name: GetAllCategories :many
SELECT id, name, created_at, updated_at FROM categories;

-- name: GetCategoryById :one
SELECT id, name, created_at, updated_at FROM categories WHERE id = ?;

-- name: CreateCategory :execresult
INSERT INTO categories(name, created_at, updated_at)
VALUES (?, NOW(), NOW());

-- name: UpdateCategory :exec
UPDATE categories SET name=? WHERE id=?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id=?;