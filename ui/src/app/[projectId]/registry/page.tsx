"use client"

import React, {useState} from 'react';
import PaginatedCards from './PaginatedCards';
import AutoBreadcrumbs from "@/app/components/AutoBreadcrumbs";
import {getLatestPrompt, getLatestPromptsInProject} from "@/services/prompt_registry";
import {useQuery} from "react-query";
import {Button, CircularProgress} from "@nextui-org/react";
import Layout from "@/app/components/Layout";
import PromptModal from "@/app/[projectId]/registry/NewPromptModal";

interface Prompt {
    prompt_name: string;
    prompt_text: string;
    prompt_version: string;
    last_edited_by: string;
    prompt_id: string;
}


export default function ViewPrompts({params}: {params: {projectId: string}}) {
    const [itemsLimit, setItemLimits] = useState(20);
    const [page, setPage] = React.useState(1);
    const [isModalOpen, setModalOpen] = useState<boolean>(false);


    const { data, error, isLoading } = useQuery(
        ['getLatestPromptsInProject', params.projectId],
        () => getLatestPromptsInProject(params.projectId, page, itemsLimit),
        {
            keepPreviousData: true, // Enable this to keep old data visible while fetching new data
            cacheTime: 3600,

        }
    );

    if (isLoading || data === undefined ) {
        return <CircularProgress aria-label="Loading..."/>
    }else if (error !== null) {
        return <p>Error loading data</p>
    }


    return (
        <Layout>

            <AutoBreadcrumbs/>
            <div className="flex justify-end mr-2">
                <Button color="primary" variant="solid" onClick={() => setModalOpen(true)}>
                    + Add Prompt
                </Button>
            </div>

            <PaginatedCards data={data} itemsPerPage={itemsLimit} page={page} setPage={setPage}/>
            <PromptModal projectId={params.projectId} isOpen={isModalOpen} onClose={() => setModalOpen(!isModalOpen)}/>
        </Layout>
    );
};

