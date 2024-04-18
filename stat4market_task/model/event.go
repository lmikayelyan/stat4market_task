package model

type Event struct {
	ID        int64  `json:"id"`
	EventType string `json:"event_type"`
	UserID    int64  `json:"user_id"`
	EventTime string `json:"event_time"`
	Payload   string `json:"payload"`
}
