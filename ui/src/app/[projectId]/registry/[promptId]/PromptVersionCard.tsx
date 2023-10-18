"use client"

import React from 'react';
import {Card, Pagination, CardHeader, CardBody, CardFooter, Button, Chip} from '@nextui-org/react';
import {Grid, Tooltip} from '@mui/material';
import LocalOfferOutlinedIcon from '@mui/icons-material/LocalOfferOutlined';
import moment from "moment/moment";
import CommitOutlinedIcon from "@mui/icons-material/CommitOutlined";

interface PromptVersionProps {
    item: PromptVersion
    handlePromptChange: (setPrompt: PromptVersion) => void
}

const PromptVersionCard: React.FC<PromptVersionProps> = ({item, handlePromptChange}) => {
    const version_created_at = moment(item.prompt_version_created_at).local().fromNow()

    return (<Card style={{minWidth: '100%'}} className="rounded-none w-100">
            <CardHeader className=" flex justify-between">
                <div>
                    <h5 className="truncate w-18 text-blue-500"><CommitOutlinedIcon/> {item.prompt_version_id.substring(0,7)}</h5>
                </div>
                <div>
                    <p className="font-semibold text-default-400 text-small">{version_created_at}</p>
                </div>

            </CardHeader>
            <CardBody className="px-3 py-0 flex flex-row justify-between">
                {/*<div className="flex gap-1">*/}
                {/*    <p className=" text-default-400 text-small">Modified by:</p>*/}
                {/*    <p className="font-semibold text-default-400 text-small">{item.prompt_version_id}</p>*/}
                {/*</div>*/}
                <div className="flex gap-1">
                    {item.prompt_tag ? <Chip
                        startContent={<LocalOfferOutlinedIcon />}
                        variant="flat"
                        color="success"
                        radius="sm"
                    >
                        {item.prompt_tag}
                    </Chip> : <p></p>}
                    {/*<p className=" text-default-400 text-small">Tag:</p>*/}
                    {/*<p className="font-semibold text-default-400 text-small">{item.prompt_tag}</p>*/}
                </div>
                <div className="flex gap-1">
                    <p className="text-default-400 text-small">Release: </p>

                    <p className="font-semibold text-default-400 text-small">test</p>
                </div>
            </CardBody>
            <CardFooter className="gap-3 flex justify-end">
                <Button onClick={() => handlePromptChange(item)} variant="flat" color="secondary" size="sm">
                    View
                </Button>

            </CardFooter>

        </Card>
    )
};

export default PromptVersionCard;
