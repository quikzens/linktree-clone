package link

type link struct {
	ID         string `bson:"id" json:"id"`
	EmailOwner string `bson:"email_owner" json:"email_owner"`
	Title      string `bson:"title" json:"title"`
	Url        string `bson:"url" json:"url"`
	IsActive   bool   `bson:"is_active" json:"is_active"`
	CreatedAt  int64  `bson:"created_at" json:"created_at"`
	UpdatedAt  int64  `bson:"updated_at" json:"updated_at"`
}

type updateLinkRequest struct {
	Title     string `json:"title" bson:"title,omitempty"`
	Url       string `json:"url" bson:"url,omitempty"`
	IsActive  bool   `json:"is_active" bson:"is_active,omitempty"`
	UpdatedAt int64  `bson:"updated_at"`
}

type userLinks struct {
	Links []string `bson:"links"`
}

type updateUserLinks struct {
	Links     []string `bson:"links"`
	UpdatedAt int64    `bson:"updated_at"`
}
