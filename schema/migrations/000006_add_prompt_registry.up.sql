create table if not exists prompt_registry_prompts
(
    prompt_id         varchar(200)                                                   not null
        constraint prompt_registry_prompts_pk
            primary key,
    prompt_name       varchar(200)                                                   not null,
    project_id        varchar(200)                                                   not null
        constraint prompt_registry_prompts_projects_project_id_fk
            references projects,
    prompt_created_at timestamp default (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'::text) not null,
    prompt_deleted    boolean   default false                                        not null,
    prompt_deleted_at timestamp
);

create table if not exists prompt_registry_prompt_versions
(
    prompt_version_id         varchar(200)                                                   not null
        constraint prompt_registry_prompt_versions_pk
            primary key,
    prompt_content            text                                                           not null,
    prompt_version_created_at timestamp default (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'::text) not null,
    prompt_id                 varchar(200)                                                   not null
        constraint prompt_versions_prompts_prompt_id_fk
            references prompt_registry_prompts,
    prompt_tag                varchar(200)
);

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

