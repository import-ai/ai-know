import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Root from '@/routes/root'
import { Article } from '@/components/Article'
import '@blocksuite/presets/themes/affine.css'
import { EditorProvider } from '@/providers/EditorProvider'
import { DevTools } from 'jotai-devtools'
import 'jotai-devtools/styles.css'
import { TestApi } from '@/components/Demo/testApi'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    children: [
      {
        path: 'demo',
        element:<TestApi/>
      },
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
    <EditorProvider>
      <RouterProvider router={router} />
      <DevTools />
    </EditorProvider>
  </React.StrictMode>,
)
