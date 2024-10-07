-- TABLE "detailMn"
BEGIN;

CREATE TABLE "detailMn" (
    "detailMnId" SERIAL PRIMARY KEY,
    -- Уникальный идентификатор для каждой детали продукта
    "ProductMnID" INT NOT NULL,
    -- Внешний ключ, ссылающийся на таблицу product
    "ChoiceName" VARCHAR(100) NOT NULL,
    -- Название параметра или выбора (например, Цвет, Размер и т.д.)
    "ChoiceValue" VARCHAR(100) NOT NULL,
    -- Значение параметра или выбора (например, Красный, Большой и т.д.)
    FOREIGN KEY ("ProductMnID") REFERENCES "productMn" ("ProductMnID") ON DELETE CASCADE -- Удаление деталей при удалении продукта (каскадное удаление)
) TABLESPACE pg_default;

COMMIT;