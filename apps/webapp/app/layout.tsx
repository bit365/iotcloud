import * as React from 'react';
import { NextAppProvider } from '@toolpad/core/nextjs';
import { AppRouterCacheProvider } from '@mui/material-nextjs/v15-appRouter';
import DashboardIcon from '@mui/icons-material/Dashboard';
import PersonIcon from '@mui/icons-material/Person';

import type { Navigation } from '@toolpad/core/AppProvider';
import { SessionProvider, signIn, signOut } from 'next-auth/react';
import { auth } from '../auth';
import theme from '../theme';

export const metadeta = {
  title: 'IotCloud',
  description: 'This is a sample app built with Toolpad Core and Next.js',
};

const NAVIGATION: Navigation = [
  {
    kind: 'header',
    title: 'Main items',
  },
  {
    segment: '',
    title: 'Dashboard',
    icon: <DashboardIcon />,
  },
  {
    segment: 'employees',
    title: 'Employees',
    icon: <PersonIcon />,
    pattern: 'employees{/:employeeId}*',
  },
];

const BRANDING = {
  title: 'IotCloud',
};


const AUTHENTICATION = {
  signIn,
  signOut,
};


export default async function RootLayout(props: { children: React.ReactNode }) {
  const session = await auth();

  return (
    <html lang="en" data-toolpad-color-scheme="light" suppressHydrationWarning>
      <body>
        <SessionProvider session={session}>
          <AppRouterCacheProvider options={{ enableCssLayer: true }}>
          
            <NextAppProvider
              navigation={NAVIGATION}
              branding={BRANDING}
              session={session}
              authentication={AUTHENTICATION}
              theme={theme}
            >
              {props.children}
            </NextAppProvider>
            
          </AppRouterCacheProvider>
        </SessionProvider>
      </body>
    </html>
  );
}
