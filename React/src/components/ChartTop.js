import React, {Component} from 'react';
import CanvasJSReact from './canvasjs.react';

var CanvasJS = CanvasJSReact.CanvasJS;
var CanvasJSChart = CanvasJSReact.CanvasJSChart;
var dataPoints =[];


class ChartTop extends Component{
    

    render(){
        const options = {
			theme: "light2",
			title: {
				text: "Top 3 Departamentos con mas Casos de Coronavirus"
			},
			data: [{
				type: "column",
				dataPoints: dataPoints
			}]
		}
        return (
            <div>
                <CanvasJSChart options = {options} 
                     onRef={ref => this.chart = ref}
                />
            </div>
            );
    }

    componentDidMount(){
		var chart = this.chart;
		fetch('http://localhost/topCasos')
		.then(function(response) {
			return response.json();
		})
		.then(function(data) {
			for (var i = 0; i < data.length; i++) {
				dataPoints.push({
					label: data[i].departamento,
					y: data[i].cantidad
				});
			}
			chart.render();
		});
	}

}




export default ChartTop;