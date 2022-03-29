package database

import (
	"log"
)

func GetCommentsByPostAndUserId(postid int, userid int) ([]Comment, error) {

	var comments []Comment
	rows, err := Db.Query(`SELECT c.comment_id,
		c.post_id,
		c.content,
		STRFTIME('%d.%m.%Y', c.datetime) as date,
		u.username, 
		CASE WHEN crl.likes IS NULL THEN 0 
			ELSE crl.likes 
			END as likes,
		CASE WHEN crd.dislikes IS NULL THEN 0 
			ELSE crd.dislikes 
			END as dislikes
	FROM comment c
	LEFT JOIN user u ON u.user_id = c.user_id
	LEFT JOIN (SELECT comment_id, count(reaction_id) as likes FROM comment_reaction WHERE type='like' GROUP BY comment_id) as crl ON c.comment_id = crl.comment_id
	LEFT JOIN (SELECT comment_id, count(reaction_id) as dislikes FROM comment_reaction WHERE type='dislike' GROUP BY comment_id) as crd ON c.comment_id = crd.comment_id
	WHERE c.post_id = ?
	ORDER BY c.datetime ASC`, postid)
	if err != nil {
		log.Println("ERROR | No comments found with post id ", postid)
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		rows.Scan(&comment.ID, &comment.Post_ID, &comment.Content, &comment.Date, &comment.Author, &comment.Likes, &comment.Dislikes)
		comment.Reaction = GetReactionByCommentAndUserID(comment.ID, userid)
		comments = append(comments, comment)
	}
	return comments, nil
}

func GetUserCommentsByPostAndUserId(postid, userid int) ([]Comment, error) {

	var comments []Comment
	rows, err := Db.Query(`SELECT c.comment_id,
		c.post_id,
		c.content,
		STRFTIME('%d.%m.%Y', c.datetime) as date,
		u.username, 
		CASE WHEN crl.likes IS NULL THEN 0 
			ELSE crl.likes 
			END as likes,
		CASE WHEN crd.dislikes IS NULL THEN 0 
			ELSE crd.dislikes 
			END as dislikes
	FROM comment c
	LEFT JOIN user u ON u.user_id = c.user_id
	LEFT JOIN (SELECT comment_id, count(reaction_id) as likes FROM comment_reaction WHERE type='like' GROUP BY comment_id) as crl ON c.comment_id = crl.comment_id
	LEFT JOIN (SELECT comment_id, count(reaction_id) as dislikes FROM comment_reaction WHERE type='dislike' GROUP BY comment_id) as crd ON c.comment_id = crd.comment_id
	WHERE c.post_id = ? AND c.user_id = ?
	ORDER BY c.datetime ASC`, postid, userid)
	if err != nil {
		log.Println("ERROR | No comments found with post id ", postid, "and user id ", userid)
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		rows.Scan(&comment.ID, &comment.Post_ID, &comment.Content, &comment.Date, &comment.Author, &comment.Likes, &comment.Dislikes)
		comment.Reaction = GetReactionByCommentAndUserID(comment.ID, userid)
		comments = append(comments, comment)
	}
	return comments, nil
}

func GetCommentById(commentid int) (Comment, error) {
	var comment Comment
	rows, err := Db.Query(`SELECT c.comment_id,
		c.post_id,
		c.content,
		STRFTIME('%d.%m.%Y', c.datetime) as date,
		u.username, 
		CASE WHEN crl.likes IS NULL THEN 0 
			ELSE crl.likes 
			END as likes,
		CASE WHEN crd.dislikes IS NULL THEN 0 
			ELSE crd.dislikes 
			END as dislikes
	FROM comment c
	LEFT JOIN user u ON u.user_id = c.user_id
	LEFT JOIN (SELECT comment_id, count(reaction_id) as likes FROM comment_reaction WHERE type='like' GROUP BY comment_id) as crl ON c.comment_id = crl.comment_id
	LEFT JOIN (SELECT comment_id, count(reaction_id) as dislikes FROM comment_reaction WHERE type='dislike' GROUP BY comment_id) as crd ON c.comment_id = crd.comment_id
	WHERE c.comment_id = ?
	ORDER BY c.datetime ASC`, commentid)
	if err != nil {
		log.Println("ERROR | No comment found with comment id ", commentid)
		return comment, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&comment.ID, &comment.Post_ID, &comment.Content, &comment.Date, &comment.Author, &comment.Likes, &comment.Dislikes)
		comment.Reaction = ""
	}
	return comment, nil
}

func GetReactionByCommentAndUserID(commentid, userid int) string {
	rows, err := Db.Query(`SELECT type
	FROM comment_reaction WHERE user_id = ? AND comment_id = ? `, userid, commentid)
	if err != nil {
		log.Println("ERROR | No likes found with userid ", userid, " and commentid ", commentid)
		return ""
	}
	defer rows.Close()

	var reaction string
	for rows.Next() {
		rows.Scan(&reaction)
	}
	return reaction
}

func GetCommentAuthorByID(commentid int) string {
	rows, err := Db.Query(`SELECT u.username
	FROM comment c, user u WHERE c.user_id=u.user_id AND comment_id = ? LIMIT 1`, commentid)
	if err != nil {
		log.Println("ERROR | No author name found for commentid ", commentid)
		return ""
	}
	defer rows.Close()

	var author string
	for rows.Next() {
		rows.Scan(&author)
	}
	return author
}
