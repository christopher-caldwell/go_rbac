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
  Button,
  Stack,
  Text,
  Title,
  Tooltip,
} from '@mantine/core'
import { useDisclosure, useLocalStorage } from '@mantine/hooks'
import { Link as RouterLink, Outlet } from '@tanstack/react-router'
import { IconHelp } from '@tabler/icons-react'

import { Link } from '../link'
import classes from './styles.module.scss'
import { useGetMe } from '@/hooks/use_get_user'
import { LS_TOKEN_KEY } from '@/api/client'

export const Layout: FC = () => {
  const [userId, setUserId] = useLocalStorage<string | null>({
    key: LS_TOKEN_KEY,
    defaultValue: null,
  })
  const [opened, { toggle }] = useDisclosure()
  const { data: user, isLoading } = useGetMe(userId)

  const isSignedIn = !isLoading && !!user

  if (isLoading) {
    return (
      <Center h='95vh'>
        <Loader />
      </Center>
    )
  }

  if (!isSignedIn || !user) {
    return (
      <Center h='95vh'>
        <Stack>
          <Title>Sign In</Title>
          <Group>
            <Button onClick={() => setUserId('123')}>User</Button>
            <Button onClick={() => setUserId('456')}>Admin</Button>
          </Group>
        </Stack>
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
                underline='never'
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
                  <Avatar style={{ cursor: 'pointer' }} size={30}>
                    {user.first_name?.charAt(0)}
                  </Avatar>
                </Menu.Target>
                <Menu.Dropdown>
                  <Menu.Label>
                    {user.first_name} {user.last_name}
                  </Menu.Label>
                  {/* <Menu.Item>
                      <Link style={{ color: 'inherit', textDecoration: 'none' }} to='/'>Settings</Link>
                    </Menu.Item> */}
                  <Menu.Item disabled={userId === '123'} onClick={() => setUserId('123')} >
                    Log in as User
                  </Menu.Item>
                  <Menu.Item disabled={userId === '456'} onClick={() => setUserId('456')} >
                    Log in as Admin
                  </Menu.Item>
                  <Menu.Item onClick={() => setUserId(null)} c='red'>
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
