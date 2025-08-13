import { FC } from 'react'
import { createLazyFileRoute } from '@tanstack/react-router'

import { DefaultErrorDisplay } from '@/components'
import { Stack, Text, Title } from '@mantine/core'

const HomePage: FC = () => {
  return (
      <Stack>
        <Title>Welcome to the RBAC Demo</Title>
        <Text>This is a demo of the RBAC system.</Text>
      </Stack>
  )
}
export const Route = createLazyFileRoute('/')({
  component: HomePage,
  errorComponent: DefaultErrorDisplay,
})
