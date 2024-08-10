import { Navbar } from '@/components/Navbar'
import { Sidebar } from '@/components/Sidebar'
import { Outlet } from 'react-router-dom'

const App = () => {
  return (
    <div className="flex h-[100dvh] w-[100dvw]">
      <Sidebar />
      {/* right side */}
      <div className="flex-1 flex flex-col overflow-hidden">
        <Navbar />
        <div className="flex-1 overflow-hidden">
          <Outlet />
        </div>
      </div>
    </div>
  )
}

export default App
