// frontend/src/utils/transform.ts
export function truncateArray<T>(arr: T[], max = 5000): T[] {
  if (arr.length <= max) return arr
  return arr.slice(0, max)
}
