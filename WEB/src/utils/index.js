import dayjs from 'dayjs'
import * as XLSX from 'xlsx'
import { DATE_FORMAT, DATETIME_FORMAT } from '@/config/constants'

/** Format a date/timestamp. Falls back to '-' when the value is empty. */
export function formatDate(value, format = DATE_FORMAT) {
  if (!value) return '-'
  const d = dayjs(value)
  return d.isValid() ? d.format(format) : '-'
}

/** Format a date/timestamp including time. */
export function formatDateTime(value, format = DATETIME_FORMAT) {
  return formatDate(value, format)
}

/** Human-friendly relative time, e.g. "3 hours ago". */
export function formatRelativeTime(value) {
  if (!value) return '-'
  const d = dayjs(value)
  if (!d.isValid()) return '-'
  const diffMinutes = dayjs().diff(d, 'minute')
  if (diffMinutes < 1) return 'just now'
  if (diffMinutes < 60) return `${diffMinutes}m ago`
  const diffHours = dayjs().diff(d, 'hour')
  if (diffHours < 24) return `${diffHours}h ago`
  const diffDays = dayjs().diff(d, 'day')
  if (diffDays < 30) return `${diffDays}d ago`
  return d.format(DATE_FORMAT)
}

/** Format a number as currency, e.g. formatCurrency(1234.5) -> "$1,234.50" */
export function formatCurrency(value, currency = 'USD', locale = 'en-US') {
  const num = Number(value)
  if (Number.isNaN(num)) return '-'
  return new Intl.NumberFormat(locale, { style: 'currency', currency }).format(num)
}

/** Format a plain number with thousand separators. */
export function formatNumber(value, locale = 'en-US') {
  const num = Number(value)
  if (Number.isNaN(num)) return '-'
  return new Intl.NumberFormat(locale).format(num)
}

/** Debounce: delay invoking `fn` until `wait` ms have passed since the last call. */
export function debounce(fn, wait = 300) {
  let timer = null
  return function debounced(...args) {
    clearTimeout(timer)
    timer = setTimeout(() => fn.apply(this, args), wait)
  }
}

/** Throttle: invoke `fn` at most once every `wait` ms. */
export function throttle(fn, wait = 300) {
  let lastTime = 0
  let timer = null
  return function throttled(...args) {
    const now = Date.now()
    const remaining = wait - (now - lastTime)
    if (remaining <= 0) {
      clearTimeout(timer)
      lastTime = now
      fn.apply(this, args)
    } else {
      clearTimeout(timer)
      timer = setTimeout(() => {
        lastTime = Date.now()
        fn.apply(this, args)
      }, remaining)
    }
  }
}

/** Copy text to the clipboard. Returns a promise<boolean> for success. */
export async function copyText(text) {
  try {
    if (navigator.clipboard && window.isSecureContext) {
      await navigator.clipboard.writeText(text)
      return true
    }
    // Fallback for non-secure contexts / older browsers
    const textarea = document.createElement('textarea')
    textarea.value = text
    textarea.style.position = 'fixed'
    textarea.style.opacity = '0'
    document.body.appendChild(textarea)
    textarea.focus()
    textarea.select()
    const success = document.execCommand('copy')
    document.body.removeChild(textarea)
    return success
  } catch {
    return false
  }
}

/** Trigger a browser download for a Blob/URL/string content. */
export function downloadFile(content, filename, mimeType = 'application/octet-stream') {
  const blob = content instanceof Blob ? content : new Blob([content], { type: mimeType })
  const url = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = filename
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  URL.revokeObjectURL(url)
}

/**
 * Export an array of plain objects to an .xlsx file.
 * @param {Array<Object>} rows
 * @param {string} filename e.g. "users.xlsx"
 * @param {string} sheetName
 */
export function exportExcel(rows, filename = 'export.xlsx', sheetName = 'Sheet1') {
  const worksheet = XLSX.utils.json_to_sheet(rows)
  const workbook = XLSX.utils.book_new()
  XLSX.utils.book_append_sheet(workbook, worksheet, sheetName)
  XLSX.writeFile(workbook, filename)
}

/**
 * Open a lightweight image preview using Element Plus's built-in viewer
 * via a temporary offscreen <el-image>. For most cases prefer binding
 * `preview-src-list` directly on an <el-image> in your template — this
 * helper is for previewing an image from JS (e.g. after upload success).
 */
export function imagePreview(url, urlList = null) {
  const list = urlList || [url]
  const event = new CustomEvent('app:image-preview', { detail: { url, list } })
  window.dispatchEvent(event)
}

/** Generate a random-ish id, handy for optimistic UI rows/keys. */
export function uid(prefix = 'id') {
  return `${prefix}_${Math.random().toString(36).slice(2, 10)}`
}
