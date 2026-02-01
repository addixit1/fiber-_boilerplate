package querybuilder

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BaseRepository provides common database operations for MGM models
type BaseRepository struct{}

// NewBaseRepository creates a new base repository instance
func NewBaseRepository() *BaseRepository {
	return &BaseRepository{}
}

// FindOptions holds options for Find queries
type FindOptions struct {
	Sort       bson.M
	Projection bson.M
	Skip       *int64
	Limit      *int64
	Populate   []string // For future population support
}

// PaginateOptions holds pagination parameters
type PaginateOptions struct {
	Page  int
	Limit int
}

// PaginateResult holds paginated query results
type PaginateResult struct {
	Data      interface{}
	Total     int64
	Page      int
	Limit     int
	TotalPage int
	NextPage  int
}

// Save creates a new document
func (r *BaseRepository) Save(ctx context.Context, model mgm.Model) error {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).CreateWithCtx(ctx, model)
}

// Find retrieves multiple documents with optional filters and options
func (r *BaseRepository) Find(ctx context.Context, model mgm.Model, results interface{}, filter bson.M, opts *FindOptions) error {
	if ctx == nil {
		ctx = context.Background()
	}

	coll := mgm.Coll(model)
	findOpts := options.Find()

	if opts != nil {
		if opts.Sort != nil {
			findOpts.SetSort(opts.Sort)
		}
		if opts.Projection != nil {
			findOpts.SetProjection(opts.Projection)
		}
		if opts.Skip != nil {
			findOpts.SetSkip(*opts.Skip)
		}
		if opts.Limit != nil {
			findOpts.SetLimit(*opts.Limit)
		}
	}

	cursor, err := coll.Find(ctx, filter, findOpts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}

// FindOne retrieves a single document
func (r *BaseRepository) FindOne(ctx context.Context, model mgm.Model, filter bson.M, opts *FindOptions) error {
	if ctx == nil {
		ctx = context.Background()
	}

	coll := mgm.Coll(model)
	findOpts := options.FindOne()

	if opts != nil {
		if opts.Sort != nil {
			findOpts.SetSort(opts.Sort)
		}
		if opts.Projection != nil {
			findOpts.SetProjection(opts.Projection)
		}
	}

	return coll.FindOne(ctx, filter, findOpts).Decode(model)
}

// FindById retrieves a document by ID
func (r *BaseRepository) FindById(ctx context.Context, model mgm.Model, id interface{}) error {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).FindByIDWithCtx(ctx, id, model)
}

// UpdateOne updates a single document
func (r *BaseRepository) UpdateOne(ctx context.Context, model mgm.Model, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).UpdateOne(ctx, filter, update)
}

// UpdateMany updates multiple documents
func (r *BaseRepository) UpdateMany(ctx context.Context, model mgm.Model, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).UpdateMany(ctx, filter, update)
}

// FindOneAndUpdate finds and updates a single document
func (r *BaseRepository) FindOneAndUpdate(ctx context.Context, model mgm.Model, filter bson.M, update bson.M, opts *options.FindOneAndUpdateOptions) error {
	if ctx == nil {
		ctx = context.Background()
	}

	coll := mgm.Coll(model)
	return coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(model)
}

// UpdateById updates a document by ID
func (r *BaseRepository) UpdateById(ctx context.Context, model mgm.Model) error {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).UpdateWithCtx(ctx, model)
}

// DeleteOne deletes a single document
func (r *BaseRepository) DeleteOne(ctx context.Context, model mgm.Model, filter bson.M) (*mongo.DeleteResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).DeleteOne(ctx, filter)
}

// DeleteMany deletes multiple documents
func (r *BaseRepository) DeleteMany(ctx context.Context, model mgm.Model, filter bson.M) (*mongo.DeleteResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).DeleteMany(ctx, filter)
}

// DeleteById deletes a document by ID
func (r *BaseRepository) DeleteById(ctx context.Context, model mgm.Model) error {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).DeleteWithCtx(ctx, model)
}

// Count counts documents matching the filter
func (r *BaseRepository) Count(ctx context.Context, model mgm.Model, filter bson.M) (int64, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).CountDocuments(ctx, filter)
}

// CountDocuments counts documents (recommended method)
func (r *BaseRepository) CountDocuments(ctx context.Context, model mgm.Model, filter bson.M) (int64, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).CountDocuments(ctx, filter)
}

// Distinct returns distinct values for a field
func (r *BaseRepository) Distinct(ctx context.Context, model mgm.Model, field string, filter bson.M) ([]interface{}, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).Distinct(ctx, field, filter)
}

