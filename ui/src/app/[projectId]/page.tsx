'use client'

import Layout from "@/app/components/Layout";
import {fetchProjects} from "@/services/projects";
import {
    Link,
    Button,
} from '@nextui-org/react';
import {useState} from "react";
import Typography from "@mui/material/Typography";
import {CircularProgress} from "@nextui-org/react";
import ConstructionIcon from '@mui/icons-material/Construction';
import MonitorHeartIcon from '@mui/icons-material/MonitorHeart';

import { DataGrid, GridColDef } from '@mui/x-data-grid';
import {useQuery} from "react-query";
import AutoBreadcrumbs from "@/app/components/AutoBreadcrumbs";
import ProjectModal from "@/app/[projectId]/NewProjectModal";

const columns: GridColDef[] = [
    { field: 'project_id', headerName: 'ID', width: 350},
    { field: 'project_name', headerName: 'Name', width: 130, renderCell: (params) => (
            <b>
                {params.value}
            </b>
        ), },
    { field: 'project_slug', headerName: 'Slug', width: 130 },
    { field: 'prompt_registry', headerName: 'Manage Prompts', width: 180, renderCell: (params) => (
            <Link href={`/${params.row.project_id}/registry`} rel="noopener">
                <Button className="text-blue-700 bg-blue-100" size={"sm"} variant="flat">
                   <ConstructionIcon/> Prompt Registry
                </Button>
            </Link>
        ) },
    { field: 'view_runs', headerName: 'Monitor', width: 180, renderCell: (params) => (
            <Link href={`/${params.row.project_id}/observe`} rel="noopener">
                <Button className="text-green-700 bg-green-100" size={"sm"} variant="flat">
                    <MonitorHeartIcon/> Monitor Runs
                </Button>
            </Link>
        ) },

    // { field: 'total_runs', headerName: 'Total Runs', width: 130 },
    // { field: 'total_tokens', headerName: 'Total Tokens', width: 130 },
    // { field: 'total_p50', headerName: 'p50 latency', width: 130 },
    // { field: 'total_p95', headerName: 'p95 latency', width: 130 },
    // { field: 'latest_run_latency', headerName: 'latest runs latency', width: 130 },

];



// @ts-ignore
export default function Projects() {
    const [page, setPage] = useState(1);
    const [isModalOpen, setModalOpen] = useState<boolean>(false);

    const rowsPerPage = 4;

    const { data: projects, error, isLoading } = useQuery(
        ['projects', page, rowsPerPage],
        () => fetchProjects(page, rowsPerPage),
        {
            keepPreviousData: true // Enable this to keep old data visible while fetching new data
        }
    );

    if (isLoading) return <CircularProgress aria-label="Loading..." />;
    if (error) return <div>Error loading data</div>;


    return (
        <Layout>
            {/* Header */}
            <AutoBreadcrumbs/>
            <div className="flex justify-between items-center">
                <p className="text-2xl">
                    Agents
                </p>
                <Button onClick={() => setModalOpen(true)} className="bg-blue-700 text-white mr-2 mb-2">
                    + Create Agent </Button>
            </div>
            <ProjectModal isOpen={isModalOpen} onClose={() => setModalOpen(!isModalOpen)} />



            {/* Table */}
            <DataGrid
                rows={projects as any[]}
                columns={columns}
                getRowId={(row) => row.project_id}
                initialState={{
                    pagination: {
                        paginationModel: { page: 0, pageSize: 5 },
                    },
                }}
                pageSizeOptions={[5, 10]}
                checkboxSelection
            />
        </Layout>
    )
}
