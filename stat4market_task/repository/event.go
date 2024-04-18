package repository

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"stat4market_task/model"
	"time"
)

type Event interface {
	Insert(ctx context.Context, event model.Event) error
}

type event struct {
	clickDb driver.Conn
}

func NewEventRepo(clickDb driver.Conn) Event {
	return &event{clickDb: clickDb}
}

func (r *event) Insert(ctx context.Context, event model.Event) error {
	queryStr := "INSERT INTO events(eventID, eventType, userID, eventTime, payload) SELECT next_id, ?, ?, ?, ? FROM sequence_table"

	err := r.clickDb.Exec(ctx, queryStr, event.EventType, event.UserID, event.EventTime, event.Payload)
	if err != nil {
		return err
	}

	err = r.clickDb.Exec(ctx, "ALTER TABLE sequence_table UPDATE next_id = next_id + 1 WHERE 1 = 1;")

	return err
}

func (r *event) getByEventTypeAndTime(ctx context.Context, eventType string, startTime, endTime time.Time) ([]model.Event, error) {
	query := `
        SELECT eventID, eventType, userID, eventTime, payload
        FROM events
        WHERE eventType = ? AND eventTime >= ? AND eventTime <= ?
    `

	rows, err := r.clickDb.Query(ctx, query, eventType, startTime, endTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var item model.Event
		err := rows.Scan(&item.ID,
			&item.EventType,
			&item.EventTime,
			&item.UserID,
			&item.Payload)

		if err != nil {
			return nil, fmt.Errorf("get.rows.Scan: %v", err)
		}

		events = append(events, item)
	}

	return events, nil
}
