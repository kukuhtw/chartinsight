// frontend/src/utils/validators.ts
export function isNonEmptyString(s?: string | null): s is string {
  return !!s && typeof s === 'string' && s.trim().length > 0
}
