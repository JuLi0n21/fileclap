package repository

import (
	"database/sql"

	"github.com/JuLi0n21/fileclap/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

type FileRepository interface {
	CreateFile(file *models.File, user *models.User) error
	GetFileByID(fileid uuid.UUID) (*models.File, error)
	GetAllFilesForUser(user *models.User) ([]*models.File, error)
	GetAllFoldersForUser(user *models.User) ([]*models.Folder, error)
	GetFilesInFolder(folder *models.Folder) ([]*models.File, error)
	UpdateFile(user *models.File) error
	DeleteFile(id uuid.UUID) error
}

type Repository struct {
	db             *sql.DB
	UserRepository UserRepository
	FileRepository FileRepository
}
