import createClient, { Middleware } from 'openapi-fetch'
import type { paths } from './types'
import { readLocalStorageValue } from '@mantine/hooks'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL as string
if (typeof API_BASE_URL !== 'string') {
  throw new Error('VITE_API_BASE_URL is not defined')
}

export const LS_TOKEN_KEY = 'token'
const tokenMiddleware: Middleware = {
  async onRequest({ request }) {
    const token = readLocalStorageValue({ key: LS_TOKEN_KEY })
    console.log('token', token)
    if (!token) return request

    request.headers.set('Authorization', `Bearer ${token}`)
    return request
  },
}

export const client = createClient<paths>({ baseUrl: API_BASE_URL })
client.use(tokenMiddleware)
