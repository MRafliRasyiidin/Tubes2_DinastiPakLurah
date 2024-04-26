import '../App.css';
import { useState } from 'react';
import Search from '../Search';
import Navbar from '../components/Navbar';
import lightLogo from "../Logo/logo-light.png";
import darkLogo from "../Logo/logo-dark.png";
import ToggleSwitch from "../components/ToggleSwitch";

function sendDataSearchType(type) {
  var dataMethod = {
    searchType: type
  }

  fetch('/type', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify(dataMethod)
  })
  .then(response => {
      if (response.ok) {
          console.log('Data sent successfully to Go backend');
      } else {
          console.error('Failed to send data to Go backend');
      }
  })
  .catch(error => {
      console.error('Error sending data to Go backend:', error);
  });
}

function App() {
  const [searchAlgorithm, setSearchAlgorithm] = useState('BFS');
  const [darkmode, setDarkMode] = useState(false); 
  const logo = darkmode ? darkLogo : lightLogo;

  const toggleAlgorithm = () => {
    setSearchAlgorithm(searchAlgorithm === 'BFS' ? 'IDS' : 'BFS');
    sendDataSearchType(searchAlgorithm);
    console.log(searchAlgorithm);
  };
  return (

    <div className={`flex flex-col items-center justify-center h-max w-auto`}>
        <Navbar darkmode={darkmode} />
        <img className="w-auto h-96 top-20" src={logo} alt="Description of the image" />
        <div className="text-center mb-10">
          <h1 className={`font-sans font-bold text-xl ${darkmode ? 'text-white' : 'text-black'}`}>
            Made By DinastiPakLurah 
          </h1>
          <h2 className={`${darkmode ? 'text-white' : 'text-black'}`} >
            made with Happy(Tears) and Joy(Pain)
          </h2>
        </div>

        <ToggleSwitch
          checked={searchAlgorithm === 'IDS'}
          onChange={toggleAlgorithm}
        />
        <Search />
        <button onClick={() => setDarkMode(!darkmode)} className={`rounded-lg fixed top-4 right-4 ${darkmode ? 'bg-white hover:bg-gray-600' : 'bg-gray-300 hover:bg-gray-700 hover:text-white'}`}>Dark Mode</button>
      </div>
  );
}

export default App;
