'use client'

import React, {useEffect, useState} from "react";
import {Chip} from "@nextui-org/react";
import {LangChainRunStep} from "@/types/langchain";

interface RunStepTreeProps {
    setCurrentRunStep: (runStep: LangChainRunStep) => void;
    runSteps: any;
    currentRunStep: any;
}

interface RunStepTreeNodeProps {
    setCurrentRunStep: (runStep: LangChainRunStep) => void;
    runSteps: any;
    currentRunStep: any;
    indentationLevel: number;
}

export const RunStepTree = ({runSteps, currentRunStep, setCurrentRunStep}: RunStepTreeProps) => {
    const nestedRunSteps = nestRunSteps(runSteps);
    return (
        <div className="flex flex-col">
            <RunStepTreeNode
                runSteps={nestedRunSteps}
                indentationLevel={1}
                currentRunStep={currentRunStep}
                setCurrentRunStep={setCurrentRunStep}
            />
        </div>
    );
};

const RunStepTreeNode = ({setCurrentRunStep, runSteps, currentRunStep, indentationLevel}: RunStepTreeNodeProps) => {
    const [formattedDates, setFormattedDates] = useState([]);
    const getChipStyles = (runType: string) => {
        switch (runType) {
            case 'chain':
                return { background: 'rgba(63, 81, 181, 0.1)', color: '#3f51b5' }; // light blue background, blue text
            case 'llm':
                return { background: 'rgba(76, 175, 80, 0.1)', color: '#4caf50' }; // light green background, green text
            case 'tool':
                return { background: 'rgba(255, 193, 7, 0.1)', color: '#ffc107' }; // light yellow background, yellow text
            default:
                return { background: 'rgba(0, 0, 0, 0.1)', color: '#000' }; // light black background, black text as default
        }
    };

    useEffect(() => {
        const dates = runSteps.map((step: any) =>
            new Date(step.event_start_time).toLocaleString()
        );
        setFormattedDates(dates);
    }, [runSteps]);
    return (
        <>
            {runSteps
                .sort((a: any, b: any) => new Date(a.event_start_time).getTime() - new Date(b.event_start_time).getTime())
                .map((runStep: LangChainRunStep, index: number) => (
                    <React.Fragment key={runStep.run_step_id}>
                        <div className="flex">
                            {Array.from({length: indentationLevel}, (_, i) => (
                                <div className="mx-2 border-r" key={i}/>
                            ))}
                            <div
                                className="group my-1 flex flex-1 cursor-pointer flex-col gap-1 rounded-sm p-2"
                                onClick={() => setCurrentRunStep(runStep)}
                            >
                                <div className="flex gap-2">
                                    <Chip radius="sm" className="line-clamp-1" style={getChipStyles(runStep.run_type)}>
                                        {runStep.run_type.toUpperCase()}
                                    </Chip>

                                    <span className="self-start rounded-sm bg-gray-100 p-1 text-xs">
                                    {runStep.run_name}
                                    </span>
                                </div>

                                <div className="flex gap-2">
                                <span className="text-xs text-gray-500">
                                    {formattedDates[index]}
                                </span>
                                </div>
                            </div>
                        </div>
                        <RunStepTreeNode
                            runSteps={runStep.children}
                            indentationLevel={indentationLevel + 1}
                            currentRunStep={currentRunStep}
                            setCurrentRunStep={setCurrentRunStep}
                        />
                    </React.Fragment>
                ))}
        </>
    );
}

export function nestRunSteps(list: any[]) {
    if (list.length === 0) return [];

    const map = new Map();
    for (const obj of list) {
        map.set(obj.run_step_id, { ...obj, children: [] });
    }

    const roots = new Map();

    for (const obj of map.values()) {
        if (obj.parent_step_id) {
            const parent = map.get(obj.parent_step_id);
            if (parent) {
                parent.children.push(obj);
            }
        } else {
            roots.set(obj.run_step_id, obj);
        }
    }

    return Array.from(roots.values());
}
