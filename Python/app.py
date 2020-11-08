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
        pool = redis.ConnectionPool(host="34.69.11.162", port=6379, password="admin",db=0,decode_responses=True)
        r = redis.Redis(connection_pool=pool)
        parsed = loads(body.decode())
        string_json="\'"+str(parsed)+"\'"
        r.lpush('casos', string_json)

                   
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