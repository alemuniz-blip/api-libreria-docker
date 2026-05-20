-- name: CreateCompra :execresult
INSERT INTO compra (usuario_id, total, estado, fecha)
VALUES (?, ?, ?, NOW());

-- name: CreateDetalleCompra :exec
INSERT INTO detalle_compra (compra_id, producto_id, cantidad, precio_unitario, subtotal)
VALUES (?, ?, ?, ?, ?);

-- name: GetDetalleCompra :many
SELECT 
    dc.compra_id,
    p.id,
    p.name,
    dc.precio_unitario,
    dc.cantidad,
    dc.subtotal
FROM detalle_compra dc
JOIN products p ON dc.producto_id = p.id
WHERE dc.compra_id = ?;

-- name: UpdateCompraTotal :exec
UPDATE compra
SET total = ?
WHERE id = ?;

-- name: GetCompraById :one
SELECT 
    id,
    usuario_id,
    fecha,
    total,
    estado
FROM compra
WHERE id = ?;

-- name: GetAllCompras :many
SELECT 
    id,
    usuario_id,
    fecha,
    total,
    estado
FROM compra;

-- name: GetComprasByUser :many
SELECT 
    id,
    usuario_id,
    fecha,
    total,
    estado
FROM compra
WHERE usuario_id = ?;