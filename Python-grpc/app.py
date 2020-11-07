#!/usr/bin/env python
import sys
import os
from flask import Flask, redirect, url_for, request, render_template, jsonify
from json import loads,  dumps
from pymongo import MongoClient
import redis
from concurrent import futures
import logging

import grpc
import helloworld_pb2
import helloworld_pb2_grpc
app = Flask(__name__)

class Greeter(helloworld_pb2_grpc.GreeterServicer):
    def SayHello(self,request,context):
        print("mensaje: %s!" % request.name )
        ## enviamos a mongo
        uri = "mongodb+srv://Admin:admin@cluster0.etlwp.mongodb.net/mydb?retryWrites=true&w=majority"
        myclient = MongoClient(uri)
        mydb = myclient["mydb"]
        mycol = mydb["casos"]
        x = mycol.insert_one(loads(request.name)) 
        ## enviamos a redis
        #redisc = redis.StrictRedis(host='localhost', port=6379,db=0,charset="utf-8", decode_responses=True)
        #parsed = loads(request.name)
        #redisc.rpush('casos', str(parsed))
        redisc = redis.Redis(
                        host='redis-16236.c238.us-central1-2.gce.cloud.redislabs.com',
                        port=16236,
                        password='aLUBDXvzv6pQVkyD5TGznWLGNdtbWFqJ')
        print(redisc)
        parsed = loads(request.name)
        redisc.rpush('casos', str(parsed))
        return helloworld_pb2.HelloReply(message = 'hola %s'%request.name)
                   
def main():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    helloworld_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()  

if __name__ == '__main__':
    try:
        logging.basicConfig()  
        main()    
    except KeyboardInterrupt:
        print('Finalizo Programa')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)