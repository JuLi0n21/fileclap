package repository

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"

	"github.com/JuLi0n21/fileclap/models"
	"github.com/JuLi0n21/fileclap/utils"
	"github.com/google/uuid"
)

func NewSQLiteRepository(dataSourceName string) (*Repository, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		db:             db,
		UserRepository: NewUserRepository(db),
		FileRepository: NewFileRepository(db),
	}

	err = repo.initDB()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// initDB creates the necessary tables in the SQLite database.
func (r *Repository) initDB() error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		salt TEXT NOT NULL
	);`
	if _, err := r.db.Exec(createUserTable); err != nil {
		return err
	}

	createFileTable := `
	CREATE TABLE IF NOT EXISTS files (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		date TEXT NOT NULL,
		size INTEGER NOT NULL,
		folder TEXT NOT NULL,
		user_id TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	if _, err := r.db.Exec(createFileTable); err != nil {
		return err
	}

	return nil
}

// Close closes the database connection.
func (r *Repository) Close() error {
	return r.db.Close()
}

type FileRepositorySQLite struct {
	db *sql.DB
}

// NewFileRepository creates a new FileRepositorySQLite instance.
func NewFileRepository(db *sql.DB) *FileRepositorySQLite {
	return &FileRepositorySQLite{db: db}
}

// UserRepositorySQLite implements UserRepository for SQLite.
type UserRepositorySQLite struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepositorySQLite instance.
func NewUserRepository(db *sql.DB) *UserRepositorySQLite {

	return &UserRepositorySQLite{db: db}
}

// CreateUser inserts a new user into the database.
func (r *UserRepositorySQLite) RegisterUser(name, email, password string) (*models.User, error) {

	u := models.NewUser(name)

	s, err := utils.GenValue(24)
	if err != nil {
		return nil, err
	}

	p, err := utils.HashPassword(s, password)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO users (id, name, email, password, salt) VALUES (?, ?, ?, ?, ?)"
	_, err = r.db.Exec(query, u.ID, u.Name, email, p, s)
	return u, err
}

func (r *UserRepositorySQLite) LoginUser(name, password string) (*models.User, error) {

	query := "SELECT * FROM users WHERE name = ? OR email = ?"
	row := r.db.QueryRow(query, name, name)

	var user models.User
	var s, p, e string

	if err := row.Scan(&user.ID, &user.Name, e, p, s); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Incorrect sign in")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(s, password, p) {
		return nil, errors.New("Incorrect sign in")
	}

	return &user, nil
}

// GetUserByID retrieves a user by their ID.
func (r *UserRepositorySQLite) GetUserByID(id uuid.UUID) (*models.User, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database.
func (r *UserRepositorySQLite) GetAllUsers() ([]*models.User, error) {
	query := "SELECT id, name FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// UpdateUser updates an existing user in the database.
func (r *UserRepositorySQLite) UpdateUser(user *models.User) error {
	query := "UPDATE users SET name = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.ID)
	return err
}

// DeleteUser deletes a user from the database.
func (r *UserRepositorySQLite) DeleteUser(id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

// CreateFile inserts a new file into the database.
func (r *FileRepositorySQLite) CreateFile(file *models.File, user *models.User) error {
	query := "INSERT INTO files (name, date, size, folder, user_id) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, file.Name, file.Date, file.Size, file.Folder, user.ID)
	return err
}

// GetFileByID retrieves a file by its ID.
func (r *FileRepositorySQLite) GetFileByID(fileid uuid.UUID) (*models.File, error) {
	query := "SELECT id, name, date, size, folder FROM files WHERE id = ?"
	row := r.db.QueryRow(query, fileid)

	var file models.File
	if err := row.Scan(&file.ID, &file.Name, &file.Date, &file.Size, &file.Folder); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // File not found
		}
		return nil, err
	}
	return &file, nil
}

// GetAllFilesForUser retrieves all files for a given user.
func (r *FileRepositorySQLite) GetAllFilesForUser(user *models.User) ([]*models.File, error) {
	query := "SELECT id, name, date, size, folder FROM files WHERE user_id = ?"
	rows, err := r.db.Query(query, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*models.File
	for rows.Next() {
		var file models.File
		if err := rows.Scan(&file.ID, &file.Name, &file.Date, &file.Size, &file.Folder); err != nil {
			return nil, err
		}
		files = append(files, &file)
	}
	return files, nil
}

func (r *FileRepositorySQLite) GetAllFoldersForUser(user *models.User) ([]*models.Folder, error) {
	query := "SELECT DISTINCT folder FROM files WHERE user_id = ?"
	rows, err := r.db.Query(query, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var folders []*models.Folder
	for rows.Next() {
		var folder *models.Folder
		if err := rows.Scan(&folder); err != nil {
			return nil, err
		}
		folders = append(folders, folder)
	}
	return folders, nil
}

// GetFilesInFolder retrieves all files in a specific folder.
func (r *FileRepositorySQLite) GetFilesInFolder(folder *models.Folder) ([]*models.File, error) {
	query := "SELECT id, name, date, size, folder FROM files WHERE folder LIKE ?"
	rows, err := r.db.Query(query, folder.Name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*models.File
	for rows.Next() {
		var file models.File
		if err := rows.Scan(&file.ID, &file.Name, &file.Date, &file.Size, &file.Folder); err != nil {
			return nil, err
		}
		files = append(files, &file)
	}
	return files, nil
}

// UpdateFile updates an existing file in the database.
func (r *FileRepositorySQLite) UpdateFile(file *models.File) error {
	query := "UPDATE files SET name = ?, date = ?, size = ?, folder = ? WHERE id = ?"
	_, err := r.db.Exec(query, file.Name, file.Date, file.Size, file.Folder, file.ID)
	return err
}

// DeleteFile deletes a file from the database.
func (r *FileRepositorySQLite) DeleteFile(id uuid.UUID) error {
	query := "DELETE FROM files WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