// Aggregate runs an aggregation pipeline
func (r *BaseRepository) Aggregate(ctx context.Context, model mgm.Model, pipeline mongo.Pipeline, results interface{}) error {
	if ctx == nil {
		ctx = context.Background()
	}

	coll := mgm.Coll(model)
	cursor, err := coll.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}

// Paginate performs pagination with aggregation pipeline
func (r *BaseRepository) Paginate(ctx context.Context, model mgm.Model, pipeline mongo.Pipeline, opts PaginateOptions) (*PaginateResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// Validate and set defaults
	if opts.Limit <= 0 {
		opts.Limit = 10
	}
	if opts.Limit > 100 {
		opts.Limit = 100
	}
	if opts.Page <= 0 {
		opts.Page = 1
	}

	// Calculate skip
	skip := (opts.Page - 1) * opts.Limit

	// Create a copy of pipeline for counting
	countPipeline := append(mongo.Pipeline{}, pipeline...)
	countPipeline = append(countPipeline, bson.D{{Key: "$count", Value: "total"}})

	// Add pagination to main pipeline
	paginatedPipeline := append(pipeline,
		bson.D{{Key: "$skip", Value: skip}},
		bson.D{{Key: "$limit", Value: opts.Limit + 1}}, // Fetch one extra to check if there's a next page
	)

	coll := mgm.Coll(model)

	// Execute count query
	var countResult []bson.M
	countCursor, err := coll.Aggregate(ctx, countPipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		return nil, err
	}
	if err := countCursor.All(ctx, &countResult); err != nil {
		return nil, err
	}

	var total int64
	if len(countResult) > 0 {
		if t, ok := countResult[0]["total"].(int32); ok {
			total = int64(t)
		} else if t, ok := countResult[0]["total"].(int64); ok {
			total = t
		}
	}

	// Execute data query
	var data []bson.M
	dataCursor, err := coll.Aggregate(ctx, paginatedPipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		return nil, err
	}
	if err := dataCursor.All(ctx, &data); err != nil {
		return nil, err
	}

	// Check if there's a next page
	nextPage := 0
	if len(data) > opts.Limit {
		nextPage = opts.Page + 1
		data = data[:opts.Limit] // Remove the extra record
	}

	totalPage := int((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &PaginateResult{
		Data:      data,
		Total:     total,
		Page:      opts.Page,
		Limit:     opts.Limit,
		TotalPage: totalPage,
		NextPage:  nextPage,
	}, nil
}

// InsertMany inserts multiple documents
func (r *BaseRepository) InsertMany(ctx context.Context, model mgm.Model, documents []interface{}) (*mongo.InsertManyResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).InsertMany(ctx, documents)
}

// BulkWrite performs bulk write operations
func (r *BaseRepository) BulkWrite(ctx context.Context, model mgm.Model, operations []mongo.WriteModel, opts *options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	return mgm.Coll(model).BulkWrite(ctx, operations, opts)
}

// FindWithPagination performs a simple find with pagination (non-aggregation)
func (r *BaseRepository) FindWithPagination(ctx context.Context, model mgm.Model, results interface{}, filter bson.M, opts PaginateOptions) (*PaginateResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	// Validate and set defaults
	if opts.Limit <= 0 {
		opts.Limit = 10
	}
	if opts.Limit > 100 {
		opts.Limit = 100
	}
	if opts.Page <= 0 {
		opts.Page = 1
	}

	// Calculate skip
	skip := int64((opts.Page - 1) * opts.Limit)
	limit := int64(opts.Limit + 1) // Fetch one extra

	// Get total count
	total, err := r.CountDocuments(ctx, model, filter)
	if err != nil {
		return nil, err
	}

	// Fetch data
	findOpts := &FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	err = r.Find(ctx, model, results, filter, findOpts)
	if err != nil {
		return nil, err
	}

	// Check for next page (this requires reflection or a different approach based on results type)
	// For simplicity, we'll calculate it based on total
	totalPage := int((total + int64(opts.Limit) - 1) / int64(opts.Limit))
	nextPage := 0
	if opts.Page < totalPage {
		nextPage = opts.Page + 1
	}

	return &PaginateResult{
		Data:      results,
		Total:     total,
		Page:      opts.Page,
		Limit:     opts.Limit,
		TotalPage: totalPage,
		NextPage:  nextPage,
	}, nil
}

// WithTimeout creates a context with timeout
func (r *BaseRepository) WithTimeout(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}
