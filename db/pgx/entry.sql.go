package db

import "context"

const createEntry = `INSERT INTO entries (account_id,amount) VALUES ($1,$2) RETURNING id,account_id,amount,created_at`

type CreateEntryParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

// CREATE
func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRow(ctx, createEntry, arg.AccountID, arg.Amount)
	var e Entry
	err := row.Scan(
		&e.ID,
		&e.AccountID,
		&e.Amount,
		&e.CreatedAt,
	)

	return e, err
}

const getEntry = `SELECT id,account_id,amount,created_at FROM entries WHERE id = $1 LIMIT 1`

// READ
func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRow(ctx, getEntry, id)
	var e Entry
	err := row.Scan(
		&e.ID,
		&e.AccountID,
		&e.Amount,
		&e.CreatedAt,
	)

	return e, err
}

const listEntries = `SELECT * FROM entries WHERE account_id = $1 ORDER BY id LIMIT $2 OFFSET $3`

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.Query(ctx, listEntries, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entries []Entry
	for rows.Next() {
		var e Entry
		if err := rows.Scan(
			&e.ID,
			&e.AccountID,
			&e.Amount,
			&e.CreatedAt,
		); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
