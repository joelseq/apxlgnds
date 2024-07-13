export type Event = {
  title: string
  description: string
  startDate: string
  endDate: string
}

type EventGroup = {
  upcoming: Event[]
  recent: Event[]
}

export type EventsResponse = {
  algs: EventGroup
  other: EventGroup
}
