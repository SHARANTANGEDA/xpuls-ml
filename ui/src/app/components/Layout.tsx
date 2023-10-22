'use client'

import styles from './Layout.module.css';
import Sidebar from './Sidebar';
import XpulsNavbar from "@/app/components/Navbar";
import * as React from "react";
import {Providers} from "@/app/providers";
import {Inter} from "next/font/google";

const inter = Inter({ subsets: ['latin'] })

export default function Layout({children,}: {
    children: React.ReactNode
}) {
    const [open, setOpen] = React.useState(true);
    const drawerWidth = 220;
    const miniDrawerWidth = 60;



    const handleDrawerToggle = () => {
        setOpen(!open);
    };
    return (
            <div className={styles.layout} style={{ display: 'flex', flexDirection: 'column' }}>
            {/* Navbar Component */}
            <XpulsNavbar
                handleDrawerToggle={handleDrawerToggle}
                drawerWidth={drawerWidth}
                miniDrawerWidth={miniDrawerWidth}
                open={open}
            />

            {/* Container for Sidebar and Main Content */}
            <div style={{ display: 'flex', flex: 1}}>
                {/* Sidebar Component */}
                <Sidebar
                    handleDrawerToggle={handleDrawerToggle}
                    drawerWidth={drawerWidth}
                    miniDrawerWidth={miniDrawerWidth}
                    open={open}
                />

                {/* Main Content */}
                <div className={styles.content} style={{ flex: 1, marginLeft: '5px' }}>
                    {children}
                </div>
            </div>
        </div>
    );
}
