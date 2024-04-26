import React, { useState, useEffect } from 'react';
import AutoCompleteInput from './components/AutoComplete.js';
import NodeGraph from './Graph.js';

function sendData(start, target, searchAlgo) {
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
  const [showGraph, setShowGraph] = useState(false);

  const handleStartChange = (value) => {
    setStart(value);
  };

  const handleTargetChange = (value) => {
    setTarget(value);
  };

//   useEffect(() => {
//     console.log(start);
//   }, [start]);

//   useEffect(() => {
//     console.log(target);
//   }, [target]);

useEffect(() => {
    console.log(showGraph);
  }, [showGraph]);

useEffect(() =>{
  if(start && target){
    setShowGraph(false);
  } 
}, [start, target]);

  const handleSearch = () => {
      if (start.length !== 0 && target.length !== 0) {
        sendData(start, target, searchAlgorithm);
        setShowGraph(true);
      }
  };
  console.log(searchAlgorithm);
  return (
    <div>
      <div className="flex flex-row items-center justify-center gap-4 ">
        <AutoCompleteInput
          placeholder="Start"
          listID="StartSuggestion"
          onChange={(value) => handleStartChange(value)}
          setStart={setStart} 
        />
        <div>
          <h1 className={`${darkmode ? 'text-white' : 'text-black'}`}>
            to
          </h1>
        </div>
        <AutoCompleteInput
          placeholder="Target"
          listID="TargetSuggestion"
          onChange={(value) => handleTargetChange(value)}
          setTarget={setTarget} 
        />
      </div>
      <div className="flex justify-center items-center mt-4">
        <button id="submitButton" type="submit" onClick={handleSearch}  className="bg-gray-400 hover:bg-gray-800 text-white font-bold py-2 px-5 rounded-xl">Search</button>
      </div>
      {showGraph && start && target && 
        <NodeGraph darkmode={darkmode} start={start} target={target} onRender={() => setShowGraph(!showGraph)} />
      }</div>
  );
}

export default Search;
