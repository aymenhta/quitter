# QUITTER
A minimalistic social media platform

## Models
Here are database tables....

### Users table
| Attributes  | Datatype     | Constraints     |
|    :----:   |    :----:    |    :----:       |
| id          | SERIAL       | PRIMARY KEY     |
| username    | VARCHAR(255) | UNIQUE NOT NULL |
| email       | VARCHAR(255) | UNIQUE NOT NULL |
| password    | CHAR(60)     | NOT NULL        |
| joined_at   | DATE         | -               |

### Posts table
| Attributes  | Datatype     | Constraints       |
|    :----:   |    :----:    |    :----:         |
| id          | SERIAL       | PRIMARY KEY       |
| content     | TEXT         | NOT NULL          |
| posted_at   | TIMESTAMP    | -                 |
| user_id     | SERIAL       | ON DELETE CASCADE |

## Tasks
    - [] Basic Crud Operations
      - [X] User Signup
      - [X] Posts Crud
      - [x] JSON Validation
    - [] JWT AuthN/Z
    - [] Advanced Crud
      - [] Posts Pagination
      - [] Reactions System (UP VOTE/DOWN VOTE)
      - [] Nested Comments System
      - [] Followers System
      - [] Write your own json-validator
    - [] Sending Emails
      - [] Email confirmations
      - [] User Activation
    - [] Notification system
    - [] Real-time chat
      - [] private 1-1 messages using websocket
      - [] private video calls using WebRTC
    - [] Building, Quality Control
    - [] Deployment & Hosting