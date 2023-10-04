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

