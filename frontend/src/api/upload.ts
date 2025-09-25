// frontend/src/api/upload.ts
import { http } from './client'
import { Dataset } from '../types/dataset'

export async function uploadFile(file: File): Promise<Dataset> {
  const form = new FormData()
  form.append('file', file)
  return http<Dataset>('/upload', { method: 'POST', body: form })
}
