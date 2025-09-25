// frontend/src/pages/ChartPage.tsx
import React, { useMemo, useState } from 'react'
import ColumnSelector from '../components/ColumnSelector'
import ChartView from '../components/ChartView'
import InsightBox from '../components/InsightBox'
import { useUploadStore } from '../hooks/useUpload'
import { useChart } from '../hooks/useChart'

export default function ChartPage() {
  const { uploadID, columns } = useUploadStore()
  const { runChart, data, loading } = useChart()

  const [colX, setColX] = useState<string>()
  const [colY, setColY] = useState<string>()
  const [groupBy, setGroupBy] = useState<string>()
  const [agg, setAgg] = useState<'avg'|'sum'|'min'|'max'>('avg')

  const valid = useMemo(() => Boolean(uploadID && colX && colY), [uploadID, colX, colY])

  const handleChart = async () => {
    if (!valid || !uploadID) return
    await runChart({
      uploadID,
      colX: colX!,
      colY: colY!,
      groupBy: groupBy || undefined,
      agg,
    })
  }

  return (
    <div className="card">
      <h2>2) Select Columns & Render</h2>

      {!uploadID && (
        <p style={{color:'#ff9393'}}>
          No dataset yet. Please upload a file on the <b>Upload</b> page first.
        </p>
      )}

      {uploadID && (
        <>
          <ColumnSelector
            columns={columns}
            label="X Axis (Category/Label)"
            value={colX}
            onChange={setColX}
          />
          <ColumnSelector
            columns={columns}
            label="Y Axis (Numeric)"
            value={colY}
            onChange={setColY}
          />
          <ColumnSelector
            columns={columns}
            label="Group By (optional)"
            value={groupBy}
            onChange={setGroupBy}
          />

          <div className="form-group" style={{marginTop:8}}>
            <label>Aggregation</label><br/>
            <select className="select" value={agg} onChange={e => setAgg(e.target.value as any)}>
              <option value="avg">Average</option>
              <option value="sum">Sum</option>
              <option value="min">Min</option>
              <option value="max">Max</option>
            </select>
          </div>

          <div className="row" style={{marginTop:12}}>
            <button className="btn" onClick={handleChart} disabled={!valid || loading}>
              {loading ? 'Processing…' : 'Render Chart'}
            </button>
          </div>
        </>
      )}

      {data && (
        <>
          <h2 style={{marginTop:16}}>Chart</h2>
          <ChartView
            x={data.x}
            y={data.y}
            xLabels={data.xLabels}
            series={data.series}
            xName={colX}
            yName={colY}
          />
          <h2>Insight</h2>
          <InsightBox text={data.insight} stats={data.stats}/>
        </>
      )}
    </div>
  )
}
