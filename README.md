# Meeting developed in Golang with GraphQL

## Run project

```bash
go mod tidy
```

```bash
make server
```

## For modifications

```bash
make graphl-generate
```

## Query graphQL playground

```bash
query GeUserById {
  user(id: "ecbbaab4-1550-4d2b-abf7-802298543e88") {
    id
    username
    meetups {
      id
      name
      description
    }
  }
}

query GetMeetups {
  Meetup {
    id
    name
    description
    user {
   id
    }
  }
}

mutation CreateMeetup {
  createMeetup(input: {name:"meet xxxxx", description:"desc xxxxx"}) {
    id
    name
    description
  }
}

mutation UpdateMeetup {
  updateMeetup(id: "3eddf233-e996-4f02-9257-b389732fee3a", input: {name: "bbbb", description:"Desc Update"}) {
    id
    name
    description
    user {
      id
      username
      email
    }
  }
}

mutation DeleteMeetup {
  deleteMeetup(id: "09acfb93-624d-4b19-96d6-9cf33b72d03d")
}

```
