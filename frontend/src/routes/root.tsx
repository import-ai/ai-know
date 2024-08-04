import { Sidebar } from '@/components/Sidebar'
import { Outlet } from 'react-router-dom'

const App = () => {
  return (
    <div style={{ display: 'flex' }}>
      <Sidebar />
      <div style={{ flex: 1, padding: '10px' }}>
        <Outlet />
      </div>
    </div>
  )
}

export default App
