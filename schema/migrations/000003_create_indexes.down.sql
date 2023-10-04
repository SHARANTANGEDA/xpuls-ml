drop index langchain_runs_project_id_index
    on public.langchain_runs (project_id);

drop index langchain_run_steps_chain_id_index
    on public.langchain_run_steps (chain_id);

