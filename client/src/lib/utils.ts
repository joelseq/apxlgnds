import type { Event } from '$lib/types'
import { formatRelative } from 'date-fns'

export function splitEvents(events: Event[]): { past: Event[]; upcoming: Event[] } {
  const now = new Date().getTime()

  const past = events
    .filter((event) => getTimeFromString(event.endDate) < now)
    .sort((a, b) => getTimeFromString(b.startDate) - getTimeFromString(a.startDate))
  const upcoming = events.filter((event) => getTimeFromString(event.endDate) >= now)

  return { past, upcoming }
}

function getTimeFromString(date: string): number {
  return new Date(date).getTime()
}

export function timeUntilEvent(event: Event): string {
  const now = new Date().getTime()
  const startDate = getTimeFromString(event.startDate)
  const endDate = getTimeFromString(event.endDate)

  // Check if the event is happening now
  if (startDate <= now && endDate >= now) {
    return 'Now'
  }

  // const diffDate = startDate - now > 0 ? startDate : endDate

  return formatRelative(startDate, new Date())
}

// Helper functions for formatting from ChatGPT

// Helper function to format time
function formatTime(date: Date) {
  let hours = date.getHours()
  let minutes: string | number = date.getMinutes()
  const period = hours >= 12 ? 'pm' : 'am'
  hours = hours % 12 || 12 // Convert to 12-hour format
  minutes = minutes < 10 ? '0' + minutes : minutes
  return `${hours}:${minutes}${period}`
}

// Days of the week
const daysOfWeek = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']

// Months of the year
const monthsOfYear = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December',
]

// Helper function to format date
function formatDate(date: Date) {
  const dayOfWeek = daysOfWeek[date.getDay()]
  const month = monthsOfYear[date.getMonth()]
  const day = date.getDate()
  return `${dayOfWeek}, ${month} ${day}`
}

export function formatDuration(startDateStr: string, endDateStr: string): string {
  const startDate = new Date(startDateStr)
  const endDate = new Date(endDateStr)
  // Extracting the start date parts
  const startDayFormatted = formatDate(startDate)

  // Formatting the start and end times
  const startTime = formatTime(startDate)
  const endTime = formatTime(endDate)

  // Check if the end date is on a different day
  const endDayFormatted = formatDate(endDate)
  if (startDayFormatted !== endDayFormatted) {
    return `${startDayFormatted}, ${startTime} – ${endDayFormatted}, ${endTime}`
  }

  // If the end date is on the same day
  return `${startDayFormatted}, ${startTime} – ${endTime}`
}
