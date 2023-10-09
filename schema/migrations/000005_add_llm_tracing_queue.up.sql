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
