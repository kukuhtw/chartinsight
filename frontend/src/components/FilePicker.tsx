// frontend/src/components/FilePicker.tsx
import React, { useRef } from 'react'

export default function FilePicker({ onPick, disabled }:{
  onPick: (file: File) => void
  disabled?: boolean
}) {
  const ref = useRef<HTMLInputElement | null>(null)
  const choose = () => ref.current?.click()
  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const f = e.target.files?.[0]
    if (f) onPick(f)
    e.target.value = ''
  }
  return (
    <div className="row">
      <button className="btn" onClick={choose} disabled={disabled}>Pilih File</button>
      <input
        ref={ref}
        type="file"
        accept=".csv,.xls,.xlsx"
        onChange={onChange}
        style={{ display: 'none' }}
      />
      <span style={{opacity:.8}}>Dukungan: .csv, .xls, .xlsx</span>
    </div>
  )
}
