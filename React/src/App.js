import logo from './logo.svg';
import './App.css';
import Chart from './components/Chart'
import ChartEdad from './components/ChartEdad'
import ChartTop from './components/ChartTop'
import ChartLast from './components/ChartLast'

function App() {
  return (
    <div className="App">
      <Chart />
      <ChartEdad />
      <ChartTop />
      <ChartLast />
    </div>
  );
}

export default App;
