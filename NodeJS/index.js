const express = require("express");
const bodyparser = require("body-parser");
const cors = require("cors");
const MongoClient = require('mongodb').MongoClient;
const redis = require('redis');
var redisClient = redis.createClient( {host:"34.69.11.162", port:6379, password:"admin", db:0, decode_responses:true} );

const uri = "mongodb+srv://Admin:admin@cluster0.etlwp.mongodb.net/mydb?retryWrites=true&w=majority";

var app = express();
app.use(cors());
app.use(bodyparser.json());
app.use(bodyparser.urlencoded({extended:true}));
var database;
var collection;

app.listen( 5000, () => {
	console.log("Server Running");
	MongoClient.connect(uri, {useNewUrlParser: true}, (error, client) => {
		if(error){
			throw error;
		}
		database = client.db("mydb");
               	collection = database.collection("casos");
               	console.log("Mongo Success!");
	});
});

app.get("/topCasos", (req, res) => {
	collection.find({}).toArray( (err, result) => {
		if (err){
			return res.status(500).send(err);
		}
		let responseJson = {};
		let tempJson = {};
		let tempArr = [];
		result.forEach( element => {
			if( ! (element['location'] in tempJson) ){
				tempJson[element['location']] = 1;
			} else {
				tempJson[element['location']] += 1;
			}
		});
		for ( const property in tempJson ) {
			tempArr.push(new Object({"departamento":property,"cantidad":tempJson[property]}));
		}
		tempArr.sort( (a,b) => (a.cantidad < b.cantidad) ? 1 : -1 );
		tempArr = tempArr.slice(0,3);
		res.send(tempArr)
	});
});

app.get("/ageCasos", (req, res) => {
        collection.find({}).toArray( (err, result) => {
                if (err){
                        return res.status(500).send(err);
                }
                let responseJson = {
			"0-10":0,
			"11-20":0,
			"21-30":0,
			"31-40":0,
			"41-50":0,
			"51-60":0,
			"61-70":0,
			"71-80":0,
			"81-90":0,
			"91-100":0
		};
                let tempArr = [];
                result.forEach( element => {
                        if( Number(element['age']) >= 0 && Number(element['age']) <= 10 ){
                                responseJson["0-10"] += 1;
                        } else if( Number(element['age']) >= 11 && Number(element['age']) <= 20 ){
                                responseJson["11-20"] += 1;
                        } else if( Number(element['age']) >= 21 && Number(element['age']) <= 30 ){
                                responseJson["21-30"] += 1;
                        } else if( Number(element['age']) >= 31 && Number(element['age']) <= 40 ){
                                responseJson["31-40"] += 1;
                        } else if( Number(element['age']) >= 41 && Number(element['age']) <= 50 ){
                                responseJson["41-50"] += 1;
                        } else if( Number(element['age']) >= 51 && Number(element['age']) <= 60 ){
                                responseJson["51-60"] += 1;
                        } else if( Number(element['age']) >= 61 && Number(element['age']) <= 70 ){
                                responseJson["61-70"] += 1;
                        } else if( Number(element['age']) >= 71 && Number(element['age']) <= 80 ){
                                responseJson["71-80"] += 1;
                        } else if( Number(element['age']) >= 81 && Number(element['age']) <= 90 ){
                                responseJson["81-90"] += 1;
                        } else if( Number(element['age']) >= 91 && Number(element['age']) <= 100 ){
                                responseJson["91-100"] += 1;
                        } else {}
                });
                for ( const property in responseJson ) {
                        tempArr.push(new Object({"rango_edades":property,"cantidad":responseJson[property]}));
                }
                res.send(tempArr)
        });
});

app.get("/allCasos", (req, res) => {
        collection.find({}).toArray( (err, result) => {
                if (err){
                        return res.status(500).send(err);
                }
                let responseJson = {};
                let tempJson = {};
                let tempArr = [];
                result.forEach( element => {
                        if( ! (element['location'] in tempJson) ){
                                tempJson[element['location']] = 1;
                        } else {
                                tempJson[element['location']] += 1;
                        }
                });
                for ( const property in tempJson ) {
                        tempArr.push(new Object({"departamento":property,"cantidad":tempJson[property]}));
                }
                tempArr.sort( (a,b) => (a.cantidad < b.cantidad) ? 1 : -1 );
                res.send(tempArr)
        });
});

app.get("/lastCaso", (req, res) => {
	redisClient.rpop("casos", (err, result) => {
		if(err){
			return res.status(500).send(err);
		}
		console.log(result);
		res.send(result)
	})
});
