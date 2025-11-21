-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProduct :one
SELECT * FROM products WHERE id = $ 1;