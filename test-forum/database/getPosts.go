package database

import (
	"log"
)

func GetPostsByCategoryId(id int) ([]Post, error) {

	var posts []Post
	rows, err := Db.Query(`SELECT p.post_id,
	p.title,
	p.content,
	CASE WHEN p.postimg IS NULL THEN "" 
	ELSE p.postimg END AS postimg,
	STRFTIME('%d.%m.%Y', p.datetime) as date,
	u.username as author,
	CASE WHEN prl.likes IS NULL THEN 0 
		ELSE prl.likes 
		END as likes,
	CASE WHEN prd.dislikes IS NULL THEN 0 
		ELSE prd.dislikes 
		END as dislikes,
	CASE WHEN c.comments IS NULL THEN 0 
		ELSE c.comments
		END as comments,
	p.status
	FROM post as p
	LEFT JOIN user as u ON u.user_id = p.user_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as likes FROM post_reaction WHERE type='like' GROUP BY post_id) as prl ON p.post_id = prl.post_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as dislikes FROM post_reaction WHERE type='dislike' GROUP BY post_id) as prd ON p.post_id = prd.post_id
	LEFT JOIN (SELECT post_id, count(comment_id) as comments FROM comment GROUP BY post_id) as c ON p.post_id = c.post_id
	LEFT JOIN post_category as pc ON p.post_id = pc.post_id
	LEFT JOIN category as ca ON pc.category_id = ca.category_id
	WHERE ca.category_id = ?`, id)
	if err != nil {
		log.Println("ERROR | No post found with category id ", id)
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Date, &post.Author, &post.Likes, &post.Dislikes, &post.Comments, &post.Status)
		post.Categories, _ = GetCategoriesByPostId(post.ID)
		posts = append(posts, post)
	}

	return posts, nil

}

func GetPostByPostAndUsedID(postid int, userid int) (Post, error) {

	var post Post

	rows, err := Db.Query(`SELECT p.post_id,
	p.title,
	p.content,
	CASE WHEN p.postimg IS NULL THEN "" 
	ELSE p.postimg END AS postimg,
	STRFTIME('%d.%m.%Y', p.datetime) as date,
	u.username as author,
	CASE WHEN prl.likes IS NULL THEN 0 
		ELSE prl.likes 
		END as likes,
	CASE WHEN prd.dislikes IS NULL THEN 0 
		ELSE prd.dislikes 
		END as dislikes,
	CASE WHEN c.comments IS NULL THEN 0 
		ELSE c.comments
		END as comments,
	p.status
	FROM post as p
	LEFT JOIN user as u ON u.user_id = p.user_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as likes FROM post_reaction WHERE type='like' GROUP BY post_id) as prl ON p.post_id = prl.post_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as dislikes FROM post_reaction WHERE type='dislike' GROUP BY post_id) as prd ON p.post_id = prd.post_id
	LEFT JOIN (SELECT post_id, count(comment_id) as comments FROM comment GROUP BY post_id) as c ON p.post_id = c.post_id
	WHERE p.post_id = ?`, postid)
	if err != nil {
		log.Println("ERROR | No post found with id ", postid)
		return post, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Date, &post.Author, &post.Likes, &post.Dislikes, &post.Comments, &post.Status)
		post.Categories, _ = GetCategoriesByPostId(postid)
		post.Reaction = GetReactionByPostAndUserID(postid, userid)
	}

	return post, nil

}

func GetAllPosts() ([]Post, error) {

	var posts []Post
	rows, err := Db.Query(`SELECT DISTINCT p.post_id,
	p.title,
	p.content,
	CASE WHEN p.postimg IS NULL THEN "" 
	ELSE p.postimg END AS postimg,
	STRFTIME('%d.%m.%Y', p.datetime) as date,
	u.username as author,
	CASE WHEN prl.likes IS NULL THEN 0 
		ELSE prl.likes 
		END as likes,
	CASE WHEN prd.dislikes IS NULL THEN 0 
		ELSE prd.dislikes 
		END as dislikes,
	CASE WHEN c.comments IS NULL THEN 0 
		ELSE c.comments
		END as comments,
	p.status
	FROM post as p
	LEFT JOIN user as u ON u.user_id = p.user_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as likes FROM post_reaction WHERE type='like' GROUP BY post_id) as prl ON p.post_id = prl.post_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as dislikes FROM post_reaction WHERE type='dislike' GROUP BY post_id) as prd ON p.post_id = prd.post_id
	LEFT JOIN (SELECT post_id, count(comment_id) as comments FROM comment GROUP BY post_id) as c ON p.post_id = c.post_id`)
	if err != nil {
		log.Println("Error getting categories")
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Date, &post.Author, &post.Likes, &post.Dislikes, &post.Comments, &post.Status)
		post.Categories, _ = GetCategoriesByPostId(post.ID)
		posts = append(posts, post)
	}

	return posts, nil

}

