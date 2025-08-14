import { client } from '@/api/client'
import { components } from '@/api/types'
import { UseMutationOptions } from '@tanstack/react-query'

export const getUnprotected = async () => {
  const { data, error } = await client.GET('/unprotected')
  if (error || !data) throw error

  return JSON.stringify(data, null, 2)
}

export const getUnprotectedQueryOptions = () => {
  return {
    queryKey: ['unprotected'],
    queryFn: () => getUnprotected(),
  }
}

export const getProtected = async () => {
  const { data, error } = await client.GET('/protected')
  if (error || !data) throw error

  return JSON.stringify(data, null, 2)
}

export const getProtectedQueryOptions = (userId: string | null) => {
  return {
    queryKey: ['protected', userId],
    queryFn: () => getProtected(),
    enabled: !!userId,
  }
}

export const getRbac = async () => {
  const { data, error } = await client.GET('/rbac')
  if (error || !data) throw error

  return JSON.stringify(data, null, 2)
}

export const getRbacQueryOptions = (userId: string | null) => {
  return {
    queryKey: ['rbac', userId],
    queryFn: () => getRbac(),
    enabled: !!userId,
  }
}

type RbacResource = components['schemas']['RbacResource']
export const postRbac = async (newRbacResource: RbacResource) => {
  const { data, error } = await client.POST('/rbac', { body: newRbacResource })
  if (error || !data) throw error

  return JSON.stringify(data, null, 2)
}

export const postRbacMutationOptions: UseMutationOptions<string, Error, RbacResource> = {
  mutationFn: (newRbacResource) => {
    return postRbac(newRbacResource)
  },
}


export const getPermissions = async () => {
  const { data, error } = await client.GET('/permissions')
  if (error || !data) throw error

  return data
}

export const getPermissionsQueryOptions = (userId: string | null) => {
  return {
    queryKey: ['permissions', userId],
    queryFn: () => getPermissions(),
    enabled: !!userId,
  }
}
