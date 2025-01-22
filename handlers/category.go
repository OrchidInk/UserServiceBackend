package handlers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	db "orchid.admin.service/db/sqlc"
	"orchid.admin.service/models"
)

func (hd *Handlers) CreateCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CategoryEn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	_, err := queries.FindByNameCategoryEn(ctx.Context(), request.CategoryNameEn)
	if err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Category already exists"})
	}

	category, err := queries.CreateCategoryEn(ctx.Context(), request.CategoryNameEn)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Category created successfully",
		"categoryId":   category.CategoryEnID,
		"categoryName": category.CategoryNameEn,
	})
}

func (hd *Handlers) CreateCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.CategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	_, err := queries.FindByNameMnCategoryMn(ctx.Context(), request.CategoryNameMn)
	if err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "category already exist"})
	}

	category, err := queries.CreateCategoryMn(ctx.Context(), request.CategoryNameMn)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Category created successfully",
		"categoryId":   category.CategoryMnID,
		"categoryName": category.CategoryNameMn,
	})
}

func (hd *Handlers) UpdateCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()
	var request models.UpdateCategoryEn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	category, err := queries.FindByCategoryEnId(ctx.Context(), request.CategoryEnId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	updatedCategory, err := queries.UpdateCategoryEn(ctx.Context(), db.UpdateCategoryEnParams{
		CategoryNameEn: request.CategoryNameEn,
		CategoryEnID:   category.CategoryEnID,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Category updated successfully",
		"categoryId":   updatedCategory.CategoryEnID,
		"categoryName": updatedCategory.CategoryNameEn,
	})
}

func (hd *Handlers) UpdateCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	var request models.UpdateCategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	category, err := queries.FindByCategoryMnId(ctx.Context(), request.CategoryMnId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "category not found"})
	}

	updatedCategory, err := queries.UpdateCategoryMn(ctx.Context(), db.UpdateCategoryMnParams{
		CategoryNameMn: request.CategoryNameMn,
		CategoryMnID:   category.CategoryMnID,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to updated"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Category updated successfully",
		"CategoryID":   updatedCategory.CategoryMnID,
		"categoryName": updatedCategory.CategoryNameMn,
	})
}

func (hd *Handlers) DeleteCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid category ID"})
	}

	_, err = queries.FindByCategoryEnId(ctx.Context(), int32(categoryID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	err = queries.DeleteCategoryById(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to delete category ID", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category deleted successfull", "category_ID": categoryID})
}

func (hd *Handlers) DeleteCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid category ID"})
	}

	_, err = queries.FindByCategoryMnId(ctx.Context(), int32(categoryID))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	err = queries.DeleteCategoryByMnId(ctx.Context(), int32(categoryID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete category"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category deleted successfully", "category_id": categoryID})
}

func (hd *Handlers) GetListCategoryEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	category, err := queries.GetListByAllCategoryEn(ctx.Context())
	if err != nil {
		slog.Error("unable to get list", slog.Any("err", "err"))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "err list"})
	}

	return ctx.JSON(category)
}

func (hd *Handlers) GetListCategoryMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	category, err := queries.GetListByAllCategoryMn(ctx.Context())
	if err != nil {
		slog.Error("unable to get list", slog.Any("err", "err"))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "err get list"})
	}

	return ctx.JSON(category)
}

