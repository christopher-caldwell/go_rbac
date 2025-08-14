import { FC } from 'react'
import { createLazyFileRoute } from '@tanstack/react-router'

import { DefaultErrorDisplay } from '@/components'
import { PermissionsCheck } from './-features/rbac'

const HomePage: FC = () => {
  return <PermissionsCheck />
}
export const Route = createLazyFileRoute('/')({
  component: HomePage,
  errorComponent: DefaultErrorDisplay,
})
