import React, { useState } from 'react';
import AutoCompleteInput from './components/AutoComplete.js';
import NodeGraph from './Graph.js';

function sendData(start, target, searchAlgo) {
  console.log("azz",start, target)

  var data = {
    startLink: start,
    targetLink: target,
    searchType: searchAlgo
  }

  fetch('/search', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
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

function Search({ darkmode, searchAlgorithm }) {
  const [start, setStart] = useState('');
  const [target, setTarget] = useState('');
  const [startTemp, setStartTemp] = useState('');
  const [targetTemp, setTargetTemp] = useState('');
  const [showGraph, setShowGraph] = useState(false);

  const handleSearch = (event) => {
    if(startTemp){
      setStart(startTemp); 
      console.log("lmao start");
    }
    if(targetTemp){
      setTarget(targetTemp);
      console.log("lmao target");

    }
    if((startTemp && targetTemp)){
      setStart(startTemp);
      setTarget(targetTemp);
      setStartTemp("");
      setTargetTemp("");
      console.log("lmao end");
    }
    console.log(start, target)
    if (start && target) {
      sendData(start, target, searchAlgorithm);
      console.log(searchAlgorithm);
      console.log(start);
      setShowGraph(true);
    }
  };

    return (
    <div className = "z-10">
      <div className="flex flex-row items-center justify-center gap-4 relative">
        <AutoCompleteInput
          placeholder="Start"
          ID = "start"
          listID="StartSuggestion"
          onChange={(value) => setStartTemp(value)}
          setStart={setStartTemp} 
        />
        <div>
          <h1 className={`${darkmode ? 'text-white' : 'text-black'}`}>
            to
          </h1>
        </div>
        <AutoCompleteInput
          placeholder="Target"
          ID = "target"
          listID="TargetSuggestion"
          onChange={(value) => setTargetTemp(value)}
          setTarget={setTargetTemp} 
        />
      </div>
      <div className="flex justify-center items-center mt-4">
        <button id="submitButton" type="submit"  onClick={handleSearch} className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-5 rounded-xl z-10">Search</button>
      </div>
      {start && target && 
        <NodeGraph darkmode={darkmode} showGraph={showGraph}start={start} target={target} />
      }</div>
  );
}

export default Search;
