// frontend/src/hooks/useUpload.ts
// frontend/src/hooks/useUpload.ts
import { useSyncExternalStore } from 'react'
import { UploadMeta } from '../types/dataset'

type State = UploadMeta

let state: State = { uploadID: '', columns: [], rows: 0 }
const listeners = new Set<() => void>()

function emit() { listeners.forEach(l => l()) }
function subscribe(cb: () => void) { listeners.add(cb); return () => listeners.delete(cb) }
function getSnapshot() { return state }

export function useUploadStore() {   // ✅ consistent export
  const snap = useSyncExternalStore(subscribe, getSnapshot)
  return {
    uploadID: snap.uploadID,
    columns: snap.columns,
    rows: snap.rows,
    setUploadMeta(meta: State) { state = meta; emit() }
  }
}
