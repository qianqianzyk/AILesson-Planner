package service

import "github.com/qianqianzyk/AILesson-Planner/internal/model"

func CreateShareResource(post *model.ShareResource) error {
	err := d.CreateShareResource(ctx, post)
	return err
}

func DeleteShareResource(resourceID, userID int) error {
	coverImg, err := d.DeleteShareResource(ctx, resourceID, userID)
	if coverImg != "" {
		DeleteObjectByUrlAsync(coverImg)
	}
	return err
}

func GetShareResourceList(resourceType int) ([]model.ShareResource, error) {
	posts, err := d.GetShareResourceList(ctx, resourceType)
	return posts, err
}

func SearchShareResource(resourceType int, keyword string) ([]model.ShareResource, error) {
	posts, err := d.SearchShareResource(ctx, resourceType, keyword)
	return posts, err
}
