#!/usr/bin/env python
import pika, sys
import os
from flask import Flask, redirect, url_for, request, render_template, jsonify
import requests
from json import loads,  dumps
from pymongo import MongoClient
app = Flask(__name__)

bod =""
def callback(ch, method, properties, body):
        print(" [x] Recibido %r" % body.decode())
        myclient = MongoClient('mongodb://mongo:27017/',username='root',password='rootpassword')
        mydb = myclient["mydb"]
        mycol = mydb["casos"]
        mydict = { "name": "Peter", "address": "Lowstreet 27" }

        x = mycol.insert_one(loads(body.decode()))            
def main(): 
    
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='rabbitmq'))
    channel = connection.channel()

    channel.queue_declare(queue='Case')




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