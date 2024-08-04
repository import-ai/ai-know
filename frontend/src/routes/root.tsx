import { Sidebar } from '@/components/Sidebar'
import { Outlet } from 'react-router-dom'

const App = () => {
  return (
    <div className="flex h-[100dvh] w-[100dvw]">
      <Sidebar />
      <div className="flex-1">
        <Outlet />
      </div>
    </div>
  )
}

export default App
