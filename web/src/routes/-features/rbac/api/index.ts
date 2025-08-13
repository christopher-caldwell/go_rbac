import { client } from '@/api/client'

export const getUnprotected = async () => {
  const { data, error } = await client.GET('/unprotected')
  if (error || !data) throw error

  return data
}

export const unprotectedQueryOptions = {
  queryKey: ['unprotected'],
  queryFn: () => getUnprotected(),
}