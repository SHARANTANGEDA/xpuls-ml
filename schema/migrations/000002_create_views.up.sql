CREATE or REPLACE VIEW view_langchain_run_filter_keys AS
SELECT
    project_id,
    ARRAY(
        SELECT DISTINCT 'runtime.' || jsonb_object_keys(runtime)
    FROM langchain_runs AS sub
    WHERE sub.project_id = main.project_id
  ) AS runtime_keys,
    ARRAY(
        SELECT DISTINCT 'labels.' || jsonb_object_keys(labels)
    FROM langchain_runs AS sub
    WHERE sub.project_id = main.project_id
  ) AS label_keys
FROM
    langchain_runs AS main
GROUP BY
    project_id;
