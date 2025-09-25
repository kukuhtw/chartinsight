// frontend/src/components/InsightBox.tsx
import React from 'react'
import { Stats } from '../types/chart'

export default function InsightBox({ text, stats }:{ text: string, stats: Stats }) {
  return (
    <div className="card" style={{marginTop:8}}>
      <div style={{marginBottom:8, fontWeight:600}}>Model Insight</div>
      <pre>{text}</pre>
      <div style={{opacity:.8, fontSize:13, marginTop:8}}>
        <div>n: {stats.n ?? 0}</div>
        <div>mean: {stats.mean?.toFixed?.(4) ?? 0}</div>
        <div>min: {stats.min?.toFixed?.(4) ?? 0}</div>
        <div>max: {stats.max?.toFixed?.(4) ?? 0}</div>
        <div>std: {stats.std?.toFixed?.(4) ?? 0}</div>
      </div>
    </div>
  )
}
