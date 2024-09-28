/**
 * Generated by orval v7.1.1 🍺
 * Do not edit manually.
 * AIKnow API
 * ## Sidebar 多级列表

列表的每一项是一个`entry`，主要字段是`id`，`type`和`title`。

- `id`: 全局唯一标识
- `type`: 合法取值为`note`，`group`或`link`
- `title`: 标题

一个`entry`下可以嵌套子`entry`，形成树形结构。

### 查询流程

1. 调用`Get Workspace`拿到 Private 和 Team Space 分别对应的最外层`entry id`。
2. 调用`Get Sub-Entries`，传参`entry id`，拿到该`entry`直接嵌套的子`entry`列表。
3. 如果子`entry`继续嵌套子`entry`（`has_sub_entries`为`true`），递归调用`Get Sub-Entries`。

### 拖拽/移动流程

调用`Update Entry`，更新`entry`的`parent`和`position`。
 * OpenAPI spec version: 1.0
 */
import * as axios from 'axios'
import type { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios'
import useSwr from 'swr'
import type { Arguments, Key, SWRConfiguration } from 'swr'
import useSWRMutation from 'swr/mutation'
import type { SWRMutationConfiguration } from 'swr/mutation'
import type {
  HandlersCreateEntryReq,
  HandlersCreateEntryResp,
  HandlersDuplicateEntryReq,
  HandlersDuplicateEntryResp,
  HandlersGetEntryResp,
  HandlersGetSubEntriesResp,
  HandlersPutEntryReq,
  HandlersPutEntryResp,
} from '.././model'

/**
 * Create an entry with the specified properties.
|      Field      | Required |      Description      |
| :-------------: | :------: | :-------------------: |
|      title      |   Yes    |  Title of new entry   |
|      type       |   Yes    |   Type of new entry   |
|     parent      |   Yes    |    Parent entry ID    |
| posistion_after |    No    | Position of new entry |

If `position_after` is empty, the new entry will be the first in parent's sub-entries. Otherwise, it's positioned after the specified sub-entry.
 * @summary Create Entry
 */
export const postApiSidebarEntries = (
  handlersCreateEntryReq: HandlersCreateEntryReq,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<HandlersCreateEntryResp>> => {
  return axios.default.post(
    `/api/sidebar/entries`,
    handlersCreateEntryReq,
    options,
  )
}

export const getPostApiSidebarEntriesMutationFetcher = (
  options?: AxiosRequestConfig,
) => {
  return (
    _: Key,
    { arg }: { arg: HandlersCreateEntryReq },
  ): Promise<AxiosResponse<HandlersCreateEntryResp>> => {
    return postApiSidebarEntries(arg, options)
  }
}
export const getPostApiSidebarEntriesMutationKey = () =>
  [`/api/sidebar/entries`] as const

export type PostApiSidebarEntriesMutationResult = NonNullable<
  Awaited<ReturnType<typeof postApiSidebarEntries>>
>
export type PostApiSidebarEntriesMutationError = AxiosError<unknown>

/**
 * @summary Create Entry
 */
export const usePostApiSidebarEntries = <
  TError = AxiosError<unknown>,
>(options?: {
  swr?: SWRMutationConfiguration<
    Awaited<ReturnType<typeof postApiSidebarEntries>>,
    TError,
    Key,
    HandlersCreateEntryReq,
    Awaited<ReturnType<typeof postApiSidebarEntries>>
  > & { swrKey?: string }
  axios?: AxiosRequestConfig
}) => {
  const { swr: swrOptions, axios: axiosOptions } = options ?? {}

  const swrKey = swrOptions?.swrKey ?? getPostApiSidebarEntriesMutationKey()
  const swrFn = getPostApiSidebarEntriesMutationFetcher(axiosOptions)

  const query = useSWRMutation(swrKey, swrFn, swrOptions)

  return {
    swrKey,
    ...query,
  }
}
/**
 * Get properties of an entry.
 * @summary Get Entry
 */
export const getApiSidebarEntriesEntryId = (
  entryId: string,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<HandlersGetEntryResp>> => {
  return axios.default.get(`/api/sidebar/entries/${entryId}`, options)
}

export const getGetApiSidebarEntriesEntryIdKey = (entryId: string) =>
  [`/api/sidebar/entries/${entryId}`] as const

export type GetApiSidebarEntriesEntryIdQueryResult = NonNullable<
  Awaited<ReturnType<typeof getApiSidebarEntriesEntryId>>
>
export type GetApiSidebarEntriesEntryIdQueryError = AxiosError<unknown>

/**
 * @summary Get Entry
 */
export const useGetApiSidebarEntriesEntryId = <TError = AxiosError<unknown>>(
  entryId: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof getApiSidebarEntriesEntryId>>,
      TError
    > & { swrKey?: Key; enabled?: boolean }
    axios?: AxiosRequestConfig
  },
) => {
  const { swr: swrOptions, axios: axiosOptions } = options ?? {}

  const isEnabled = swrOptions?.enabled !== false && !!entryId
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getGetApiSidebarEntriesEntryIdKey(entryId) : null))
  const swrFn = () => getApiSidebarEntriesEntryId(entryId, axiosOptions)

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  )

  return {
    swrKey,
    ...query,
  }
}
/**
 * Update properties of an entry.
|      Field      | Required |      Description      |
| :-------------: | :------: | :-------------------: |
|      title      |    No    |  Title of the entry   |
|     parent      |    No    |    Parent entry ID    |
| posistion_after |    No    | Position of the entry |

> - If `title` is non-empty, update title of the entry.
> - If `parent` is non-empty, move the entry to the specified parent entry.
>     - If `position_after` is empty, the new entry will be the first in parent's sub-entries. Otherwise, it's positioned after the specified sub-entry.
 * @summary Update Entry
 */
