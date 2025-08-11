package ent

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Client is the client that holds all ent builders.
type Client struct {
	db *sql.DB
	// Temple is the client for interacting with the Temple builders.
	Temple *TempleClient
	// GoshuinCollection is the client for interacting with the GoshuinCollection builders.
	GoshuinCollection *GoshuinCollectionClient
}

// NewClient creates a new client configured with the given options.
func NewClient() *Client {
	return &Client{
		Temple:            NewTempleClient(nil),
		GoshuinCollection: NewGoshuinCollectionClient(nil),
	}
}

// NewClientWithDB creates a new client with database connection.
func NewClientWithDB(db *sql.DB) *Client {
	return &Client{
		db:                db,
		Temple:            NewTempleClient(db),
		GoshuinCollection: NewGoshuinCollectionClient(db),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

// TempleClient is a client for the Temple schema.
type TempleClient struct {
	db *sql.DB
}

// NewTempleClient returns a client for the Temple from the given config.
func NewTempleClient(db *sql.DB) *TempleClient {
	return &TempleClient{db: db}
}

// Create returns a builder for creating a Temple entity.
func (c *TempleClient) Create() *TempleCreate {
	return &TempleCreate{db: c.db}
}

// Query returns a query builder for Temple.
func (c *TempleClient) Query() *TempleQuery {
	return &TempleQuery{db: c.db}
}

// Get returns a Temple entity by its id.
func (c *TempleClient) Get(ctx context.Context, id int) (*Temple, error) {
	if c.db == nil {
		// ダミーデータを返す
		return &Temple{
			ID:        id,
			Name:      "浅草寺",
			NameEn:    "Senso-ji Temple",
			Latitude:  35.7148,
			Longitude: 139.7967,
		}, nil
	}

	query := `
		SELECT id, name, name_en, description, description_en, latitude, longitude, 
		       address, phone, website, instagram, twitter, opening_hours, 
		       goshuin_fee, goshuin_office, is_active, created_at, updated_at
		FROM temples WHERE id = ?
	`

	var temple Temple
	var createdAt, updatedAt time.Time
	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&temple.ID, &temple.Name, &temple.NameEn, &temple.Description, &temple.DescriptionEn,
		&temple.Latitude, &temple.Longitude, &temple.Address, &temple.Phone,
		&temple.Website, &temple.Instagram, &temple.Twitter, &temple.OpeningHours,
		&temple.GoshuinFee, &temple.GoshuinOffice, &temple.IsActive, &createdAt, &updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get temple: %v", err)
	}

	temple.CreatedAt = createdAt.Format(time.RFC3339)
	temple.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &temple, nil
}

// GoshuinCollectionClient is a client for the GoshuinCollection schema.
type GoshuinCollectionClient struct {
	db *sql.DB
}

// NewGoshuinCollectionClient returns a client for the GoshuinCollection from the given config.
func NewGoshuinCollectionClient(db *sql.DB) *GoshuinCollectionClient {
	return &GoshuinCollectionClient{db: db}
}

// Create returns a builder for creating a GoshuinCollection entity.
func (c *GoshuinCollectionClient) Create() *GoshuinCollectionCreate {
	return &GoshuinCollectionCreate{db: c.db}
}

// Query returns a query builder for GoshuinCollection.
func (c *GoshuinCollectionClient) Query() *GoshuinCollectionQuery {
	return &GoshuinCollectionQuery{db: c.db}
}

// Get returns a GoshuinCollection entity by its id.
func (c *GoshuinCollectionClient) Get(ctx context.Context, id int) (*GoshuinCollection, error) {
	if c.db == nil {
		// ダミーデータを返す
		return &GoshuinCollection{
			ID:       id,
			TempleID: 1,
			ImageURL: "https://example.com/goshuin.jpg",
			Notes:    "Beautiful goshuin stamp",
		}, nil
	}

	query := `
		SELECT id, temple_id, image_url, notes, collected_at, created_at, updated_at
		FROM goshuin_collections WHERE id = ?
	`

	var collection GoshuinCollection
	var collectedAt, createdAt, updatedAt time.Time
	err := c.db.QueryRowContext(ctx, query, id).Scan(
		&collection.ID, &collection.TempleID, &collection.ImageURL, &collection.Notes,
		&collectedAt, &createdAt, &updatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get goshuin collection: %v", err)
	}

	collection.CollectedAt = collectedAt.Format(time.RFC3339)
	collection.CreatedAt = createdAt.Format(time.RFC3339)
	collection.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &collection, nil
}

// UpdateOneID returns a builder for updating a GoshuinCollection entity.
func (c *GoshuinCollectionClient) UpdateOneID(id int) *GoshuinCollectionUpdateOneID {
	return &GoshuinCollectionUpdateOneID{db: c.db, id: id}
}

// DeleteOneID returns a builder for deleting a GoshuinCollection entity.
func (c *GoshuinCollectionClient) DeleteOneID(id int) *GoshuinCollectionDeleteOneID {
	return &GoshuinCollectionDeleteOneID{db: c.db, id: id}
}

// Temple entity
type Temple struct {
	ID            int     `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	NameEn        string  `json:"name_en,omitempty"`
	Description   string  `json:"description,omitempty"`
	DescriptionEn string  `json:"description_en,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	Address       string  `json:"address,omitempty"`
	Phone         string  `json:"phone,omitempty"`
	Website       string  `json:"website,omitempty"`
	Instagram     string  `json:"instagram,omitempty"`
	Twitter       string  `json:"twitter,omitempty"`
	OpeningHours  string  `json:"opening_hours,omitempty"`
	GoshuinFee    string  `json:"goshuin_fee,omitempty"`
	GoshuinOffice string  `json:"goshuin_office,omitempty"`
	IsActive      bool    `json:"is_active,omitempty"`
	CreatedAt     string  `json:"created_at,omitempty"`
	UpdatedAt     string  `json:"updated_at,omitempty"`
}

// GoshuinCollection entity
type GoshuinCollection struct {
	ID          int    `json:"id,omitempty"`
	TempleID    int    `json:"temple_id,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	Notes       string `json:"notes,omitempty"`
	CollectedAt string `json:"collected_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// TempleQuery is a query builder for Temple.
type TempleQuery struct {
	db *sql.DB
}

// All returns all temples.
func (tq *TempleQuery) All(ctx context.Context) ([]*Temple, error) {
	if tq.db == nil {
		// ダミーデータを返す
		return []*Temple{
			{
				ID:        1,
				Name:      "浅草寺",
				NameEn:    "Senso-ji Temple",
				Latitude:  35.7148,
				Longitude: 139.7967,
			},
			{
				ID:        2,
				Name:      "明治神宮",
				NameEn:    "Meiji Shrine",
				Latitude:  35.6764,
				Longitude: 139.6993,
			},
		}, nil
	}

	query := `
		SELECT id, name, name_en, description, description_en, latitude, longitude, 
		       address, phone, website, instagram, twitter, opening_hours, 
		       goshuin_fee, goshuin_office, is_active, created_at, updated_at
		FROM temples WHERE is_active = TRUE
		ORDER BY name
	`

	rows, err := tq.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query temples: %v", err)
	}
	defer rows.Close()

	var temples []*Temple
	for rows.Next() {
		var temple Temple
		var createdAt, updatedAt time.Time
		err := rows.Scan(
			&temple.ID, &temple.Name, &temple.NameEn, &temple.Description, &temple.DescriptionEn,
			&temple.Latitude, &temple.Longitude, &temple.Address, &temple.Phone,
			&temple.Website, &temple.Instagram, &temple.Twitter, &temple.OpeningHours,
			&temple.GoshuinFee, &temple.GoshuinOffice, &temple.IsActive, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan temple: %v", err)
		}

		temple.CreatedAt = createdAt.Format(time.RFC3339)
		temple.UpdatedAt = updatedAt.Format(time.RFC3339)
		temples = append(temples, &temple)
	}

	return temples, nil
}

// Where adds a predicate to the query.
func (tq *TempleQuery) Where(conds ...interface{}) *TempleQuery {
	return tq
}

// Only returns a single temple.
func (tq *TempleQuery) Only(ctx context.Context) (*Temple, error) {
	temples, err := tq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(temples) == 0 {
		return nil, fmt.Errorf("no temples found")
	}
	if len(temples) > 1 {
		return nil, fmt.Errorf("multiple temples found")
	}
	return temples[0], nil
}

// GoshuinCollectionQuery is a query builder for GoshuinCollection.
type GoshuinCollectionQuery struct {
	db *sql.DB
}

// All returns all goshuin collections.
func (gcq *GoshuinCollectionQuery) All(ctx context.Context) ([]*GoshuinCollection, error) {
	if gcq.db == nil {
		return []*GoshuinCollection{}, nil
	}

	query := `
		SELECT id, temple_id, image_url, notes, collected_at, created_at, updated_at
		FROM goshuin_collections
		ORDER BY collected_at DESC
	`

	rows, err := gcq.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query goshuin collections: %v", err)
	}
	defer rows.Close()

	var collections []*GoshuinCollection
	for rows.Next() {
		var collection GoshuinCollection
		var collectedAt, createdAt, updatedAt time.Time
		err := rows.Scan(
			&collection.ID, &collection.TempleID, &collection.ImageURL, &collection.Notes,
			&collectedAt, &createdAt, &updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan goshuin collection: %v", err)
		}

		collection.CollectedAt = collectedAt.Format(time.RFC3339)
		collection.CreatedAt = createdAt.Format(time.RFC3339)
		collection.UpdatedAt = updatedAt.Format(time.RFC3339)
		collections = append(collections, &collection)
	}

	return collections, nil
}

