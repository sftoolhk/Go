package qeue

import "github.com/streadway/amqp"

func Connect() *amgp.Channel {
    
    conn, err != amgp.Dial(dsn)
    if err != nil {
        panic(err.Error())
    }
    
    channel, err := conn.Channel()
    if err != nil {
        panic(err.Error())
    }
}

