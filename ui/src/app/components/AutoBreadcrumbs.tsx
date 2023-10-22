'use client'

import {Breadcrumbs, Typography} from '@mui/material';
import Link from 'next/link';
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import { usePathname } from 'next/navigation'
import HomeIcon from "@mui/icons-material/Home";
import { AiFillHome } from 'react-icons/ai';

function AutoBreadcrumbs() {
    const pathname = usePathname()

    const pathnames = pathname.split('/').filter(x => x);

    return (
            <Breadcrumbs separator={<NavigateNextIcon fontSize="small"/>}>
                <Link href="/">
                    <div className="flex items-center justify-center align-center">
                        <AiFillHome className="text-sky-800 mr-1" size={18}/>
                        <p className="text-sky-800 font-semibold"> Home</p>
                        {/*<HomeIcon className="mr-2" />*/}
                    </div>
                    {/*<Typography color="inherit"> <HomeIcon/> Home</Typography>*/}
                </Link>
                {pathnames.map((value, index) => {
                    const last = index === pathnames.length - 1;
                    const to = `/${pathnames.slice(0, index + 1).join('/')}`;

                    return last ? (
                        <p className="font-bold text-blue-700"  key={to}>
                            {value}
                        </p>
                    ) : (
                        <Link href={to} key={to} passHref>
                            <p >{value}</p>
                        </Link>
                    );
                })}
            </Breadcrumbs>
    );
}

export default AutoBreadcrumbs;
