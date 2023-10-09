create table if not exists projects
(
    project_id          varchar(200)                                                        not null
    constraint projects_pk
    primary key,
    project_name        varchar(200)                                                        not null,
    project_created_at  timestamp      default (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'::text) not null,
    project_deleted     boolean        default false                                        not null,
    project_deleted_at  timestamp,
    project_slug        varchar(200)                                                        not null
    constraint projects_pk2
    unique,
    project_description varchar(2000),
    project_tags        varchar(200)[] default '{}'::character varying[]                    not null
    );

create table if not exists langchain_runs
(
    chain_id              varchar(200)                                                   not null
    constraint langchain_chains_pk
    primary key,
    project_id            varchar(200)                                                   not null
    constraint langchain_chains_projects_project_id_fk
    references projects,
    model_info            jsonb,
    labels                jsonb     default '{}'::jsonb                                  not null,
    runtime               jsonb,
    chain_tracked_at      timestamp default (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'::text) not null,
    first_step_start_time timestamp,
    last_step_end_time    timestamp,
    total_tokens          integer,
    prompt_tokens         integer,
    completion_tokens     integer
    );

create table if not exists langchain_run_steps
(
    run_step_id                     varchar(200)                                                   not null
    constraint langchain_runs_pk
    primary key,
    run_name                        varchar(200)                                                   not null,
    run_type                        varchar(200)                                                   not null,
    execution_order                 integer                                                        not null,
    prompt_template                 text,
    prompt_template_input_variables varchar(250)[],
    prompt_input                    text,
    prompt_chat_history             text,
    prompt_agent_scratchpad         text,
    prompt_output                   text,
    event_start_time                timestamp                                                      not null,
    event_end_time                  timestamp,
    parent_step_id                  varchar(200),
    token_usage                     jsonb,
    begin_json                      jsonb,
    end_json                        jsonb,
    tracked_at                      timestamp default (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'::text) not null,
    chain_id                        varchar(200)
    constraint langchain_run_steps_langchain_runs_chain_id_fk
    references langchain_runs,
    prompt_content                  text
    );

create index if not exists langchain_run_steps_chain_id_index
    on langchain_run_steps (chain_id);

create index if not exists langchain_runs_project_id_index
    on langchain_runs (project_id);

create table if not exists schema_migrations
(
    version bigint  not null
    primary key,
    dirty   boolean not null
);

create table if not exists compute_environments
(
    env_name       varchar(250) not null
    constraint compute_environments_pk
    primary key,
    env_properties jsonb
    );

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
    references prompt_registry_prompts
    );

create table if not exists prompt_registry_deployments
(
    prompt_deployment_id       varchar(200)                                                      not null
    constraint prompt_registry_deployments_pk
    primary key,
    deployment_env             varchar(250)                                                      not null
    constraint prompt_deployments_compute_environments_env_name_fk
    references compute_environments,
    deployment_mode            varchar(200) default 'ROLLOUT'::character varying                 not null,
    deployment_properties      jsonb        default '{}'::jsonb                                  not null,
    deployment_created_at      timestamp    default (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'::text) not null,
    deployment_last_updated_at timestamp,
    prompt_id                  varchar(200)                                                      not null
    constraint prompt_deployments_prompt_registry_prompts_prompt_id_fk
    references prompt_registry_prompts,
    deployment_active          boolean      default true                                         not null,
    deployment_deleted         boolean      default false                                        not null,
    deployment_deleted_at      timestamp
    );

create table if not exists llm_tracing_queue
(
    queue_item_id         varchar(200)                                                      not null
    constraint llm_tracing_queue_pk
    primary key,
    queue_item_type       varchar(200) default 'LANGCHAIN_RUN'::character varying           not null,
    item_stored_at        timestamp    default (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'::text) not null,
    params                jsonb        default '{}'::jsonb                                  not null,
    data                  jsonb        default '{}'::jsonb                                  not null,
    queue_processing_code varchar(200),
    queue_item_processed  boolean      default false                                        not null
    );
