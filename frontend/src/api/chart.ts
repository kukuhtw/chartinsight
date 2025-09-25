// frontend/src/api/chart.ts
// frontend/src/api/chart.ts
import { http } from './client'
import { ChartResponse } from '../types/chart'

export type Params = {
  uploadID: string
  colX: string
  colY: string
  groupBy?: string
  agg?: 'avg'|'sum'|'min'|'max'
}

export async function runChartAPI(params: Params): Promise<ChartResponse> {
  return http<ChartResponse>('/chart', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(params),
  })
}
