import React, { useState } from 'react';
import { Chip, TextField, Modal, Box, Typography } from '@mui/material';
import {Button} from "@nextui-org/react";
import { useRecoilState } from 'recoil';
import { useMutation } from 'react-query';
import {isProjectSlugAvailable} from "@/services/projects";
import {convertToAlphanumericWithUnderscore} from "@/utils/common"

interface ProjectModalProps {
    isOpen: boolean;
    onClose: () => void;
}

const ProjectModal: React.FC<ProjectModalProps> = ({ isOpen, onClose }) => {
    const [projectName, setProjectName] = useState<string>('');
    const [projectSlug, setProjectSlug] = useState("");
    const [slugError, setErrorForSlug] = useState("");
    const [tag, setTag] = useState<string>('');
    const [tags, setTags] = useState<string[]>([]);

    const validateSlugMutation = useMutation(isProjectSlugAvailable, {
        onSuccess: (isValid) => {
            // Handle success. You can set any state or perform side effects based on the API response here.
            // Assuming the API returns { isValid: true/false, message: 'error message' }
            if (isValid) {
                setErrorForSlug("Slug must be unique");
            } else {
                setErrorForSlug(''); // Clear the error
            }
        },
        onError: (error) => {
            // Handle error. For example, set an error message.
            setErrorForSlug('Failed to validate slug');
        },
    });

    const addTag = () => {
        if (tag && !tags.includes(tag)) {
            setTags(prev => [...prev, tag]);
        }
        setTag('');
    };

    const removeTag = (tagToRemove: string) => {
        setTags(prev => prev.filter(t => t !== tagToRemove));
    };

    const updateProjectName = (projectName: string) => {
        setProjectName(projectName);

        setProjectSlug(convertToAlphanumericWithUnderscore(projectName))
        // validateSlugMutation.mutate(projectSlug)
    };

    const handleSubmit = () => {
        // Handle the submission logic here (e.g., API call to create a new project)
        console.log('Project Name:', projectName);
        console.log('Tags:', tags);
        onClose();
    };

    return (
        <Modal open={isOpen} onClose={onClose}>
            <Box sx={{ width: 400, padding: 3, bgcolor: 'background.paper', margin: 'auto', marginTop: '10%' }}>
                <p className="text-2xl text-black">
                    Create New Project
                </p>
                <TextField
                    fullWidth
                    margin="normal"
                    label="Project Name"
                    value={projectName}
                    onChange={(e) => updateProjectName(e.target.value)}
                    onBlur={() => validateSlugMutation.mutate(projectSlug)}
                />
                <TextField
                    fullWidth
                    margin="normal"
                    label="Project Slug"
                    value={projectSlug}
                    onChange={(e) => setProjectSlug(e.target.value)}
                    onBlur={() =>  validateSlugMutation.mutate(projectSlug)}
                    error={Boolean(slugError)}
                    helperText={slugError}
                />
                <TextField
                    fullWidth
                    margin="normal"
                    label="Add Tag"
                    value={tag}
                    onChange={(e) => setTag(e.target.value)}
                    onKeyPress={(e) => e.key === 'Enter' && addTag()}
                />
                {tags.map((t, index) => (
                    <Chip key={index} label={t} onDelete={() => removeTag(t)} style={{ marginRight: 8, marginTop: 8 }} />
                ))}
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

export default ProjectModal;
