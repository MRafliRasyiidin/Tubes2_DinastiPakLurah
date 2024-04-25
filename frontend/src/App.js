import './App.css';
import { useState } from 'react';
import AutoCompleteInput from "./AutoComplete.js";
import lightLogo from "./logo-light.png";
import darkLogo from "./logo-dark.png"

function App() {
  const [darkmode, setDarkMode] = useState(false); 
  const logo = darkmode? darkLogo: lightLogo;
  return (
    <div class= {`flex flex-col items-center justify-center h-screen ${darkmode ? 'bg-black' : 'bg-white'}`}>
      <img class = "w-auto h-96 top-20" src={logo} alt="Description of the image" />
      <div class = "text-center mb-10">
        <h1 class ={`font-sans font-bold text-xl ${darkmode ? 'text-white' : 'text-black'}`}>
          Made By DinastiPakLurah 
        </h1>
        <h2 class = {`${darkmode? 'text-white' : 'text-black'}`} >
          made with happy(tears) and joy(pain)
        </h2>
      </div>
      <div class = "flex flex-row items-center justify-center gap-4 ">
        <AutoCompleteInput  placeholder = "Start" listID = "StartSuggestion"></AutoCompleteInput> 
        <div>
          <h1 class = {`${darkmode ? 'text-white' : 'text-black'}`}>
            to
          </h1>
        </div>
        <AutoCompleteInput  placeholder = "Target" listID = "TargetSuggestion"></AutoCompleteInput> 
      </div>
      <div class = "flex justify-center items-center mt-4">
        <button id="submitButton" type="submit" class = "bg-gray-400 hover:bg-gray-800 text-white font-bold py-2 px-5 rounded-xl">Search</button>
      </div>
      <button onClick={() => setDarkMode(!darkmode)} class = {`rounded-lg fixed top-4 right-4 ${darkmode ? 'bg-white hover:bg-gray-600' : 'bg-gray-300 hover:bg-gray-700 hover:text-white'}`}>Dark Mode</button>
    </div>
  );
}

export default App;
