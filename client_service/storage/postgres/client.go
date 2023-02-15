package postgres

import (
	"eCommerce/eCommerce"
	"errors"
	"time"
)

// AddClient ...
func (stg Postgres) AddClient(id string, entity *eCommerce.CreateClientRequest) error {

	_, err := stg.db.Exec(`INSERT INTO clients 
	(
		id,
		title,
		body,
		author_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		id,
		entity.Content.Title,
		entity.Content.Body,
		entity.AuthorId,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetClientByID ...
func (stg Postgres) GetClientByID(id string) (*clients.GetClientByIDResponse, error) {
	res := &clients.GetClientByIDResponse{
		Content: &clients.Content{},
		Author:  &clients.GetClientByIDResponse_Author{},
	}
	var deletedAt *time.Time
	var updatedAt, authorUpdatedAt *string
	err := stg.db.QueryRow(`SELECT 
		ar.id,
		ar.title,
		ar.body,
		ar.created_at,
		ar.updated_at,
		ar.deleted_at,
		au.id,
		au.fullname,
		au.created_at,
		au.updated_at
    FROM clients AS ar JOIN author AS au ON ar.author_id = au.id WHERE ar.id = $1`, id).Scan(
		&res.Id,
		&res.Content.Title,
		&res.Content.Body,
		&res.CreatedAt,
		&updatedAt,
		&deletedAt,
		&res.Author.Id,
		&res.Author.Fullname,
		&res.Author.CreatedAt,
		&authorUpdatedAt,
	)
	if err != nil {
		return res, err
	}

	if updatedAt != nil {
		res.UpdatedAt = *updatedAt
	}

	if authorUpdatedAt != nil {
		res.Author.UpdatedAt = *authorUpdatedAt
	}

	if deletedAt != nil {
		return res, errors.New("clients not found")
	}

	return res, err
}

// GetClientList ...
func (stg Postgres) GetClientList(offset, limit int, search string) (*clients.GetClientListResponse, error) {
	resp := &clients.GetClientListResponse{
		Clients: make([]*clients.Client, 0),
	}

	rows, err := stg.db.Queryx(`SELECT
	id,
	title,
	body,
	author_id,
	created_at,
	updated_at
	FROM clients WHERE deleted_at IS NULL AND ((title ILIKE '%' || $1 || '%') OR (body ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		a := &clients.Client{
			Content: &clients.Content{},
		}

		var updatedAt *string

		err := rows.Scan(
			&a.Id,
			&a.Content.Title,
			&a.Content.Body,
			&a.AuthorId,
			&a.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return resp, err
		}

		if updatedAt != nil {
			a.UpdatedAt = *updatedAt
		}

		resp.Clients = append(resp.Clients, a)
	}

	return resp, err
}

// UpdateClient ...
func (stg Postgres) UpdateClient(entity *clients.UpdateClientRequest) error {
	if entity.Content == nil {
		entity.Content = &clients.Content{}
	}
	res, err := stg.db.NamedExec("UPDATE clients  SET title=:t, body=:b, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": entity.Id,
		"t":  entity.Content.Title,
		"b":  entity.Content.Body,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("clients not found")
}

// DeleteClient ...
func (stg Postgres) DeleteClient(id string) error {
	res, err := stg.db.Exec("UPDATE clients SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("clients not found")
}
