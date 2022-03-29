package database

import (
	"errors"
	"log"
	"time"
)

func GetNotificationsByUserId(userid int) ([]Notification, error) {
	//Returning top 5 notifications, ordered by priority, first unread and then read ones.
	var notifications []Notification
	rows, err := Db.Query(`SELECT notification_id, user_id, content, link,
		priority, STRFTIME('%d.%m.%Y', datetime) as date, read
	FROM notification WHERE user_id = ? AND priority >= 0 AND read = 0
	UNION 
	SELECT notification_id, user_id, content, link,
		priority, STRFTIME('%d.%m.%Y', datetime) as date, read
	FROM notification WHERE user_id = ? AND priority >= 0 AND read > 0
	ORDER BY read asc, priority desc`, userid, userid)

	if err != nil {
		log.Println("ERROR | No notifications found with userid ", userid)
		return notifications, err
	}
	defer rows.Close()

	for rows.Next() {
		if len(notifications) < 5 {
			var notification Notification
			rows.Scan(&notification.ID, &notification.User_ID, &notification.Content, &notification.Link, &notification.Priority, &notification.Date, &notification.Read)
			notifications = append(notifications, notification)
		}
	}
	return notifications, nil
}

func AddNotificationForRole(notification, link string, priority int, role string) error {
	var userlist []int
	rows, err := Db.Query(`SELECT u.user_id FROM user u, role r
			WHERE r.role_id = u.role_id AND r.role = ?`, role)
	if err != nil {
		log.Println("ERROR | Impossible to retrieve userlist with role: ", role)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		rows.Scan(&id)
		userlist = append(userlist, id)
	}

	if len(userlist) == 0 {
		log.Println("ERROR | No users found with role", role)
		return errors.New("ERROR | No users found with role " + role)
	}
	for i := range userlist {
		err := AddNotificationForUser(notification, link, priority, userlist[i])
		if err != nil {
			log.Println("ERROR | Adding notification failed, user_id=", userlist[i])
		}
	}

	return nil
}

func AddNotificationForUser(notification, link string, priority, userid int) error {
	statement, err := Db.Prepare("INSERT INTO notification (user_id, content, link, priority, datetime, read) VALUES (?, ?, ?, ?, ?, ?);")
	if err != nil {
		log.Println("ERROR | Adding notification failed")
		return err
	}
	defer statement.Close()
	statement.Exec(userid, notification, link, priority, time.Now(), 0)
	return nil
}
