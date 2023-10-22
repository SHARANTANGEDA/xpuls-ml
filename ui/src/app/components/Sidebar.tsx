// src/app/components/Sidebar.tsx
import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import Drawer from '@mui/material/Drawer';
import IconButton from '@mui/material/IconButton';
import {Divider} from "@nextui-org/react";
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import {ChevronRightIcon} from "@nextui-org/shared-icons";
import {ListItemButton} from "@mui/material";import Link from 'next/link'
import {Image, NavbarBrand} from "@nextui-org/react";
import {products, socials} from "@/utils/feature_groups";
import { SocialIcon } from 'react-social-icons'
import {BsChevronLeft, BsChevronRight} from 'react-icons/bs';

interface XpulsSidebarProps {
    handleDrawerToggle: () => void;
    drawerWidth: number;
    miniDrawerWidth: number;
    open: boolean;
}

export default function Sidebar({handleDrawerToggle, drawerWidth, miniDrawerWidth, open}: XpulsSidebarProps) {
    const [sopen, setOpen] = React.useState(true);

    const bottomProps = open ? "": ""
    return (
        <Drawer
                variant="permanent"
                sx={{
                    width: open ? drawerWidth : miniDrawerWidth,
                    flexShrink: 0,
                    '& .MuiDrawer-paper': {
                        width: open ? drawerWidth : miniDrawerWidth,
                        boxSizing: 'border-box',
                        // marginTop: '4rem',
                        marginBottom: '0rem',
                    },
                }}
            >
            {open ? (<div onClick={() => window.location.href = "/"}
                          className="flex flex-row justify-center align-center items-center mt-2"
                           style={{ cursor: 'pointer'}}>
                <Image style={{marginRight: '16px'}}
                       width={42}
                       height={42}
                       src="/xpulsai.png"
                       alt="XpulsAI Logo"
                />
                <div className="flex flex-row"><p className="text-transparent bg-clip-text text-3xl font-bold
                 bg-gradient-to-r from-cyan-600 to-blue-600 ">xpuls.</p> <p className="text-transparent bg-clip-text text-3xl font-bold
                 bg-gradient-to-r from-indigo-600 to-pink-600">ai</p></div>

            </div>): (
                <div className="flex flex-row justify-center items-center">
                    <Image style={{marginRight: '16px'}}
                           width={52}
                           height={52}
                           src="/xpulsai.png"
                           alt="XpulsAI Logo"
                    />
                </div>
            )}
            <ListItem>
                <Divider className="my-2" />
            </ListItem>


            <List>
                {products.map((item, index) => (
                    <ListItem className="flex " key={item.name}  disablePadding
                              sx={{ display: 'block', color: 'black' }}>
                        <Link href={item.url}>

                        <ListItemButton>
                            <ListItemIcon>
                                {item.icon}
                            </ListItemIcon>
                            {/*<ListItemText primary={item.name} sx={{ opacity: open ? 1 : 0, fontWeight: "bold" }} />*/}
                            <p className="font-semibold" style={{ opacity: open ? 1 : 0}}>{item.name}</p>
                        </ListItemButton>
                        </Link>
                    </ListItem>

                ))}
                </List>
            <ListItem>
                <Divider className="my-4  " />
            </ListItem>


            <List className="fixed bottom-0 w-full" style={{ position: "absolute", bottom: "0" }}>
                <ListItem>
                    <Divider className="my-4  " />
                </ListItem>
                {open? (
                    <List className="flex relative flex-row bottom-0 w-full">
                        {
                            socials.map((item, index) => (
                                <ListItemButton target="_blank" href={item.url}  className="hover:bg-transparent"> <ListItemIcon className="hover:bg-transparent"> {item.icon}</ListItemIcon></ListItemButton>
                            ))
                        }
                    </List>
                ): (<List className="flex flex-col bottom-0 w-full">
                    {
                        socials.map((item, index) => (
                            <ListItemButton target="_blank" href={item.url} className="hover:bg-transparent"> <ListItemIcon className="hover:bg-transparent"> {item.icon}</ListItemIcon></ListItemButton>
                        ))
                    }

                </List>)}
                <ListItem>
                    <Divider className="my-4  " />
                </ListItem>

                <ListItemButton onClick={handleDrawerToggle} >
                    <ListItemIcon className="flex-row justify-start">
                        {open? <BsChevronLeft size={24}/> : <BsChevronRight size={24}/>}
                    </ListItemIcon>
                </ListItemButton>
            </List>


            </Drawer>
    );
};
