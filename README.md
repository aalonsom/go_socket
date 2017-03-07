# go_socket
Programs using sockets and concurrency with Go

This example shows how to handle events, which implies that the process has
to be waiting for arriving the event. In this case, there is a server and a 
a client. A number of later can be executed for interacting in parallel with the
server. The code is simple initially and additional functionalities are 
included. 

The interaction of a server with a set of clients is relied on sockets. The
client sends a message to the server, which replies with a simple text. There
are several versions of the server. 

The server has to be first, before starting clients. When the server is aborted,
all the clients are also finished. A client can be stopped, while the 
rest of the system can proceed. 

Server
------
This is the simplest code. The server opens a socket, binding it
for waiting at a port. Then it is waiting for listening from a client. When this
connects, then a socket connection is created for leting this client and the 
server to interact messages. A server can have a number of sockets connections
for different codes

When a socket connection is started, then the code creates a goroutine for
communicating with the client. If there are several sockets, the there should
be a corresponding goroutine (echoServer). It may happen that a client could send
a number of messages to this socket. Then, the server will handle them. If this
process may take some time to complete it. Given that the messages from the
connection socket is handled sequentially, so if some of this process is slow,
there could be a delay for handling them. 

(Note: the provided code is based on 
https://jan.newmarch.name/go/socket/chapter-socket.html. Thanks)

Server Concurrency
------------------
This version relies on adding a new goroutine (handler) when a message is 
received from a connection socket. So, if several messages have been received 
from a client, then can be handled in concurrency. A goroutine handles each of 
message. 

ServerConcPool
--------------
Although goroutines is efficient, some times the performance when creating a 
number of goroutines in a pool, waiting for a message. Initially, a set of
goroutines are created and waiting. When a message is received is redirected
one of them. 

ServerConcPoolDyn
-----------------
In the previous sample it may occur that all the goroutines are occupied, and 
the received message can be blocked, waiting for finishing one of them when
handling a message. This example can create a goroutine if the pool is 
exhausted
