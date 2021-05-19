
-- +migrate Up
CREATE
OR REPLACE FUNCTION fetch_author(
    v_name varchar,
    v_limit int,
    v_offset int
)
    RETURNS TABLE
            (
                id         uuid,
                name       varchar,
                created_at timestamp,
                updated_at timestamp
            )
    LANGUAGE sql
AS
$function$
select a.id,
       a.name,
       a.created_at,
       a.updated_at
from author a
where a.deleted_at is null
  and case when v_name is not null then name ilike concat('%', v_name, '%') else true end
    limit v_limit offset v_offset
    $function$;

-- +migrate Down
