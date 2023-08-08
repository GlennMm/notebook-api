# Notebook System Design

### User Requirements
  - User authentication
  - CRUD for notes 
  - Share notes 

### Models

- User
  - fullname
  - display name
  - password [hashed]
  - avatar
  - email

- Note
  - user_id
  - text
  - permission_id
  - state [draft/published]

- Note Permission
  - note_id
  - users_ids

### Functional requirements
- User auth
  - login
  - signup
  - password changed
  - user crud [update should be allowed to user own profilre only]

- User auth
  - CRUD for notes
    - Update should be allowed to user who have permission 
    - Delete should be allowed to user who creted the note 

- Share notes

------

## Technologies used

Programming language : Golang
Web framework : gorrila mux
Orm : 
<!-- Programming language : Golang
Programming language : Golang
Programming language : Golang -->