func (hd *Handlers) GetCategoriesWithSubCategoriesEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	rows, err := queries.GetCategoriesWithSubCategories(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to fetch categories"})
	}

	categoryMap := make(map[int32]models.CategoryEn)

	for _, row := range rows {
		if _, exists := categoryMap[row.CategoryEnID]; !exists {
			categoryMap[row.CategoryEnID] = models.CategoryEn{
				CategoryEnID:   row.CategoryEnID,
				CategoryNameEn: row.CategoryNameEn,
			}
		}

		if row.SubCategoryIDEn.Valid {
			subCategory := models.SubCategoryEn{
				SubCategoryIDEn:   row.SubCategoryIDEn.Int32,
				SubCategoryNameEN: row.SubCategoryNameEn.String,
				CategoryEnID:      row.CategoryEnID,
			}

			categoryMap[row.CategoryEnID] = appendToSubCategoriesEn(categoryMap[row.CategoryEnID], subCategory)
		}
	}

	var categoryList []fiber.Map
	for _, category := range categoryMap {
		categoryList = append(categoryList, fiber.Map{
			"categoryEnId":   category.CategoryEnID,
			"categoryNameEn": category.CategoryNameEn,
			"subcategories":  category.SubCategories,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(categoryList)
}

func (hd *Handlers) GetCategoriesWithSubCategoriesMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	rows, err := queries.GetCategoriesWithSubCategoriesMn(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Failed to fetch categories"})
	}

	categoryMap := make(map[int32]models.CategoryMn)

	for _, row := range rows {
		if _, exists := categoryMap[row.CategoryMnID]; !exists {
			categoryMap[row.CategoryMnID] = models.CategoryMn{
				CategoryMnID:   row.CategoryMnID,
				CategoryNameMn: row.CategoryNameMn,
			}
		}

		if row.SubCategoryIDMn.Valid {
			subCategory := models.SubCategoryMn{
				SubCategoryIDMn:   row.SubCategoryIDMn.Int32,
				SubCategoryNameMn: row.SubCategoryNameMn.String,
				CategoryMnID:      row.CategoryMnID,
			}

			categoryMap[row.CategoryMnID] = appendToSubCategoriesMn(categoryMap[row.CategoryMnID], subCategory)
		}
	}

	var categoryList []fiber.Map
	for _, category := range categoryMap {
		categoryList = append(categoryList, fiber.Map{
			"categoryMnId":   category.CategoryMnID,
			"categoryNameMn": category.CategoryNameMn,
			"subcategories":  category.SubCategories,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(categoryList)
}

func appendToSubCategoriesEn(category models.CategoryEn, subCategory models.SubCategoryEn) models.CategoryEn {
	if category.SubCategories == nil {
		category.SubCategories = []models.SubCategoryEn{}
	}
	category.SubCategories = append(category.SubCategories, subCategory)
	return category
}

func appendToSubCategoriesMn(category models.CategoryMn, subcategory models.SubCategoryMn) models.CategoryMn {
	if category.SubCategories == nil {
		category.SubCategories = []models.SubCategoryMn{}
	}
	category.SubCategories = append(category.SubCategories, subcategory)
	return category
}

func (hd *Handlers) GetCategoriesWithSubCategoriesAndProductsEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	// Fetch data from the database
	rows, err := queries.GetCategoriesWithSubCategoriesAndProductsEn(ctx.Context())
	if err != nil {
		slog.Error("Error fetching categories with subcategories and products", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Database error"})
	}

	// Organize data into structured models
	categories := make(map[int32]*models.CategoryWithSubCategoriesAndProductsEn)

	for _, row := range rows {
		// If category does not exist, initialize it
		if _, exists := categories[row.CategoryEnID]; !exists {
			categories[row.CategoryEnID] = &models.CategoryWithSubCategoriesAndProductsEn{
				CategoryEnID:    row.CategoryEnID,
				CategoryNameEn:  row.CategoryNameEn,
				SubcategoriesEn: []models.SubCategoryWithProductsEn{},
			}
		}

		// Find or create the subcategory
		subcategories := categories[row.CategoryEnID].SubcategoriesEn
		var subcategory *models.SubCategoryWithProductsEn
		for i := range subcategories {
			if subcategories[i].SubCategoryIDEn == row.SubCategoryIDEn.Int32 {
				subcategory = &subcategories[i]
				break
			}
		}

		if subcategory == nil {
			categories[row.CategoryEnID].SubcategoriesEn = append(categories[row.CategoryEnID].SubcategoriesEn, models.SubCategoryWithProductsEn{
				SubCategoryIDEn:   row.SubCategoryIDEn.Int32,
				SubCategoryNameEn: row.SubCategoryNameEn.String,
				Products:          []models.ProductEn{},
			})
			subcategory = &categories[row.CategoryEnID].SubcategoriesEn[len(categories[row.CategoryEnID].SubcategoriesEn)-1]
		}

		// Add product to the subcategory
		if row.ProductEnID.Valid {
			subcategory.Products = append(subcategory.Products, models.ProductEn{
				ProductEnID:   row.ProductEnID.Int32,
				ProductNameEn: row.ProductNameEn.String,
				PriceEn:       row.PriceEn.String,
				StockQuantity: row.StockQuantity.Int32,
				ImagesPathEn:  row.ImagesPathEn.String,
			})
		}
	}

	result := make([]*models.CategoryWithSubCategoriesAndProductsEn, 0, len(categories))
	for _, category := range categories {
		result = append(result, category)
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (hd *Handlers) GetCategoriesWithSubCategoriesAndProductsMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	rows, err := queries.GetCategoriesWithSubCategoriesAndProductMn(ctx.Context())
	if err != nil {
		slog.Error("Error fetching categories with subcategories and products", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	categories := make(map[int32]*models.CategoryWithSubCategoriesAndProductsMn)
	for _, row := range rows {
		if _, exists := categories[row.CategoryMnID]; !exists {
			categories[row.CategoryMnID] = &models.CategoryWithSubCategoriesAndProductsMn{
				CategoryMnID:    row.CategoryMnID,
				CategoryNameMn:  row.CategoryNameMn,
				SubCategoriesMn: []models.SubCategoryWithProductsMn{},
			}
		}

		subCategories := categories[row.CategoryMnID].SubCategoriesMn
		var subCategory *models.SubCategoryWithProductsMn
		for i := range subCategories {
			if subCategories[i].SubCategoryIDMn == row.SubCategoryIDMn.Int32 {
				subCategory = &subCategories[i]
				break
			}
		}

		if subCategory == nil {
			categories[row.CategoryMnID].SubCategoriesMn = append(categories[row.CategoryMnID].SubCategoriesMn, models.SubCategoryWithProductsMn{
				SubCategoryIDMn:   row.ProductMnID.Int32,
				SubCategoryNameMn: row.SubCategoryNameMn.String,
				Products:          []models.ProductMn{},
			})
			subCategory = &categories[row.CategoryMnID].SubCategoriesMn[len(categories[row.CategoryMnID].SubCategoriesMn)-1]
		}

		if row.ProductMnID.Valid {
			subCategory.Products = append(subCategory.Products, models.ProductMn{
				ProductMnID:   row.ProductMnID.Int32,
				ProductNameMn: row.ProductNameMn.String,
				PriceMn:       row.PriceMn.String,
				StockQuantity: row.StockQuantity.Int32,
				ImagesPathMn:  row.ImagesPathMn.String,
			})
		}
	}

	result := make([]*models.CategoryWithSubCategoriesAndProductsMn, 0, len(categories))
	for _, cateogory := range categories {
		result = append(result, cateogory)
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (hd *Handlers) FindSubCategoriesAndProductsByCategoryIDEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	// Parse the Category ID from the request
	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		slog.Error("unable to parse category id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid category ID"})
	}

	// Execute the query
	rows, err := queries.FindSubCategoriesAndProductsByCategoryIDEn(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to find subcategories and products", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Database error"})
	}

	// Structure the result
	result := make(map[string]interface{})
	result["CategoryID"] = categoryID
	result["SubCategories"] = []map[string]interface{}{}

	subCategoryMap := make(map[int32]map[string]interface{})

	for _, row := range rows {
		// Group products under each subcategory
		if _, exists := subCategoryMap[row.SubCategoryIDEn.Int32]; !exists {
			subCategoryMap[row.SubCategoryIDEn.Int32] = map[string]interface{}{
				"SubCategoryID":   row.SubCategoryIDEn,
				"SubCategoryName": row.SubCategoryNameEn,
				"Products":        []map[string]interface{}{},
			}
			result["SubCategories"] = append(result["SubCategories"].([]map[string]interface{}), subCategoryMap[row.SubCategoryIDEn.Int32])
		}

		// Add product to subcategory
		if row.ProductEnID.Valid {
			subCategoryMap[row.SubCategoryIDEn.Int32]["Products"] = append(
				subCategoryMap[row.SubCategoryIDEn.Int32]["Products"].([]map[string]interface{}),
				map[string]interface{}{
					"ProductID":     row.ProductEnID.Int32,
					"ProductName":   row.ProductNameEn.String,
					"Price":         row.PriceEn.String,
					"StockQuantity": row.StockQuantity.Int32,
					"ImagesPath":    row.ImagesPathEn.String,
				},
			)
		}
	}

	// Return the structured response
	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (hd *Handlers) FindSubCategoriesAndProductsByCategoryIDMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		slog.Error("unable to parse category id", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	rows, err := queries.FindSubCategoriesAndProductsByCategoryIDMn(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to find subcategories and products", slog.Any("Err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	result := make(map[string]interface{})
	result["CategoryID"] = categoryID
	result["SubCategories"] = []map[string]interface{}{}

	subCategoryMap := make(map[int32]map[string]interface{})

	for _, row := range rows {
		if _, exists := subCategoryMap[row.SubCategoryIDMn.Int32]; !exists {
			subCategoryMap[row.SubCategoryIDMn.Int32] = map[string]interface{}{
				"SubCategoryID":   row.SubCategoryIDMn,
				"SubCategoryName": row.SubCategoryNameMn,
				"Products":        []map[string]interface{}{},
			}
			result["SubCategories"] = append(result["SubCategories"].([]map[string]interface{}), subCategoryMap[row.SubCategoryIDMn.Int32])
		}

		if row.ProductMnID.Valid {
			subCategoryMap[row.SubCategoryIDMn.Int32]["Products"] = append(
				subCategoryMap[row.SubCategoryIDMn.Int32]["Products"].([]map[string]interface{}),
				map[string]interface{}{
					"ProductID":     row.ProductMnID.Int32,
					"ProductName":   row.ProductNameMn.String,
					"Price":         row.PriceMn.String,
					"StockQuantity": row.StockQuantity.Int32,
					"ImagesPath":    row.ImagesPathMn.String,
				},
			)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
