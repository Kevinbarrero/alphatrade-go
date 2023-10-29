-- name: CreateStrategy :one
INSERT INTO strategies (
  user_id,
  name,
  indicators,
  buyconditions,
  sellconditions
) VALUES
($1, $2, $3, $4, $5) RETURNING *;

-- name: GetStrategiesByUser :many
select * from strategies
where user_id = $1;

-- name: DeleteStrategy :exec
delete from strategies
where id = $1;
