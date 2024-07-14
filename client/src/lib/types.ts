export type Event = {
  title: string
  description: string
  startDate: string
  endDate: string
  metadata: EventMetadata
}

type EventMetadata = {
  day?: string
  region?: string
  reddit_url?: string
  battlefy_url?: string
}

type EventGroup = {
  upcoming: Event[]
  recent: Event[]
}

export type EventsResponse = {
  algs: EventGroup
  other: EventGroup
}
