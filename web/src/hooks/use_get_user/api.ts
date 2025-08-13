import { client } from '@/api/client'

export const getMe = async () => {
  const { data, error } = await client.GET('/me')
  if (error || !data) throw error

  return data
}

export const getMeQueryOptions = (userId: string | null) => ({
  queryKey: ['me', userId],
  queryFn: () => getMe(),
  enabled: !!userId,
})