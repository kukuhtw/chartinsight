// frontend/src/api/client.ts
// const baseURL = import.meta.env.VITE_API_BASE_URL?.replace(/\/+$/,'') || ''
// frontend/src/api/client.ts
const baseURL = '/api'; // selalu proxy ke /api
// ... sisanya tetap

export async function http<T>(path: string, opts?: RequestInit): Promise<T> {
  const res = await fetch(baseURL + path, {
    ...opts,
    headers: {
      'Accept': 'application/json',
      ...(opts?.headers || {}),
    }
  })
  const text = await res.text()
  let json: any
  try { json = text ? JSON.parse(text) : {} } catch { throw new Error(text || 'Invalid JSON') }
  if (!res.ok) {
    throw new Error(json?.error || `HTTP ${res.status}`)
  }
  return json as T
}
