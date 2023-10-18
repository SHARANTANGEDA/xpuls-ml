import React, {useEffect, useState} from 'react';
import { Select, MenuItem, Chip, TextField, Modal, Box, Typography, InputLabel,
    FormControl } from '@mui/material';
import {Button, CircularProgress} from "@nextui-org/react";
import { SelectItem} from "@nextui-org/react";
import {useMutation, useQuery} from "react-query";
import {fetchProjects} from "@/services/projects";
import {addNewPrompt} from "@/services/prompt_registry";

interface ProjectModalProps {
    isOpen: boolean;
    projectId: string;
    onClose: () => void;
}

const PromptModal: React.FC<ProjectModalProps> = ({ isOpen, projectId, onClose }) => {
    const [promptName, setPromptName] = useState<string>('');
    const [selectedProject, setSelectedProject] = useState<string>(projectId);

    const useAddPrompt = useMutation(addNewPrompt, {
        onSuccess: () => {
            window.location.reload()
            onClose();
        },
        onError: (error) => {
            console.error('Error adding prompt:', error);
        },
        onSettled: () => {
            // This will run whether the mutation is successful or fails
            console.log('Finished adding prompt');
        }
    });

    const handleSubmit = () => {
        // Handle the submission logic here (e.g., API call to create a new project)
        useAddPrompt.mutate({project_id: selectedProject, prompt_name: promptName })

    };

    const { data: projects, error, isLoading } = useQuery(
        ['projects'],
        () => fetchProjects(1, 1000),
        {
            keepPreviousData: true // Enable this to keep old data visible while fetching new data
        }
    );

    if (isLoading || error || projects === undefined) {
        return <CircularProgress aria-label="Loading..." />
    }


    return (
        <Modal open={isOpen} onClose={onClose}>
            <Box sx={{ width: 400, padding: 3, bgcolor: 'background.paper', margin: 'auto', marginTop: '10%' }}>
                <p className="text-black text-2xl">
                    Add New Prompt
                </p>
                <TextField required={true}
                    fullWidth
                    margin="normal"
                    label="Prompt Name"
                    value={promptName}
                    onChange={(e) => setPromptName(e.target.value)}
                />
                <FormControl required={true} className="max-w-xl" sx={{ m: 1, minWidth: 250 }}>

                <InputLabel id="simple-select-helper-label">Select Project</InputLabel>
                <Select
                    className="max-w-xl"
                    labelId="simple-select-outlined-label"
                    id="simple-select-outlined"
                    value={selectedProject}
                    onChange={(e) => setSelectedProject(e.target.value)}
                    label="Select Project"
                >
                    {projects.map((project) => (
                        <MenuItem key={project.project_id} value={project.project_id}>
                            {project.project_name}
                        </MenuItem>
                    ))}
                </Select>
                </FormControl>


                <Box mt={2} display="flex" justifyContent="space-between">
                    <Button color="danger" variant="bordered" onClick={onClose}>
                        Close
                    </Button>
                    <Button color="primary" onClick={handleSubmit}>
                        Create
                    </Button>
                </Box>
            </Box>
        </Modal>
    );
};

export default PromptModal;
