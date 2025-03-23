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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	// Validate subcategory.
	_, err := queries.FindBySCategoryIdEn(ctx.Context(), request.SCategoryEnID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid subCategoryEnID"})
	}

	// Convert price values.
	price, err := decimal.NewFromString(request.PriceEn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid price format"})
	}
	costPrice, err := decimal.NewFromString(request.CostPriceEn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cost price"})
	}
	retailPrice, err := decimal.NewFromString(request.RetailPriceEn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid retail price"})
	}
	var mainImage string
	if len(request.ImagesPathEn) > 0 {
		mainImage = request.ImagesPathEn[0]
	}
	// Create the main product record.
	createdProduct, err := queries.CreateProductEn(ctx.Context(), db.CreateProductEnParams{
		ProductNameEn:         request.ProductNameEn,
		SCategoryIdEn:         request.SCategoryEnID,
		PriceEn:               price.String(),
		StockQuantity:         request.StockQuantity,
		ImagesPathEn:          mainImage, // Main product field; images will be added separately.
		DescriptionEn:         request.DescriptionEn,
		BrandEn:               request.BrandEn,
		ManufacturedCountryEn: request.ManufacturedCountryEn,
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Insert multiple images.
	for _, imagePath := range request.ImagesPathEn {
		// Skip empty image paths.
		if imagePath == "" {
			continue
		}
		_, err = queries.CreateProductImageEn(ctx.Context(), db.CreateProductImageEnParams{
			ProductEnID: createdProduct.ProductEnID,
			ImagePath:   imagePath,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to insert product image",
				"details": err.Error(),
			})
		}
	}

	// Process product colors.
	for _, colorId := range request.ColorIds {
		_, err := queries.FindByColorId(ctx.Context(), colorId)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Color not found", "colorId": colorId})
		}
		_, err = queries.InsertProductEnColor(ctx.Context(), db.InsertProductEnColorParams{
			ProductEnID: createdProduct.ProductEnID,
			ColorId:     colorId,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert color link"})
		}
	}

	// Process product sizes.
	for _, sizeId := range request.SizeIds {
		_, err := queries.FindByIdSize(ctx.Context(), sizeId)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Size not found", "sizeId": sizeId})
		}
		_, err = queries.InsertProductEnSize(ctx.Context(), db.InsertProductEnSizeParams{
			ProductEnID: createdProduct.ProductEnID,
			SizeId:      sizeId,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to insert size link"})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Product created successfully",
		"productID": createdProduct.ProductEnID,
	})
}

