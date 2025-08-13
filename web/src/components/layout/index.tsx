import { FC } from 'react'
import {
  ActionIcon,
  Anchor,
  // Anchor,
  AppShell,
  Avatar,
  Burger,
  Center,
  Divider,
  Group,
  Loader,
  Menu,
  Text,
  Tooltip,
} from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'
import { Link as RouterLink, Outlet } from '@tanstack/react-router'
import { IconHelp } from '@tabler/icons-react'

import { Link } from '../link'
import classes from './styles.module.scss'

const isLoaded = true
const isSignedIn = true
const user = {
  firstName: 'John',
  lastName: 'Doe',
  imageUrl: 'https://via.placeholder.com/150',
  fullName: 'John Doe',
}

export const Layout: FC = () => {
  const [opened, { toggle }] = useDisclosure()
  // const { isLoaded, isSignedIn, signOut } = useAuth()
  // const { user } = useUser()

  if (!isLoaded) {
    return (
      <Center h='95vh'>
        <Loader />
      </Center>
    )
  }

  if (!isSignedIn || !user) {
    return (
      <Center h='95vh'>
        <h1>SIGN IN</h1>
      </Center>
    )
  }

  return (
      <AppShell
        header={{ height: 60 }}
        navbar={{ width: 300, breakpoint: 'sm', collapsed: { desktop: true, mobile: !opened } }}
        padding='md'
      >
        <AppShell.Header>
          <Group h='100%' px='md'>
            <Burger opened={opened} onClick={toggle} hiddenFrom='sm' size='sm' />
            <Group justify='space-between' style={{ flex: 1 }}>
              <Group visibleFrom='sm'>
                <Link
                  to='/'
                  activeProps={{
                    className: classes.activeLink,
                  }}
                  underline="never"
                  
                >
                  Home
                </Link>
              </Group>
              <Group>
                <RouterLink to='/'>
                  <Tooltip label='Help' openDelay={700}>
                    <ActionIcon variant='transparent'>
                      <IconHelp stroke={1.5} />
                    </ActionIcon>
                  </Tooltip>
                </RouterLink>
                {/* <Link to='/settings'>
                  <ActionIcon variant='transparent'>
                    <IconSettings stroke={1.5} />
                  </ActionIcon>
                </Link> */}
                <Menu>
                  <Menu.Target>
                    <Avatar src={user.imageUrl} style={{ cursor: 'pointer' }} size={30}>
                      {user.firstName?.charAt(0)}
                    </Avatar>
                  </Menu.Target>
                  <Menu.Dropdown>
                    <Menu.Label>{user.fullName}</Menu.Label>
                    {/* <Menu.Item>
                      <Link style={{ color: 'inherit', textDecoration: 'none' }} to='/'>Settings</Link>
                    </Menu.Item> */}
                    <Menu.Item onClick={() => {}} c='red'>
                      Logout
                    </Menu.Item>
                  </Menu.Dropdown>
                </Menu>
                <Divider orientation='vertical' />
                <Anchor href='https://www.rbac.com' target='_blank'>
                  <Text>RBAC Demo</Text>
                </Anchor>
              </Group>
            </Group>
          </Group>
        </AppShell.Header>

        <AppShell.Main bg='gray.2'>
          <Outlet />
        </AppShell.Main>
      </AppShell>
  )
}

