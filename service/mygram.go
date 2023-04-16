package service

import (
	"MyGram/model"
)

type MyGramService interface {
	UserRegister(user model.User) (res model.User, err error)
	UserLogin(user model.User) (res model.User, err error)

	//Photo Endpoint
	PhotoGetAll() (res []model.Photo, err error)
	PhotoGet(photo model.Photo) (res model.Photo, err error)
	PhotoCreate(photo model.Photo) (res model.Photo, err error)
	PhotoUpdate(photo model.Photo) (res model.Photo, err error)
	PhotoDelete(photo model.Photo) (err error)
	PhotoAuthorization(photo model.Photo) (res model.Photo, err error)

	//Comment Endpoint
	CommentGetAll(comment model.Comment) (res []model.Comment, err error)
	CommentGet(comment model.Comment) (res model.Comment, err error)
	CommentCreate(comment model.Comment) (res model.Comment, err error)
	CommentUpdate(comment model.Comment) (res model.Comment, err error)
	CommentDelete(comment model.Comment) (err error)
	CommentAuthorization(comment model.Comment) (res model.Comment, err error)

	//Social Media Endpoint
	SocialMediaGetAll() (res []model.SocialMedia, err error)
	SocialMediaGet(socialmedia model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaCreate(socialmedia model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaUpdate(socialmedia model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaDelete(socialmedia model.SocialMedia) (err error)
	SocialMediaAuthorization(model.SocialMedia) (res model.SocialMedia, err error)
}

// User Endpoint
func (s *Service) UserRegister(user model.User) (res model.User, err error) {
	res, err = s.repo.UserRegister(user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) UserLogin(user model.User) (res model.User, err error) {
	res, err = s.repo.UserLogin(user)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Photo Endpoint
func (s *Service) PhotoGetAll() (res []model.Photo, err error) {
	res, err = s.repo.PhotoGetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoGet(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoGet(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoCreate(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoCreate(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoUpdate(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoUpdate(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoDelete(photo model.Photo) (err error) {
	err = s.repo.PhotoDelete(photo)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) PhotoAuthorization(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoAuthorization(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Comment Endpoint
func (s *Service) CommentGetAll(comment model.Comment) (res []model.Comment, err error) {
	res, err = s.repo.CommentGetAll(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentGet(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentGet(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentCreate(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentCreate(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentUpdate(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentUpdate(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentDelete(comment model.Comment) (err error) {
	err = s.repo.CommentDelete(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CommentAuthorization(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentAuthorization(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

// SocialMedia Endpoint
func (s *Service) SocialMediaGetAll() (res []model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaGetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaGet(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaGet(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaCreate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaCreate(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaUpdate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaUpdate(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaDelete(socialmedia model.SocialMedia) (err error) {
	err = s.repo.SocialMediaDelete(socialmedia)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SocialMediaAuthorization(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaAuthorization(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}
