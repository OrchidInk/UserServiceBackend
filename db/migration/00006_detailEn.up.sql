-- TABLE "detailEn"
BEGIN;

CREATE TABLE "detailEn" (
    "detailEnId" SERIAL PRIMARY KEY,       -- Уникальный идентификатор для каждой детали продукта
    "ProductEnID" INT NOT NULL,              -- Внешний ключ, ссылающийся на таблицу product
    "ChoiceName" VARCHAR(100) NOT NULL,    -- Название параметра или выбора (например, Цвет, Размер и т.д.)
    "ChoiceValue" VARCHAR(100) NOT NULL,   -- Значение параметра или выбора (например, Красный, Большой и т.д.)
    FOREIGN KEY ("ProductEnID") REFERENCES "productEn"("ProductEnID") ON DELETE CASCADE -- Удаление деталей при удалении продукта (каскадное удаление)
) TABLESPACE pg_default;

COMMIT;
