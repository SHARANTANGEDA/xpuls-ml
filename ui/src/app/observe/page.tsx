'use client'

import Layout from "@/app/components/Layout";
import {fetchProjects} from "@/services/projects";
import {
    Table,
    Pagination,
    TableHeader,
    TableColumn,
    TableRow,
    TableCell,
    getKeyValue,
    Link,
    Button, Divider
} from '@nextui-org/react';
import {useEffect, useState} from "react";
import {Breadcrumbs, TableBody} from "@mui/material";
import Typography from "@mui/material/Typography";
import {CircularProgress} from "@nextui-org/react";



import { DataGrid, GridColDef } from '@mui/x-data-grid';
import {useQuery} from "react-query";
import AutoBreadcrumbs from "@/app/components/AutoBreadcrumbs";
import ProjectModal from "@/app/observe/NewProjectModal";

const columns: GridColDef[] = [
    { field: 'project_id', headerName: 'ID', width: 350},
    { field: 'project_name', headerName: 'Name', width: 130, renderCell: (params) => (
            <b>
                {params.value}
            </b>
        ), },
    { field: 'project_slug', headerName: 'Slug', width: 130 },
    { field: 'open', headerName: 'View', width: 180, renderCell: (params) => (
            <Link href={`/observe/${params.row.project_id}`} rel="noopener">
                <Button size={"sm"} color="primary" variant="flat">
                    view runs
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

    const pages = 2

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
            {/* Breadcrumbs */}
            {/*<Breadcrumbs size="mini">*/}
            {/*    <Breadcrumbs.Item>Home</Breadcrumbs.Item>*/}
            {/*    <Breadcrumbs.Item>Projects</Breadcrumbs.Item>*/}
            {/*    <Breadcrumbs.Item>Observe</Breadcrumbs.Item>*/}
            {/*</Breadcrumbs>*/}

            {/* Header */}
            <AutoBreadcrumbs/>
            <ProjectModal isOpen={isModalOpen} onClose={() => setModalOpen(false)} />


            <div className="flex justify-between items-center">
                <Typography variant="h5" gutterBottom>
                    Agents
                </Typography>
                {/*<Button onClick={() => setModalOpen(true)} radius="full" className="bg-gradient-to-tr from-blue-600 to-green-600 text-white shadow-lg">*/}
                {/*    + New*/}
                {/*</Button>*/}
            </div>

            {/* Table */}
            <DataGrid
                rows={projects}
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
