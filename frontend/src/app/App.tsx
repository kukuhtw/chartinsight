// frontend/src/app/App.tsx
import React from 'react'
import { Outlet, Link, useLocation } from 'react-router-dom'

export default function App() {
  const { pathname } = useLocation()
  return (
    <div className="container">
      <header className="row" style={{justifyContent:'space-between', marginBottom:16}}>
        <Link to="/"><h1>CSV/XLS → Chart + Insight</h1></Link>
        <nav className="row">
          <Link to="/" style={{padding:'8px 12px', borderRadius:8, background: pathname==='/'?'#1a2447':'transparent'}}>Upload</Link>
          <Link to="/chart" style={{padding:'8px 12px', borderRadius:8, background: pathname==='/chart'?'#1a2447':'transparent'}}>Chart</Link>
        </nav>
      </header>
      <Outlet />
      <footer style={{marginTop:24, opacity:.7}}>© {new Date().getFullYear()} csvxlchart</footer>
    </div>
  )
}
