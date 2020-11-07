#!/usr/bin/env python
import pika, sys
import os
from flask import Flask, redirect, url_for, request, render_template, jsonify
import requests
from json import loads,  dumps
from pymongo import MongoClient
import redis
import pymongo

bod =""
def callback(ch, method, properties, body):
        print(" [x] Recibido %r" % body.decode())

        #insercion base de datos mongo

        myclient = pymongo.MongoClient("mongodb+srv://Admin:admin@cluster0.etlwp.mongodb.net/mydb?retryWrites=true&w=majority")
        db = myclient['mydb']
        collection = db['casos']
        x = collection.insert_one(loads(body.decode()))
        print(x)
        #insercion base de datos redis
        #redisc = redis.StrictRedis(host='localhost', port=6379,db=0,charset="utf-8", decode_responses=True)
        #parsed = loads(body.decode())
        #redisc.rpush('casos', str(parsed))

        #recuperado = redis.mget('caso')
        #print ("tipo de variable de recuperado: "+str(recuperado))

                   
def main(): 
    
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='mu-rabbit-rabbitmq.project.svc.cluster.local',port=5672))
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