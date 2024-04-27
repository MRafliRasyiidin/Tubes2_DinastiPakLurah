import React, { useState } from 'react';
import AutoCompleteInput from './components/AutoComplete.js';
import NodeGraph from './Graph.js';

function Search({ darkmode, searchAlgorithm }) {
  const [start, setStart] = useState('');
  const [target, setTarget] = useState('');
  const [startTemp, setStartTemp] = useState('');
  const [targetTemp, setTargetTemp] = useState('');
  const [showGraph, setShowGraph] = useState(false);

  const handleSearch = () => {
    if(startTemp){
      setStart(startTemp); 
    }
    if(targetTemp){
      setTarget(targetTemp);
    }
    if((startTemp && targetTemp)){
      setStart(startTemp);
      setTarget(targetTemp);
      setStartTemp("");
      setTargetTemp("");
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
