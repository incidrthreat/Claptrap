package provider

type Provider interface {
	Connect() bool
	Reconnect() bool
	IsAlive() bool
	ListenForEvents()
	GetEvents() <-chan Event
	AutoJoinAllChannel() bool

	// Get Information
	GetEmailByUsername(username string) string
	GetUserIDByUsername(username string) string

	// Actions
	MessagePublic(channelID, message string) bool
	MessageUser(userID, message string) bool
	InviteUser(userID, channelID string) bool
	KickUser(userID, channelID string) bool
	RevokeUserActiveSessions(userID string) bool
	UpdateUserMessage(postID, message string) bool
	DeleteMessage(event Event) bool
	ReplaceMessagePlaceholders(event Event, message string) string
}
