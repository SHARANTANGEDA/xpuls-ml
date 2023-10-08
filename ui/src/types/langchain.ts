export type LangChainRunStep = {
    run_step_id: string;
    run_name: string;
    run_type: string;
    execution_order: number;
    prompt_template?: string | null;
    prompt_template_input_variables?: string[] | null;
    prompt_input?: string | null;
    prompt_chat_history?: string | null;
    prompt_agent_scratchpad?: string | null;
    prompt_output?: string | null;
    event_start_time: Date;
    event_end_time: Date;
    parent_step_id: string | null;
    token_usage: any; // jsonb might be better represented as a specific object type if its structure is known
    begin_json: any; // jsonb might be better represented as a specific object type if its structure is known
    end_json: any; // jsonb might be better represented as a specific object type if its structure is known
    tracked_at: Date;
    chain_id: string;
    prompt_content?: string | null;
    children?: Array<any> | null;
};
