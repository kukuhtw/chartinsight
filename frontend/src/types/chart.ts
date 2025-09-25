// frontend/src/types/chart.ts
export type ChartRequest = {
  uploadID: string
  colX: string
  colY: string
  groupBy?: string
  agg?: 'avg'|'sum'|'min'|'max'
}

export type Stats = {
  n: number
  min: number
  max: number
  mean: number
  std: number
}

export type Series = {
  name: string
  data: number[]
}

export type ChartResponse = {
  // legacy
  x?: string[]
  y?: number[]

  // grouped
  xLabels?: string[]
  series?: Series[]

  stats: Stats
  insight: string
  echartsOption?: Record<string, any>
}
