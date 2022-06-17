package user

type user struct {
	ID        string   `bson:"id"`
	Username  string   `bson:"username"`
	Email     string   `bson:"email"`
	Links     []string `bson:"links"`
	CreatedAt int64    `bson:"created_at"`
	UpdatedAt int64    `bson:"updated_at"`
}

type userResponse struct {
	ID        string         `bson:"id" json:"id"`
	Username  string         `bson:"username" json:"username"`
	Email     string         `bson:"email" json:"email"`
	Links     []linkResponse `bson:"links" json:"links"`
	CreatedAt int64          `bson:"created_at" json:"created_at"`
	UpdatedAt int64          `bson:"updated_at" json:"updated_at"`
}

type linkResponse struct {
	ID        string `bson:"id" json:"id"`
	Title     string `bson:"title" json:"title"`
	Url       string `bson:"url" json:"url"`
	IsActive  bool   `bson:"is_active" json:"is_active"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
}
