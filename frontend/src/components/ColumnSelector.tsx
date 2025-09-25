// frontend/src/components/ColumnSelector.tsx
// frontend/src/components/ColumnSelector.tsx
import React from 'react'

export default function ColumnSelector({
  columns, label, value, onChange
}:{
  columns: string[]
  label: string
  value?: string
  onChange: (v: string) => void
}) {
  return (
    <div className="form-group" style={{marginBottom:12}}>
      <label>{label}</label><br/>
      <select
        className="select"
        value={value || ''}
        onChange={(e) => onChange(e.target.value)}
      >
        <option value="" disabled>Choose column…</option>
        {columns.map(c => <option key={c} value={c}>{c}</option>)}
      </select>
    </div>
  )
}
