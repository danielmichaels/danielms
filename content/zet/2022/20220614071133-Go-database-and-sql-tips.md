+++
title = "Go database and sql tips"
categories = ["zet"]
tags = ["zet"]
slug = "go-database-and-sql-tips"
date = "2022-06-14 00:00:00 +0000 UTC"
draft = "false"
ShowToc = "true"
mermaid = "true"
+++

# Go database and sql tips

Note to self.


```go
func (s store) ItemUpdate(id int, name, description string, image []byte) error {
	stmt := `
	UPDATE items
	SET
		name = $1,
		description = $2,
		image = $3,
		updated_at = datetime('now')
	WHERE
		id = $4
	RETURNING id`
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var c Item
	args := []interface{}{name, description, image, id}
	err := s.DB.QueryRowContext(ctx, stmt, args...).Scan(&c.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return errors.New("item does not exist")
		default:
			return err
		}
	}
	return nil
}
```

This code works but this **really** tripped me up...

The `stmt` block uses `$1` numbering, instead of the usual `?`. I stupidly
thought the `$n` numbering meant something when executing the `args...` inside
the `QueryRowContext` method.

**It does not**. It is simply there for you, to help count the number of
variable statements inside the query, AFAIK. This cost me an hour or so because
I set the `id = $1` so as to map the `ItemUpdate` parameters with the query.

I hope I remember this in the future! I don't like ORM's but this definitely would 
not have tripped me up. 

Tags:

    #go #database #failure
