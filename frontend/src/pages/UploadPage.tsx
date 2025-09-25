// frontend/src/pages/UploadPage.tsx
import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import FilePicker from '../components/FilePicker'
import { uploadFile } from '../api/upload'
import { useUploadStore } from '../hooks/useUpload'

export default function UploadPage() {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const nav = useNavigate()
  const { setUploadMeta } = useUploadStore()

  const onPick = async (file: File) => {
    setError(null)
    setLoading(true)
    try {
      const res = await uploadFile(file)
      setUploadMeta({
        uploadID: res.uploadID,
        columns: res.columns,
        rows: res.rows
      })
      nav('/chart')
    } catch (e: any) {
      setError(e?.message ?? 'Gagal upload')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="card">
      <h2>1) Upload CSV / XLSX</h2>
      <p>Unggah file data Anda. Server akan membaca header & menyiapkan dataset.</p>
      <FilePicker onPick={onPick} disabled={loading}/>
      {loading && <p>Uploading...</p>}
      {error && <p style={{color:'#ff9393'}}>{error}</p>}
      <p style={{opacity:.8, fontSize:13}}>Format disarankan: CSV dengan header baris pertama, atau XLSX sheet pertama.</p>
    </div>
  )
}
