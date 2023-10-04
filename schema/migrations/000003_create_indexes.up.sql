create index if not exists langchain_runs_project_id_index
    on public.langchain_runs (project_id);

create index if not exists langchain_run_steps_chain_id_index
    on public.langchain_run_steps (chain_id);

