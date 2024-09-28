import { useGetApiSidebarEntriesEntryId, useGetApiWorkspace } from '@/api'

export const TestApi = () => {
  const { data } = useGetApiWorkspace()
  useGetApiSidebarEntriesEntryId()
  return <div>{JSON.stringify(data?.data)}</div>
}
