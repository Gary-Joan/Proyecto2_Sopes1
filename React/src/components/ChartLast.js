import React, { Component } from 'react';
import './App.css';
import CanvasJSReact from './canvasjs.react';

var CanvasJS = CanvasJSReact.CanvasJS;
var CanvasJSChart = CanvasJSReact.CanvasJSChart;
var dataPoints =[];

class ChartLast extends Component{

  constructor(props){
    super(props);
    this.loadLastCase = this.loadLastCase.bind(this);
    this.state= {lastCaseData: {name: '',
    location: '',
    age: 0,
    infectedtype : '',
    state : ''}}
  }
    

    render(){
        return (
            <div>
              <h1>Ultimo dato ingresado</h1>
              <button class="new-button" onClick={this.loadLastCase}>Cargar Ultimo</button>
              <p>{ this.state.lastCaseData.name ? "Nombre: " + this.state.lastCaseData.name : '' }</p>
			  <p>{ this.state.lastCaseData.location ? "Ubicacion: " + this.state.lastCaseData.location : '' }</p>
			  <p>{ this.state.lastCaseData.age ? "Edad: " + this.state.lastCaseData.age : '' }</p>
			  <p>{ this.state.lastCaseData.infectedtype ? "Tipo de Infeccion: " + this.state.lastCaseData.infectedtype : '' }</p>
			  <p>{ this.state.lastCaseData.state ? "Estado: " + this.state.lastCaseData.state : '' }</p> 	
			</div>
            );
    }

    loadLastCase(){
      fetch('http://35.223.84.133/lastCaso')
      .then(function(response) {
        
        return response.json();
      })
      .then(data => {
        let newLastCaseData = JSON.parse(data.replaceAll("'","\""));
        console.log(newLastCaseData)
        this.setState({lastCaseData : newLastCaseData})
      }).catch(function(error) {
        console.log('No hay mas datos:');
      })
    }

    componentDidMount(){
		
	}

}





export default ChartLast;