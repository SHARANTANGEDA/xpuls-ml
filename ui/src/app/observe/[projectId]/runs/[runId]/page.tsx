'use client'

import Layout from "@/app/components/Layout";
import {useEffect, useState} from "react";
import Typography from "@mui/material/Typography";

import "@/app/globals.css"
import {RunStepTree} from "@/app/components/TreeComponent/ObservationTree";
import ReactMarkdown from 'react-markdown';
import {useQuery} from "react-query";
import {fetchLangChainFilterKeys, fetchLangChainRunSteps} from "@/services/langchain_runs";
import {CircularProgress} from "@nextui-org/react";
import {Divider} from "@nextui-org/react";
import {LangChainRunStep} from "@/types/langchain";
import AutoBreadcrumbs from "@/app/components/AutoBreadcrumbs";


export default function LangChainRunView({ params }: { params: { projectId: string, runId: string } }) {

    const [currentRunStep, setCurrentRunStep] = useState<LangChainRunStep | null>(null);


    const { data: runSteps, error, isLoading } = useQuery(
        ['fetchLangChainRunSteps', params.projectId, params.runId],
        () => fetchLangChainRunSteps(params.projectId, params.runId),
        {
            keepPreviousData: true, // Enable this to keep old data visible while fetching new data
            cacheTime: 3600,

        }
    );


    if (isLoading) return <CircularProgress aria-label="Loading..." />;
    if (error) return <div>Error loading data</div>;

    if (runSteps !== null && !isLoading && !error && currentRunStep === null) {
        setCurrentRunStep(runSteps[0])
    }



    return (
        <Layout>

            {/* Header */}
            <AutoBreadcrumbs/>

            <Typography variant="h5" gutterBottom>
                Run:
            </Typography>

            {/* Table */}
            <div style={{display: 'flex', width: '100%'}}>
                <div style={{ width: '25%' }}>
                    <RunStepTree
                        runSteps={runSteps}
                        currentRunStep={currentRunStep}
                        setCurrentRunStep={setCurrentRunStep}
                    />

                </div>
                {currentRunStep !== null ?
                (<  div style={{ margin: '20px', padding: '20px', border: '1px solid #ddd', width: '75%' }}>
                    <ReactMarkdown>{`# *${currentRunStep.run_name} (${currentRunStep.run_type})*`}</ReactMarkdown>
                    <h2></h2>
                    <br/>
                    <p><b>Event Start Time:</b> {new Date(currentRunStep.event_start_time).toLocaleString()}</p>
                    <p><b>Event End Time:</b> {new Date(currentRunStep.event_end_time).toLocaleString()}</p>
                    <p><b>Total Tokens:</b> {currentRunStep.token_usage?.total_tokens}</p>

                    {currentRunStep.prompt_template && (
                        <div>
                            <Divider style={{marginTop: '15px', marginBottom: '15px'}}/>
                            <h3><b>Prompt Template:</b></h3>
                            <ReactMarkdown>{currentRunStep.prompt_template}</ReactMarkdown>
                        </div>
                    )}
                    {currentRunStep.prompt_content && (
                        <div>
                            <Divider style={{marginTop: '15px', marginBottom: '15px'}}/>
                            <h3><b>Prompt Content:</b></h3>
                            <ReactMarkdown>{currentRunStep.prompt_content}</ReactMarkdown>
                        </div>
                    )}
                    {currentRunStep.prompt_input && (
                        <div>
                            <Divider style={{marginTop: '15px', marginBottom: '15px'}}/>
                            <h3><b>Prompt Input:</b></h3>
                            <ReactMarkdown>{currentRunStep.prompt_input}</ReactMarkdown>
                        </div>
                    )}
                    {currentRunStep.prompt_chat_history && (
                        <div>
                            <Divider style={{marginTop: '15px', marginBottom: '15px'}}/>
                            <h3><b>Prompt Chat History:</b></h3>
                            <ReactMarkdown>{currentRunStep.prompt_chat_history}</ReactMarkdown>
                        </div>
                    )}
                    {currentRunStep.prompt_agent_scratchpad && (
                        <div>
                            <Divider style={{marginTop: '15px', marginBottom: '15px'}}/>
                            <h3><b>Prompt Agent Scratchpad:</b></h3>
                            <ReactMarkdown>{currentRunStep.prompt_agent_scratchpad}</ReactMarkdown>
                        </div>
                    )}
                    {currentRunStep.prompt_output && (
                        <div>
                            <Divider style={{marginTop: '15px', marginBottom: '15px'}}/>
                            <h3><b>Prompt Output:</b></h3>
                            <ReactMarkdown>{currentRunStep.prompt_output}</ReactMarkdown>
                        </div>
                    )}
                </div>) : null}

            </div>




        </Layout>
    )
}
