"use client"

import React, {useEffect, useState} from 'react';
import MDEditor from '@uiw/react-md-editor';
import Layout from "@/app/components/Layout";
import AutoBreadcrumbs from "@/app/components/AutoBreadcrumbs";
import {Button, CircularProgress, Tooltip} from '@nextui-org/react';
import {useQuery} from "react-query";
import {getLatestPrompt, getPromptVersions} from "@/services/prompt_registry";
import PromptVersionCard from './PromptVersionCard';
import CommitOutlinedIcon from '@mui/icons-material/CommitOutlined';


export default function  PromptPage({ params }: { params: { projectId: string, promptId: string } })  {

    const [value, setValue] = useState("**Hello world!!!**");

    const [showOlderPrompt, setShowOlderPrompt] = useState(false)
    const [showPrompt, setShowPrompt] = useState<PromptVersion | null>(null)


    const { data: promptVersions, error, isLoading } = useQuery(
        ['getPromptVersions', params.projectId, params.promptId],
        () => getPromptVersions(params.projectId, params.promptId),
        {
            keepPreviousData: true, // Enable this to keep old data visible while fetching new data
            cacheTime: 3600,

        }
    );

    const { data: latestPrompt, error: latestPromptError,
        isLoading: promptLoading } = useQuery(
        ['getLatestPrompt', params.projectId, params.promptId],
        () => getLatestPrompt(params.projectId, params.promptId),
        {
            keepPreviousData: true, // Enable this to keep old data visible while fetching new data
            cacheTime: 3600,

        }
    );

    useEffect(() => {
        if (latestPrompt !== undefined) {
            setShowPrompt(latestPrompt);
            setValue(latestPrompt.prompt_content)
        }
    }, [latestPrompt]);

    const handlePromptChange = (olderPromptVersion: PromptVersion) => {
        setShowOlderPrompt(true)
        setShowPrompt(olderPromptVersion)
        setValue(olderPromptVersion.prompt_content)

    };

    if (promptLoading || latestPrompt === undefined || showPrompt === null) {
        return <CircularProgress aria-label="Loading..."/>
    }else if (latestPromptError !== null) {
        return <p>Error loading data</p>
    }

    return (
        <Layout >
            <div className="flex flex-row divide-x">
                <div className="basis-3/4 max-w-[75%]">
                    <AutoBreadcrumbs/>

                    <div className="flex flex-row justify-between mb-3">
                        <div className="flex flex-row items-center justify-center gap-4">
                            <p className="text-2xl" >{showPrompt.prompt_name}</p>

                            <Tooltip color="foreground" content="Click to copy 📋" placement="top">
                            <p className="text-blue-500"><CommitOutlinedIcon/> {showPrompt.prompt_version_id.substring(0,7)}</p>
                            </Tooltip>
                        </div>
                        <Button color="primary" variant="shadow">
                            Save
                        </Button>
                    </div>
                    <MDEditor highlightEnable={true} style={{"minHeight": "50%", fontSize: "10pt"}}
                        value={value} onChange={(newValue) => setValue(
                        newValue !== undefined ? newValue: value)} data-color-mode="light"
                    />
                    <div className="flex flex-row justify-end mt-3">
                        <Button color="primary" variant="shadow">
                            Save
                        </Button>
                    </div>
                </div>
                <div className="basis-1/4 max-w-[25%] border-l-4 h-screen">
                    {isLoading ? <CircularProgress aria-label="Loading..." />: null}
                    {!isLoading && (error || promptVersions === undefined) ? <div>Error loading data</div>: null}
                    {!isLoading && !error && promptVersions !== undefined ? promptVersions.map(
                        (item, index) => (
                            <PromptVersionCard key={index} item={item}
                                                             handlePromptChange={handlePromptChange} />
                        )): null}
                </div>

            </div>


            {/*<MDEditor.Markdown source={value} style={{ whiteSpace: 'pre-wrap' }} />*/}

        </Layout>
    );
};

