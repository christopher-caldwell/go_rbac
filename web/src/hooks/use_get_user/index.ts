import { useQuery } from '@tanstack/react-query'
import { getMeQueryOptions } from './api'

export const useGetMe = (userId: string | null) => {
  return useQuery(getMeQueryOptions(userId))
}
