"use client"
// components/PaginatedCards.tsx

import React from 'react';
import { Card, Pagination, CardHeader, CardBody, CardFooter, Button } from '@nextui-org/react';
import { Grid, Tooltip } from '@mui/material';
import ReactMarkdown from 'react-markdown';
import moment from "moment";
import CommitOutlinedIcon from "@mui/icons-material/CommitOutlined";

interface PaginatedCardsProps {
    data: PromptVersion[];
    setPage: (page: number) => void;
    itemsPerPage: number;
    page: number;
}

const PaginatedCards: React.FC<PaginatedCardsProps> = ({ data, setPage, itemsPerPage = 10,
                                                           page=1 }) => {

    const handleChange = (page: number) => {
        setPage(page);
    };

    const paginatedData = data.slice((page - 1) * itemsPerPage, page * itemsPerPage);

    return (
        <div>
            <div className="grid auto-rows-auto grid-flow-row-dense grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4" >
                {paginatedData.map((item, index) => (

                        <Card key={index} className="max-w-[340px]">
                            <CardHeader className="flex justify-between">
                                    <div>
                                        <h5 className="font-semibold truncate w-48">{item.prompt_name}</h5>
                                    </div>
                                    <div>
                                        <Button onClick={() => (window.location.href=`/${item.project_id}/registry/${item.prompt_id}`)} className="col-end-1" color="primary" variant={"flat"} size="sm">
                                            Open
                                        </Button>
                                    </div>

                            </CardHeader>
                            <CardBody className="px-3 py-0 text-small text-default-400">
                                <div className="row-span-full flex gap-1">
                                    <p className="text-default-400 text-small">ID: </p>
                                    <Tooltip title="Click to copy 📋" placement="top">
                                        <p className="truncate font-semibold text-default-400 text-small cursor-pointer"
                                           onClick={() => navigator.clipboard.writeText(item.prompt_id)}>
                                            {item.prompt_id} 📋</p>
                                    </Tooltip>
                                </div>
                                <ReactMarkdown className="line-clamp-2">{item.prompt_content}</ReactMarkdown>
                            </CardBody>
                            <CardFooter className="gap-3 flex justify-between">
                                <div className="flex gap-1">
                                    <p className="font-semibold text-small text-blue-500"
                                    ><CommitOutlinedIcon/>  {item.prompt_version_id.substring(0, 7)}</p>
                                </div>
                                <div>
                                    <p className="font-semibold text-default-400 text-small">{moment(item.prompt_version_created_at).local().fromNow()}</p>
                                </div>

                                {/*<div className="flex gap-1">*/}
                                {/*    <p className="text-default-400 text-small">Last Modified by: </p>*/}

                                {/*    <p className="font-semibold text-default-400 text-small">{item.prompt_version_created_at}</p>*/}
                                {/*</div>*/}

                            </CardFooter>

                        </Card>
                ))}
            </div>

            <div className="flex justify-center mt-4">
                <Pagination initialPage={page} onChange={handleChange}  total={data.length}/>
            </div>
        </div>
    )
};

export default PaginatedCards;
