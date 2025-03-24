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
	CategoryIDSTR := ctx.Params("id")
	CategoryId, err := strconv.Atoi(CategoryIDSTR)
	if err != nil {
		slog.Error("unable to parse id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err})
	}

	var request models.UpdateCategoryEn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	_, err = queries.FindByCategoryEnId(ctx.Context(), int32(CategoryId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}

	updatedCategory, err := queries.UpdateCategoryEn(ctx.Context(), db.UpdateCategoryEnParams{
		CategoryNameEn: request.CategoryNameEn,
		CategoryEnID:   int32(CategoryId),
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
	CategoryID := ctx.Params("id")
	CategoryId, err := strconv.Atoi(CategoryID)

	var request models.UpdateCategoryMn
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid request"})
	}

	_, err = queries.FindByCategoryMnId(ctx.Context(), int32(CategoryId))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "category not found"})
	}

	updatedCategory, err := queries.UpdateCategoryMn(ctx.Context(), db.UpdateCategoryMnParams{
		CategoryNameMn: request.CategoryNameMn,
		CategoryMnID:   int32(CategoryId),
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

	categoryMap := make(map[int32]*models.CategoryEn)

	for _, row := range rows {
		cat, exists := categoryMap[row.CategoryEnID]
		if !exists {
			cat = &models.CategoryEn{
				CategoryEnID:   row.CategoryEnID,
				CategoryNameEn: row.CategoryNameEn,
				SubCategories:  []models.SubCategoryEn{},
			}
			categoryMap[row.CategoryEnID] = cat
		}

		if row.SubCategoryIDEn.Valid {
			var subCat *models.SubCategoryEn
			for i, sc := range cat.SubCategories {
				if sc.SubCategoryIDEn == row.SubCategoryIDEn.Int32 {
					subCat = &cat.SubCategories[i]
					break
				}
			}
			if subCat == nil {
				newSubCat := models.SubCategoryEn{
					SubCategoryIDEn:   row.SubCategoryIDEn.Int32,
					SubCategoryNameEN: row.SubCategoryNameEn.String,
					CategoryEnID:      row.CategoryEnID,
					SCategories:       []models.SCategoryEn{},
				}
				cat.SubCategories = append(cat.SubCategories, newSubCat)
				subCat = &cat.SubCategories[len(cat.SubCategories)-1]
			}
			if row.SCategoryIdEn.Valid {
				newSCategory := models.SCategoryEn{
					SCategoryIdEn:   row.SCategoryIdEn.Int32,
					SCategoryNameEn: row.SCategoryNameEn.String,
					SubCategoryIDEn: row.SubCategoryIDEn.Int32,
				}
				subCat.SCategories = append(subCat.SCategories, newSCategory)
			}
		}
	}

	// Build a list for the response.
	var categoryList []fiber.Map
	for _, cat := range categoryMap {
		categoryList = append(categoryList, fiber.Map{
			"categoryEnId":   cat.CategoryEnID,
			"categoryNameEn": cat.CategoryNameEn,
			"subcategories":  cat.SubCategories,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(categoryList)
}

func (hd *Handlers) GetCategoriesWithSubCategoriesMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	rows, err := queries.GetCategoriesWithSubCategoriesMn(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"err": "Failed to fetch categories"})
	}

	categoryMap := make(map[int32]*models.CategoryMn)

	for _, row := range rows {
		cat, exists := categoryMap[row.CategoryMnID]
		if !exists {
			cat = &models.CategoryMn{
				CategoryMnID:   row.CategoryMnID,
				CategoryNameMn: row.CategoryNameMn,
				SubCategories:  []models.SubCategoryMn{},
			}
			categoryMap[row.CategoryMnID] = cat
		}

		if row.SubCategoryIDMn.Valid {
			var subCat *models.SubCategoryMn = nil
			for i := range cat.SubCategories {
				if cat.SubCategories[i].SubCategoryIDMn == row.SubCategoryIDMn.Int32 {
					subCat = &cat.SubCategories[i]
					break
				}
			}
			if subCat == nil {
				newSubCat := models.SubCategoryMn{
					SubCategoryIDMn:   row.SubCategoryIDMn.Int32,
					SubCategoryNameMn: row.SubCategoryNameMn.String,
					CategoryMnID:      row.CategoryMnID,
					SCategories:       []models.SCategoryMn{},
				}
				cat.SubCategories = append(cat.SubCategories, newSubCat)
				subCat = &cat.SubCategories[len(cat.SubCategories)-1]
			}
			// If the row also contains an sCategory, append it.
			if row.SCategoryIdMn.Valid {
				newSCategory := models.SCategoryMn{
					SCategoryIdMn:   row.SCategoryIdMn.Int32,
					SCategoryNameMn: row.SCategoryNameMn.String,
					SubCategoryIDMn: row.SubCategoryIDMn.Int32,
				}
				subCat.SCategories = append(subCat.SCategories, newSCategory)
			}
		}
	}

	// Build the response list.
	var categoryList []fiber.Map
	for _, cat := range categoryMap {
		categoryList = append(categoryList, fiber.Map{
			"categoryMnId":   cat.CategoryMnID,
			"categoryNameMn": cat.CategoryNameMn,
			"subcategories":  cat.SubCategories,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(categoryList)
}

// func appendToSubCategoriesEn(category models.CategoryEn, subCategory models.SubCategoryEn) models.CategoryEn {
// 	if category.SubCategories == nil {
// 		category.SubCategories = []models.SubCategoryEn{}
// 	}
// 	category.SubCategories = append(category.SubCategories, subCategory)
// 	return category
// }

// func appendToSubCategoriesMn(category models.CategoryMn, subcategory models.SubCategoryMn) models.CategoryMn {
// 	if category.SubCategories == nil {
// 		category.SubCategories = []models.SubCategoryMn{}
// 	}
// 	category.SubCategories = append(category.SubCategories, subcategory)
// 	return category
// }

func (hd *Handlers) GetCategoriesWithSubCategoriesAndProductsEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	rows, err := queries.GetCategoriesWithSubCategoriesAndProductsEn(ctx.Context())
	if err != nil {
		slog.Error("Error fetching categories with subcategories and products", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": "Database error"})
	}

	// Use a map to group categories by ID.
	categoryMap := make(map[int32]*models.CategoryWithSubCategoriesAndProductsEn)
	for _, row := range rows {
		// Get or create the category.
		cat, exists := categoryMap[row.CategoryEnID]
		if !exists {
			cat = &models.CategoryWithSubCategoriesAndProductsEn{
				CategoryEnID:    row.CategoryEnID,
				CategoryNameEn:  row.CategoryNameEn,
				SubcategoriesEn: []models.SubCategoryWithProductsEn{},
			}
			categoryMap[row.CategoryEnID] = cat
		}

		// Process subcategory if present.
		if row.SubCategoryIDEn.Valid {
			var subcat *models.SubCategoryWithProductsEn
			// Search for existing subcategory in this category.
			for i := range cat.SubcategoriesEn {
				if cat.SubcategoriesEn[i].SubCategoryIDEn == row.SubCategoryIDEn.Int32 {
					subcat = &cat.SubcategoriesEn[i]
					break
				}
			}
			// If not found, create a new subcategory.
			if subcat == nil {
				newSubCat := models.SubCategoryWithProductsEn{
					SubCategoryIDEn:   row.SubCategoryIDEn.Int32,
					SubCategoryNameEn: row.SubCategoryNameEn.String,
					// Initialize SCategories with an empty slice.
					SCategories: []models.SCategoryEn{},
				}
				cat.SubcategoriesEn = append(cat.SubcategoriesEn, newSubCat)
				subcat = &cat.SubcategoriesEn[len(cat.SubcategoriesEn)-1]
			}

			// Now, if an sCategory is present in this row, nest the product under that sCategory.
			if row.SCategoryIdEn.Valid {
				var sCatPtr *models.SCategoryEn
				// Search for the sCategory within the subcategory.
				for i := range subcat.SCategories {
					if subcat.SCategories[i].SCategoryIdEn == row.SCategoryIdEn.Int32 {
						sCatPtr = &subcat.SCategories[i]
						break
					}
				}
				// If not found, create it.
				if sCatPtr == nil {
					newSCategory := models.SCategoryEn{
						SCategoryIdEn:   row.SCategoryIdEn.Int32,
						SCategoryNameEn: row.SCategoryNameEn.String,
						SubCategoryIDEn: row.SubCategoryIDEn.Int32,
						Products:        []models.ProductEn{},
					}
					subcat.SCategories = append(subcat.SCategories, newSCategory)
					sCatPtr = &subcat.SCategories[len(subcat.SCategories)-1]
				}
				// If a product is present, add it to the sCategory's Products.
				if row.ProductEnID.Valid {
					newProduct := models.ProductEn{
						ProductEnID:   row.ProductEnID.Int32,
						ProductNameEn: row.ProductNameEn.String,
						PriceEn:       row.PriceEn.String,
						StockQuantity: row.StockQuantity.Int32,
						ImagesPathEn:  row.ImagesPathEn,
					}
					sCatPtr.Products = append(sCatPtr.Products, newProduct)
				}
			} else {
				// Optionally, if no sCategory is provided but a product exists,
				// you might want to add it directly under the subcategory.
				// For example:
				// if row.ProductEnID.Valid {
				//     newProduct := models.ProductEn{...}
				//     subcat.Products = append(subcat.Products, newProduct)
				// }
			}
		}
	}

	// Convert categoryMap to a slice.
	var result []*models.CategoryWithSubCategoriesAndProductsEn
	for _, cat := range categoryMap {
		result = append(result, cat)
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

	// Create a map to group categories by CategoryMnID.
	categoryMap := make(map[int32]*models.CategoryWithSubCategoriesAndProductsMn)

	for _, row := range rows {
		// Get or create the category.
		cat, exists := categoryMap[row.CategoryMnID]
		if !exists {
			cat = &models.CategoryWithSubCategoriesAndProductsMn{
				CategoryMnID:    row.CategoryMnID,
				CategoryNameMn:  row.CategoryNameMn,
				SubCategoriesMn: []models.SubCategoryWithProductsMn{},
			}
			categoryMap[row.CategoryMnID] = cat
		}

		// Process subcategory if present.
		if row.SubCategoryIDMn.Valid {
			var subcat *models.SubCategoryWithProductsMn
			// Look for an existing subcategory within the category.
			for i := range cat.SubCategoriesMn {
				if cat.SubCategoriesMn[i].SubCategoryIDMn == row.SubCategoryIDMn.Int32 {
					subcat = &cat.SubCategoriesMn[i]
					break
				}
			}
			// If not found, create a new subcategory using the subcategory columns.
			if subcat == nil {
				newSubCat := models.SubCategoryWithProductsMn{
					SubCategoryIDMn:   row.SubCategoryIDMn.Int32,
					SubCategoryNameMn: row.SubCategoryNameMn.String,
					SCategories:       []models.SCategoryMn{},
				}
				cat.SubCategoriesMn = append(cat.SubCategoriesMn, newSubCat)
				subcat = &cat.SubCategoriesMn[len(cat.SubCategoriesMn)-1]
			}

			// Process sCategory if present.
			if row.SCategoryIdMn.Valid {
				var sCatPtr *models.SCategoryMn
				// Look for an existing sCategory within this subcategory.
				for i := range subcat.SCategories {
					if subcat.SCategories[i].SCategoryIdMn == row.SCategoryIdMn.Int32 {
						sCatPtr = &subcat.SCategories[i]
						break
					}
				}
				// If not found, create a new sCategory using the sCategory columns.
				if sCatPtr == nil {
					newSCategory := models.SCategoryMn{
						SCategoryIdMn:   row.SCategoryIdMn.Int32,
						SCategoryNameMn: row.SCategoryNameMn.String,
						SubCategoryIDMn: row.SubCategoryIDMn.Int32,
						Products:        []models.ProductMn{},
					}
					subcat.SCategories = append(subcat.SCategories, newSCategory)
					sCatPtr = &subcat.SCategories[len(subcat.SCategories)-1]
				}

				// If a product is present, add it to the sCategory's Products slice.
				if row.ProductMnID.Valid {
					newProduct := models.ProductMn{
						ProductMnID:   row.ProductMnID.Int32,
						ProductNameMn: row.ProductNameMn.String,
						PriceMn:       row.PriceMn.String,
						StockQuantity: row.StockQuantity.Int32,
						ImagesPathMn:  row.ImagesPathMn,
					}
					sCatPtr.Products = append(sCatPtr.Products, newProduct)
				}
			}
		}
	}

	// Convert the category map to a slice.
	var result []*models.CategoryWithSubCategoriesAndProductsMn
	for _, cat := range categoryMap {
		result = append(result, cat)
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (hd *Handlers) FindSubCategoriesAndProductsByCategoryIDEn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	// Parse the Category ID from the URL parameter.
	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		slog.Error("unable to parse category id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid category ID"})
	}

	// Execute the SQLC query.
	rows, err := queries.FindSubCategoriesAndProductsByCategoryIDEn(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to find subcategories and products", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	// Prepare the result map.
	result := make(map[string]interface{})
	result["CategoryID"] = categoryID
	result["SubCategories"] = []map[string]interface{}{}

	// Map to group subcategories by their ID.
	subCategoryMap := make(map[int32]map[string]interface{})

	for _, row := range rows {
		// Process only if the row has a valid subcategory.
		if row.SubCategoryIDEn.Valid {
			subcatID := row.SubCategoryIDEn.Int32

			// If this subcategory isn't already in our map, create a new entry.
			if _, exists := subCategoryMap[subcatID]; !exists {
				subCategoryMap[subcatID] = map[string]interface{}{
					"SubCategoryID":   subcatID,
					"SubCategoryName": row.SubCategoryNameEn.String,
					"SCategories":     []map[string]interface{}{},
				}
				// Append the new subcategory to the result.
				result["SubCategories"] = append(result["SubCategories"].([]map[string]interface{}), subCategoryMap[subcatID])
			}

			// Process the sCategory (third level) if present.
			if row.SCategoryIdEn.Valid {
				sCatID := row.SCategoryIdEn.Int32

				// Get the current list of sCategories under this subcategory.
				sCategoriesSlice := subCategoryMap[subcatID]["SCategories"].([]map[string]interface{})
				var sCategoryFound map[string]interface{}
				for _, sCat := range sCategoriesSlice {
					if sCat["SCategoryIdEn"].(int32) == sCatID {
						sCategoryFound = sCat
						break
					}
				}
				// If no sCategory exists, create a new one.
				if sCategoryFound == nil {
					sCategoryFound = map[string]interface{}{
						"SCategoryIdEn":   sCatID,
						"SCategoryNameEn": row.SCategoryNameEn.String,
						"Products":        []map[string]interface{}{},
					}
					subCategoryMap[subcatID]["SCategories"] = append(sCategoriesSlice, sCategoryFound)
				}
				// Process the product if present.
				if row.ProductEnID.Valid {
					newProduct := map[string]interface{}{
						"ProductEnID":   row.ProductEnID.Int32,
						"ProductName":   row.ProductNameEn.String,
						"Price":         row.PriceEn.String,
						"StockQuantity": row.StockQuantity.Int32,
						"ImagesPath":    row.ImagesPathEn,
					}
					// Append the product to the sCategory's Products array.
					sCategoryFound["Products"] = append(sCategoryFound["Products"].([]map[string]interface{}), newProduct)
				}
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (hd *Handlers) FindSubCategoriesAndProductsByCategoryIDMn(ctx *fiber.Ctx) error {
	queries, _, _ := hd.queries()

	// Parse the Category ID from the URL parameter.
	categoryIDStr := ctx.Params("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		slog.Error("unable to parse category id", slog.Any("err", err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Invalid category ID"})
	}

	// Execute the SQLC query.
	rows, err := queries.FindSubCategoriesAndProductsByCategoryIDMn(ctx.Context(), int32(categoryID))
	if err != nil {
		slog.Error("unable to find subcategories and products", slog.Any("err", err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err})
	}

	// Prepare the result map.
	result := make(map[string]interface{})
	result["CategoryID"] = categoryID
	result["SubCategories"] = []map[string]interface{}{}

	// Map to group subcategories by their ID.
	subCategoryMap := make(map[int32]map[string]interface{})

	for _, row := range rows {
		// Process only if the row has a valid subcategory.
		if row.SubCategoryIDMn.Valid {
			subcatID := row.SubCategoryIDMn.Int32

			// If this subcategory isn't already in our map, create a new entry.
			if _, exists := subCategoryMap[subcatID]; !exists {
				subCategoryMap[subcatID] = map[string]interface{}{
					"SubCategoryID":   subcatID,
					"SubCategoryName": row.SubCategoryNameMn.String,
					"SCategories":     []map[string]interface{}{},
				}
				// Append the new subcategory to the result.
				result["SubCategories"] = append(result["SubCategories"].([]map[string]interface{}), subCategoryMap[subcatID])
			}

			// Process the sCategory (third level) if present.
			if row.SCategoryIdMn.Valid {
				sCatID := row.SCategoryIdMn.Int32

				// Get the current list of sCategories under this subcategory.
				sCategoriesSlice := subCategoryMap[subcatID]["SCategories"].([]map[string]interface{})
				var sCategoryFound map[string]interface{}
				for _, sCat := range sCategoriesSlice {
					if sCat["SCategoryIdMn"].(int32) == sCatID {
						sCategoryFound = sCat
						break
					}
				}
				// If no sCategory exists, create a new one.
				if sCategoryFound == nil {
					sCategoryFound = map[string]interface{}{
						"SCategoryIdMn":   sCatID,
						"SCategoryNameMn": row.SCategoryNameMn.String,
						"Products":        []map[string]interface{}{},
					}
					subCategoryMap[subcatID]["SCategories"] = append(sCategoriesSlice, sCategoryFound)
				}
				// Process the product if present.
				if row.ProductMnID.Valid {
					newProduct := map[string]interface{}{
						"ProductMnID":   row.ProductMnID.Int32,
						"ProductName":   row.ProductNameMn.String,
						"Price":         row.PriceMn.String,
						"StockQuantity": row.StockQuantity.Int32,
						"ImagesPath":    row.ImagesPathMn,
					}
					// Append the product to the sCategory's Products array.
					sCategoryFound["Products"] = append(sCategoryFound["Products"].([]map[string]interface{}), newProduct)
				}
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
