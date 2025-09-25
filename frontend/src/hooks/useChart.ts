// frontend/src/hooks/useChart.ts
import { useState } from 'react'
import { runChartAPI } from '../api/chart'
import { ChartResponse } from '../types/chart'

type Params = {
  colX: string
  colY: string
  groupBy?: string
}

export function useChart() {
  const [data, setData] = useState<ChartResponse | null>(null)
  const [loading, setLoading] = useState(false)

  async function runChart(params: Params) {
    setLoading(true)
    try {
      const result = await runChartAPI(params)
      setData(result)
    } finally {
      setLoading(false)
    }
  }

  return { runChart, data, loading }
}
