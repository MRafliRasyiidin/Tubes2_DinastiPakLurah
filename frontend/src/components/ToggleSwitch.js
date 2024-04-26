import React from 'react'
import ReactSwitch from 'react-switch';

function ToggleSwitch({checked, onChange, leftInfo, rightInfo, info, scolor = "#2E51A2",color="#D1D5DB" }) {
  return (
    <div className="mb-4">
        <p className = "text-center">{info}</p>
        <div className = "flex flex-row">
            <span className="mr-3">{leftInfo}</span>
            <ReactSwitch
                checked={checked}
                onChange={onChange}
                onColor= {scolor}
                offColor= {color} 
                checkedIcon={false}
                uncheckedIcon={false}
            />
            <span className="ml-3">{rightInfo}</span>
        </div>
    </div>
  )
}

export default ToggleSwitch