// CreateProductMn creates a new Mongolian product with multiple images, colors, and sizes.
func (hd *Handlers) CreateProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CreateProductMnRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid request body"})
	}

	// Validate subcategory.
	_, err := queries.FindBySCategoryIdMn(ctx.Context(), request.SCategoryMnID)
	if err != nil {
		slog.Error("subcategory does not exist", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "invalid subCategoryMnID"})
	}

	// Convert price values.
	price, err := decimal.NewFromString(request.PriceMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid price format"})
	}
	costPrice, err := decimal.NewFromString(request.CostPriceMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid cost price format"})
	}
	retailPrice, err := decimal.NewFromString(request.RetailPriceMn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid retail price format"})
	}
	var mainImage string
	if len(request.ImagesPathMn) > 0 {
		mainImage = request.ImagesPathMn[0]
	}

	createdProduct, err := queries.CreateProductMn(ctx.Context(), db.CreateProductMnParams{
		ProductNameMn:         request.ProductNameMn,
		SCategoryIdMn:         request.SCategoryMnID,
		PriceMn:               price.String(),
		StockQuantity:         request.StockQuantity,
		ImagesPathMn:          mainImage,
		DescriptionMn:         request.DescriptionMn,
		BrandMn:               request.BrandMn,
		ManufacturedCountryMn: request.ManufacturedCountryMn,
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
		RetailPriceMn:         retailPrice.String(),
		WarehouseStockMn:      request.WarehouseStockMn,
	})
	if err != nil {
		slog.Error("unable to create product mn", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}

	// Insert multiple images.
	for _, imagePath := range request.ImagesPathMn {
		if imagePath == "" {
			continue
		}
		_, err = queries.CreateProductImageMn(ctx.Context(), db.CreateProductImageMnParams{
			ProductMnID: createdProduct.ProductMnID,
			ImagePath:   imagePath,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err":     "Failed to insert product image",
				"details": err.Error(),
			})
		}
	}

	// Insert color links.
	for _, colorId := range request.ColorIds {
		_, err := queries.InsertProductMnColor(ctx.Context(), db.InsertProductMnColorParams{
			ProductMnID: createdProduct.ProductMnID,
			ColorId:     colorId,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to insert color link"})
		}
	}

	// Insert size links.
	for _, sizeId := range request.SizeIds {
		_, err := queries.InsertProductMnSize(ctx.Context(), db.InsertProductMnSizeParams{
			ProductMnID: createdProduct.ProductMnID,
			SizeId:      sizeId,
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to insert size link"})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Product created successfully",
		"productID": createdProduct.ProductMnID,
	})
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

func (hd *Handlers) GetProductEnWithDetailsByID(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	idStr := ctx.Params("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product id"})
	}

	// Call the query that aggregates color and size details.
	product, err := queries.GetProductEnWithAllColorsAndSizesByID(ctx.Context(), int32(productID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(product)
}

func (hd *Handlers) GetProductMnWithDetailsByID(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	ProductIdStr := ctx.Params("id")
	ProductId, err := strconv.Atoi(ProductIdStr)
	if err != nil {
		slog.Error("unable to product Id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	Product, err := queries.GetProductMnWithAllColorsAndSizesByID(ctx.Context(), int32(ProductId))
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

func (hd *Handlers) UpdateSProductEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	ProductEnIdStr := ctx.Params("id")
	ProductEnId, err := strconv.Atoi(ProductEnIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "cannot parse id"})
	}

	var rqst models.UpdateSProductEn
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "cannot parse "})
	}

	_, err = queries.FindByProductIdEn(ctx.Context(), int32(ProductEnId))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Product not found"})
	}

	updatedSProduct, err := queries.UpdateSProductEn(ctx.Context(), db.UpdateSProductEnParams{
		StockQuantity: rqst.StockQuantity,
		ProductEnID:   int32(ProductEnId),
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Err": "Cannot update product stock quantity"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"msgg":          "successfully",
		"Updated Stock": updatedSProduct.StockQuantity,
		"ProductEnID":   updatedSProduct.ProductEnID,
	})
}

func (hd *Handlers) UpdateSProductMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	ProductMnIdStr := ctx.Params("id")
	ProductMnId, err := strconv.Atoi(ProductMnIdStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Err": "Cannot parse id"})
	}

	var rqst models.UpdateSProductMn
	if err := ctx.BodyParser(&rqst); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "parse error"})
	}

	_, err = queries.FindByProductIdMn(ctx.Context(), int32(ProductMnId))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Err": "not found"})
	}

	updatedSPRoductMn, err := queries.UpdateSProductMn(ctx.Context(), db.UpdateSProductMnParams{StockQuantity: rqst.StockQuantity, ProductMnID: int32(ProductMnId)})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Msgg":          "successfully",
		"Updated Stock": updatedSPRoductMn.StockQuantity,
		"productId":     ProductMnId,
	})
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

//		// Return the product with details
//		return ctx.Status(fiber.StatusOK).JSON(product)
//	}
//
// GET endpoint that joins the product with Color and Size details for English products.
func (hd *Handlers) GetProductEnWithDetails(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	// This query should join "productEn" with "Color" and "Size".
	products, err := queries.GetProductEnWithAllColorsAndSizes(ctx.Context())
	if err != nil {
		slog.Error("unable to fetch products with details", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (hd *Handlers) GetProductMnWithDetails(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	products, err := queries.GetProductMnWithAllColorsAndSizes(ctx.Context())
	if err != nil {
		slog.Error("unable to fetch Mongolian products with details", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}
