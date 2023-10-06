'use client'

import Layout from "@/app/components/Layout";
import {fetchProjects} from "@/services/projects";
import {
    Link,
} from '@nextui-org/react';
import {useEffect, useState} from "react";
import {Breadcrumbs, TableBody, Checkbox, Tooltip} from "@mui/material";
import Typography from "@mui/material/Typography";
import {CircularProgress, Chip} from "@nextui-org/react";
import {DataGrid, GridColDef, GridToolbar} from '@mui/x-data-grid';
import {useQuery} from "react-query";
import {fetchLangChainFilterKeys, fetchLangChainRuns} from "@/services/langchain_runs";
import FilterPanel from "@/app/observe/[projectId]/FilterPanel";
import moment from "moment";
import "@/app/globals.css"
import {WhiteBackgroundTooltip} from "@/app/components/WhiteBackgrounToolTip";


const columns: GridColDef[] = [
    { field: 'chain_id', headerName: 'ID', width: 150, renderCell: (params) => (
            <Link href={`/observe/${params.row.project_id}/runs/${params.row.chain_id}`} rel="noopener">
                {params.value}
            </Link>
        ),},
    { field: 'chain_tracked_at', headerName: 'Tracked At ⏱️', width: 150, renderCell: (params) => {
            // Getting the labels from params.value
            const tracked_at = moment(params.row.chain_tracked_at).local()
            return (<div>{tracked_at.fromNow()}</div>)
        }},

    { field: 'duration', headerName: 'Run Duration ⏳', width: 150, renderCell: (params) => {
            // Getting the labels from params.value
            const startTime =  new Date(Date.parse(params.row.first_step_start_time)).getTime();
            const endTime =  new Date(Date.parse(params.row.last_step_end_time)).getTime();
            const duration = endTime - startTime
            return (<div>{!isNaN(duration) ? `${duration/1000} secs`: 'N/A'}</div>)
        }
    },

    // { field: 'first_step_start_time', headerName: 'Run Started At', width: 150},
    // { field: 'last_step_end_time', headerName: 'Run Ends At', width: 150},
    { field: 'total_tokens', headerName: 'Total Tokens', width: 120, },
    { field: 'prompt_tokens', headerName: 'Prompt Tokens', width: 120, },
    { field: 'completion_tokens', headerName: 'Completion Tokens', width: 100, },
    {
        field: 'model_info', // Replace with the actual column name
        headerName: 'Model Info', // Replace with your column's header
        width: 200, // Adjust this as needed
        renderCell: (params) => {
            const chips = params.row.model_info !== null ? Object.entries(params.row.model_info).map(
                ([key, value], index) => (typeof value === 'string') ? (
                    <Chip
                        key={index}
                        variant="bordered"
                        size="sm"
                    >{`${value}`}</Chip>
                ): null): ''
            return (<WhiteBackgroundTooltip title={<div>{chips}</div>} enterDelay={500}>
                <div style={{display: 'flex', flexWrap: 'wrap', gap: '4px', background: 'white'}}>
                    {chips}
                </div>
            </WhiteBackgroundTooltip>)
        },
    },{
        field: 'labels', // Replace with the actual column name
        headerName: 'Labels', // Replace with your column's header
        width: 200, // Adjust this as needed
        renderCell: (params) => {
            const chips = params.row.labels !== null ? Object.entries(params.row.labels).map(
                ([key, value], index) => (typeof value === 'string') ? (
                    <Chip
                        key={index}
                        color="primary"
                        variant="flat"
                        size="sm"
                    >{`${value}`}</Chip>
                ): null): ''
            return (<WhiteBackgroundTooltip style={{background: 'white'}} title={<div>{chips}</div>} enterDelay={500}>
                <div style={{display: 'flex', flexWrap: 'wrap', gap: '4px'}}>
                    {chips}
                </div>
            </WhiteBackgroundTooltip>)
        },
    },{
        field: 'runtime', // Replace with the actual column name
        headerName: 'Compute', // Replace with your column's header
        width: 200, // Adjust this as needed
        renderCell: (params) => {
            const chips = params.row.runtime !== null ? Object.entries(params.row.runtime).map(
                ([key, value], index) => (typeof value === 'string') ? (
                    <Chip
                        key={index}
                        variant="bordered"
                        size="sm"
                    >{`${value}`}</Chip>
                ): null): ''
            return (<WhiteBackgroundTooltip style={{background: 'white'}} title={<div>{chips}</div>} enterDelay={500}>
                <div style={{display: 'flex', flexWrap: 'wrap', gap: '4px', background: 'white'}}>
                    {chips}
                </div>
            </WhiteBackgroundTooltip>)
        },
    },
    // { field: 'labels', headerName: 'Labels', width: 250, renderCell: (params) => {
    //         // Getting the labels from params.value
    //         const labels = {...params.row.labels, ...params.row.runtime}
    //         return (
    //             <div style={{ display: 'flex', flexWrap: 'wrap', gap: '4px' }}>
    //                 {Object.entries(labels).map(([key, value], index) => (
    //                     <Chip
    //                         key={index}
    //                         label={`${value}`}
    //                         variant="outlined"
    //                         size="small"
    //                     />
    //                 ))}
    //             </div>
    //         );
    //     }
    // },

    // { field: 'total_runs', headerName: 'Total Runs', width: 130 },
    // { field: 'total_tokens', headerName: 'Total Tokens', width: 130 },
    // { field: 'total_p50', headerName: 'p50 latency', width: 130 },
    // { field: 'total_p95', headerName: 'p95 latency', width: 130 },
    // { field: 'latest_run_latency', headerName: 'latest runs latency', width: 130 },

];



