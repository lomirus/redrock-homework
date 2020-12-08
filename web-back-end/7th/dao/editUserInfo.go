package dao

import "errors"

func EditUsername(id int, newUsername string) error {
	_, err := db.Exec("UPDATE `users` SET `username` = ? WHERE `id` = ?", newUsername, id)
	if err != nil {
		return errors.New("failed to update username")
	}
	return nil
}
func EditPassword(id int, newPassword string) error {
	_, err := db.Exec("UPDATE `users` SET `password` = ? WHERE `id` = ?", newPassword, id)
	if err != nil {
		return err
	}
	return nil
}
func EditBio(id int, newBio string) error {
	_, err := db.Exec("UPDATE `users` SET `bio` = ? WHERE `id` = ?", newBio, id)
	if err != nil {
		return err
	}
	return nil
}
