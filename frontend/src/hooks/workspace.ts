import { useGetApiWorkspace } from '@/api'

export const useEnties = () => {
  const { data } = useGetApiWorkspace()
  return data?.data.entries
}

export const usePrivateEnties = () => {
  const data = useEnties()
  return data?.private
}
export const useTeamEnties = () => {
  const data = useEnties()
  return data?.team
}
