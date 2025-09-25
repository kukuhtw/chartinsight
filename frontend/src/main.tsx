// frontend/src/main.tsx
import React from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import App from './app/App'
import UploadPage from './pages/UploadPage'
import ChartPage from './pages/ChartPage'
import './index.css'

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      { index: true, element: <UploadPage /> },
      { path: 'chart', element: <ChartPage /> }
    ]
  }
])

createRoot(document.getElementById('root')!).render(<RouterProvider router={router} />)
