import { FC, PropsWithChildren } from 'react'
import { createTheme, MantineProvider } from '@mantine/core'

const theme = createTheme({})

export const ThemeProvider: FC<PropsWithChildren> = ({ children }) => {
  return <MantineProvider theme={theme}>{children}</MantineProvider>
}
