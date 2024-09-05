package database

import (
	"database/sql"
)

// AddBan inserisce un nuovo record di ban nel database.
func (db *appdbimpl) AddBan(b Ban) (Ban, error) {
	_, err := db.c.Exec(`INSERT INTO bans (banId, bannedId, userId) VALUES (?, ?, ?)`, b.BanId, b.BannedId, b.UserId)
	if err != nil {
		return b, err
	}
	return b, nil
}

// RemoveBan rimuove un record di ban esistente dal database.
func (db *appdbimpl) RemoveBan(b Ban) error {
	res, err := db.c.Exec(`DELETE FROM bans WHERE banId=? AND userId=? AND bannedId=?`, b.BanId, b.UserId, b.BannedId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrBanNotFound
	}
	return nil
}

// UpdateBanStatus aggiorna lo stato di un ban esistente.
func (db *appdbimpl) UpdateBanStatus(status int, followerId uint64, userId uint64) error {
	res, err := db.c.Exec(`UPDATE followers SET banStatus=? WHERE followerId=? AND userId=?`, status, followerId, userId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrBanNotFound
	}
	return nil
}

// CheckIfBanned verifica se un utente è stato bannato.
func (db *appdbimpl) CheckIfBanned(r uint64, u uint64) (bool, error) {
	var res bool
	if err := db.c.QueryRow(`SELECT EXISTS(SELECT 1 FROM bans WHERE bannedId = ? AND userId = ?)`, r, u).Scan(&res); err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
		return false, err
	}
	return res, nil
}

// GetBan recupera un record di ban basato sugli ID utente e token.
func (db *appdbimpl) GetBan(u User, token uint64) (Ban, error) {
	var ban Ban
	if err := db.c.QueryRow(`SELECT banId, bannedId, userId FROM bans WHERE bannedId = ? AND userId = ?`, u.Id, token).Scan(&ban.BanId, &ban.BannedId, &ban.UserId); err != nil {
		if err == sql.ErrNoRows {
			return ban, ErrBanNotFound
		}
		return ban, err
	}
	return ban, nil
}

// GetBanById recupera un record di ban basato sull'ID del ban.
func (db *appdbimpl) GetBanById(b Ban) (Ban, error) {
	var ban Ban
	if err := db.c.QueryRow(`SELECT banId, bannedId, userId FROM bans WHERE banId = ?`, b.BanId).Scan(&ban.BanId, &ban.BannedId, &ban.UserId); err != nil {
		if err == sql.ErrNoRows {
			return ban, ErrBanNotFound
		}
		return ban, err
	}
	return ban, nil
}
