import { useState, useEffect } from 'react';
import axios from 'axios';

const AutoCompleteInput = ({ label, placeholder, listID, onChange, setStart, setTarget }) => {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);

  useEffect(() => {
    const loadSuggestions = async () => {
      try {
        const response = await axios.get('https://en.wikipedia.org/w/api.php?origin=*', {
          params: {
            action: 'opensearch',
            format: 'json',
            search: searchTerm
          }
        });
        setSearchResults(response.data[1]);
      } catch (error) {
        console.error('Error fetching suggestions:', error);
      }
    };

    if (searchTerm) {
      loadSuggestions();
    } else {
      setSearchResults([]);
    }
  }, [searchTerm]);

  const handleChange = (event) => {
    setSearchTerm(event.target.value);
    onChange(event.target.value); // Call the onChange function passed from props
    if (label === "Start") {
      setStart(event.target.value); // Call setStart if label is "Start"
    } else if (label === "Target") {
      setTarget(event.target.value); // Call setTarget if label is "Target"
    }
  };

  return (
    <form className="flex flex-col gap-4">
      <div className="flex items-center">
        <label>{label}</label>
        <input
          className="bg-gray-400 hover:bg-gray-700 w-48 h-10 rounded-xl w-30 placeholder-white"
          type="text"
          value={searchTerm}
          onChange={handleChange}
          placeholder={placeholder}
          list={listID}
        />
      </div>
      <datalist id={listID}>
        {searchResults.map((suggestion, index) => (
          <option key={index} value={suggestion} />
        ))}
      </datalist>
    </form>
  );
};

export default AutoCompleteInput;
