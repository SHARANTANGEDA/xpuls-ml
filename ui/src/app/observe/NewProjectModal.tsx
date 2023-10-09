import React, { useState } from 'react';
import { Chip, TextField, Modal, Box, Typography } from '@mui/material';
import {Button} from "@nextui-org/react";

interface ProjectModalProps {
    isOpen: boolean;
    onClose: () => void;
}

const ProjectModal: React.FC<ProjectModalProps> = ({ isOpen, onClose }) => {
    const [projectName, setProjectName] = useState<string>('');
    const [tag, setTag] = useState<string>('');
    const [tags, setTags] = useState<string[]>([]);

    const addTag = () => {
        if (tag && !tags.includes(tag)) {
            setTags(prev => [...prev, tag]);
        }
        setTag('');
    };

    const removeTag = (tagToRemove: string) => {
        setTags(prev => prev.filter(t => t !== tagToRemove));
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
                <Typography variant="h6" gutterBottom>
                    Create New Project
                </Typography>
                <TextField
                    fullWidth
                    margin="normal"
                    label="Project Name"
                    value={projectName}
                    onChange={(e) => setProjectName(e.target.value)}
                />
                <TextField
                    fullWidth
                    margin="normal"
                    label="Project Slug"
                    value={projectName}
                    onChange={(e) => setProjectName(e.target.value)}
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
