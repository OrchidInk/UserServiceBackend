-- -- name: CreateProductMn :one
-- INSERT INTO
--     "productMn" (
--         "CategoryMnID",
--         "PriceMn",
--         "StockQuantity",
--         "ImagesPathMn"
--     )
-- VALUES
--     (
--         "CategoryMnID" - sqlc.arg ('CategoryMnID'),
--         "PriceMn" = sqlc.arg ('PriceMn'),
--         "StockQuantity" = sqlc.arg ('StockQuantity'),
--         "ImagesPathMn" = sqlc.arg ('ImagesPathMn')
--     ) RETURNING *;
-- -- name: GetListProductMn :many
-- SELECT
--     *
-- FROM
--     "productMn"
-- ORDER BY
--     "Created_At" DESC;
-- -- name: UpdateByProductMnPrice :one
-- UPDATE "productMn"
-- SET
--     "PriceMn" = sqlc.arg ('PriceMn')
-- WHERE
--     "ProductMnID" = sqlc.arg ('ProductMnID') RETURNING *;
-- -- name: UpdateByProductMnStockQuantity :one
-- UPDATE "productMn"
-- SET
--     "StockQuantity" = sqlc.arg ('StockQuantity')
-- WHERE
--     "ProductMnID" = sqlc.arg ('ProductMnID') RETURNING *;
-- -- name: UpdateByMnImagePath :one
-- UPDATE "productMn"
-- SET
--     "ImagesPathMn" = sqlc.arg ('ImagesPathMn')
-- WHERE
--     "ProductMnID" = sqlc.arg ('ProductMnID') RETURNING *;
-- -- name: DeleteByProductMnId :exec
-- DELETE FROM "productMn"
-- WHERE
--     "ProductMnID" = sqlc.arg ('ProductMnID');
-- -- name: FilterByProductMnName :many
-- SELECT
--     *
-- FROM
--     "productMn"
-- WHERE
--     "ProductMnName" ILIKE '%' || sqlc.arg ('ProductMnName') || '%' -- Case-insensitive search for partial match
-- ORDER BY
--     "Created_At" DESC;