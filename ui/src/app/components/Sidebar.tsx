// src/app/components/Sidebar.tsx
import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import Drawer from '@mui/material/Drawer';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import InboxIcon from '@mui/icons-material/MoveToInbox';
import MailIcon from '@mui/icons-material/Mail';
import {ChevronRightIcon} from "@nextui-org/shared-icons";
import {ListItemButton} from "@mui/material";
import {products} from "@/app/components/products";
import Link from 'next/link'

interface XpulsSidebarProps {
    handleDrawerToggle: () => void;
    drawerWidth: number;
    miniDrawerWidth: number;
    open: boolean;
}

export default function Sidebar({handleDrawerToggle, drawerWidth, miniDrawerWidth, open}: XpulsSidebarProps) {

    return (
        <Drawer
                variant="permanent"
                sx={{
                    width: open ? drawerWidth : miniDrawerWidth,
                    flexShrink: 0,
                    '& .MuiDrawer-paper': {
                        width: open ? drawerWidth : miniDrawerWidth,
                        boxSizing: 'border-box',
                        marginTop: '4rem',
                        marginBottom: '0rem',
                    },
                }}
            >


            <List>
                {products.map((item, index) => (
                    <ListItem key={item.name}  disablePadding
                              sx={{ display: 'block', color: 'black' }}>
                        <Link href={item.url}>

                        <ListItemButton>
                            <ListItemIcon>
                                {item.icon}
                            </ListItemIcon>
                            <ListItemText primary={item.name} sx={{ opacity: open ? 1 : 0 }} />
                        </ListItemButton>
                        </Link>
                    </ListItem>

                ))}
                </List>

            <ListItem button onClick={handleDrawerToggle} sx={{
                position: 'absolute', // Absolute positioning
                bottom: '5rem', // Anchor to the bottom
                left: 0, // Anchor to the left
            }}>
                <ListItemIcon>
                    {open? <ChevronLeftIcon /> : <ChevronRightIcon />}
                </ListItemIcon>
            </ListItem>
            </Drawer>
    );
};
