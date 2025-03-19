package dao

import (
	"context"
	"github.com/qianqianzyk/AILesson-Planner/internal/model"
	"go.uber.org/zap"
	"time"
)

func (d *Dao) CreateFile(ctx context.Context, file *model.File) error {
	result := d.orm.WithContext(ctx).Model(&model.File{}).Create(file)
	return result.Error
}

func (d *Dao) DeleteFilesByIDs(ctx context.Context, ids []int, userID int) error {
	result := d.orm.WithContext(ctx).
		Where("id IN (?) AND user_id = ?", ids, userID).
		Delete(&model.File{})
	return result.Error
}

func (d *Dao) CompleteDelFile(ctx context.Context, ids []int, userID int) error {
	result := d.orm.WithContext(ctx).
		Unscoped().
		Where("id IN (?) AND user_id = ?", ids, userID).
		Delete(&model.File{})
	return result.Error
}

func (d *Dao) GetFileByNameAndParentID(ctx context.Context, userID int, name string, parentID int) (*model.File, error) {
	var file model.File
	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("user_id = ? AND name = ? AND parent_id = ?", userID, name, parentID).First(&file)
	return &file, result.Error
}

func (d *Dao) GetFileListByParentID(ctx context.Context, userID, parentID, pageNum, pageSize int) ([]model.File, *int64, error) {
	var files []model.File
	var sum int64
	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("user_id = ? AND parent_id = ?", userID, parentID).
		Count(&sum).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&files)
	return files, &sum, result.Error
}

func (d *Dao) GetFileList(ctx context.Context, parentID, pageNum, pageSize int) ([]model.File, *int64, error) {
	var files []model.File
	var sum int64
	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("parent_id = ?", parentID).
		Count(&sum).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&files)
	return files, &sum, result.Error
}

func (d *Dao) GetFilesByParentID(ctx context.Context, parentID int) ([]model.File, error) {
	var files []model.File
	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("parent_id = ?", parentID).Find(&files)
	return files, result.Error
}

func (d *Dao) GetFileListByType(ctx context.Context, userID, fileType, pageNum, pageSize int) ([]model.File, *int64, error) {
	var files []model.File
	var sum int64
	result := d.orm.WithContext(ctx).Model(&model.File{}).
		Where("user_id = ? AND f_type = ?", userID, fileType).
		Count(&sum).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&files)
	return files, &sum, result.Error
}

func (d *Dao) GetFileByID(ctx context.Context, id, userID int) (*model.File, error) {
	var file *model.File
	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("id = ? AND user_id = ?", id, userID).First(&file)
	if result.Error != nil {
		return nil, result.Error
	}
	return file, nil
}

func (d *Dao) GetFile(ctx context.Context, id int) (*model.File, error) {
	var file *model.File
	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("id = ?", id).First(&file)
	if result.Error != nil {
		return nil, result.Error
	}
	return file, nil
}

func (d *Dao) GetRecycleFileByUserID(ctx context.Context, userID int) ([]model.File, error) {
	var files []model.File
	result := d.orm.WithContext(ctx).
		Unscoped().
		Where("user_id = ? AND deleted_at IS NOT NULL", userID).
		Find(&files)
	return files, result.Error
}

func (d *Dao) GetCollectFileByUserID(ctx context.Context, userID, pageNum, pageSize int) ([]model.File, *int64, error) {
	var files []model.File
	var sum int64
	result := d.orm.WithContext(ctx).Model(model.File{}).
		Where("user_id = ? AND is_dir IS False AND is_collect IS NOT False", userID).
		Count(&sum).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&files)
	return files, &sum, result.Error
}

func (d *Dao) UpdateFile(ctx context.Context, file *model.File) error {
	result := d.orm.WithContext(ctx).Save(&file)
	return result.Error
}

func (d *Dao) UpdateFilesCollect(ctx context.Context, files []*model.File) error {
	fileIDs := make([]int, len(files))
	for i, file := range files {
		fileIDs[i] = int(file.ID)
	}

	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("id IN ?", fileIDs).Updates(map[string]interface{}{"is_collect": files[0].IsCollect})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Dao) UpdateFilesMove(ctx context.Context, files []*model.File) error {
	fileIDs := make([]int, len(files))
	for i, file := range files {
		fileIDs[i] = int(file.ID)
	}

	result := d.orm.WithContext(ctx).Model(&model.File{}).Where("id IN ?", fileIDs).Updates(map[string]interface{}{"parent_id": files[0].ParentID})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Dao) UpdateDirectoryAndSubFilesCollectStatus(ctx context.Context, dir *model.File) error {
	if dir.ParentID == 0 {
		result := d.orm.WithContext(ctx).Exec("UPDATE files SET is_collect = ? WHERE id = ?", dir.IsCollect, dir.ID)
		if result.Error != nil {
			return result.Error
		}
	}

	query := `
    WITH RECURSIVE SubFiles AS (
        SELECT id
        FROM files
        WHERE parent_id = ?

        UNION ALL

        SELECT f.id
        FROM files f
        INNER JOIN SubFiles sf ON f.parent_id = sf.id
    )
    UPDATE files
    SET is_collect = ?
    WHERE id IN (SELECT id FROM SubFiles);
    `

	result := d.orm.WithContext(ctx).Exec(query, dir.ID, dir.IsCollect)
	return result.Error
}