// Where adds a predicate to the query.
func (gcq *GoshuinCollectionQuery) Where(conds ...interface{}) *GoshuinCollectionQuery {
	return gcq
}

// Only returns a single goshuin collection.
func (gcq *GoshuinCollectionQuery) Only(ctx context.Context) (*GoshuinCollection, error) {
	collections, err := gcq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(collections) == 0 {
		return nil, fmt.Errorf("no goshuin collections found")
	}
	if len(collections) > 1 {
		return nil, fmt.Errorf("multiple goshuin collections found")
	}
	return collections[0], nil
}

// TempleCreate is a builder for creating a Temple entity.
type TempleCreate struct {
	db     *sql.DB
	temple *Temple
}

// SetName sets the name field.
func (tc *TempleCreate) SetName(name string) *TempleCreate {
	if tc.temple == nil {
		tc.temple = &Temple{}
	}
	tc.temple.Name = name
	return tc
}

// SetNameEn sets the name_en field.
func (tc *TempleCreate) SetNameEn(nameEn string) *TempleCreate {
	if tc.temple == nil {
		tc.temple = &Temple{}
	}
	tc.temple.NameEn = nameEn
	return tc
}

// SetLatitude sets the latitude field.
func (tc *TempleCreate) SetLatitude(lat float64) *TempleCreate {
	if tc.temple == nil {
		tc.temple = &Temple{}
	}
	tc.temple.Latitude = lat
	return tc
}

