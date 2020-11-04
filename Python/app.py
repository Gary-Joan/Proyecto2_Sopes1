#!/usr/bin/env python
import pika, sys
import os
from flask import Flask, redirect, url_for, request, render_template, jsonify
import requests
from json import loads,  dumps
from pymongo import MongoClient
import redis
app = Flask(__name__)

bod =""
def callback(ch, method, properties, body):
        print(" [x] Recibido %r" % body.decode())

        #insercion base de datos mongo
        myclient = MongoClient('mongodb://localhost:27017/',username='root',password='rootpassword')
        mydb = myclient["mydb"]
        mycol = mydb["casos"]
        x = mycol.insert_one(loads(body.decode())) 
        #insercion base de datos redis
        redisc = redis.StrictRedis(host='localhost', port=6379,db=0,charset="utf-8", decode_responses=True)
        parsed = loads(body.decode())
        redisc.rpush('casos', str(parsed))

        #recuperado = redis.mget('caso')
        #print ("tipo de variable de recuperado: "+str(recuperado))

                   
def main(): 
    
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
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