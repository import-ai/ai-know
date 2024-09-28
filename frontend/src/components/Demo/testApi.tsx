import {
  useGetApiSidebarEntriesEntryId,
  useGetApiSidebarEntriesEntryIdSubEntries,
  useGetApiWorkspace,
} from '@/api'
import { usePrivateEnties } from '@/hooks/workspace'

export const TestApi = () => {
  const privateId = usePrivateEnties()
  const { data: sidebarData } = useGetApiSidebarEntriesEntryId(privateId ?? '')
  const { data: subdata } = useGetApiSidebarEntriesEntryIdSubEntries(
    privateId ?? '',
  )
  return (
    <div>
      {JSON.stringify(sidebarData?.data)}
      <div>{JSON.stringify(subdata?.data)}</div>
    </div>
  )
}
