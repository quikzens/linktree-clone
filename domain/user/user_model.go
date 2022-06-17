package user

type user struct {
	ID        string   `bson:"id"`
	Username  string   `bson:"username"`
	Email     string   `bson:"email"`
	Links     []string `bson:"links"`
	CreatedAt int64    `bson:"created_at"`
	UpdatedAt int64    `bson:"updated_at"`
}
