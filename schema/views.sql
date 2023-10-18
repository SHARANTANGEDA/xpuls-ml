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


create or replace view view_prompt_registry_versions
            (prompt_version_id, prompt_content, prompt_version_created_at, prompt_id, prompt_tag, prompt_name,
             project_id, prompt_created_at, prompt_deleted, prompt_deleted_at)
as
SELECT prpv.prompt_version_id,
       prpv.prompt_content,
       prpv.prompt_version_created_at,
       prpv.prompt_id,
       prpv.prompt_tag,
       prp.prompt_name,
       prp.project_id,
       prp.prompt_created_at,
       prp.prompt_deleted,
       prp.prompt_deleted_at
FROM prompt_registry_prompt_versions prpv
         LEFT JOIN prompt_registry_prompts prp ON prp.prompt_id::text = prpv.prompt_id::text;

create or replace view view_prompt_registry_latest_version
            (prompt_version_id, prompt_content, prompt_version_created_at, prompt_id, prompt_tag, prompt_name,
             project_id, prompt_created_at, prompt_deleted, prompt_deleted_at)
as
SELECT DISTINCT ON (prpv.prompt_id) prpv.prompt_version_id,
    prpv.prompt_content,
    prpv.prompt_version_created_at,
    prp.prompt_id,
    prpv.prompt_tag,
    prp.prompt_name,
    prp.project_id,
    prp.prompt_created_at,
    prp.prompt_deleted,
    prp.prompt_deleted_at
FROM prompt_registry_prompt_versions prpv
    LEFT JOIN prompt_registry_prompts prp ON prp.prompt_id::text = prpv.prompt_id::text
ORDER BY prpv.prompt_id, prpv.prompt_version_created_at DESC;