func (d *Dao) GetDirectoryStats(ctx context.Context, dirID int64) (*model.DirectoryStats, error) {
	var stats model.DirectoryStats

	sql := `
	WITH RECURSIVE FileTree AS (
	    SELECT id, is_dir, size, parent_id
	    FROM files
	    WHERE id = ?
	    UNION ALL
	    SELECT f.id, f.is_dir, f.size, f.parent_id
	    FROM files f
	    INNER JOIN FileTree ft ON f.parent_id = ft.id
	)
	SELECT 
	    SUM(CASE WHEN is_dir = 0 THEN 1 ELSE 0 END) AS file_count,
	    SUM(CASE WHEN is_dir = 1 THEN 1 ELSE 0 END) AS directory_count,
	    SUM(CASE WHEN is_dir = 0 THEN size ELSE 0 END) AS total_size
	FROM FileTree;
	`

	err := d.orm.WithContext(ctx).Raw(sql, dirID).Scan(&stats).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (d *Dao) RecoverFile(ctx context.Context, ids []int, userID int) error {
	sql := "UPDATE files SET deleted_at = NULL WHERE id IN (?) AND user_id = ?"
	result := d.orm.WithContext(ctx).Exec(sql, ids, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Dao) CleanupExpiredFiles(ctx context.Context) error {
	expiredTime := time.Now().AddDate(0, 0, -10)

	var files []model.File
	if err := d.orm.WithContext(ctx).
		Unscoped().
		Where("deleted_at IS NOT NULL AND deleted_at < ?", expiredTime).
		Find(&files).Error; err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir {
			if err := d.deleteDirectory(ctx, int(file.ID)); err != nil {
				zap.L().Warn("无法删除目录", zap.Error(err))
				continue
			}
		}

		if err := d.orm.WithContext(ctx).Unscoped().
			Where("id = ?", file.ID).
			Delete(&model.File{}).Error; err != nil {
			zap.L().Warn("无法删除文件", zap.Error(err))
		}
	}
	return nil
}

func (d *Dao) deleteDirectory(ctx context.Context, parentID int) error {
	return d.orm.WithContext(ctx).Exec(`
		WITH RECURSIVE all_files AS (
			SELECT id FROM files WHERE id = ?
			UNION ALL
			SELECT f.id FROM files f
			INNER JOIN all_files af ON f.parent_id = af.id
		)
		DELETE FROM files WHERE id IN (SELECT id FROM all_files);
	`, parentID).Error
}

func (d *Dao) CreateLink(ctx context.Context, link *model.ShareLink) error {
	result := d.orm.WithContext(ctx).Model(&model.ShareLink{}).Create(link)
	return result.Error
}

func (d *Dao) GetLinkByUrl(ctx context.Context, url string) (*model.ShareLink, error) {
	var link *model.ShareLink
	result := d.orm.WithContext(ctx).Model(&model.ShareLink{}).Where("link = ?", url).First(&link)
	return link, result.Error
}

func (d *Dao) CheckFilesExistenceByIDs(ctx context.Context, fileIDs []int, userID int) ([]*model.File, map[int]bool, error) {
	existsMap := make(map[int]bool)
	var files []*model.File

	files, err := d.GetFilesByIDsAndUserID(ctx, fileIDs, userID)
	if err != nil {
		return nil, nil, err
	}

	for _, file := range files {
		existsMap[int(file.ID)] = true
	}

	for _, fileID := range fileIDs {
		if _, found := existsMap[fileID]; !found {
			existsMap[fileID] = false
		}
	}

	return files, existsMap, nil
}

func (d *Dao) GetFilesByIDs(ctx context.Context, fileIDs []int) ([]model.File, error) {
	var files []model.File
	err := d.orm.WithContext(ctx).Model(model.File{}).
		Where("id IN (?) AND deleted_at IS NULL", fileIDs).
		Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (d *Dao) GetFilesByIDsAndUserID(ctx context.Context, fileIDs []int, userID int) ([]*model.File, error) {
	var files []*model.File
	err := d.orm.WithContext(ctx).Model(model.File{}).
		Where("id IN (?) AND user_id = ? AND deleted_at IS NULL", fileIDs, userID).
		Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}
