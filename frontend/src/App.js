import logo from './logo.svg';
import './App.css';
import Dropdown from 'react-dropdown';
import 'react-dropdown/style.css';

function App() {
  const options = [
    'e-mail', 'username', 'ip'
  ];
  const defaultOption = options[0];
  return (
    <div className="App">
      <header className="App-header">
        <h1>PwN-me</h1>
        <p id="smalltext">HIBP at home..</p>
        <div id="box">
          <h2 id="blacktext">Enter your:</h2>
          <Dropdown options={options} value={defaultOption} placeholder="Select an option"/>
          <input type="text" id="inputbox" placeholder="Enter here" />
          <button id="pwn">pwn me</button>
          <p id="smalltextt">NOTE: The app is not shipped with real databreaches. Only a dummy databreach. Refer to the README for users and emails.</p>
        </div>
      </header>
    </div>
  );
}

export default App;

