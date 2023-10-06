'use client'

import {NextUIProvider} from '@nextui-org/react'
import {QueryClient, QueryClientProvider} from 'react-query';
import {BrowserRouter as Router} from 'react-router-dom';

const queryClient = new QueryClient();

export function Providers({children}: { children: React.ReactNode }) {
    return (
            <QueryClientProvider client={queryClient}>
                <NextUIProvider>
                    {children}
                </NextUIProvider>
            </QueryClientProvider>

    )
}