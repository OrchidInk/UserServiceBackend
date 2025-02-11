package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateProductEnRequest
	if err := ctx.BodyParser(&request); err != nil {
		slog.Error("unable to parse request body", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request body"})
	}

	// Check if the subCategoryEnID exists
	_, err := queries.FindBySCategoryIdEn(ctx.Context(), request.SCategoryEnID)
	if err != nil {
		slog.Error("subcategory does not exist", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid subCategoryEnID"})
	}

	// Convert price to decimal
	price, err := decimal.NewFromString(request.PriceEn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid price format"})
	}
	if price.Exponent() < -2 || price.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Price exceeds allowed range (maximum: 9999999999.99)",
		})
	}

	costPrice, err := decimal.NewFromString(request.CostPriceEn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid price format"})
	}
	if costPrice.Exponent() < -2 || costPrice.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Price exceeds allowed range (maximum: 9999999999.99)",
		})
	}

	retailPrice, err := decimal.NewFromString(request.RetailPriceEn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid price format"})
	}
	if retailPrice.Exponent() < -2 || retailPrice.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Price exceeds allowed range (maximum: 9999999999.99)",
		})
	}

	// Insert product
	createProduct, err := queries.CreateProductEn(ctx.Context(), db.CreateProductEnParams{
		ProductNameEn:         request.ProductNameEn,
		SCategoryIdEn:         request.SCategoryEnID,
		PriceEn:               price.String(),
		StockQuantity:         request.StockQuantity,
		ImagesPathEn:          request.ImagesPathEn,
		DescriptionEn:         request.DescriptionEn,
		BrandEn:               request.BrandEn,
		ManufacturedCountryEn: request.ManufacturedCountryEn,
		ColorEn:               request.ColorEn,
		SizeEn:                request.SizeEn,
		PenOutputEn:           request.PenOutputEn,
		FeaturesEn:            request.FeaturesEn,
		MaterialEn:            request.MaterialEn,
		StapleSizeEn:          request.StapleSizeEn,
		CapacityEn:            request.CapacityEn,
		WeightEn:              request.WeightEn,
		ThicknessEn:           request.ThicknessEn,
		PackagingEn:           request.PackagingEn,
		ProductCodeEn:         request.ProductCodeEn,
		CostPriceEn:           costPrice.String(),
		RetailPriceEn:         retailPrice.String(),
		WarehouseStockEn:      request.WarehouseStockEn,
	})
	if err != nil {
		slog.Error("unable to create product en", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Product created successfully",
		"product ID": createProduct.ProductEnID,
	})
}

func (hd *Handlers) CreateProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateProductMnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err := queries.FindBySCategoryIdMn(ctx.Context(), request.SCategoryMnID)
	if err != nil {
		slog.Error("subCategory does not exist", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid subCategoryMNID"})
	}

	price, err := decimal.NewFromString(request.PriceMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid price format"})
	}

	if price.Exponent() < -2 || price.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Price exceeds allowed range (maximum: 9999999999.99)",
		})
	}

	costPrice, err := decimal.NewFromString(request.CostPriceMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid cost price format"})
	}

	if costPrice.Exponent() < -2 || costPrice.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Price exceeds allowed range (maximum: 9999999999.99)",
		})
	}

	retailPrice, err := decimal.NewFromString(request.RetailPriceMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid cost price format"})
	}

	if retailPrice.Exponent() < -2 || retailPrice.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Price exceeds allowed range (maximum: 9999999999.99)",
		})
	}

	createProduct, err := queries.CreateProductMn(ctx.Context(), db.CreateProductMnParams{
		ProductNameMn:         request.ProductNameMn,
		SCategoryIdMn:         request.SCategoryMnID,
		PriceMn:               price.String(),
		StockQuantity:         request.StockQuantity,
		ImagesPathMn:          request.ImagesPathMn,
		DescriptionMn:         request.DescriptionMn,
		BrandMn:               request.BrandMn,
		ManufacturedCountryMn: request.ManufacturedCountryMn,
		ColorMn:               request.ColorMn,
		SizeMn:                request.SizeMn,
		PenOutputMn:           request.PenOutputMn,
		FeaturesMn:            request.FeaturesMn,
		MaterialMn:            request.MaterialMn,
		StapleSizeMn:          request.StapleSizeMn,
		CapacityMn:            request.CapacityMn,
		WeightMn:              request.WeightMn,
		ThicknessMn:           request.ThicknessMn,
		PackagingMn:           request.PackagingMn,
		ProductCodeMn:         request.ProductCodeMn,
		CostPriceMn:           costPrice.String(),
		RetailPriceMn:         request.RetailPriceMn,
		WarehouseStockMn:      request.WarehouseStockMn,
	})
	if err != nil {
		slog.Error("unable to create product request", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "product create successfully", "product Id": createProduct.ProductMnID})
}

