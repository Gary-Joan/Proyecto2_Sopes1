#!/usr/bin/env python
import pika, sys, os
import os
from flask import Flask, redirect, url_for, request, render_template, jsonify
import requests
from json import loads,  dumps
app = Flask(__name__)





@app.route('/')
def hello():
    return "<h2>SERVIDOR 2</h2>"


def main():
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()

    channel.queue_declare(queue='Case')

    def callback(ch, method, properties, body):
        print(" [x] Recibido %r" % body.decode())

    channel.basic_consume(queue='Case', on_message_callback=callback, auto_ack=True)

    print(' [*] Esperando mensajes. Presione CTRL+C para terminar')
    channel.start_consuming()

    

if __name__ == '__main__':
    try:
        main()
        
    except KeyboardInterrupt:
        print('Finalizo Programa')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)