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
  reddit?: RedditMetadata
  battlefy_url?: string
}

type RedditMetadata = {
  title: string
  url: string
}

type EventGroup = {
  upcoming: Event[]
  recent: Event[]
}

export type EventsResponse = {
  algs: EventGroup
  other: EventGroup
}