export const putApiSidebarEntriesEntryId = (
  entryId: string,
  handlersPutEntryReq: HandlersPutEntryReq,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<HandlersPutEntryResp>> => {
  return axios.default.put(
    `/api/sidebar/entries/${entryId}`,
    handlersPutEntryReq,
    options,
  )
}

export const getPutApiSidebarEntriesEntryIdMutationFetcher = (
  entryId: string,
  options?: AxiosRequestConfig,
) => {
  return (
    _: Key,
    { arg }: { arg: HandlersPutEntryReq },
  ): Promise<AxiosResponse<HandlersPutEntryResp>> => {
    return putApiSidebarEntriesEntryId(entryId, arg, options)
  }
}
export const getPutApiSidebarEntriesEntryIdMutationKey = (entryId: string) =>
  [`/api/sidebar/entries/${entryId}`] as const

export type PutApiSidebarEntriesEntryIdMutationResult = NonNullable<
  Awaited<ReturnType<typeof putApiSidebarEntriesEntryId>>
>
export type PutApiSidebarEntriesEntryIdMutationError = AxiosError<unknown>

/**
 * @summary Update Entry
 */
export const usePutApiSidebarEntriesEntryId = <TError = AxiosError<unknown>>(
  entryId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof putApiSidebarEntriesEntryId>>,
      TError,
      Key,
      HandlersPutEntryReq,
      Awaited<ReturnType<typeof putApiSidebarEntriesEntryId>>
    > & { swrKey?: string }
    axios?: AxiosRequestConfig
  },
) => {
  const { swr: swrOptions, axios: axiosOptions } = options ?? {}

  const swrKey =
    swrOptions?.swrKey ?? getPutApiSidebarEntriesEntryIdMutationKey(entryId)
  const swrFn = getPutApiSidebarEntriesEntryIdMutationFetcher(
    entryId,
    axiosOptions,
  )

  const query = useSWRMutation(swrKey, swrFn, swrOptions)

  return {
    swrKey,
    ...query,
  }
}
/**
 * Delete an entry and all its sub-entries.
 * @summary Delete Entry
 */
export const deleteApiSidebarEntriesEntryId = (
  entryId: string,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<unknown>> => {
  return axios.default.delete(`/api/sidebar/entries/${entryId}`, options)
}

export const getDeleteApiSidebarEntriesEntryIdMutationFetcher = (
  entryId: string,
  options?: AxiosRequestConfig,
) => {
  return (_: Key, __: { arg: Arguments }): Promise<AxiosResponse<unknown>> => {
    return deleteApiSidebarEntriesEntryId(entryId, options)
  }
}
export const getDeleteApiSidebarEntriesEntryIdMutationKey = (entryId: string) =>
  [`/api/sidebar/entries/${entryId}`] as const

export type DeleteApiSidebarEntriesEntryIdMutationResult = NonNullable<
  Awaited<ReturnType<typeof deleteApiSidebarEntriesEntryId>>
>
export type DeleteApiSidebarEntriesEntryIdMutationError = AxiosError<unknown>

/**
 * @summary Delete Entry
 */
export const useDeleteApiSidebarEntriesEntryId = <TError = AxiosError<unknown>>(
  entryId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof deleteApiSidebarEntriesEntryId>>,
      TError,
      Key,
      Arguments,
      Awaited<ReturnType<typeof deleteApiSidebarEntriesEntryId>>
    > & { swrKey?: string }
    axios?: AxiosRequestConfig
  },
) => {
  const { swr: swrOptions, axios: axiosOptions } = options ?? {}

  const swrKey =
    swrOptions?.swrKey ?? getDeleteApiSidebarEntriesEntryIdMutationKey(entryId)
  const swrFn = getDeleteApiSidebarEntriesEntryIdMutationFetcher(
    entryId,
    axiosOptions,
  )

  const query = useSWRMutation(swrKey, swrFn, swrOptions)

  return {
    swrKey,
    ...query,
  }
}
/**
 * Duplicate an entry.
|      Field      | Required |      Description      |
| :-------------: | :------: | :-------------------: |
|      title      |    No    |  Title of new entry   |
|     parent      |   Yes    |    Parent entry ID    |
| posistion_after |    No    | Position of new entry |

If `title` is empty, it will default to the old entry’s title.
If `position_after` is empty, the new entry will be the first in parent's sub-entries. Otherwise, it's positioned after the specified sub-entry.
 * @summary Duplicate Entry
 */
