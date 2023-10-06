import React, { useState } from 'react';
import {Button, Input, Select, SelectItem} from "@nextui-org/react";
import {useQuery} from "react-query";
import {fetchLangChainFilterKeys} from "@/services/langchain_runs";
import {AiFillPlusCircle} from "react-icons/ai";

const filterKeys = [
    "runtime.library_version",
    "runtime.runtime",
    "runtime.sdk_version",
    "runtime.py_implementation",
    "runtime.runtime_version",
    "runtime.platform",
    "runtime.library",
    "runtime.langchain_version"
];



const conditions = [
    {conditionName: "Equals", conditionValue: "="},
    {conditionName: "Contains", conditionValue: "ILIKE"},
    {conditionName: "Not Contains", conditionValue: "NOT ILIKE"},
]

interface FilterBarProps {
    projectId: string
}
export default function FilterPanel(
    {projectId}: FilterBarProps
) {
    const [filters, setFilters] = useState([
        { key: '', condition: '', value: '' }
    ]);

    const { data: filterOptions, error: filterOptionsError, isLoading: filterOptionsLoading } = useQuery(
        ['getLangChainFilterOptions', projectId],
        () => fetchLangChainFilterKeys(projectId, 1, 1000),
        {
            keepPreviousData: true, // Enable this to keep old data visible while fetching new data
            cacheTime: Infinity,

        }
    );

    let filterLabels: any[] = []
    if (filterOptions !== null && filterOptions !== undefined) {
        filterLabels = [...filterOptions.runtime_keys, ...filterOptions.label_keys].map(label => ({ label, value: label }))
    }

    const handleApplyFilter = () => {
        // Convert each filter to a part of a SQL query and join them
        const sqlQuery = filters.map(
            filter => `${filter.key} ${filter.condition} '${filter.value}'`
        ).join(' AND ');
    };

    const handleAddFilter = () => {
        setFilters([...filters, { key: '', condition: '', value: '' }]);
    };

    const handleChange = (index: number, field: string, value: string) => {
        const newFilters: any[] = [...filters];
        newFilters[index][field] = value;
        setFilters(newFilters);
    };

    return (
        <div>
            <h5>Advanced Filters: [Under Construction] </h5>
            {filters.map((filter, index) => (
                <div style={{display: 'flex'}} key={index}>
                    <Select
                        color={"primary"}
                        size={"sm"}
                        items={filterLabels}
                        placeholder="Select Label"
                        className="max-w-xs"
                        style={{marginRight: '5px', minWidth: '200px'}}
                    >
                        {(filter) => <SelectItem style={{ color: 'black' }} key={filter.value} value={filter.value}>{filter.label}</SelectItem>}
                    </Select>
                    <Select
                        color={"success"}
                        size={"sm"}
                        items={conditions}
                        defaultSelectedKeys={[conditions[0].conditionValue]}
                        placeholder="Condition"
                        style={{maxWidth: '125px', }}

                        // onChange={(e) => handleChange(index, 'condition', e.target.value)}
                    >
                        {(condition)=> <SelectItem style={{ color: 'black' }} key={condition.conditionValue} value={condition.conditionValue}>
                            {condition.conditionName}</SelectItem>}
                    </Select>
                    {/*<SearchableSelect*/}
                    {/*    projectId={projectId}*/}
                    {/*    labelKey={filter.key}*/}
                    {/*    condition={filter.condition}*/}
                    {/*    index={index}*/}
                    {/*    handleChange={handleChange}*/}
                    {/*/>*/}
                    <Input
                        style={{minWidth: '200px', height: '100%'}}
                        color={"primary"}
                        placeholder="Value"
                        value={filter.value}
                        onChange={(e) => handleChange(index, 'value', e.target.value)}
                    />

                </div>


            ))}
            <div style={{display: 'flex', justifyContent: 'center'}} >
                <div style={{marginRight: '15px'}}>
                    <Button color="primary" variant="flat" isIconOnly onClick={handleAddFilter}><AiFillPlusCircle/></Button>
                </div>
                {/*<Button color="primary" variant="ghost" onClick={handleApplyFilter}>Apply Filters</Button>*/}

            </div>
            </div>
    );
}
