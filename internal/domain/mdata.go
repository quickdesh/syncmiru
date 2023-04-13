package domain

import (
	"context"
	"time"
)

type MangaDataRepo interface {
	Store(ctx context.Context, mdata *MangaData) (*MangaData, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, mdata *MangaData) (*MangaData, error)
	ListMangaData(ctx context.Context, apiKey string) ([]MangaData, error)
	GetMangaDataByApiKey(ctx context.Context, apiKey string) (*MangaData, error)
}

type MangaData struct {
	ID         int         `json:"id"`
	Manga      []Manga     `json:"manga"`
	Extensions []Extension `json:"extensions"`
	Categories []Category  `json:"categories"`
	UserApiKey *APIKey     `json:"user_api_key,omitempty"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type Manga struct {
	Source       int64     `json:"source"`
	URL          string    `json:"url"`
	Favourite    bool      `json:"favourite"`
	Title        string    `json:"title"`
	Artist       string    `json:"artist"`
	Author       string    `json:"author"`
	Description  string    `json:"description"`
	Genre        []string  `json:"genre"`
	Status       int       `json:"status"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	DateAdded    int64     `json:"dateAdded"`
	Viewer       int       `json:"viewer"`
	Chapters     []Chapter `json:"chapters"`
	Categories   []int     `json:"categories"`
	ViewerFlags  int       `json:"viewer_flags"`
	History      []History `json:"history"`
}

type Chapter struct {
	Id            int64  `json:"id"`
	MangaId       int64  `json:"mangaId"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Scanlator     string `json:"scanlator"`
	Read          bool   `json:"read"`
	Bookmark      bool   `json:"bookmark"`
	LastPageRead  int64  `json:"lastPageRead"`
	DateFetch     int64  `json:"dateFetch"`
	DateUpload    int64  `json:"dateUpload"`
	ChapterNumber int    `json:"chapterNumber"`
	SourceOrder   int    `json:"sourceOrder"`
	MangaUrl      string `json:"mangaUrl"`
	MangaSource   int64  `json:"mangaSource"`
}

type History struct {
	URL          string `json:"url"`
	LastRead     int64  `json:"lastRead"`
	ReadDuration int    `json:"readDuration"`
}

type Extension struct {
	Name     string `json:"name"`
	SourceID int64  `json:"sourceId"`
}

type Category struct {
	Name  string `json:"name"`
	Flags int    `json:"flags"`
	Order int    `json:"order"`
}