// @ts-ignore
export default function LangChainRuns({ params }: { params: { projectId: string } }) {

    const [page, setPage] = useState(1);
    const [rowsPerPage, setRowsPerPage] = useState(4);
    const [searchText, setSearchText] = useState('');
    const [filterLabels, setFilterLabels] = useState({});
    const [filters, setFilters] = useState({});
    const [hasMoreData, setHasMoreData] = useState(true);



    // const { data: langChainRuns, error, isLoading } = useQuery(
    //     ['getLangChainRuns', page, rowsPerPage],
    //     () => fetchLangChainRuns(params.projectId, page, rowsPerPage),
    //     {
    //         keepPreviousData: true // Enable this to keep old data visible while fetching new data
    //     }
    // );

    const { data: langChainRuns, error, isLoading } = useQuery(
        ['getLangChainRuns', page, rowsPerPage],
        async () => {
            const data = await fetchLangChainRuns(params.projectId, page, rowsPerPage);
            if (!data || data.length < rowsPerPage) {
                setHasMoreData(false);
            }
            return data;
        },
        {
            keepPreviousData: true
        }
    );

    const { data: filterOptions, error: filterOptionsError, isLoading: filterOptionsLoading } = useQuery(
        ['getLangChainFilterOptions', params.projectId, page, rowsPerPage],
        () => fetchLangChainFilterKeys(params.projectId, page, rowsPerPage),
        {
            keepPreviousData: true, // Enable this to keep old data visible while fetching new data
            cacheTime: Infinity,

        }
    );




    const handleSearchChange = (e: any) => {
        setSearchText(e.target.value);
    };


    if (isLoading) return <CircularProgress aria-label="Loading..." />;
    if (error) return <div>Error loading data</div>;

    const handleFilterChange = (name: string, value: string) => {
        setFilters({
            ...filters,
            [name]: value,
        });
    };

    const applyFilters = () => {
        // apply the filters to the DataGrid here
    };


    return (
        <Layout>

            {/* Header */}
            <Typography variant="h5" gutterBottom>
                Runs
            </Typography>

            {/* Table */}
            <div style={{display: 'flex', width: '100%'}}>
                <div style={{ flexGrow: 1, width: '65%' }}>
                <DataGrid

                    rows={langChainRuns}
                    columns={columns}
                    getRowId={(row) => row.chain_id}
                    initialState={{
                        pagination: {
                            paginationModel: { page: 1, pageSize: 10 },
                        },
                        filter: {
                            filterModel: {
                                items: [],
                                quickFilterExcludeHiddenColumns: true,
                            },
                        },
                    }}
                    pageSizeOptions={[10, 15, 20, 5]}

                    onStateChange={(state) => {
                        if (state.pagination.page !== page) {
                            if (hasMoreData) {
                                setPage(state.pagination.page);
                            }
                            // setPage(state.pagination.page);
                        }
                        if (state.pagination.pageSize !== rowsPerPage) {
                            setRowsPerPage(state.pagination.pageSize);
                        }
                    }}
                    // onPageChange={(params) => {
                    //     if (hasMoreData || params.page + 1 < page) {
                    //         setPage(params.page + 1);
                    //     }
                    // }}
                    // onPageSizeChange={(params) => setRowsPerPage(params.pageSize)}
                    checkboxSelection
                    slots={{ toolbar: GridToolbar }}
                    slotProps={{
                        toolbar: {
                            showQuickFilter: true,
                        },
                    }}
                />
                </div>
                {/* Search and Filter Section */}
                <div style={{ padding: '10px', width: '35%' }}>
                    {/*<Input*/}
                    {/*    clearable*/}
                    {/*    placeholder="Search"*/}
                    {/*    value={searchText}*/}
                    {/*    onChange={handleSearchChange}*/}
                    {/*/>*/}
                    {/*<FilterPanel projectId={params.projectId}/>*/}

                </div>
            </div>




        </Layout>
    )
}
