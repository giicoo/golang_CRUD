SELECT
  *
FROM
  author
WHERE
  author_id = @param1 OR
  name = @param2