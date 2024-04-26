import React from 'react'
import ReactSwitch from 'react-switch';

function ToggleSwitch({checked, onChange, leftInfo, rightInfo, info}) {
  return (
    <div className="mb-4">
        <p className = "text-center">{info}</p>
        <div className = "flex flex-row">
            <span className="mr-3">{leftInfo}</span>
            <ReactSwitch
                checked={checked}
                onChange={onChange}
                onColor="#2E51A2"
                offColor="#D1D5DB"
                checkedIcon={false}
                uncheckedIcon={false}
            />
            <span className="ml-3">{rightInfo}</span>
        </div>
    </div>
  )
}

export default ToggleSwitch


