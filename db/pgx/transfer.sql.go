package db

import "context"

const createTransfer = `
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var t Transfer
	err := row.Scan(
		&t.ID,
		&t.FromAccountID,
		&t.ToAccountID,
		&t.Amount,
		&t.CreatedAt,
	)

	return t, err
}

const getTransfer = `SELECT * FROM transfers WHERE ID = $1 LIMIT 1`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRow(ctx, getTransfer, id)
	var t Transfer
	err := row.Scan(
		&t.ID,
		&t.FromAccountID,
		&t.ToAccountID,
		&t.Amount,
		&t.CreatedAt,
	)

	return t, err
}

const listTransfer = `SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE 
    from_account_id = $1 OR
    to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransfer struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfer(ctx context.Context, arg ListTransfer) ([]Transfer, error) {
	rows, err := q.db.Query(ctx, listTransfer, arg.FromAccountID, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transfers []Transfer
	for rows.Next() {
		var t Transfer
		if err := rows.Scan(
			&t.ID,
			&t.FromAccountID,
			&t.ToAccountID,
			&t.Amount,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}
		transfers = append(transfers, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return transfers, nil
}
