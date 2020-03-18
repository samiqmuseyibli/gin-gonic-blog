package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Wishlist struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	PostID    uint64    `gorm:"size:100;not null" json:"post_id"`
	Author    User      `json:"author"`
	UserID    uint32    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (w *Wishlist) Prepare() {
	w.UserID = uint32(w.UserID)
	w.PostID = uint64(w.PostID)
	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()
}

func (w *Wishlist) AddToWishlist(db *gorm.DB) (*Wishlist, error) {
	var err error
	err = db.Debug().Model(&Wishlist{}).Create(&w).Error
	if err != nil {
		return &Wishlist{}, err
	}
	if w.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", w.UserID).Take(&w.Author).Error
		if err != nil {
			return &Wishlist{}, err
		}
	}
	return w, nil
}

func (w *Wishlist) Validate() map[string]string {

	var err error

	var errorMessages = make(map[string]string)

	if w.PostID < 1 {
		err = errors.New("Required Post Id")
		errorMessages["Required_post"] = err.Error()

	}
	if w.UserID < 1 {
		err = errors.New("Required valid User Id")
		errorMessages["Required_author"] = err.Error()
	}
	return errorMessages
}
