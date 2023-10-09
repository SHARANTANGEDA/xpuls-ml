'use client'

import {Breadcrumbs, Typography} from '@mui/material';
import Link from 'next/link';
import NavigateNextIcon from '@mui/icons-material/NavigateNext';
import { usePathname } from 'next/navigation'
import HomeIcon from "@mui/icons-material/Home";

function AutoBreadcrumbs() {
    const pathname = usePathname()

    const pathnames = pathname.split('/').filter(x => x);

    return (
            <Breadcrumbs separator={<NavigateNextIcon fontSize="small"/>}>
                <Link href="/">
                    <div className="flex items-center justify-center">
                        <HomeIcon className="mr-2" />
                        <span className="text-current">Home</span>
                    </div>
                    {/*<Typography color="inherit"> <HomeIcon/> Home</Typography>*/}
                </Link>
                {pathnames.map((value, index) => {
                    const last = index === pathnames.length - 1;
                    const to = `/${pathnames.slice(0, index + 1).join('/')}`;

                    return last ? (
                        <Typography color="textPrimary" key={to}>
                            {value}
                        </Typography>
                    ) : (
                        <Link href={to} key={to} passHref>
                            <Typography color="inherit">{value}</Typography>
                        </Link>
                    );
                })}
            </Breadcrumbs>
    );
}

export default AutoBreadcrumbs;
