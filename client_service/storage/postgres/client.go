package postgres

import "eCommerce/eCommerce"

// AddArticle ...
func (stg Postgres) AddArticle(id string, entity *eCommerce.CreateArticleRequest) error {
	_, err := stg.GetAuthorByID(entity.AuthorId)
	if err != nil {
		return err
	}

	if entity.Content == nil {
		entity.Content = &blogpost.Content{}
	}

	_, err = stg.db.Exec(`INSERT INTO article 
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