export const postApiSidebarEntriesEntryIdDuplicate = (
  entryId: string,
  handlersDuplicateEntryReq: HandlersDuplicateEntryReq,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<HandlersDuplicateEntryResp>> => {
  return axios.default.post(
    `/api/sidebar/entries/${entryId}/duplicate`,
    handlersDuplicateEntryReq,
    options,
  )
}

export const getPostApiSidebarEntriesEntryIdDuplicateMutationFetcher = (
  entryId: string,
  options?: AxiosRequestConfig,
) => {
  return (
    _: Key,
    { arg }: { arg: HandlersDuplicateEntryReq },
  ): Promise<AxiosResponse<HandlersDuplicateEntryResp>> => {
    return postApiSidebarEntriesEntryIdDuplicate(entryId, arg, options)
  }
}
export const getPostApiSidebarEntriesEntryIdDuplicateMutationKey = (
  entryId: string,
) => [`/api/sidebar/entries/${entryId}/duplicate`] as const

export type PostApiSidebarEntriesEntryIdDuplicateMutationResult = NonNullable<
  Awaited<ReturnType<typeof postApiSidebarEntriesEntryIdDuplicate>>
>
export type PostApiSidebarEntriesEntryIdDuplicateMutationError =
  AxiosError<unknown>

/**
 * @summary Duplicate Entry
 */
export const usePostApiSidebarEntriesEntryIdDuplicate = <
  TError = AxiosError<unknown>,
>(
  entryId: string,
  options?: {
    swr?: SWRMutationConfiguration<
      Awaited<ReturnType<typeof postApiSidebarEntriesEntryIdDuplicate>>,
      TError,
      Key,
      HandlersDuplicateEntryReq,
      Awaited<ReturnType<typeof postApiSidebarEntriesEntryIdDuplicate>>
    > & { swrKey?: string }
    axios?: AxiosRequestConfig
  },
) => {
  const { swr: swrOptions, axios: axiosOptions } = options ?? {}

  const swrKey =
    swrOptions?.swrKey ??
    getPostApiSidebarEntriesEntryIdDuplicateMutationKey(entryId)
  const swrFn = getPostApiSidebarEntriesEntryIdDuplicateMutationFetcher(
    entryId,
    axiosOptions,
  )

  const query = useSWRMutation(swrKey, swrFn, swrOptions)

  return {
    swrKey,
    ...query,
  }
}
/**
 * Get sub-entries of an entry.
 * @summary Get Sub-Entries
 */
export const getApiSidebarEntriesEntryIdSubEntries = (
  entryId: string,
  options?: AxiosRequestConfig,
): Promise<AxiosResponse<HandlersGetSubEntriesResp>> => {
  return axios.default.get(
    `/api/sidebar/entries/${entryId}/sub_entries`,
    options,
  )
}

export const getGetApiSidebarEntriesEntryIdSubEntriesKey = (entryId: string) =>
  [`/api/sidebar/entries/${entryId}/sub_entries`] as const

export type GetApiSidebarEntriesEntryIdSubEntriesQueryResult = NonNullable<
  Awaited<ReturnType<typeof getApiSidebarEntriesEntryIdSubEntries>>
>
export type GetApiSidebarEntriesEntryIdSubEntriesQueryError =
  AxiosError<unknown>

/**
 * @summary Get Sub-Entries
 */
export const useGetApiSidebarEntriesEntryIdSubEntries = <
  TError = AxiosError<unknown>,
>(
  entryId: string,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof getApiSidebarEntriesEntryIdSubEntries>>,
      TError
    > & { swrKey?: Key; enabled?: boolean }
    axios?: AxiosRequestConfig
  },
) => {
  const { swr: swrOptions, axios: axiosOptions } = options ?? {}

  const isEnabled = swrOptions?.enabled !== false && !!entryId
  const swrKey =
    swrOptions?.swrKey ??
    (() =>
      isEnabled ? getGetApiSidebarEntriesEntryIdSubEntriesKey(entryId) : null)
  const swrFn = () =>
    getApiSidebarEntriesEntryIdSubEntries(entryId, axiosOptions)

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  )

  return {
    swrKey,
    ...query,
  }
}
