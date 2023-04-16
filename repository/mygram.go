package repository

import (
	"MyGram/model"
	"time"
)

type MyGramRepo interface {
	//User Endpoint
	UserRegister(model.User) (res model.User, err error)
	UserLogin(model.User) (res model.User, err error)

	//Photo Endpoint
	PhotoGetAll() (res []model.Photo, err error)
	PhotoGet(model.Photo) (res model.Photo, err error)
	PhotoCreate(model.Photo) (res model.Photo, err error)
	PhotoUpdate(model.Photo) (res model.Photo, err error)
	PhotoDelete(model.Photo) (err error)
	PhotoAuthorization(model.Photo) (res model.Photo, err error)

	// //Comment Endpoint
	CommentGetAll(comment model.Comment) (res []model.Comment, err error)
	CommentGet(model.Comment) (res model.Comment, err error)
	CommentCreate(model.Comment) (res model.Comment, err error)
	CommentUpdate(model.Comment) (res model.Comment, err error)
	CommentDelete(model.Comment) (err error)
	CommentAuthorization(model.Comment) (res model.Comment, err error)

	// //Social Media Endpoint
	SocialMediaGetAll() (res []model.SocialMedia, err error)
	SocialMediaGet(model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaCreate(model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaUpdate(model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaDelete(model.SocialMedia) (err error)
	SocialMediaAuthorization(model.SocialMedia) (res model.SocialMedia, err error)
}

// User Endpoint
func (r Repo) UserRegister(user model.User) (res model.User, err error) {

	err = r.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r Repo) UserLogin(user model.User) (res model.User, err error) {

	err = r.db.Debug().Where("username = ?", user.Username).Take(&res).Error
	if err != nil {
		err = r.db.Debug().Where("email = ?", user.Email).Take(&res).Error
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

// Photo Endpoint
func (r Repo) PhotoGetAll() (res []model.Photo, err error) {
	err = r.db.Model(&model.Photo{}).Preload("Comments").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoGet(photo model.Photo) (res model.Photo, err error) {
	err = r.db.Model(&model.Photo{}).Preload("Comments").First(&res, photo.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoCreate(photo model.Photo) (res model.Photo, err error) {
	err = r.db.Create(&photo).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoUpdate(photo model.Photo) (res model.Photo, err error) {
	err = r.db.First(&res, photo.ID).Error

	if err != nil {
		return res, err
	}

	err = r.db.Model(&res).Where("id = ?", photo.ID).Updates(map[string]interface{}{"title": photo.Title, "caption": photo.Caption, "photo_url": photo.PhotoUrl, "updated_at": time.Now()}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoDelete(photo model.Photo) (err error) {
	err = r.db.First(&photo, photo.ID).Error

	if err != nil {
		return err
	}

	r.db.Delete(&photo, photo.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) PhotoAuthorization(photo model.Photo) (res model.Photo, err error) {
	err = r.db.Select("user_id").First(&res, photo.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// Comment Endpoint
func (r Repo) CommentGetAll(comment model.Comment) (res []model.Comment, err error) {
	err = r.db.Where("photo_id <> ?", comment.PhotoID).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentGet(comment model.Comment) (res model.Comment, err error) {
	err = r.db.Model(&model.Comment{}).First(&res, comment.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentCreate(comment model.Comment) (res model.Comment, err error) {
	err = r.db.Create(&comment).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentUpdate(comment model.Comment) (res model.Comment, err error) {
	err = r.db.First(&res, comment.ID).Error

	if err != nil {
		return res, err
	}

	err = r.db.Model(&res).Where("id = ?", comment.ID).Updates(map[string]interface{}{"message": comment.Message, "updated_at": time.Now()}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentDelete(comment model.Comment) (err error) {
	err = r.db.First(&comment, comment.ID).Error

	if err != nil {
		return err
	}

	r.db.Delete(&comment, comment.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CommentAuthorization(comment model.Comment) (res model.Comment, err error) {
	err = r.db.Select("user_id").First(&res, comment.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

// SocialMedia Endpoint
func (r Repo) SocialMediaGetAll() (res []model.SocialMedia, err error) {
	err = r.db.Model(&model.SocialMedia{}).Preload("User").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaGet(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.Model(&model.SocialMedia{}).Preload("User").First(&res, socialmedia.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaCreate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.Create(&socialmedia).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaUpdate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.First(&res, socialmedia.ID).Error

	if err != nil {
		return res, err
	}

	err = r.db.Model(&res).Where("id = ?", socialmedia.ID).Updates(map[string]interface{}{"name": socialmedia.Name, "social_media_url": socialmedia.SocialMediaUrl, "updated_at": time.Now()}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaDelete(socialmedia model.SocialMedia) (err error) {
	err = r.db.First(&socialmedia, socialmedia.ID).Error

	if err != nil {
		return err
	}

	r.db.Delete(&socialmedia, socialmedia.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) SocialMediaAuthorization(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.Select("user_id").First(&res, socialmedia.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
