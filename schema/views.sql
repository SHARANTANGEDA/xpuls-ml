create view mv_langchain_run_filter_keys(key) as
SELECT DISTINCT 'runtime.'::text || jsonb_object_keys(langchain_runs.runtime) AS key
FROM langchain_runs
UNION
SELECT DISTINCT 'labels.'::text || jsonb_object_keys(langchain_runs.labels) AS key
FROM langchain_runs;
