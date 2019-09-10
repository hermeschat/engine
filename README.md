# Hermes
### What is Hermes
Hermes is a scalable, GRPC based messaging service.
### Technology Stack
Nats-Streaming <br>
MongoDB <br>
Redis <br>
### How Hermes Works
Hermes uses Nats-Streaming as pub/sub service and uses nats concept of channels as 
a way of communication between two users. Each new event 
will be published to a channel and the eventHandler specific to the event (there are only two event handlers)
will deliver the event to the users using GRPC streaming.
#### Flow of new message
1. User creates a connection to hermes event buff with a valid JWT token.
2. Hermes subscribe user to user-discovery<br>
3. User sends a new message to Hermes<br>
4. Hermes finds target channel of the message (based on recipient, or if message itself has a channel)
5. Hermes makes sure that all members of target channel are subscribed to channel.
6. Hermes based on strategy and channel type decides whether to save the message or not.
7. Hermes sends message into channel (publishes message into nats).
8. other users that are subscribed to channel receive the message.
### Installation
```
   git clone github.com/hermeschat/engine
   cd hermes
   docker-compose up .
```
### License 
Hermes uses MIT licenese for more information read licence file.