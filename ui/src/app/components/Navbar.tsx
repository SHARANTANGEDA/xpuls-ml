import React from "react";

import {Navbar, NavbarBrand, NavbarContent, NavbarItem, Link, Button, NavbarMenuToggle} from "@nextui-org/react";

import './Navbar.module.css'
import {Image} from "@nextui-org/react";

// Define the props type
interface XpulsNavbarProps {
    handleDrawerToggle: () => void;
    drawerWidth: number;
    miniDrawerWidth: number;
    open: boolean;
}


export default function XpulsNavbar({ handleDrawerToggle, drawerWidth, miniDrawerWidth, open }: XpulsNavbarProps) {
    return (
        <Navbar isBordered maxWidth={"full"} position={"sticky"}
                style={{
                    marginLeft: open ? drawerWidth : miniDrawerWidth ,
                    zIndex: 1600, // Set a high z-index value here
                    // width: `calc(100% - ${open ? drawerWidth : miniDrawerWidth}px)` // Set the width

                }}>
            {/*<Toolbar>*/}
            {/*<NavbarMenuToggle color="inherit"*/}
            {/*                  aria-label="open drawer"*/}
            {/*                  onClick={handleDrawerToggle}*/}
            {/*                  edge="start">*/}
            {/*    */}

            {/*</NavbarMenuToggle>*/}


            {/*</Toolbar>*/}
            <NavbarBrand>
                {/*<IconButton*/}
                {/*    color="inherit"*/}
                {/*    aria-label="open drawer"*/}
                {/*    onClick={handleDrawerToggle}*/}
                {/*    edge="start"*/}
                {/*>*/}
                {/*    <MenuIcon />*/}
                {/*</IconButton>*/}
                <Image style={{marginRight: '30px'}}
                    width={36}
                    height={36}
                    src="/favicon.ico"
                    alt="XpulsAI Logo"
                />

                <h2 className="font-bold text-inherit">Xpuls.ai</h2>
            </NavbarBrand>
            {/*<NavbarContent className="hidden sm:flex gap-4" justify="center">*/}
            {/*    <NavbarItem>*/}
            {/*        <Link color="foreground" href="#">*/}
            {/*            Features*/}
            {/*        </Link>*/}
            {/*    </NavbarItem>*/}
            {/*    <NavbarItem isActive>*/}
            {/*        <Link href="#" aria-current="page">*/}
            {/*            Customers*/}
            {/*        </Link>*/}
            {/*    </NavbarItem>*/}
            {/*    <NavbarItem>*/}
            {/*        <Link color="foreground" href="#">*/}
            {/*            Integrations*/}
            {/*        </Link>*/}
            {/*    </NavbarItem>*/}
            {/*</NavbarContent>*/}
            {/*<NavbarContent justify="end">*/}
            {/*    <NavbarItem className="hidden lg:flex">*/}
            {/*        <Link href="#">Login</Link>*/}
            {/*    </NavbarItem>*/}
            {/*    <NavbarItem>*/}
            {/*        /!*<Button as={Link} color="primary" href="#" variant="flat">*!/*/}
            {/*        /!*    Sign Up*!/*/}
            {/*        /!*</Button>*!/*/}
            {/*    </NavbarItem>*/}
            {/*</NavbarContent>*/}

        </Navbar>
    );
}
