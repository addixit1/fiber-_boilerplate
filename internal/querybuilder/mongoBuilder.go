package querybuilder

import "go.mongodb.org/mongo-driver/bson"

type Builder struct {
	filter bson.M
}

func New() *Builder {
	return &Builder{filter: bson.M{}}
}

// Eq adds an equality filter
func (b *Builder) Eq(key string, value any) *Builder {
	b.filter[key] = value
	return b
}

// Ne adds a not equal filter
func (b *Builder) Ne(key string, value any) *Builder {
	b.filter[key] = bson.M{"$ne": value}
	return b
}

// Gt adds a greater than filter
func (b *Builder) Gt(key string, value any) *Builder {
	b.filter[key] = bson.M{"$gt": value}
	return b
}

// Gte adds a greater than or equal filter
func (b *Builder) Gte(key string, value any) *Builder {
	b.filter[key] = bson.M{"$gte": value}
	return b
}

// Lt adds a less than filter
func (b *Builder) Lt(key string, value any) *Builder {
	b.filter[key] = bson.M{"$lt": value}
	return b
}

// Lte adds a less than or equal filter
func (b *Builder) Lte(key string, value any) *Builder {
	b.filter[key] = bson.M{"$lte": value}
	return b
}

// In adds an $in operator filter
func (b *Builder) In(key string, values []any) *Builder {
	b.filter[key] = bson.M{"$in": values}
	return b
}

// Nin adds a not in filter
func (b *Builder) Nin(key string, values []any) *Builder {
	b.filter[key] = bson.M{"$nin": values}
	return b
}

// Regex adds a regex filter with case-insensitive option
func (b *Builder) Regex(key string, value string) *Builder {
	b.filter[key] = bson.M{"$regex": value, "$options": "i"}
	return b
}

// Exists adds a field existence check
func (b *Builder) Exists(key string, exists bool) *Builder {
	b.filter[key] = bson.M{"$exists": exists}
	return b
}

// Or combines multiple filters with OR logic
func (b *Builder) Or(filters ...bson.M) *Builder {
	b.filter["$or"] = filters
	return b
}

// And combines multiple filters with AND logic
func (b *Builder) And(filters ...bson.M) *Builder {
	b.filter["$and"] = filters
	return b
}

// Build returns the final bson.M filter
func (b *Builder) Build() bson.M {
	return b.filter
}
