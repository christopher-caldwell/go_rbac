import { FC, useEffect, useState } from 'react'
import { Button, Card, Group, Stack, Title, Code } from '@mantine/core'
import { useLocalStorage } from '@mantine/hooks'
import { LS_TOKEN_KEY } from '@/api/client'
import { useMutation, useQuery } from '@tanstack/react-query'
import {
  getPermissionsQueryOptions,
  getProtectedQueryOptions,
  getRbacQueryOptions,
  getUnprotectedQueryOptions,
  postRbacMutationOptions,
} from './api'
import { authorizer } from './api/enforcer'

export const PermissionsCheck: FC = () => {
  const [userId] = useLocalStorage({ key: LS_TOKEN_KEY })
  const [resultDisplay, setResultDisplay] = useState<string | undefined>(undefined)
  const [canPostRbac, setCanPostRbac] = useState<boolean>(false)

  const { data: permissionsData } = useQuery(getPermissionsQueryOptions(userId))
  const { data: unprotectedData } = useQuery(getUnprotectedQueryOptions())
  const { data: protectedData } = useQuery(getProtectedQueryOptions(userId))
  const { data: rbacData } = useQuery(getRbacQueryOptions(userId))
  const { mutateAsync: postRbac } = useMutation(postRbacMutationOptions)

  useEffect(() => {
    if (!permissionsData) return
    authorizer.setPermission({
      read: ['data1', 'data2'],
      write: ['data1'],
    })
    authorizer.setUser(userId)

    console.log('perm', authorizer.getPermission())
    authorizer.can('write', 'data1').then((can) => {
      console.log('can', can)
      setCanPostRbac(can)
    })
  }, [permissionsData, userId])

  const handleRequestChange = async (path: string) => {
    if (path === 'postRbac') {
      try {
        const res = await postRbac({
          message: 'test',
          something_secret: 'test',
          something_very_secret: 'test',
        })
        setResultDisplay(res)
      } catch (e) {
        console.error('error in catch', e)
        setResultDisplay(JSON.stringify(e, null, 2))
      }
    } else {
      switch (path) {
        case 'unprotected':
          setResultDisplay(unprotectedData)
          break
        case 'protected':
          setResultDisplay(protectedData)
          break
        case 'getRbac':
          setResultDisplay(rbacData)
          break
      }
    }
  }

  if (!userId) return null
  return (
    <Stack>
      <Title>Permissions Check</Title>
      <Group mb='xl'>
        <Button onClick={() => handleRequestChange('unprotected')}>Unprotected</Button>
        <Button onClick={() => handleRequestChange('protected')}>Protected</Button>
        <Button onClick={() => handleRequestChange('getRbac')}>Get Rbac</Button>
        <Button color={canPostRbac ? 'green' : 'red'} onClick={() => handleRequestChange('postRbac')}>
          Post Rbac
        </Button>
      </Group>
      <Card>
        <Title mb='xl' order={3}>
          Result
        </Title>
        <Code>{resultDisplay ?? 'No result'}</Code>
      </Card>
    </Stack>
  )
}