func GetPostsByUserID(id int) ([]Post, error) {

	var posts []Post

	rows, err := Db.Query(`SELECT p.post_id,
	p.title,
	p.content,
	CASE WHEN p.postimg IS NULL THEN "" 
	ELSE p.postimg END AS postimg,
	STRFTIME('%d.%m.%Y', p.datetime) as date,
	u.username as author,
	CASE WHEN prl.likes IS NULL THEN 0 
		ELSE prl.likes 
		END as likes,
	CASE WHEN prd.dislikes IS NULL THEN 0 
		ELSE prd.dislikes 
		END as dislikes,
	CASE WHEN c.comments IS NULL THEN 0 
		ELSE c.comments
		END as comments,
	p.status
	FROM post as p
	LEFT JOIN user as u ON u.user_id = p.user_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as likes FROM post_reaction WHERE type='like' GROUP BY post_id) as prl ON p.post_id = prl.post_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as dislikes FROM post_reaction WHERE type='dislike' GROUP BY post_id) as prd ON p.post_id = prd.post_id
	LEFT JOIN (SELECT post_id, count(comment_id) as comments FROM comment GROUP BY post_id) as c ON p.post_id = c.post_id
	WHERE u.user_id = ?`, id)
	if err != nil {
		log.Println("ERROR | No post found with category id ", id)
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Date, &post.Author, &post.Likes, &post.Dislikes, &post.Comments, &post.Status)
		post.Categories, _ = GetCategoriesByPostId(post.ID)
		posts = append(posts, post)
	}

	return posts, nil

}

func GetLikedPostsByUserID(userid int) ([]Post, error) {

	var posts []Post

	rows, err := Db.Query(`SELECT DISTINCT p.post_id,
	p.title,
	p.content,
	CASE WHEN p.postimg IS NULL THEN "" 
	ELSE p.postimg END AS postimg,
	STRFTIME('%d.%m.%Y', p.datetime) as date,
	u.username as author,
	CASE WHEN prl.likes IS NULL THEN 0 
		ELSE prl.likes 
		END as likes,
	CASE WHEN prd.dislikes IS NULL THEN 0 
		ELSE prd.dislikes 
		END as dislikes,
	CASE WHEN c.comments IS NULL THEN 0 
		ELSE c.comments
		END as comments,
	p.status
	FROM post as p
	LEFT JOIN user as u ON u.user_id = p.user_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as likes FROM post_reaction WHERE type='like' GROUP BY post_id) as prl ON p.post_id = prl.post_id
	LEFT JOIN (SELECT post_id, count(reaction_id) as dislikes FROM post_reaction WHERE type='dislike' GROUP BY post_id) as prd ON p.post_id = prd.post_id
	LEFT JOIN (SELECT post_id, count(comment_id) as comments FROM comment GROUP BY post_id) as c ON p.post_id = c.post_id
	INNER JOIN (SELECT post_id FROM post_reaction WHERE user_id = ?
				UNION
				SELECT post_id FROM comment WHERE user_id = ?) as cr ON p.post_id = cr.post_id`, userid, userid)
	if err != nil {
		log.Println("ERROR | No post found liked/commented by user id ", userid)
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Date, &post.Author, &post.Likes, &post.Dislikes, &post.Comments, &post.Status)
		post.Categories, _ = GetCategoriesByPostId(post.ID)
		post.Reaction = GetReactionByPostAndUserID(post.ID, userid)
		post.UserComments, _ = GetUserCommentsByPostAndUserId(post.ID, userid)
		posts = append(posts, post)
	}

	return posts, nil

}

func GetReactionByPostAndUserID(postid, userid int) string {

	rows, err := Db.Query(`SELECT type
	FROM post_reaction WHERE user_id = ? AND post_id = ? `, userid, postid)
	if err != nil {
		log.Println("ERROR | No likes found with userid ", userid)
		return ""
	}
	defer rows.Close()

	var reaction string
	for rows.Next() {
		rows.Scan(&reaction)
	}
	return reaction

}

func GetCategoriesByPostId(id int) ([]string, error) {
	var categories []string
	rows, err := Db.Query(`SELECT c.title
		FROM category c,
			post_category pc
		WHERE c.category_id = pc.category_id
			AND pc.post_id = ?`, id)
	if err != nil {
		log.Println("Error getting categories")
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		rows.Scan(&category)
		categories = append(categories, category)
	}
	return categories, nil
}
