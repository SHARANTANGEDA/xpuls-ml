create or replace view view_langchain_run_filter_keys(project_id, runtime_keys, label_keys) as
SELECT main.project_id,
       ARRAY(SELECT DISTINCT 'runtime.'::text || jsonb_object_keys(sub.runtime)
             FROM langchain_runs sub
             WHERE sub.project_id::text = main.project_id::text) AS runtime_keys,
       ARRAY(SELECT DISTINCT 'labels.'::text || jsonb_object_keys(sub.labels)
             FROM langchain_runs sub
             WHERE sub.project_id::text = main.project_id::text) AS label_keys
FROM langchain_runs main
GROUP BY main.project_id;