func (hd *Handlers) DeleteProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindByProductIdEn(ctx.Context(), int32(ProductID))
	if err != nil {
		slog.Error("unable to find product", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid find product id"})
	}

	err = queries.DeleteByProductEnId(ctx.Context(), int32(ProductID))
	if err != nil {
		slog.Error("unable to delete productEn id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted product ID"})
}

func (hd *Handlers) DeleteProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductId, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	_, err = queries.FindByProductIdMn(ctx.Context(), int32(ProductId))
	if err != nil {
		slog.Error("unable to find product ID", slog.Any("err", err))
	}

	err = queries.DeleteByProductMnId(ctx.Context(), int32(ProductId))
	if err != nil {
		slog.Error("unable to find product ID", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully deleted product "})
}

func (hd *Handlers) DeductProductStockEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid product ID"})
	}

	var request struct {
		QuantityPurchased int32 `json:"quantityPurchased"`
	}
	if err := ctx.BodyParser(&request); err != nil || request.QuantityPurchased <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid quantity"})
	}

	updatedProduct, err := queries.DeductSockQuantityByProductEnID(ctx.Context(), db.DeductSockQuantityByProductEnIDParams{
		ProductEnID:       int32(ProductID),
		QuantityPurchased: request.QuantityPurchased,
	})
	if err != nil {
		slog.Error("unable to deduct stock", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to deduct stock", "error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Stock deducted successfully",
		"productId": updatedProduct.ProductEnID,
		"stockLeft": updatedProduct.StockQuantity,
	})
}

func (hd *Handlers) DeductProductStockMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIDStr := ctx.Params("id")
	ProductID, err := strconv.Atoi(ProductIDStr)
	if err != nil {
		slog.Error("unable to parse id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	var request struct {
		QuantityPurchased int32 `json:"quantityPurchased"`
	}
	if err := ctx.BodyParser(&request); err != nil || request.QuantityPurchased <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid quantity"})
	}

	updatedProduct, err := queries.DeductSockQuantityByProductMnID(ctx.Context(), db.DeductSockQuantityByProductMnIDParams{
		ProductMnID:       int32(ProductID),
		QuantityPurchased: request.QuantityPurchased,
	})
	if err != nil {
		slog.Error("unable to deduct stock", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to deduct stock", "error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Stock deducted successfully",
		"productId": updatedProduct.ProductMnID,
		"stockLeft": updatedProduct.StockQuantity,
	})
}

func (hd *Handlers) GetProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	Product, err := queries.GetListProductEn(ctx.Context())
	if err != nil {
		slog.Error("unable to product get list")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(Product)
}

func (hd *Handlers) GetProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	Product, err := queries.GetListProductMn(ctx.Context())
	if err != nil {
		slog.Error("unable to product get list")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(Product)
}

