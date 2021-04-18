import axios from 'axios'
import useSWR, { mutate } from 'swr'
import { env } from 'config'

// TODO: Add auth interceptor
const client = axios.create({ baseURL: env.API_HOST })

const fetcher = (options = {}) => (endpoint = '') =>
  client.get(endpoint, options)
const fetch = (endpoint = '', options = {}) =>
  useSWR(endpoint, fetcher(options))

const creator = (endpoint = '', body = {}, options = {}) =>
  client.post(endpoint, body, options)
const create = (endpoint = '', body = {}, options = {}) =>
  mutate(endpoint, creator(endpoint, body, options))

const updator = (endpoint = '', body = {}, options = {}) =>
  client.put(endpoint, body, options)
const update = (endpoint = '', body = {}, options = {}) =>
  mutate(endpoint, updator(endpoint, body, options))

const patcher = (endpoint = '', body = {}, options = {}) =>
  client.patch(endpoint, body, options)
const patch = (endpoint = '', body = {}, options = {}) =>
  mutate(endpoint, patcher(endpoint, body, options))

const deletor = (endpoint = '', options = {}) =>
  client.delete(endpoint, options)
const remove = (endpoint = '', options = {}) =>
  mutate(endpoint, deletor(endpoint, options), false)
const deactivate = (endpoint = '', options = {}) =>
  mutate(endpoint, deletor(endpoint, options))

export default {
  client,
  fetcher,
  fetch,
  creator,
  create,
  updator,
  update,
  patcher,
  patch,
  deletor,
  remove,
  deactivate,
}
