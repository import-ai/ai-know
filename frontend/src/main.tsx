import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Root from '@/routes/root'
import { Article } from '@/components/Article'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    children: [
      {
        path: 'article/:id',
        element: <Article />,
      },
      {
        path: '/',
        element: (
          <div>
            <h1>Welcome</h1>
            <p>Select an article from the sidebar.</p>
          </div>
        ),
      },
    ],
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