// SetLongitude sets the longitude field.
func (tc *TempleCreate) SetLongitude(lng float64) *TempleCreate {
	if tc.temple == nil {
		tc.temple = &Temple{}
	}
	tc.temple.Longitude = lng
	return tc
}

// Save saves the temple to the database.
func (tc *TempleCreate) Save(ctx context.Context) (*Temple, error) {
	if tc.db == nil {
		// ダミー実装
		if tc.temple == nil {
			tc.temple = &Temple{}
		}
		tc.temple.ID = 3 // 仮のID
		return tc.temple, nil
	}

	if tc.temple == nil {
		tc.temple = &Temple{}
	}

	query := `
		INSERT INTO temples (name, name_en, description, description_en, latitude, longitude, 
		                    address, phone, website, instagram, twitter, opening_hours, 
		                    goshuin_fee, goshuin_office, is_active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := tc.db.ExecContext(ctx, query,
		tc.temple.Name, tc.temple.NameEn, tc.temple.Description, tc.temple.DescriptionEn,
		tc.temple.Latitude, tc.temple.Longitude, tc.temple.Address, tc.temple.Phone,
		tc.temple.Website, tc.temple.Instagram, tc.temple.Twitter, tc.temple.OpeningHours,
		tc.temple.GoshuinFee, tc.temple.GoshuinOffice, tc.temple.IsActive,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create temple: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %v", err)
	}

	tc.temple.ID = int(id)
	return tc.temple, nil
}

// GoshuinCollectionCreate is a builder for creating a GoshuinCollection entity.
type GoshuinCollectionCreate struct {
	db         *sql.DB
	collection *GoshuinCollection
}

// SetTempleID sets the temple_id field.
func (gcc *GoshuinCollectionCreate) SetTempleID(id int) *GoshuinCollectionCreate {
	if gcc.collection == nil {
		gcc.collection = &GoshuinCollection{}
	}
	gcc.collection.TempleID = id
	return gcc
}

// SetImageURL sets the image_url field.
func (gcc *GoshuinCollectionCreate) SetImageURL(url string) *GoshuinCollectionCreate {
	if gcc.collection == nil {
		gcc.collection = &GoshuinCollection{}
	}
	gcc.collection.ImageURL = url
	return gcc
}

// SetNotes sets the notes field.
func (gcc *GoshuinCollectionCreate) SetNotes(notes string) *GoshuinCollectionCreate {
	if gcc.collection == nil {
		gcc.collection = &GoshuinCollection{}
	}
	gcc.collection.Notes = notes
	return gcc
}

// Save saves the goshuin collection to the database.
func (gcc *GoshuinCollectionCreate) Save(ctx context.Context) (*GoshuinCollection, error) {
	if gcc.db == nil {
		// ダミー実装
		if gcc.collection == nil {
			gcc.collection = &GoshuinCollection{}
		}
		gcc.collection.ID = 1 // 仮のID
		return gcc.collection, nil
	}

	if gcc.collection == nil {
		gcc.collection = &GoshuinCollection{}
	}

	query := `
		INSERT INTO goshuin_collections (temple_id, image_url, notes, collected_at)
		VALUES (?, ?, ?, NOW())
	`

	result, err := gcc.db.ExecContext(ctx, query,
		gcc.collection.TempleID, gcc.collection.ImageURL, gcc.collection.Notes,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create goshuin collection: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id: %v", err)
	}

	gcc.collection.ID = int(id)
	return gcc.collection, nil
}

// GoshuinCollectionUpdateOneID is a builder for updating a GoshuinCollection entity.
type GoshuinCollectionUpdateOneID struct {
	db         *sql.DB
	id         int
	collection *GoshuinCollection
}

// SetImageURL sets the image_url field.
func (gcu *GoshuinCollectionUpdateOneID) SetImageURL(url string) *GoshuinCollectionUpdateOneID {
	if gcu.collection == nil {
		gcu.collection = &GoshuinCollection{}
	}
	gcu.collection.ImageURL = url
	return gcu
}

// SetNotes sets the notes field.
func (gcu *GoshuinCollectionUpdateOneID) SetNotes(notes string) *GoshuinCollectionUpdateOneID {
	if gcu.collection == nil {
		gcu.collection = &GoshuinCollection{}
	}
	gcu.collection.Notes = notes
	return gcu
}

// Save saves the updated goshuin collection to the database.
func (gcu *GoshuinCollectionUpdateOneID) Save(ctx context.Context) (*GoshuinCollection, error) {
	if gcu.db == nil {
		// ダミー実装
		if gcu.collection == nil {
			gcu.collection = &GoshuinCollection{}
		}
		gcu.collection.ID = gcu.id
		return gcu.collection, nil
	}

	if gcu.collection == nil {
		gcu.collection = &GoshuinCollection{}
	}

	query := `
		UPDATE goshuin_collections 
		SET image_url = ?, notes = ?, updated_at = NOW()
		WHERE id = ?
	`

	_, err := gcu.db.ExecContext(ctx, query,
		gcu.collection.ImageURL, gcu.collection.Notes, gcu.id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update goshuin collection: %v", err)
	}

	gcu.collection.ID = gcu.id
	return gcu.collection, nil
}

// GoshuinCollectionDeleteOneID is a builder for deleting a GoshuinCollection entity.
type GoshuinCollectionDeleteOneID struct {
	db *sql.DB
	id int
}

// Exec executes the delete operation.
func (gcd *GoshuinCollectionDeleteOneID) Exec(ctx context.Context) error {
	if gcd.db == nil {
		// ダミー実装：何もしない
		return nil
	}

	query := `DELETE FROM goshuin_collections WHERE id = ?`

	_, err := gcd.db.ExecContext(ctx, query, gcd.id)
	if err != nil {
		return fmt.Errorf("failed to delete goshuin collection: %v", err)
	}

	return nil
}
