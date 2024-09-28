import { useGetApiWorkspace } from "@/api"

export const TestApi = () => {
  const {data} = useGetApiWorkspace()
  return <div>
  
  {JSON.stringify(data)}</div>
}
