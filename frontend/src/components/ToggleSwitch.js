import React from 'react'
import ReactSwitch from 'react-switch';

function ToggleSwitch({checked, onChange}) {
  return (
    <div className="flex items-center mb-4">
        <span className="mr-3">IDS</span>
        <ReactSwitch
            checked={checked}
            onChange={onChange}
            onColor="#2E51A2"
            offColor="#D1D5DB"
            checkedIcon={false}
            uncheckedIcon={false}
        />
        <span className="ml-3">BFS</span>
    </div>
  )
}

export default ToggleSwitch


