### 1. make a correlated query into a non-correlated query

- a correlated query:

  ```sql
  SELECT _
  FROM some_table
  WHERE relevant_field IN
  (
  SELECT relevant_field
  FROM some_table
  GROUP BY relevant_field
  HAVING COUNT(_) > 1
  )

  ```

- a non-correlated query

  ```sql
  SELECT *
  FROM some_table
  WHERE relevant_field IN
  (
      SELECT * FROM
      (
          SELECT relevant_field
          FROM some_table
          GROUP BY relevant_field
          HAVING COUNT(*) > 1
      ) AS subquery
  )
  ```
