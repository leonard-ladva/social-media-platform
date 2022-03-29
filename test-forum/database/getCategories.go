package database

import (
	"log"
)

func GetCategories() ([]Category, error) {
	var categories []Category
	rows, err := Db.Query(`SELECT * FROM category c`)
	if err != nil {
		log.Println("ERROR | No categories found")
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		rows.Scan(&category.ID, &category.Title, &category.Description, &category.Image)
		category.Checked = ""
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategoriesForPostEdit(postid int) ([]Category, error) {
	var categories []Category
	rows, err := Db.Query(`SELECT c.category_id, c.title, c.description, c.img_link, 
		CASE WHEN pc.postcat_id IS NULL THEN "" ELSE "checked" END as checked
		FROM category c
		LEFT JOIN (SELECT * FROM post_category WHERE post_id = ?) pc ON pc.category_id = c.category_id`, postid)
	if err != nil {
		log.Println("ERROR | No categories found")
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		rows.Scan(&category.ID, &category.Title, &category.Description, &category.Image, &category.Checked)
		categories = append(categories, category)
	}

	return categories, nil

}

func GetCategoryIdByTitle(title string) (int, error) {
	var category_id int
	rows, err := Db.Query(`SELECT c.category_id FROM category c WHERE c.title = ? LIMIT 1`, title)
	if err != nil {
		log.Println("ERROR | Category ID not found for category: ", title)
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&category_id)
	}
	return category_id, nil
}
