// frontend/src/components/ChartView.tsx
import React, { useEffect, useRef } from 'react'
import * as echarts from 'echarts'
import { Series } from '../types/chart'

export default function ChartView({
  x, y, xLabels, series, xName, yName
}:{
  x?: string[]; y?: number[];
  xLabels?: string[]; series?: Series[];
  xName?: string; yName?: string;
}) {
  const ref = useRef<HTMLDivElement | null>(null)

  useEffect(() => {
    if (!ref.current) return
    const chart = echarts.init(ref.current)

    let opt: echarts.EChartsOption

    if (xLabels && series && series.length > 0) {
      opt = {
        tooltip: { trigger: 'axis' },
        legend: { top: 0 },
        xAxis: { type: 'category', data: xLabels, name: xName },
        yAxis: { type: 'value', name: yName },
        series: series.map(s => ({
          name: s.name,
          type: 'line',
          smooth: true,
          data: s.data
        }))
      }
    } else {
      opt = {
        tooltip: {},
        xAxis: { type: 'category', data: x || [], name: xName },
        yAxis: { type: 'value', name: yName },
        series: [{
          type: 'line',
          smooth: true,
          data: y || []
        }]
      }
    }

    chart.setOption(opt)
    const onResize = () => chart.resize()
    window.addEventListener('resize', onResize)
    return () => { window.removeEventListener('resize', onResize); chart.dispose() }
  }, [x, y, xLabels, series, xName, yName])

  return (
    <div
      ref={ref}
      style={{
        width: '100%',
        height: 420,
        borderRadius: 12,
        overflow: 'hidden',
        background:'#0f1530'
      }}
    />
  )
}
