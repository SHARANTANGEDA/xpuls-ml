'use client'

import {NextUIProvider} from '@nextui-org/react'
import {QueryClient, QueryClientProvider} from 'react-query';
import {RecoilRoot} from "recoil";

const queryClient = new QueryClient();

export function Providers({children}: { children: React.ReactNode }) {
    return (
        <RecoilRoot>
            <QueryClientProvider client={queryClient}>
                <NextUIProvider>
                    {children}
                </NextUIProvider>
            </QueryClientProvider>
        </RecoilRoot>

    )
}
