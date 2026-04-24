package argument

type AddReaction struct {
	PostID    string `json:"post_id"`
	EmojiName string `json:"emoji_name"`
}
