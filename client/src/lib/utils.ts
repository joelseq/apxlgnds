import type { Event } from '$lib/types'

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
  const now = new Date()
  const startDate = new Date(event.startDate)
  const endDate = new Date(event.endDate)

  // Check if the event is happening now
  if (startDate <= now && endDate >= now) {
    return 'Now'
  }

  const diffDate = startDate.getTime() - now.getTime() > 0 ? startDate : endDate

  const diff = diffDate.getTime() - now.getTime()
  let suffix = 'left'
  if (diff < 0) {
    suffix = 'ago'
  }

  const absDiff = Math.abs(diff)

  // Calculate the difference in days, hours, minutes
  const days = Math.floor(absDiff / (1000 * 60 * 60 * 24))
  const hours = Math.floor((absDiff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
  const minutes = Math.floor((absDiff % (1000 * 60 * 60)) / (1000 * 60))

  if (days > 0) {
    if (hours > 0) {
      return `${days}d ${hours}h ${suffix}`
    } else {
      return `${days}d ${suffix}`
    }
  } else if (hours > 0) {
    if (minutes > 0) {
      return `${hours}h ${minutes}m ${suffix}`
    } else {
      return `${hours}h ${suffix}`
    }
  } else {
    return `${minutes}m ${suffix}`
  }
}
