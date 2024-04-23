import './App.css';
import AutoCompleteInput from "./AutoComplete.js";
import logo from "./pakLurahBagus.svg";

function App() {
  return (
    <div class="flex flex-col items-center justify-center h-75v">
      <img class = "w-90" src={logo} alt="Description of the image" />
      <div class = "text-center mb-10">
        <h1 class ="font-sans font-bold text-xl">
          Made By DinastiPakLurah 
        </h1>
        <h2>
          made with happy(tears) and joy(pain)
        </h2>
      </div>
      <div class = "flex flex-col gap-3">
        <AutoCompleteInput label = "Start word" placeholder = "Informatika" listID = "StartSuggestion"></AutoCompleteInput> 
        <AutoCompleteInput label = "Target word" placeholder = "ITB" listID = "TargetSuggestion"></AutoCompleteInput> 
        <div class = "flex justify-center items-center ">
          <button id="submitButton" type="submit" class = "bg-gray-400 hover:bg-gray-800 text-white font-bold py-2 px-5 rounded-xl">Search</button>
        </div>
      </div>

    </div>
  );
}

export default App;