func (hd *Handlers) FindProductId(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	ProductIdSTR := ctx.Params("id")
	ProductId, err := strconv.Atoi(ProductIdSTR)
	if err != nil {
		slog.Error("unable to find product id", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	Product, err := queries.FindByProductIdEn(ctx.Context(), int32(ProductId))
	if err != nil {
		slog.Error("unable to find id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(Product)
}

func (hd *Handlers) FindProductIdMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIdStr := ctx.Params("id")
	ProductId, err := strconv.Atoi(ProductIdStr)
	if err != nil {
		slog.Error("unable to product Id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	Product, err := queries.FindByProductIdMn(ctx.Context(), int32(ProductId))
	if err != nil {
		slog.Error("unable to find product Id", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(Product)
}

func (hd *Handlers) UpdateProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	ProductIdSTR := ctx.Params("id")
	ProductId, err := strconv.Atoi(ProductIdSTR)
	if err != nil {
		slog.Error("unable to convert id", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
	}

	var req db.UpdateProductEnParams
	if err := ctx.BodyParser(&req); err != nil {
		slog.Error("failed to parse request body", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	req.ProductEnID = int32(ProductId)

	UpdatedProduct, err := queries.UpdateProductEn(ctx.Context(), req)
	if err != nil {
		slog.Error("failed to update product", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update product", "details": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(UpdatedProduct)
}

func (hd *Handlers) UpdateProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	ProductIdStr := ctx.Params("id")
	ProductId, err := strconv.Atoi(ProductIdStr)
	if err != nil {
		slog.Error("unable to convert id", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
	}

	var req db.UpdateProductMnParams
	if err := ctx.BodyParser(&req); err != nil {
		slog.Error("failed to parse request body", slog.Any("Err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	req.ProductMnID = int32(ProductId)

	UpdatedProduct, err := queries.UpdateProductMn(ctx.Context(), req)
	if err != nil {
		slog.Error("failed to update product", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update product", "details": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(UpdatedProduct)
}

// func (hd *Handlers) GetProductWithDetailsEn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	// Fetch product data with associated details
// 	rows, err := queries.GetProductWithDetailsEn(ctx.Context())
// 	if err != nil {
// 		slog.Error("Unable to fetch products with details", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product details"})
// 	}

// 	// Map products with details
// 	productMap := make(map[int32]*models.ProductWithDetailsEn)
// 	for _, row := range rows {
// 		if _, exists := productMap[row.ProductEnID]; !exists {
// 			// Add product if not already in the map
// 			productMap[row.ProductEnID] = &models.ProductWithDetailsEn{
// 				ProductEnID:     row.ProductEnID,
// 				ProductNameEn:   row.ProductNameEn,
// 				SubCategoryIDEn: row.SubCategoryIDEn,
// 				PriceEn:         row.PriceEn,
// 				StockQuantity:   row.StockQuantity,
// 				ImagesPathEn:    row.ImagesPathEn,
// 				CreatedAt:       row.CreatedAt.Time,
// 				UpdatedAt:       row.UpdatedAt.Time,
// 				Details:         []models.DetailEn{},
// 			}
// 		}

// 		if row.DetailEnId.Valid {
// 			productMap[row.ProductEnID].Details = append(productMap[row.ProductEnID].Details, models.DetailEn{
// 				DetailEnID:  row.DetailEnId.Int32,
// 				ChoiceName:  row.ChoiceName.String,
// 				ChoiceValue: row.ChoiceValue.String,
// 			})
// 		}
// 	}

// 	// Convert map to slice for response
// 	products := make([]models.ProductWithDetailsEn, 0, len(productMap))
// 	for _, product := range productMap {
// 		products = append(products, *product)
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(products)
// }

// func (hd *Handlers) GetProductWithDetailsMn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	rows, err := queries.GetProductWithDetailMn(ctx.Context())
// 	if err != nil {
// 		slog.Error("unable to fetch products with details", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to fetch product with details"})
// 	}

// 	productMap := make(map[int32]*models.ProductWithDetailsMn)
// 	for _, row := range rows {
// 		if _, exists := productMap[row.ProductMnID]; !exists {
// 			productMap[row.ProductMnID] = &models.ProductWithDetailsMn{
// 				ProductMnID:     row.ProductMnID,
// 				ProductNameMn:   row.ProductNameMn,
// 				SubCategoryIDMn: row.SubCategoryIDMn,
// 				PriceMn:         row.PriceMn,
// 				StockQuantity:   row.StockQuantity,
// 				ImagesPathMn:    row.ImagesPathMn,
// 				CreatedAt:       row.CreatedAt.Time,
// 				UpdatedAt:       row.UpdatedAt.Time,
// 				Details:         []models.DetailMn{},
// 			}
// 		}
// 		if row.DetailMnId.Valid {
// 			productMap[row.ProductMnID].Details = append(productMap[row.ProductMnID].Details, models.DetailMn{
// 				DetailMnID:  row.DetailMnId.Int32,
// 				ChoiceName:  row.ChoiceName.String,
// 				ChoiceValue: row.ChoiceValue.String,
// 			})
// 		}
// 	}

// 	products := make([]models.ProductWithDetailsMn, 0, len(productMap))
// 	for _, product := range productMap {
// 		products = append(products, *product)
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(products)
// }

// func (hd *Handlers) FindByProductWithDetailsByIDEn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	productIDStr := ctx.Params("id")
// 	productID, err := strconv.Atoi(productIDStr)
// 	if err != nil {
// 		slog.Error("Invalid product ID", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
// 	}

// 	rows, err := queries.FindProductWithDetailsByIDEn(ctx.Context(), int32(productID))
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			slog.Error("Product not found", slog.Int("ProductID", productID))
// 			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
// 		}
// 		slog.Error("Database error while finding product", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product details"})
// 	}

// 	var product *models.ProductWithDetailsEn
// 	for _, row := range rows {
// 		if product == nil {
// 			product = &models.ProductWithDetailsEn{
// 				ProductEnID:     row.ProductEnID,
// 				ProductNameEn:   row.ProductNameEn,
// 				SubCategoryIDEn: row.SubCategoryIDEn,
// 				PriceEn:         row.PriceEn,
// 				StockQuantity:   row.StockQuantity,
// 				ImagesPathEn:    row.ImagesPathEn,
// 				CreatedAt:       row.CreatedAt.Time,
// 				UpdatedAt:       row.UpdatedAt.Time,
// 				Details:         []models.DetailEn{},
// 			}
// 		}

// 		if row.DetailEnId.Valid { // Ensure the detail is not null
// 			product.Details = append(product.Details, models.DetailEn{
// 				DetailEnID:  row.DetailEnId.Int32,
// 				ChoiceName:  row.ChoiceName.String,
// 				ChoiceValue: row.ChoiceValue.String,
// 			})
// 		}
// 	}

// 	if product == nil {
// 		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(product)
// }

// func (hd *Handlers) FindByProductWithDetailsByIDMn(ctx *fiber.Ctx) error {
// 	queries, _, _ := hd.queries()

// 	// Parse the product ID from URL
// 	productIDSTR := ctx.Params("id")
// 	productID, err := strconv.Atoi(productIDSTR)
// 	if err != nil {
// 		slog.Error("Unable to parse product ID", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid product ID"})
// 	}

// 	// Query the product and its details
// 	rows, err := queries.FindProductWithDetailsByIDMn(ctx.Context(), int32(productID))
// 	if err != nil {
// 		slog.Error("Database query error", slog.Any("err", err))
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
// 	}

// 	// Build the product response
// 	var product *models.ProductWithDetailsMn
// 	for _, row := range rows {
// 		if product == nil {
// 			product = &models.ProductWithDetailsMn{
// 				ProductMnID:     row.ProductMnID,
// 				ProductNameMn:   row.ProductNameMn,
// 				SubCategoryIDMn: row.SubCategoryIDMn,
// 				PriceMn:         row.PriceMn,
// 				StockQuantity:   row.StockQuantity,
// 				ImagesPathMn:    row.ImagesPathMn,
// 				CreatedAt:       row.CreatedAt.Time,
// 				UpdatedAt:       row.UpdatedAt.Time,
// 				Details:         []models.DetailMn{},
// 			}
// 		}

// 		if row.DetailMnId.Valid {
// 			product.Details = append(product.Details, models.DetailMn{
// 				DetailMnID:  row.DetailMnId.Int32,
// 				ChoiceName:  row.ChoiceName.String,
// 				ChoiceValue: row.ChoiceValue.String,
// 			})
// 		}
// 	}

// 	// Handle product not found
// 	if product == nil {
// 		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
// 	}

// 	// Return the product with details
// 	return ctx.Status(fiber.StatusOK).JSON(product)
// }
