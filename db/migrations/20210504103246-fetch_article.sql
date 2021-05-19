-- +migrate Up
CREATE
    OR REPLACE FUNCTION fetch_article(
    v_title varchar,
    v_author_id uuid,
    v_from_date date,
    v_until_date date,
    v_limit int,
    v_offset int
)
    RETURNS TABLE
            (
                id         uuid,
                title      varchar,
                content    varchar,
                author     json,
                created_at timestamp,
                updated_at timestamp
            )
    LANGUAGE sql
AS
$function$
select ar.id,
       ar.title,
       ar.content,
       json_build_object(
               'id', a.id,
               'name', a.name
           ) as author,
       ar.created_at,
       ar.updated_at
from article ar
         join author a on ar.author_id = a.id and a.deleted_at is null
where ar.deleted_at is null
  and case when v_title is not null then title ilike concat('%', v_title, '%') else true end
  and case when v_author_id is not null then author_id = v_author_id else true end
  and case when v_from_date is not null then date(ar.created_at) >= v_from_date else true end
  and case when v_until_date is not null then date(ar.created_at) <= v_until_date else true end
limit v_limit offset v_offset
$function$;


-- +migrate Down